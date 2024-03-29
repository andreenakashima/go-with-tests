package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeating a character 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}
