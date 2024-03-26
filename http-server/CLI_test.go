package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/andreenakashima/go-with-tests/http-server"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartCalled     bool
	StartCalledWith int

	FinishedCalled     bool
	FinishedCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishedCalledWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartCalledWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartCalledWith)
		}
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}
	})
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}

}
