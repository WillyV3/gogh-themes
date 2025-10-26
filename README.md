# gogh-themes

**361 professional terminal color schemes for Go applications**

A zero-dependency Go package providing 361 terminal color themes sourced from the [Gogh](https://github.com/Gogh-Co/Gogh) collection. All themes are compiled into pure Go code for maximum performance - no runtime YAML parsing required.

## Installation

```bash
go get github.com/willyv3/gogh-themes
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/willyv3/gogh-themes"
)

func main() {
    // Get all 361 themes
    allThemes := goghthemes.All()
    fmt.Printf("Total themes: %d\n", len(allThemes))

    // Get a specific theme
    theme, ok := goghthemes.Get("Dracula")
    if ok {
        fmt.Printf("Background: %s\n", theme.Background)
        fmt.Printf("Foreground: %s\n", theme.Foreground)

        // Primary colors
        fmt.Printf("Blue: %s\n", theme.Blue)
        fmt.Printf("Red: %s\n", theme.Red)

        // Bright colors (NEW in v1.1.0!)
        fmt.Printf("BrightBlue: %s\n", theme.BrightBlue)
        fmt.Printf("BrightRed: %s\n", theme.BrightRed)
    }

    // Get all theme names (sorted alphabetically)
    names := goghthemes.Names()
    for _, name := range names {
        fmt.Println(name)
    }
}
```

## Theme Structure

Each theme provides **full 16-color ANSI support**:

```go
type Theme struct {
    Name       string // Theme name
    Background string // Background color (hex)
    Foreground string // Foreground/text color (hex)

    // Primary colors (ANSI 0-7)
    Black      string // ANSI color 0
    Red        string // ANSI color 1
    Green      string // ANSI color 2
    Yellow     string // ANSI color 3
    Blue       string // ANSI color 4
    Magenta    string // ANSI color 5
    Cyan       string // ANSI color 6
    White      string // ANSI color 7

    // Bright colors (ANSI 8-15)
    BrightBlack   string // ANSI color 8
    BrightRed     string // ANSI color 9
    BrightGreen   string // ANSI color 10
    BrightYellow  string // ANSI color 11
    BrightBlue    string // ANSI color 12
    BrightMagenta string // ANSI color 13
    BrightCyan    string // ANSI color 14
    BrightWhite   string // ANSI color 15
}
```

## Popular Themes

Some of the included themes:
- Dracula
- Nord
- Gruvbox Dark
- Catppuccin (Mocha, Frappé, Latte, Macchiato)
- Tokyo Night (Standard, Storm, Light)
- Solarized (Dark, Light)
- Monokai Pro
- One Dark
- Ayu (Dark, Mirage, Light)
- And 340+ more!

## Features

- **Zero dependencies** - Pure Go, no runtime parsing
- **361 themes** - Professional color schemes
- **Full 16-color ANSI support** - All primary + bright colors
- **Fast** - All themes compiled into binary
- **Small** - ~200KB of Go code
- **Type-safe** - Full Go struct definitions
- **Well-documented** - Clear API with examples
- **Lipgloss integration** - Optional subpackage for Bubble Tea TUIs

## Lipgloss Integration (Bubble Tea TUIs)

Want to use these themes in [Bubble Tea](https://github.com/charmbracelet/bubbletea) applications? Import the `lipgloss` subpackage for pre-wrapped `lipgloss.Color` types:

```bash
go get github.com/willyv3/gogh-themes/lipgloss
```

```go
import (
    "github.com/charmbracelet/lipgloss"
    lipglossthemes "github.com/willyv3/gogh-themes/lipgloss"
)

// Get theme with lipgloss.Color types
theme, _ := lipglossthemes.Get("Dracula")

// Use directly in styles
titleStyle := lipgloss.NewStyle().
    Foreground(theme.BrightBlue).
    Background(theme.Background).
    Bold(true)
```

See [lipgloss/README.md](lipgloss/README.md) for complete documentation and examples.

## API

### `All() map[string]Theme`
Returns a map of all 361 themes keyed by theme name.

### `Get(name string) (Theme, bool)`
Returns a specific theme by name. Case-sensitive.

### `Names() []string`
Returns all theme names sorted alphabetically.

## Credits

Color schemes sourced from [Gogh](https://github.com/Gogh-Co/Gogh) - a collection of terminal color schemes for GNOME Terminal, Pantheon Terminal, Tilix, and XFCE4 Terminal.

## License

MIT

## Projects Using gogh-themes

- [gittui](https://github.com/willyv3/gittui) - GitHub profile TUI with 364 themes
