package backend

import (
	"context"
	"embed"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func TestEncryptCancel(t *testing.T) {
	assert := assert.New(t)
	app := getTestApp()

	mockSaveFileDialog("", nil)
	_, err := app.Encrypt([]string{}, "")

	assert.Nil(err)
}

func TestEncrypt(t *testing.T) {
	assert := assert.New(t)
	app := getTestApp()
	outputDir, _ := os.MkdirTemp("", "sl_test_output")
	outputPath := filepath.Join(outputDir, "e.sla")
	input, _ := os.CreateTemp("", "sl_test_input")
	inputs := []string{input.Name()}
	pwd := "123456789"
	_ = os.WriteFile(input.Name(), []byte("testing"), 0776)
	defer os.Remove(input.Name())
	defer os.RemoveAll(outputDir)

	mockSaveFileDialog(outputPath, nil)
	mockEventsEmit()
	dTypeChan := mockMessageDialog()
	_, err := app.Encrypt(inputs, pwd)
	dialogType := <-dTypeChan
	_, outputNotExistErr := os.Stat(outputPath)

	assert.Nil(err)
	assert.Equal(dialogType, runtime.InfoDialog)
	assert.Nil(outputNotExistErr)
}

func TestEncryptFail(t *testing.T) {
	assert := assert.New(t)
	app := getTestApp()
	outputDir, _ := os.MkdirTemp("", "sl_test_output")
	outputPath := filepath.Join(outputDir, "e.sla")
	input, _ := os.CreateTemp("", "sl_test_input")
	inputs := []string{input.Name()}
	pwd := "short"
	_ = os.WriteFile(input.Name(), []byte("testing"), 0776)
	defer os.Remove(input.Name())
	defer os.RemoveAll(outputDir)

	mockSaveFileDialog(outputPath, nil)
	mockEventsEmit()
	dTypeChan := mockMessageDialog()
	_, err := app.Encrypt(inputs, pwd)
	dialogType := <-dTypeChan
	_, outputNotExistErr := os.Stat(outputPath)

	assert.Nil(err)
	assert.Equal(dialogType, runtime.ErrorDialog)
	assert.NotNil(outputNotExistErr)
}

func getTestApp() *App {
	app, _ := NewApp([]byte{}, embed.FS{})
	app.ctx = context.Background()
	return app
}

func mockMessageDialog() chan runtime.DialogType {
	dType := make(chan runtime.DialogType, 2)
	MessageDialog = func(ctx context.Context, dialogOptions runtime.MessageDialogOptions) (string, error) {
		dType <- dialogOptions.Type
		return "", nil
	}
	return dType
}

func mockSaveFileDialog(path string, err error) {
	SaveFileDialog = func(ctx context.Context, dialogOptions runtime.SaveDialogOptions) (string, error) {
		return path, err
	}
}

func mockEventsEmit() chan string {
	event := make(chan string, 10000)
	EventsEmit = func(ctx context.Context, eventName string, optionalData ...interface{}) {
		event <- eventName
	}
	return event
}
