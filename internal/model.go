package internal

import (
	"time"

	"github.com/bjess9/pr-pilot/internal/config"
	"github.com/bjess9/pr-pilot/internal/github"
	gh "github.com/google/go-github/v55/github"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg struct {
	err error
}

func (e errMsg) Error() string {
	return e.err.Error()
}

type refreshMsg struct{}

type model struct {
	table  table.Model
	prs    []*gh.PullRequest
	loaded bool
	err    error
	token  string
}

func InitialModel(token string) model {
	columns := createTableColumns()
	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(15),
	)
	t.Focus()

	return model{table: t, token: token}
}

func refreshCmd() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(60 * time.Second)
		return refreshMsg{}
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.fetchPRs, refreshCmd())
}

func (m model) fetchPRs() tea.Msg {
	cfg, err := config.LoadConfig()
	if err != nil {
		return errMsg{err}
	}
	prs, err := github.FetchOpenPRs(cfg.Repos, m.token)
	if err != nil {
		return errMsg{err}
	}
	return prs
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.table, cmd = m.table.Update(msg)

	switch msg := msg.(type) {
	case []*gh.PullRequest:
		m.loaded = true
		m.prs = msg
		rows := createTableRows(m.prs)
		m.table.SetRows(rows)
		return m, tea.Batch(cmd, refreshCmd())

	case refreshMsg:
		return m, tea.Batch(m.fetchPRs, refreshCmd())

	case errMsg:
		m.err = msg.err

	case error:
		m.err = msg

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.loaded && len(m.table.Rows()) > 0 {
				idx := m.table.Cursor()
				if idx >= 0 && idx < len(m.prs) {
					pr := m.prs[idx]
					prURL := pr.GetHTMLURL()
					if prURL != "" {
						return m, openURLCmd(prURL)
					}
				}
			}
		}
	}

	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return errorView(m.err)
	}
	if !m.loaded {
		return loadingView()
	}

	m.table.SetStyles(tableStyles())

	return baseStyle.Render(
		m.table.View() + "\n" + helpStyle.Render("↑/↓: Navigate  •  Enter: Open PR  •  q: Quit"),
	)
}
