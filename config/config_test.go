package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("TODO_PORT", fmt.Sprint(wantPort))
	got, err := New()
	if err != nil {
		t.Errorf("Error: %+v\n", err)
	}

	if got.Port != wantPort {
		t.Errorf("got %d, want %d", got.Port, wantPort)
	}

	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("got %s, want %s", got.Env, wantEnv)
	}
}
