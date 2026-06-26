package translate

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
)

// ── IsPlaceholder ─────────────────────────────────────────────────────────────

func TestIsPlaceholder(t *testing.T) {
	cases := []struct {
		name  string
		value string
		extra []string
		want  bool
	}{
		{"empty", "", nil, true},
		{"spaces only", "   ", nil, true},
		{"TODO colon space", "TODO: translate", nil, true},
		{"TODO no space", "TODO:translate", nil, true},
		{"TODO with text", "TODO: needs review", nil, true},
		{"FIXME", "FIXME: something", nil, true},
		{"FIXME no space", "FIXME:urgent", nil, true},
		{"real value", "Hello World", nil, false},
		{"real value spanish", "Hola mundo", nil, false},
		{"custom placeholder match", "PENDING", []string{"PENDING"}, true},
		{"custom placeholder non-match", "Other text", []string{"PENDING"}, false},
		{"empty extra placeholder ignored", "real", []string{""}, false},
		{"leading space trimmed TODO", "  TODO: translate  ", nil, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPlaceholder(tc.value, tc.extra...)
			if got != tc.want {
				t.Errorf("IsPlaceholder(%q, %v) = %v, want %v", tc.value, tc.extra, got, tc.want)
			}
		})
	}
}

// ── Protector ─────────────────────────────────────────────────────────────────

func TestProtector_RoundTrip(t *testing.T) {
	cases := []struct {
		name   string
		words  []string
		input  string
	}{
		{"no placeholders", nil, "Hello World"},
		{"empty string", nil, ""},
		{"rails named", nil, "Welcome %{name}!"},
		{"handlebars", nil, "Hello {{user}}"},
		{"printf named", nil, "Value: %<count>d"},
		{"printf s", nil, "Error: %s"},
		{"printf d", nil, "Count: %d"},
		{"printf f", nil, "Rate: %08.2f"},
		{"mixed placeholders", nil, "Hello %{name}, you have %d messages"},
		{"custom word", []string{"NeverBounce"}, "Validate with NeverBounce today"},
		{"custom word multiple", []string{"NeverBounce"}, "NeverBounce and NeverBounce again"},
		{"longest first precedence", []string{"Requiems API", "API"}, "Use Requiems API not just API"},
		{"custom + placeholder", []string{"IPstack"}, "IPstack returns %{ip} data"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewProtector(tc.words)
			protected := p.Protect(tc.input)
			restored := protected.Restore(protected.Text)
			if restored != tc.input {
				t.Errorf("round-trip failed:\n  input:    %q\n  protect:  %q\n  restored: %q", tc.input, protected.Text, restored)
			}
		})
	}
}

func TestProtector_TokensNotInOutput(t *testing.T) {
	p := NewProtector(nil)
	protected := p.Protect("Hello %{name}, you have %d messages")
	if protected.Text == "Hello %{name}, you have %d messages" {
		t.Error("Protect should have replaced placeholders with tokens")
	}
	// Tokens look like XTOKENnX; original patterns should not appear raw.
	for _, fragment := range []string{"%{name}", "%d"} {
		if contains(protected.Text, fragment) {
			t.Errorf("protected text still contains raw placeholder %q", fragment)
		}
	}
}

func TestProtector_LongestFirstPrecedence(t *testing.T) {
	// "Requiems API" should be tokenized as a unit, not as "API" leaving "Requiems "
	p := NewProtector([]string{"API", "Requiems API"})
	protected := p.Protect("Requiems API is great")
	restored := protected.Restore(protected.Text)
	if restored != "Requiems API is great" {
		t.Errorf("longest-first failed: got %q", restored)
	}
	// "Requiems API" should appear as single token, not two.
	if contains(protected.Text, "Requiems") {
		t.Errorf("protected text still contains 'Requiems' — longest-first not applied: %q", protected.Text)
	}
}

func TestProtector_EmptyCustomWords(t *testing.T) {
	p := NewProtector([]string{"", "", ""})
	input := "Hello %{name}"
	protected := p.Protect(input)
	if protected.Restore(protected.Text) != input {
		t.Error("empty custom words broke round-trip")
	}
}

// ── FindMissing ───────────────────────────────────────────────────────────────

func makeEntry(key, value, shortPath string) locale.Entry {
	return locale.Entry{Key: key, Value: value, ShortPath: shortPath}
}

func TestFindMissing_AbsentInTarget(t *testing.T) {
	src := []locale.Entry{makeEntry("en.foo.bar", "Hello", "en/foo.en.yml")}
	got := FindMissing(src, nil, "en", "es", "/root", "")
	if len(got) != 1 {
		t.Fatalf("expected 1 missing entry, got %d", len(got))
	}
	if got[0].TargetKey != "es.foo.bar" {
		t.Errorf("TargetKey = %q, want %q", got[0].TargetKey, "es.foo.bar")
	}
}

func TestFindMissing_PresentInTarget(t *testing.T) {
	src := []locale.Entry{makeEntry("en.foo.bar", "Hello", "en/foo.en.yml")}
	tgt := []locale.Entry{makeEntry("es.foo.bar", "Hola", "es/foo.es.yml")}
	got := FindMissing(src, tgt, "en", "es", "/root", "")
	if len(got) != 0 {
		t.Errorf("expected 0 missing entries, got %d", len(got))
	}
}

func TestFindMissing_PlaceholderTargetTreatedAsAbsent(t *testing.T) {
	src := []locale.Entry{makeEntry("en.foo.bar", "Hello", "en/foo.en.yml")}
	tgt := []locale.Entry{makeEntry("es.foo.bar", "TODO: translate", "es/foo.es.yml")}
	got := FindMissing(src, tgt, "en", "es", "/root", "")
	if len(got) != 1 {
		t.Errorf("placeholder target should be treated as absent, got %d", len(got))
	}
}

func TestFindMissing_CustomSkeletonPlaceholder(t *testing.T) {
	src := []locale.Entry{makeEntry("en.foo.bar", "Hello", "en/foo.en.yml")}
	tgt := []locale.Entry{makeEntry("es.foo.bar", "PENDING", "es/foo.es.yml")}
	got := FindMissing(src, tgt, "en", "es", "/root", "PENDING")
	if len(got) != 1 {
		t.Errorf("custom placeholder should be treated as absent, got %d", len(got))
	}
}

func TestFindMissing_EmptySourceValueSkipped(t *testing.T) {
	src := []locale.Entry{makeEntry("en.foo.bar", "", "en/foo.en.yml")}
	got := FindMissing(src, nil, "en", "es", "/root", "")
	if len(got) != 0 {
		t.Errorf("empty source value should be skipped, got %d", len(got))
	}
}

func TestFindMissing_TargetFilePathComputed(t *testing.T) {
	src := []locale.Entry{makeEntry("en.tools.foo", "Bar", "en/tools.en.yml")}
	got := FindMissing(src, nil, "en", "es", "/myroot", "")
	if len(got) == 0 {
		t.Fatal("expected 1 missing entry")
	}
	want := filepath.Join("/myroot", "config", "locales", "es", "tools.es.yml")
	if got[0].TargetFile != want {
		t.Errorf("TargetFile = %q, want %q", got[0].TargetFile, want)
	}
}

// ── Cache ─────────────────────────────────────────────────────────────────────

func TestCache_GetFromEmpty(t *testing.T) {
	c := &Cache{Entries: make(map[string]map[string]string)}
	if v, ok := c.Get("en", "Hello", "es"); ok {
		t.Errorf("expected cache miss, got %q", v)
	}
}

func TestCache_SetThenGet(t *testing.T) {
	c := &Cache{Entries: make(map[string]map[string]string)}
	c.Set("en", "Hello", "es", "Hola")
	v, ok := c.Get("en", "Hello", "es")
	if !ok || v != "Hola" {
		t.Errorf("Get() = (%q, %v), want (\"Hola\", true)", v, ok)
	}
}

func TestCache_SrcLangIsolation(t *testing.T) {
	// D2 regression: same text from different source languages must not collide.
	c := &Cache{Entries: make(map[string]map[string]string)}
	c.Set("en", "Hello", "es", "Hola")
	c.Set("fr", "Hello", "es", "Hola (from fr)")

	v, _ := c.Get("en", "Hello", "es")
	if v != "Hola" {
		t.Errorf("en→es: got %q, want %q", v, "Hola")
	}
	v2, _ := c.Get("fr", "Hello", "es")
	if v2 != "Hola (from fr)" {
		t.Errorf("fr→es: got %q, want %q", v2, "Hola (from fr)")
	}
}

func TestCache_TargetLangIsolation(t *testing.T) {
	c := &Cache{Entries: make(map[string]map[string]string)}
	c.Set("en", "Hello", "es", "Hola")
	c.Set("en", "Hello", "fr", "Bonjour")

	v1, _ := c.Get("en", "Hello", "es")
	v2, _ := c.Get("en", "Hello", "fr")
	if v1 != "Hola" || v2 != "Bonjour" {
		t.Errorf("target lang isolation failed: es=%q fr=%q", v1, v2)
	}
}

func TestCache_SaveAndReload(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cache.json")

	c := &Cache{path: path, Entries: make(map[string]map[string]string)}
	c.Set("en", "Hello", "es", "Hola")
	c.Set("en", "Goodbye", "fr", "Au revoir")
	if err := c.Save(); err != nil {
		t.Fatalf("Save: %v", err)
	}

	c2, err := LoadCache(path)
	if err != nil {
		t.Fatalf("LoadCache: %v", err)
	}
	v, ok := c2.Get("en", "Hello", "es")
	if !ok || v != "Hola" {
		t.Errorf("reload: en/Hello/es = (%q, %v), want (\"Hola\", true)", v, ok)
	}
	v2, ok2 := c2.Get("en", "Goodbye", "fr")
	if !ok2 || v2 != "Au revoir" {
		t.Errorf("reload: en/Goodbye/fr = (%q, %v), want (\"Au revoir\", true)", v2, ok2)
	}
}

func TestLoadCache_NonExistentFile(t *testing.T) {
	c, err := LoadCache("/nonexistent/path/cache.json")
	if err != nil {
		t.Fatalf("expected no error for missing file, got %v", err)
	}
	if len(c.Entries) != 0 {
		t.Errorf("expected empty cache, got %d entries", len(c.Entries))
	}
}

func TestLoadCache_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.json")
	os.WriteFile(path, []byte("not json {{{"), 0o644)
	_, err := LoadCache(path)
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

func TestLoadCache_ValidJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cache.json")
	data, _ := json.Marshal(map[string]map[string]string{
		"en\x00Hello": {"es": "Hola"},
	})
	os.WriteFile(path, data, 0o644)
	c, err := LoadCache(path)
	if err != nil {
		t.Fatalf("LoadCache: %v", err)
	}
	v, ok := c.Get("en", "Hello", "es")
	if !ok || v != "Hola" {
		t.Errorf("expected (\"Hola\", true), got (%q, %v)", v, ok)
	}
}

// ── CharCount ─────────────────────────────────────────────────────────────────

func TestCharCount(t *testing.T) {
	entries := []MissingEntry{
		{Entry: locale.Entry{Value: "Hello"}},    // 5 ASCII runes
		{Entry: locale.Entry{Value: "héllo"}},    // 5 runes, 6 UTF-8 bytes
		{Entry: locale.Entry{Value: "世界"}},      // 2 runes
	}
	got := CharCount(entries)
	want := 5 + 5 + 2
	if got != want {
		t.Errorf("CharCount = %d, want %d", got, want)
	}
}

func TestCharCount_Empty(t *testing.T) {
	if n := CharCount(nil); n != 0 {
		t.Errorf("CharCount(nil) = %d, want 0", n)
	}
}

// ── LoadDictionaryFile ────────────────────────────────────────────────────────

func TestLoadDictionaryFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "dict.txt")
	content := "# Brand names\nNeverBounce\n\nIPstack\n# comment\nAPI Ninjas\n"
	os.WriteFile(path, []byte(content), 0o644)

	words, err := LoadDictionaryFile(path)
	if err != nil {
		t.Fatalf("LoadDictionaryFile: %v", err)
	}
	want := []string{"NeverBounce", "IPstack", "API Ninjas"}
	if !equalStringSlice(words, want) {
		t.Errorf("words = %v, want %v", words, want)
	}
}

func TestLoadDictionaryFile_NonExistent(t *testing.T) {
	_, err := LoadDictionaryFile("/nonexistent/dict.txt")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestLoadDictionaryFile_OnlyCommentsAndBlanks(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.txt")
	os.WriteFile(path, []byte("# comment\n\n# another\n"), 0o644)
	words, err := LoadDictionaryFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(words) != 0 {
		t.Errorf("expected empty slice, got %v", words)
	}
}

// ── targetFilePath ────────────────────────────────────────────────────────────

func TestTargetFilePath(t *testing.T) {
	cases := []struct {
		root, shortPath, srcLang, tgtLang, want string
	}{
		{"/r", "en/tools.en.yml", "en", "es", "/r/config/locales/es/tools.es.yml"},
		{"/r", "en/admin.en.yml", "en", "fr", "/r/config/locales/fr/admin.fr.yml"},
		{"/r", "en/shared.en.yml", "en", "es", "/r/config/locales/es/shared.es.yml"},
	}
	for _, tc := range cases {
		got := targetFilePath(tc.root, tc.shortPath, tc.srcLang, tc.tgtLang)
		if got != tc.want {
			t.Errorf("targetFilePath(%q, %q, %q, %q) = %q, want %q",
				tc.root, tc.shortPath, tc.srcLang, tc.tgtLang, got, tc.want)
		}
	}
}

// ── helpers ───────────────────────────────────────────────────────────────────

func contains(s, sub string) bool {
	return len(sub) > 0 && len(s) >= len(sub) &&
		func() bool {
			for i := 0; i <= len(s)-len(sub); i++ {
				if s[i:i+len(sub)] == sub {
					return true
				}
			}
			return false
		}()
}

func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	ac, bc := make([]string, len(a)), make([]string, len(b))
	copy(ac, a)
	copy(bc, b)
	sort.Strings(ac)
	sort.Strings(bc)
	for i := range ac {
		if ac[i] != bc[i] {
			return false
		}
	}
	return true
}
