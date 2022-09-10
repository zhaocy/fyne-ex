package themex

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/zhaocy/fyne-theme/font"
)

var themex *ThemeX

func init() {
	themex = New()
}

type Font string

type ThemeX struct {
	font Font
}

func New() *ThemeX {
	theme := new(ThemeX)

	return theme
}

var _ fyne.Theme = (*ThemeX)(nil)

func (*ThemeX) Font(t fyne.TextStyle) fyne.Resource {
	if t.Italic {
		return theme.DefaultTheme().Font(t)
	}
	return font.ResourceMSYHTTC
}

func (*ThemeX) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(c, v)
}

func (*ThemeX) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*ThemeX) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
