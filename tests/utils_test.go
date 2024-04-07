package tests

import (
	"redirectfy/internal/utils"
	"testing"
)

func TestValidaNomeDeUsuario(t *testing.T) {
	got := utils.ValidaNomeDeUsuario("-vmgFQgkzqEcwnguY3MmLxf4HpOh88SuLyYDNfmSoviX4LHgPXC7QfZ78yuIQjuPTcyLhnlAKWj47xJ_7a0ny-IWdVUGsAkNfIqusbhSEW3j4wpDPBqEjaWm")
	expected := true

	if got != expected {
		t.Errorf("expected=%t got=%t", expected, got)
	}
}
