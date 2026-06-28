package source

import (
	"testing"
)

func TestFragmentVarName_SimpleAccessor(t *testing.T) {
	got := fragmentVarName("user.name", map[string]int{})
	if got != "name" {
		t.Errorf("expected %q, got %q", "name", got)
	}
}

func TestFragmentVarName_ChainedMethod(t *testing.T) {
	got := fragmentVarName("log.response_time_ms", map[string]int{})
	if got != "response_time_ms" {
		t.Errorf("expected %q, got %q", "response_time_ms", got)
	}
}

func TestFragmentVarName_SanitizesOperators(t *testing.T) {
	// Defense layer: even if "||" somehow reaches fragmentVarName, output must be a valid Ruby identifier.
	got := fragmentVarName("||", map[string]int{})
	if got != "value" {
		t.Errorf("expected fallback %q, got %q", "value", got)
	}
}

func TestFragmentVarName_Dedup(t *testing.T) {
	seen := map[string]int{"count": 1}
	got := fragmentVarName("order.count", seen)
	if got != "count2" {
		t.Errorf("expected %q, got %q", "count2", got)
	}
}

// parseFragmentLine complexity tests for boolean operators.

func TestParseFragmentLine_OrOperatorIsComplex(t *testing.T) {
	line := `<td><%= log.response_time_ms || "—" %>ms</td>`
	relSlash := "app/views/dashboard/overview/index.html.erb"
	result := parseFragmentLine(line, line, relSlash, "/abs/index.html.erb", 10)
	if result == nil {
		t.Fatal("expected a FragmentLine result, got nil")
	}
	if !result.Complex {
		t.Errorf("expected Complex=true for || expression, got false. ComplexReason=%q", result.ComplexReason)
	}
	if result.PatchedLine != "" {
		t.Errorf("expected empty PatchedLine for complex result, got %q", result.PatchedLine)
	}
}

func TestParseFragmentLine_AndOperatorIsComplex(t *testing.T) {
	line := `<td><%= user.active? && user.verified? %>verified status</td>`
	relSlash := "app/views/admin/users/show.html.erb"
	result := parseFragmentLine(line, line, relSlash, "/abs/show.html.erb", 5)
	if result == nil {
		return // nil when text is too short — that's OK
	}
	if !result.Complex {
		t.Errorf("expected Complex=true for && expression, got false. ComplexReason=%q", result.ComplexReason)
	}
}

func TestParseFragmentLine_SimpleExprNotComplex(t *testing.T) {
	// No HTML tags — only interleaved ERB and plain text.
	line := `  Response time: <%= log.response_time_ms %>ms processed`
	relSlash := "app/views/dashboard/overview/index.html.erb"
	result := parseFragmentLine(line, line, relSlash, "/abs/index.html.erb", 10)
	if result == nil {
		t.Fatal("expected a FragmentLine result, got nil")
	}
	if result.Complex {
		t.Errorf("expected Complex=false for simple accessor, got true. ComplexReason=%q", result.ComplexReason)
	}
	if result.PatchedLine == "" {
		t.Errorf("expected non-empty PatchedLine for non-complex result")
	}
}
