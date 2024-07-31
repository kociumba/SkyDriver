package styles

import "github.com/charmbracelet/lipgloss"

var (
	padding = lipgloss.NewStyle().Padding(1).Margin(1)
	Faint   = lipgloss.NewStyle().Faint(true)

	ProductStyle      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFFFFF")).Inherit(padding)
	ProductIDStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#A500E0")).Inherit(padding)
	SellPriceStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#F7E701")).Inherit(padding)
	BuyPriceStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#F1B101")).Inherit(padding)
	DiffStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#F39F04")).Inherit(padding)
	WeeklychangeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F16100")).Inherit(padding)

	PredictionUP   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#3BD100")).Inherit(padding)
	PredictionDown = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#DD0000")).Inherit(padding)
)
