# gogh-themes/lipgloss

**Lipgloss integration for gogh-themes - Ready-to-use color themes for Bubble Tea TUIs**

This subpackage wraps all 361 gogh-themes colors as `lipgloss.Color` types, making them instantly usable in [Bubble Tea](https://github.com/charmbracelet/bubbletea) terminal user interfaces without any conversion.

## Installation

```bash
go get github.com/willyv3/gogh-themes/lipgloss
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/charmbracelet/lipgloss"
    lipglossthemes "github.com/willyv3/gogh-themes/lipgloss"
)

func main() {
    // Get a theme - all colors are lipgloss.Color ready!
    theme, _ := lipglossthemes.Get("Dracula")

    // Use directly in styles - no conversion needed
    titleStyle := lipgloss.NewStyle().
        Foreground(theme.Blue).
        Background(theme.Background).
        Bold(true).
        Padding(1, 2)

    errorStyle := lipgloss.NewStyle().
        Foreground(theme.BrightRed).
        Bold(true)

    successStyle := lipgloss.NewStyle().
        Foreground(theme.BrightGreen)

    // Render with styles
    fmt.Println(titleStyle.Render("Welcome to My TUI"))
    fmt.Println(errorStyle.Render("[ERROR] Error occurred"))
    fmt.Println(successStyle.Render("[OK] Success"))
}
```

## Complete Example - Themed TUI Component

```go
package main

import (
    "github.com/charmbracelet/lipgloss"
    lipglossthemes "github.com/willyv3/gogh-themes/lipgloss"
)

// ThemedUI creates a consistent UI using a gogh theme
type ThemedUI struct {
    theme         lipglossthemes.Theme
    titleStyle    lipgloss.Style
    subtitleStyle lipgloss.Style
    errorStyle    lipgloss.Style
    successStyle  lipgloss.Style
    borderStyle   lipgloss.Style
}

func NewThemedUI(themeName string) *ThemedUI {
    theme, _ := lipglossthemes.Get(themeName)

    return &ThemedUI{
        theme: theme,
        titleStyle: lipgloss.NewStyle().
            Foreground(theme.BrightBlue).
            Bold(true).
            Underline(true),
        subtitleStyle: lipgloss.NewStyle().
            Foreground(theme.Cyan),
        errorStyle: lipgloss.NewStyle().
            Foreground(theme.BrightRed).
            Bold(true),
        successStyle: lipgloss.NewStyle().
            Foreground(theme.BrightGreen).
            Bold(true),
        borderStyle: lipgloss.NewStyle().
            BorderStyle(lipgloss.RoundedBorder()).
            BorderForeground(theme.Magenta).
            Padding(1, 2),
    }
}

func (ui *ThemedUI) RenderDashboard(title, content string) string {
    titleText := ui.titleStyle.Render(title)
    contentText := content // Apply styles as needed

    return ui.borderStyle.Render(
        lipgloss.JoinVertical(lipgloss.Left, titleText, "", contentText),
    )
}
```

## Using All 16 Colors

All themes include **full 16-color ANSI support**:

```go
theme, _ := lipglossthemes.Get("Gruvbox")

// Primary colors (0-7)
lipgloss.NewStyle().Foreground(theme.Black)
lipgloss.NewStyle().Foreground(theme.Red)
lipgloss.NewStyle().Foreground(theme.Green)
lipgloss.NewStyle().Foreground(theme.Yellow)
lipgloss.NewStyle().Foreground(theme.Blue)
lipgloss.NewStyle().Foreground(theme.Magenta)
lipgloss.NewStyle().Foreground(theme.Cyan)
lipgloss.NewStyle().Foreground(theme.White)

// Bright colors (8-15) - Perfect for emphasis!
lipgloss.NewStyle().Foreground(theme.BrightBlack)   // Gray
lipgloss.NewStyle().Foreground(theme.BrightRed)     // Error highlights
lipgloss.NewStyle().Foreground(theme.BrightGreen)   // Success highlights
lipgloss.NewStyle().Foreground(theme.BrightYellow)  // Warnings
lipgloss.NewStyle().Foreground(theme.BrightBlue)    // Titles
lipgloss.NewStyle().Foreground(theme.BrightMagenta) // Special elements
lipgloss.NewStyle().Foreground(theme.BrightCyan)    // Links
lipgloss.NewStyle().Foreground(theme.BrightWhite)   // Max contrast
```

## API

### `All() map[string]Theme`
Returns all 361 themes as a map of theme names to lipgloss-wrapped Themes.

### `Get(name string) (Theme, bool)`
Returns a lipgloss-wrapped theme by name. Case-sensitive. Returns false if not found.

### `Names() []string`
Returns all theme names sorted alphabetically.

### `FromTheme(t goghthemes.Theme) Theme`
Converts a standard `goghthemes.Theme` to a lipgloss-wrapped `Theme`.

## Popular Themes

Try these popular themes in your TUI:
- **Dracula** - Dark purple theme with high contrast
- **Gruvbox** - Retro groove colors
- **Tokyo Night** - Modern dark blue
- **Nord** - Arctic, north-bluish color palette
- **One Dark** - Atom's iconic dark theme
- **Monokai** - Classic Sublime Text theme
- **Solarized Dark/Light** - Precision colors
- And 354 more!

## Why Use This?

- **Zero boilerplate** - Colors are pre-wrapped as `lipgloss.Color`
- **Instant theming** - Switch themes with one line
- **Professional colors** - 361 curated terminal themes
- **Full ANSI support** - All 16 colors available
- **Type-safe** - Full Go type checking
- **Zero runtime overhead** - Compile-time conversion

## Core Package

This subpackage depends on the core `github.com/willyv3/gogh-themes` package. If you don't need lipgloss integration, use the core package for a zero-dependency solution.

## License

MIT - Same as parent package
