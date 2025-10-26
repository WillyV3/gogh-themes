# gogh-themes Examples

Interactive examples demonstrating gogh-themes usage in real applications.

## Available Examples

### [bubbletea-demo](./bubbletea-demo)

**Interactive theme picker for Bubble Tea TUIs**

Demonstrates:
- Integrating gogh-themes/lipgloss in Bubble Tea apps
- Theme cycling and fuzzy search
- All 16 ANSI colors in action
- Dynamic style updates
- Dual-view overlay pattern

**Run it:**
```bash
cd bubbletea-demo
go run main.go
```

**Features:**
- 361 professional themes
- Press `t` to cycle themes
- Press `/` to search themes
- Live preview with multiple styled components

---

## Running Examples

Each example is a standalone Go module:

```bash
# Navigate to example directory
cd <example-name>

# Run directly
go run main.go

# Or build and run
go build
./<example-name>
```

## Using in Your Project

Examples use the `replace` directive during development:

```go
// go.mod
replace github.com/willyv3/gogh-themes => ../..
```

In your own projects, use the published version:

```bash
go get github.com/willyv3/gogh-themes/lipgloss@latest
```

## Contributing Examples

Have a cool use case? Examples welcome!

Guidelines:
- Self-contained (own go.mod)
- Clear README with usage
- Demonstrates specific features
- Clean, commented code
- Works with latest gogh-themes release

## More Resources

- [Main Package](https://github.com/willyv3/gogh-themes)
- [Lipgloss Docs](https://github.com/charmbracelet/lipgloss)
- [Bubble Tea Tutorial](https://github.com/charmbracelet/bubbletea/tree/master/tutorials)
