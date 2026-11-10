package config

import "testing"

func TestGetEnv(t *testing.T) {
	t.Run("returns the environment value when set", func(t *testing.T) {
		t.Setenv("APP_TEST_KEY", "from-env")

		if got := getEnv("APP_TEST_KEY", "fallback"); got != "from-env" {
			t.Fatalf("expected %q, got %q", "from-env", got)
		}
	})

	t.Run("falls back when the value is empty", func(t *testing.T) {
		t.Setenv("APP_TEST_KEY", "")

		if got := getEnv("APP_TEST_KEY", "fallback"); got != "fallback" {
			t.Fatalf("expected %q, got %q", "fallback", got)
		}
	})
}
