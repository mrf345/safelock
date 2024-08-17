package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type customTheme struct {
	fyne.Theme
	dark bool
}

func (t *customTheme) Color(color fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if t.dark {
		return t.Theme.Color(color, theme.VariantDark)
	} else {
		return t.Theme.Color(color, theme.VariantLight)
	}
}

func (t *customTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return t.Theme.Size(name) - 2
	}

	return t.Theme.Size(name)
}

func GetCustomTheme(dark bool) fyne.Theme {
	t := &customTheme{Theme: theme.DefaultTheme()}
	t.dark = dark
	return t
}
