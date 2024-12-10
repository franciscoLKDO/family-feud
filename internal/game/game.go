package game

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/franciscolkdo/family-feud/internal/family"
	"github.com/franciscolkdo/family-feud/internal/keymap"
	"github.com/franciscolkdo/family-feud/internal/round"
	"github.com/franciscolkdo/family-feud/internal/score"
	"github.com/franciscolkdo/family-feud/internal/style"
)

var _ tea.Model = Model{}

type models int

const (
	blueFamily models = iota
	redFamily
	totalScore
	rounds
)

type Phase int

const (
	FaceOff Phase = iota
	FullGuess
	Steal
)

type Model struct {
	Rounds     []round.Config
	roundCount int
	Models     map[models]tea.Model
	keyMap     keymap.KeyMap

	isFaceOff bool

	Width int
}

func (m Model) isQuitMsg(msg tea.Msg) bool {
	if msg, ok := msg.(tea.KeyMsg); ok {
		if key.Matches(msg, m.keyMap.Quit) {
			return true
		}
	}
	return false
}

func (m *Model) faceOff(msg tea.KeyMsg) tea.Cmd {
	if m.isFaceOff {
		m.isFaceOff = false
		if key.Matches(msg, m.keyMap.RedFamily) {
			return family.OnFamilySelection(family.Red)
		}
		if key.Matches(msg, m.keyMap.BlueFamily) {
			return family.OnFamilySelection(family.Blue)
		}

	}
	return nil
}

func (m *Model) nextRound() tea.Cmd {
	if m.roundCount < len(m.Rounds) {
		m.roundCount++
		m.Models[rounds] = round.New(m.Rounds[m.roundCount])
		m.Models[totalScore] = score.New()
	}
	m.isFaceOff = true
	return tea.Batch(family.OnFamilySelection(family.None))
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, model := range m.Models {
		cmds = append(cmds, model.Init())
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	// Check if quit is called
	if m.isQuitMsg(msg) {
		return m, tea.Quit
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
	case tea.KeyMsg:
		if key.Matches(msg, m.keyMap.ResetFaceOff) {
			m.isFaceOff = true
			cmds = append(cmds, family.OnFamilySelection(family.None))
		}
		if key.Matches(msg, m.keyMap.BlueFamily) || key.Matches(msg, m.keyMap.RedFamily) {
			cmds = append(cmds, m.faceOff(msg))
		}
		if key.Matches(msg, m.keyMap.GoodChoice) {
			cmds = append(cmds, round.OnGoodChoice(msg.String()))
		}
		if key.Matches(msg, m.keyMap.WrongChoice) {
			cmds = append(cmds, round.OnWrongChoice())
		}
		if key.Matches(msg, m.keyMap.WinRound) {
			cmds = append(cmds, score.OnWinRound())
		}
		if key.Matches(msg, m.keyMap.NextRound) {
			cmds = append(cmds, m.nextRound())
		}
	case score.WinRoundScoreMsg:
		cmds = append(cmds, family.OnFamilyWin(msg.Value))
	case round.ResultMsg:
		if msg.Status == round.Success {
			cmds = append(cmds, score.OnScore(msg.Points))
		} else {
			cmds = append(cmds, family.OnFamilyFail())
		}
	}
	for k, model := range m.Models {
		var cmd tea.Cmd
		m.Models[k], cmd = model.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var s strings.Builder

	left := lipgloss.Place(lipgloss.Width(m.Models[blueFamily].View()), lipgloss.Height(m.Models[rounds].View()), lipgloss.Left, lipgloss.Top, m.Models[blueFamily].View(), lipgloss.WithWhitespaceBackground(style.RootStyle.GetBackground()))
	center := lipgloss.Place(lipgloss.Width(m.Models[rounds].View()), lipgloss.Height(m.Models[rounds].View()), lipgloss.Left, lipgloss.Top, m.Models[rounds].View(), lipgloss.WithWhitespaceBackground(style.RootStyle.GetBackground()))
	right := lipgloss.Place(lipgloss.Width(m.Models[redFamily].View()), lipgloss.Height(m.Models[rounds].View()), lipgloss.Left, lipgloss.Top, m.Models[redFamily].View(), lipgloss.WithWhitespaceBackground(style.RootStyle.GetBackground()))
	body := lipgloss.JoinHorizontal(lipgloss.Top, left, center, right)

	points := lipgloss.Place(lipgloss.Width(body), lipgloss.Height(m.Models[totalScore].View()), lipgloss.Center, lipgloss.Center, m.Models[totalScore].View(), lipgloss.WithWhitespaceBackground(style.RootStyle.GetBackground()))

	s.WriteString(lipgloss.JoinVertical(lipgloss.Top, points, body))
	return m.center(s.String())
}

func (m Model) center(content string) string {
	return lipgloss.Place(m.Width, lipgloss.Height(content), lipgloss.Center, lipgloss.Center, content, lipgloss.WithWhitespaceBackground(style.DarkGray))
}

func New(cfg Config) tea.Model {
	m := Model{
		Rounds:     cfg.Rounds,
		roundCount: 0,
		Models: map[models]tea.Model{
			blueFamily: family.NewBlueFamily(cfg.BlueFamily),
			redFamily:  family.NewRedFamily(cfg.RedFamily),
			totalScore: score.New(),
		},
		isFaceOff: true,

		keyMap: keymap.DefaultKeyMap(),
	}
	m.Models[rounds] = round.New(m.Rounds[m.roundCount])
	return m
}
