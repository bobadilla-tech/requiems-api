package translate

import (
	"strings"
	"testing"
)

func TestProtect_PlaceholderRails(t *testing.T) {
	p := NewProtector(nil)
	pr := p.Protect("Hello %{name}!")
	if strings.Contains(pr.Text, "%{name}") {
		t.Error("expected %{name} to be replaced with token")
	}
	if pr.Restore(pr.Text) != "Hello %{name}!" {
		t.Errorf("round-trip failed: got %q", pr.Restore(pr.Text))
	}
}

func TestProtect_PlaceholderHandlebars(t *testing.T) {
	p := NewProtector(nil)
	pr := p.Protect("Click {{action}} now")
	if strings.Contains(pr.Text, "{{action}}") {
		t.Error("expected {{action}} to be replaced with token")
	}
	if pr.Restore(pr.Text) != "Click {{action}} now" {
		t.Errorf("round-trip failed: got %q", pr.Restore(pr.Text))
	}
}

func TestProtect_PrintfPositional(t *testing.T) {
	p := NewProtector(nil)
	cases := []string{
		"Count: %d items",
		"Rate: %08.2f%%",
		"Status: %s",
	}
	for _, s := range cases {
		pr := p.Protect(s)
		restored := pr.Restore(pr.Text)
		if restored != s {
			t.Errorf("Protect/Restore(%q) = %q, want original", s, restored)
		}
	}
}

func TestProtect_CustomWords(t *testing.T) {
	p := NewProtector([]string{"Requiems"})
	pr := p.Protect("Use Requiems to validate data")
	if strings.Contains(pr.Text, "Requiems") {
		t.Error("expected custom word to be replaced with token")
	}
	if pr.Restore(pr.Text) != "Use Requiems to validate data" {
		t.Errorf("round-trip failed: got %q", pr.Restore(pr.Text))
	}
}

func TestProtect_CustomWordLongestFirst(t *testing.T) {
	// "Pay Button" (longer) must be matched as one token, not split into
	// "Pay" + " " + "Button" by the shorter custom word "Pay".
	p := NewProtector([]string{"Pay", "Pay Button"})
	pr := p.Protect("Click Pay Button to proceed")
	if strings.Contains(pr.Text, "XTOKEN1X") {
		t.Error("expected only one token for 'Pay Button', got two (longest-first sort failed)")
	}
	if pr.Restore(pr.Text) != "Click Pay Button to proceed" {
		t.Errorf("round-trip failed: got %q", pr.Restore(pr.Text))
	}
}

func TestProtect_CustomWordBeforePlaceholder(t *testing.T) {
	// Custom word "TOKEN" appears in the original text and also as a substring
	// of the generated token names (XTOKEN0X, XTOKEN1X, …).
	// Correct processing order: custom words first, then standard placeholders.
	// If standard placeholders ran first, the generated "XTOKEN0X" token would
	// contain "TOKEN" and the subsequent custom-word pass would corrupt it.
	p := NewProtector([]string{"TOKEN"})
	pr := p.Protect("TOKEN is %{count}")
	if pr.Restore(pr.Text) != "TOKEN is %{count}" {
		t.Errorf("round-trip failed: got %q (suspected token order corruption)", pr.Restore(pr.Text))
	}
}

func TestRestore_MultipleTokens(t *testing.T) {
	p := NewProtector(nil)
	pr := p.Protect("Hello %{first} and %{last}!")
	if pr.Restore(pr.Text) != "Hello %{first} and %{last}!" {
		t.Errorf("multi-placeholder round-trip failed: got %q", pr.Restore(pr.Text))
	}
}

func TestRestore_NoOp(t *testing.T) {
	p := NewProtector(nil)
	const s = "Plain text with no placeholders"
	pr := p.Protect(s)
	if pr.Text != s {
		t.Errorf("expected text unchanged, got %q", pr.Text)
	}
	if pr.Restore(pr.Text) != s {
		t.Errorf("restore of unchanged text failed: got %q", pr.Restore(pr.Text))
	}
}
