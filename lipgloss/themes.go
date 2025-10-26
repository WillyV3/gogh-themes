package lipgloss

import (
	"github.com/charmbracelet/lipgloss"
	goghthemes "github.com/willyv3/gogh-themes"
)

// Theme wraps all colors as lipgloss.Color for direct use in Bubble Tea TUIs.
// All color fields are ready to use with lipgloss styles like:
//
//	theme, _ := lipglossthemes.Get("Dracula")
//	style := lipgloss.NewStyle().Foreground(theme.Blue).Background(theme.Background)
type Theme struct {
	Name       string
	Background lipgloss.Color
	Foreground lipgloss.Color

	// Primary colors (ANSI 0-7)
	Black   lipgloss.Color
	Red     lipgloss.Color
	Green   lipgloss.Color
	Yellow  lipgloss.Color
	Blue    lipgloss.Color
	Magenta lipgloss.Color
	Cyan    lipgloss.Color
	White   lipgloss.Color

	// Bright colors (ANSI 8-15)
	BrightBlack   lipgloss.Color
	BrightRed     lipgloss.Color
	BrightGreen   lipgloss.Color
	BrightYellow  lipgloss.Color
	BrightBlue    lipgloss.Color
	BrightMagenta lipgloss.Color
	BrightCyan    lipgloss.Color
	BrightWhite   lipgloss.Color
}

// All returns all 361 themes with lipgloss.Color wrapped values.
func All() map[string]Theme {
	themes := make(map[string]Theme, 361)
	for name, t := range goghthemes.All() {
		themes[name] = FromTheme(t)
	}
	return themes
}

// Get returns a lipgloss-wrapped theme by name (case-sensitive).
// Returns false if the theme doesn't exist.
func Get(name string) (Theme, bool) {
	t, ok := goghthemes.Get(name)
	if !ok {
		return Theme{}, false
	}
	return FromTheme(t), true
}

// Names returns all 361 theme names sorted alphabetically.
func Names() []string {
	return goghthemes.Names()
}

// FromTheme converts a standard goghthemes.Theme to a lipgloss-wrapped Theme.
// This is useful if you want to work with both representations.
func FromTheme(t goghthemes.Theme) Theme {
	return Theme{
		Name:       t.Name,
		Background: lipgloss.Color(t.Background),
		Foreground: lipgloss.Color(t.Foreground),

		// Primary colors
		Black:   lipgloss.Color(t.Black),
		Red:     lipgloss.Color(t.Red),
		Green:   lipgloss.Color(t.Green),
		Yellow:  lipgloss.Color(t.Yellow),
		Blue:    lipgloss.Color(t.Blue),
		Magenta: lipgloss.Color(t.Magenta),
		Cyan:    lipgloss.Color(t.Cyan),
		White:   lipgloss.Color(t.White),

		// Bright colors
		BrightBlack:   lipgloss.Color(t.BrightBlack),
		BrightRed:     lipgloss.Color(t.BrightRed),
		BrightGreen:   lipgloss.Color(t.BrightGreen),
		BrightYellow:  lipgloss.Color(t.BrightYellow),
		BrightBlue:    lipgloss.Color(t.BrightBlue),
		BrightMagenta: lipgloss.Color(t.BrightMagenta),
		BrightCyan:    lipgloss.Color(t.BrightCyan),
		BrightWhite:   lipgloss.Color(t.BrightWhite),
	}
}
