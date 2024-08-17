package gui

import (
	"errors"
	"testing"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestCancelProcessAction(t *testing.T) {
	assert := assert.New(t)
	called := make(chan bool, 1)
	action := Action{Cancel: func() { called <- true }}

	action.cancelProcess()

	assert.Nil(action.Cancel)
	assert.True(<-called)
}

func TestLayoutAfterHandleError(t *testing.T) {
	assert := assert.New(t)
	progErr := errors.New("something went wrong")
	window := test.NewTempWindow(t, container.NewVBox())
	action := Action{
		Footer:     container.NewVBox(),
		BusyFooter: container.NewVBox(),
		Window:     window,
	}

	action.handleProgressErr(progErr)

	assert.True(action.BusyFooter.Hidden)
	assert.False(action.Footer.Hidden)
}

func TestValidatePassword(t *testing.T) {
	assert := assert.New(t)
	action := &Action{}
	validPwd := "testing12345"
	invalidPwd := "test"

	assert.Nil(action.validatePassword(validPwd))
	assert.NotNil(action.validatePassword(invalidPwd))
}

func TestLayoutAfterHandleSuccess(t *testing.T) {
	assert := assert.New(t)
	window := test.NewTempWindow(t, container.NewVBox())
	action := &Action{
		Footer:     container.NewVBox(),
		BusyFooter: container.NewVBox(),
		ProgLabel:  widget.NewLabel("test"),
		ProgBar:    widget.NewProgressBar(),
		Window:     window,
	}

	action.handleSuccessDialog("test").Hide()

	assert.True(action.BusyFooter.Hidden)
	assert.False(action.Footer.Hidden)
	assert.Equal(action.ProgLabel.Text, "")
	assert.Equal(action.ProgBar.Value, 0.0)
}

func TestUpdateProgress(t *testing.T) {
	assert := assert.New(t)
	status := "encrypting stuff"
	percent := 25.8
	action := &Action{
		ProgLabel: widget.NewLabel("test"),
		ProgBar:   widget.NewProgressBar(),
	}

	action.ProgBar.Max = 100.0
	action.updateProgressBar(status, percent)

	assert.Equal(action.ProgLabel.Text, status)
	assert.Equal(action.ProgBar.Value, percent)
}

func TestSwitchDark(t *testing.T) {
	assert := assert.New(t)
	app := test.NewTempApp(t)
	light := widget.NewToolbarAction(nil, func() {})
	dark := widget.NewToolbarAction(nil, func() {})
	action := &Action{
		App:     app,
		ToolBar: widget.NewToolbar(light, dark),
	}

	action.switchDark()

	assert.Equal(app.Settings().Theme(), GetCustomTheme(true))
	assert.False(light.ToolbarObject().Visible())
	assert.True(dark.ToolbarObject().Visible())
}

func TestSwitchLight(t *testing.T) {
	assert := assert.New(t)
	app := test.NewTempApp(t)
	light := widget.NewToolbarAction(nil, func() {})
	dark := widget.NewToolbarAction(nil, func() {})
	action := &Action{
		App:     app,
		ToolBar: widget.NewToolbar(light, dark),
	}

	action.switchLight()

	assert.Equal(app.Settings().Theme(), GetCustomTheme(false))
	assert.False(dark.ToolbarObject().Visible())
	assert.True(light.ToolbarObject().Visible())
}
