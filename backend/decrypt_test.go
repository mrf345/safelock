package backend

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
	tempDir, _ := os.MkdirTemp("", "sl_test_output")
	inputPath, outputPath, encErr := getEncryptedPath(tempDir, pwd)
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
	tempDir, _ := os.MkdirTemp("", "sl_test_output")
	inputPath, outputPath, encErr := getEncryptedPath(tempDir, pwd)
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

func getEncryptedPath(dir, pwd string) (inputPath, outputPath string, err error) {
	var outputFile *os.File

	if outputFile, err = os.CreateTemp(dir, "e.sla"); err != nil {
		return
	}

	outputPath = outputFile.Name()
	inputPath = filepath.Join(dir, "test.txt")

	if err = os.WriteFile(inputPath, []byte("testing"), 0776); err != nil {
		return
	}

	ctx := context.TODO()
	if err = safelock.New().Encrypt(ctx, []string{inputPath}, outputFile, pwd); err != nil {
		panic(err)
	}

	os.Remove(inputPath)
	return
}

func mockOpenDirectoryDialog(path string, err error) {
	OpenDirectoryDialog = func(ctx context.Context, dialogOptions runtime.OpenDialogOptions) (string, error) {
		return path, err
	}
}
