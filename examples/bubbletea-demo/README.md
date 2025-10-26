# Bubble Tea Theme Demo

Interactive demo showcasing gogh-themes/lipgloss integration with Bubble Tea.

## Features

- **361 Professional Themes** - All themes from the Gogh collection
- **Theme Cycling** - Press `t` to cycle through themes
- **Fuzzy Search** - Press `/` to search and select themes
- **Live Preview** - See all 16 ANSI colors in action
- **Multiple Components** - Titles, status messages, borders, spinner

## Quick Start

```bash
# Run from this directory
go run main.go

# Or build and run
go build
./bubbletea-demo
```

## Keyboard Controls

- `t` - Cycle to next theme
- `/` - Open theme search (type to filter)
- `↑`/`↓` or `k`/`j` - Navigate theme list
- `Enter` - Select theme
- `Esc` - Cancel search
- `q` or `Ctrl+C` - Quit

## What It Demonstrates

### Theme Integration

```go
import lipglossthemes "github.com/willyv3/gogh-themes/lipgloss"

// Get a theme
theme, _ := lipglossthemes.Get("Dracula")

// Use colors directly
titleStyle := lipgloss.NewStyle().
    Foreground(theme.BrightBlue).
    Bold(true)
```

### All 16 ANSI Colors

The demo uses all 16 colors from each theme:

**Primary Colors (0-7):**
- Black, Red, Green, Yellow, Blue, Magenta, Cyan, White

**Bright Colors (8-15):**
- BrightBlack, BrightRed, BrightGreen, BrightYellow
- BrightBlue, BrightMagenta, BrightCyan, BrightWhite

### UI Components

- **Titles** (BrightBlue, Bold)
- **Subtitles** (Cyan)
- **Spinner** (BrightMagenta, Animated)
- **Status Messages**:
  - Success (BrightGreen)
  - Error (BrightRed)
  - Warning (BrightYellow)
- **Borders** (Magenta, Rounded)
- **Help Text** (BrightBlack/Gray)
- **Theme Badge** (BrightCyan on Background)

## Try These Themes

Press `/` and search for:
- `dracula` - High contrast purple theme
- `nord` - Arctic blue palette
- `gruvbox` - Retro warm colors
- `tokyo` - Tokyo Night dark blue
- `monokai` - Sublime Text classic
- `solarized` - Precision colors
- `catppuccin` - Pastel themes
- `one dark` - Atom's iconic theme

## Architecture

The example demonstrates:

- **Clean State Management** - Theme state separate from UI state
- **Dual View Pattern** - Main view + overlay picker
- **Live Filtering** - Real-time fuzzy search
- **Dynamic Styling** - All styles recreated on theme change
- **Keyboard Navigation** - Context-sensitive keybindings

## Use in Your Project

This example can serve as a template for:

1. **Adding theme support** to your Bubble Tea app
2. **Building theme pickers** with fuzzy search
3. **Implementing dual-view** overlays/modals
4. **Applying consistent colors** across components

Copy the theme management pattern and adapt to your needs!

## Dependencies

- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/bubbles` - UI components
- `github.com/charmbracelet/lipgloss` - Styling
- `github.com/willyv3/gogh-themes/lipgloss` - Theme colors

## Source

Main package: [github.com/willyv3/gogh-themes](https://github.com/willyv3/gogh-themes)
