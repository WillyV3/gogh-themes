package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	lipglossthemes "github.com/willyv3/gogh-themes/lipgloss"
)

type errMsg error

type model struct {
	// Original demo components
	spinner  spinner.Model
	quitting bool
	err      error

	// Theme system
	currentTheme   lipglossthemes.Theme
	themeName      string
	allThemeNames  []string
	themeIndex     int

	// Theme picker state
	pickerOpen     bool
	filterInput    textinput.Model
	filteredThemes []string
	selectedIdx    int

	// Window dimensions
	width  int
	height int
}

// Key bindings
var (
	quitKeys = key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	)
	themeToggleKey = key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "cycle theme"),
	)
	themePickerKey = key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "search themes"),
	)
)

func initialModel() model {
	// Initialize spinner
	s := spinner.New()
	s.Spinner = spinner.Dot

	// Initialize theme system - start with Dracula
	allNames := lipglossthemes.Names()
	theme, _ := lipglossthemes.Get("Dracula")

	// Find Dracula's index
	themeIdx := 0
	for i, name := range allNames {
		if name == "Dracula" {
			themeIdx = i
			break
		}
	}

	// Initialize filter input for theme picker
	ti := textinput.New()
	ti.Placeholder = "Search themes..."
	ti.CharLimit = 50
	ti.Width = 40

	m := model{
		spinner:        s,
		currentTheme:   theme,
		themeName:      "Dracula",
		allThemeNames:  allNames,
		themeIndex:     themeIdx,
		filterInput:    ti,
		filteredThemes: allNames,
		selectedIdx:    0,
	}

	// Apply initial theme to spinner
	m.applyThemeToSpinner()

	return m
}

func (m *model) applyThemeToSpinner() {
	m.spinner.Style = lipgloss.NewStyle().Foreground(m.currentTheme.BrightMagenta)
}

func (m *model) cycleTheme() {
	m.themeIndex = (m.themeIndex + 1) % len(m.allThemeNames)
	m.themeName = m.allThemeNames[m.themeIndex]
	theme, _ := lipglossthemes.Get(m.themeName)
	m.currentTheme = theme
	m.applyThemeToSpinner()
}

func (m *model) selectTheme(name string) {
	theme, ok := lipglossthemes.Get(name)
	if !ok {
		return
	}
	m.currentTheme = theme
	m.themeName = name

	// Update theme index
	for i, n := range m.allThemeNames {
		if n == name {
			m.themeIndex = i
			break
		}
	}

	m.applyThemeToSpinner()
}

func (m *model) filterThemes() {
	filter := strings.ToLower(m.filterInput.Value())
	if filter == "" {
		m.filteredThemes = m.allThemeNames
		return
	}

	// Simple fuzzy matching - substring match
	filtered := []string{}
	for _, name := range m.allThemeNames {
		if strings.Contains(strings.ToLower(name), filter) {
			filtered = append(filtered, name)
		}
	}
	m.filteredThemes = filtered

	// Reset selection if out of bounds
	if m.selectedIdx >= len(m.filteredThemes) {
		m.selectedIdx = 0
	}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		// Handle theme picker keys
		if m.pickerOpen {
			switch {
			case key.Matches(msg, key.NewBinding(key.WithKeys("esc"))):
				m.pickerOpen = false
				m.filterInput.Blur()
				return m, nil

			case key.Matches(msg, key.NewBinding(key.WithKeys("enter"))):
				if len(m.filteredThemes) > 0 && m.selectedIdx < len(m.filteredThemes) {
					m.selectTheme(m.filteredThemes[m.selectedIdx])
				}
				m.pickerOpen = false
				m.filterInput.Blur()
				m.filterInput.SetValue("")
				m.filterThemes()
				return m, nil

			case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
				if m.selectedIdx > 0 {
					m.selectedIdx--
				}
				return m, nil

			case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
				if m.selectedIdx < len(m.filteredThemes)-1 {
					m.selectedIdx++
				}
				return m, nil

			default:
				// Update filter input
				var cmd tea.Cmd
				m.filterInput, cmd = m.filterInput.Update(msg)
				m.filterThemes()
				m.selectedIdx = 0
				return m, cmd
			}
		}

		// Handle main view keys
		switch {
		case key.Matches(msg, quitKeys):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, themeToggleKey):
			m.cycleTheme()
			return m, nil

		case key.Matches(msg, themePickerKey):
			m.pickerOpen = true
			m.filterInput.Focus()
			m.selectedIdx = 0
			return m, textinput.Blink

		}
		return m, nil

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	if m.pickerOpen {
		return m.renderThemePicker()
	}

	return m.renderMainView()
}

func (m model) renderMainView() string {
	// Define styles using current theme
	titleStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.BrightBlue).
		Bold(true).
		Padding(1, 0)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.Cyan)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.currentTheme.Magenta).
		Padding(1, 2).
		Width(60)

	statusStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.BrightGreen)

	errorStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.BrightRed).
		Bold(true)

	warningStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.BrightYellow)

	helpStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.BrightBlack).
		Padding(1, 0)

	themeInfoStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.BrightCyan).
		Background(m.currentTheme.Background).
		Padding(0, 1)

	// Build content
	title := titleStyle.Render("Bubble Tea Theme Demo")
	subtitle := subtitleStyle.Render("Powered by gogh-themes/lipgloss")

	content := fmt.Sprintf(
		"%s\n\n%s  %s\n\n%s\n%s\n%s\n\n%s",
		subtitle,
		m.spinner.View(),
		"Loading with style...",
		statusStyle.Render("[OK] Themes working perfectly"),
		errorStyle.Render("[ERROR] This would be an error"),
		warningStyle.Render("[WARN] This would be a warning"),
		themeInfoStyle.Render(fmt.Sprintf("Current Theme: %s (%d/%d)", m.themeName, m.themeIndex+1, len(m.allThemeNames))),
	)

	box := boxStyle.Render(content)

	help := helpStyle.Render(
		"t: cycle theme  •  /: search themes  •  q: quit",
	)

	if m.quitting {
		return fmt.Sprintf("\n%s\n\n%s\n\n%s\n\n", title, box, help)
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", title, box, help)
}

func (m model) renderThemePicker() string {
	// Picker styles
	headerStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.BrightBlue).
		Bold(true).
		Padding(1, 0)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.currentTheme.Magenta).
		Padding(1, 2).
		Width(50)

	selectedStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.BrightGreen).
		Bold(true).
		Background(m.currentTheme.Background)

	normalStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.Foreground)

	helpStyle := lipgloss.NewStyle().
		Foreground(m.currentTheme.BrightBlack).
		Padding(1, 0)

	// Header
	header := headerStyle.Render("Theme Picker")

	// Filter input
	filterView := m.filterInput.View()

	// Theme list (show max 10)
	maxDisplay := 10
	startIdx := m.selectedIdx
	if startIdx > len(m.filteredThemes)-maxDisplay {
		startIdx = len(m.filteredThemes) - maxDisplay
	}
	if startIdx < 0 {
		startIdx = 0
	}

	var themeList strings.Builder
	themeList.WriteString(fmt.Sprintf("\nShowing %d of %d themes:\n\n", len(m.filteredThemes), len(m.allThemeNames)))

	for i := startIdx; i < len(m.filteredThemes) && i < startIdx+maxDisplay; i++ {
		theme := m.filteredThemes[i]
		if i == m.selectedIdx {
			themeList.WriteString(selectedStyle.Render("> " + theme))
		} else {
			themeList.WriteString(normalStyle.Render("  " + theme))
		}
		themeList.WriteString("\n")
	}

	content := fmt.Sprintf("%s\n\n%s", filterView, themeList.String())
	box := boxStyle.Render(content)

	help := helpStyle.Render(
		"↑/↓: navigate  •  enter: select  •  esc: cancel",
	)

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", header, box, help)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
