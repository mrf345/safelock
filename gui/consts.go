package gui

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
)

const (
	Version    = "1.0"
	Github     = "https://github.com/mrf345/safelock"
	DismissMsg = "cancel_form"
)

var (
	About = strings.Join([]string{
		fmt.Sprintf("Safelock version %s\n", Version),
		"This tool is a free open-source project licensed under",
		"Mozilla Public License version 2.0. For any suggestions,",
		"future updates or bug reports please visit:",
	}, "\n")
)

var DefaultSize = fyne.Size{Height: 420, Width: 550}
var logoSize = fyne.Size{Height: 280, Width: 250}
var formSize = fyne.Size{Height: 400, Width: 400}
