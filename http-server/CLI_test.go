package poker

import "testing"

func TestCLI(t *testing.T) {
	PlayerStore := &StubPlayerStore{}
	cli := &CLI{PlayerStore}
	cli.PlayPoker()

	if len(PlayerStore.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}
}
