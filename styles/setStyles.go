package styles

import "github.com/charmbracelet/lipgloss"

var (
	padding = lipgloss.NewStyle().Padding(1).Margin(1)
	Faint   = lipgloss.NewStyle().Faint(true)

	ProductStyle      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFFFFF")).Inherit(padding)
	ProductIDStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#A500E0")).Inherit(padding)
	SellPriceStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#F7B801")).Inherit(padding)
	BuyPriceStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#F18701")).Inherit(padding)
	DiffStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#F35B04")).Inherit(padding)
	WeeklychangeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Inherit(padding)
)
