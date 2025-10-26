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
    theme, ok := goghthemes.Get("Tokyo Night")
    if ok {
        fmt.Printf("Background: %s\n", theme.Background)
        fmt.Printf("Foreground: %s\n", theme.Foreground)
        fmt.Printf("Blue: %s\n", theme.Blue)
    }

    // Get all theme names (sorted alphabetically)
    names := goghthemes.Names()
    for _, name := range names {
        fmt.Println(name)
    }
}
```

## Theme Structure

Each theme provides:

```go
type Theme struct {
    Name       string // Theme name
    Background string // Background color (hex)
    Foreground string // Foreground/text color (hex)
    Black      string // ANSI color 0
    Red        string // ANSI color 1
    Green      string // ANSI color 2
    Yellow     string // ANSI color 3
    Blue       string // ANSI color 4
    Magenta    string // ANSI color 5
    Cyan       string // ANSI color 6
    White      string // ANSI color 7
    Gray       string // ANSI color 8 (bright black)
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

- ⚡ **Zero dependencies** - Pure Go, no runtime parsing
- 🎨 **361 themes** - Professional color schemes
- 🚀 **Fast** - All themes compiled into binary
- 📦 **Tiny** - ~150KB of Go code
- 🔒 **Type-safe** - Full Go struct definitions
- 📝 **Well-documented** - Clear API with examples

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
