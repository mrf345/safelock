package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/mrf345/safelock-cli/safelock"
	"github.com/stretchr/testify/assert"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func TestDecryptCancel(t *testing.T) {
	assert := assert.New(t)
	app := getTestApp()

	mockOpenDirectoryDialog("", nil)
	_, err := app.Decrypt("", "")

	assert.Nil(err)
}

func TestDecrypt(t *testing.T) {
	assert := assert.New(t)
	app := getTestApp()
	pwd := "123456789"
	tempDir, encErr := os.MkdirTemp("", "sl_test_output")
	inputPath, outputPath, _ := getEncryptedPath(tempDir, pwd)
	defer os.RemoveAll(tempDir)

	assert.Nil(encErr)

	mockOpenDirectoryDialog(tempDir, nil)
	mockEventsEmit()
	dTypeChan := mockMessageDialog()
	_, err := app.Decrypt(outputPath, pwd)
	dialogType := <-dTypeChan
	_, outputNotExistErr := os.Stat(inputPath)

	assert.Nil(err)
	assert.Equal(dialogType, runtime.InfoDialog)
	assert.Nil(outputNotExistErr)
}

func TestDecryptFail(t *testing.T) {
	assert := assert.New(t)
	app := getTestApp()
	pwd := "123456789"
	tempDir, encErr := os.MkdirTemp("", "sl_test_output")
	inputPath, outputPath, _ := getEncryptedPath(tempDir, pwd)
	defer os.RemoveAll(tempDir)

	assert.Nil(encErr)

	mockOpenDirectoryDialog(tempDir, nil)
	mockEventsEmit()
	dTypeChan := mockMessageDialog()
	_, err := app.Decrypt(outputPath, "wrong pass")
	dialogType := <-dTypeChan
	_, outputNotExistErr := os.Stat(inputPath)

	assert.Nil(err)
	assert.Equal(dialogType, runtime.ErrorDialog)
	assert.NotNil(outputNotExistErr)
}

func getEncryptedPath(dir, pwd string) (string, string, error) {
	outputPath := filepath.Join(dir, "e.sla")
	inputPath := filepath.Join(dir, "test.txt")
	_ = os.WriteFile(inputPath, []byte("testing"), 0776)
	ctx := context.TODO()
	err := safelock.New().Encrypt(ctx, []string{inputPath}, outputPath, pwd)
	os.Remove(inputPath)
	return inputPath, outputPath, err
}

func mockOpenDirectoryDialog(path string, err error) {
	OpenDirectoryDialog = func(ctx context.Context, dialogOptions runtime.OpenDialogOptions) (string, error) {
		return path, err
	}
}
