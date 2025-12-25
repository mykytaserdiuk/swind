package scene

import "testing"

func TestMatchPercent_ASCII(t *testing.T) {
	s := &GameScene{text: "HELLO"}

	p, ok := s.matchPercent("HEXXO")
	if !ok {
		t.Fatalf("expected ok=true")
	}
	// H E X X O vs H E L L O -> matches H,E,O = 3 of 5 -> 60%
	if p != 60 {
		t.Fatalf("expected 60 got %d", p)
	}
}

func TestMatchPercent_FullMatch(t *testing.T) {
	s := &GameScene{text: "WACHALA"}
	p, ok := s.matchPercent("WACHALA")
	if !ok || p != 100 {
		t.Fatalf("expected 100%% match, got %d ok=%v", p, ok)
	}
}

func TestMatchPercent_Multibyte(t *testing.T) {
	s := &GameScene{text: "Привет"}
	// change one rune
	p, ok := s.matchPercent("Приект")
	if !ok {
		t.Fatalf("expected ok=true for equal-length multibyte strings")
	}
	// Пр и в е т vs Пр и е к т -> matches: П,р,и,?,? maybe 3 of 6 -> 50%
	// Let's compute expected by manual count: positions 0:P same,1:р same,2:и same,3:в vs е diff,4:е vs к diff,5:т same => matches 4/6 => 66%
	if p != 66 {
		t.Fatalf("expected 66 got %d", p)
	}
}

func TestMatchPercent_LengthMismatch(t *testing.T) {
	s := &GameScene{text: "SHORT"}
	_, ok := s.matchPercent("TOO LONG")
	if ok {
		t.Fatalf("expected ok=false for length mismatch")
	}
}
