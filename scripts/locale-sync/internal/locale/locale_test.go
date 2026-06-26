package locale

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// ── WriteEntries / UpsertEntries ──────────────────────────────────────────────

func TestWriteEntries_NewFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "en.yml")
	changed, err := WriteEntries(path, map[string]string{"en.foo.bar": "Hello"})
	if err != nil {
		t.Fatalf("WriteEntries: %v", err)
	}
	if !changed {
		t.Error("changed = false, want true for new file")
	}
	entries, _ := parseFile(path, dir)
	if !hasEntry(entries, "en.foo.bar", "Hello") {
		t.Errorf("entry en.foo.bar=Hello not found in written file; entries: %v", entries)
	}
}

func TestWriteEntries_SkipsExistingKey(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "en.yml")
	// Write initial value.
	WriteEntries(path, map[string]string{"en.foo": "Original"})
	// Try to overwrite — WriteEntries should skip existing keys.
	changed, err := WriteEntries(path, map[string]string{"en.foo": "Changed"})
	if err != nil {
		t.Fatalf("WriteEntries: %v", err)
	}
	if changed {
		t.Error("changed = true, want false when key already exists")
	}
	entries, _ := parseFile(path, dir)
	if !hasEntry(entries, "en.foo", "Original") {
		t.Error("existing value was overwritten by WriteEntries")
	}
}

func TestWriteEntries_NestedKey(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "en.yml")
	_, err := WriteEntries(path, map[string]string{"en.a.b.c.d": "Deep"})
	if err != nil {
		t.Fatalf("WriteEntries: %v", err)
	}
	entries, _ := parseFile(path, dir)
	if !hasEntry(entries, "en.a.b.c.d", "Deep") {
		t.Errorf("deep nested key not found; entries: %v", entries)
	}
}

func TestWriteEntries_MultipleKeys(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "en.yml")
	_, err := WriteEntries(path, map[string]string{
		"en.x": "X value",
		"en.y": "Y value",
	})
	if err != nil {
		t.Fatalf("WriteEntries: %v", err)
	}
	entries, _ := parseFile(path, dir)
	if !hasEntry(entries, "en.x", "X value") || !hasEntry(entries, "en.y", "Y value") {
		t.Error("not all keys written")
	}
}

func TestUpsertEntries_NilOverwrite_BehavesLikeWrite(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "en.yml")
	changed, err := UpsertEntries(path, map[string]string{"en.foo": "Hello"}, nil)
	if err != nil {
		t.Fatalf("UpsertEntries: %v", err)
	}
	if !changed {
		t.Error("changed = false for new key with nil shouldOverwrite")
	}
}

func TestUpsertEntries_OverwritesPlaceholder(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "en.yml")
	// First write a placeholder.
	WriteEntries(path, map[string]string{"en.msg": "TODO: translate"})
	// Now upsert with overwrite function that allows TODO: values.
	changed, err := UpsertEntries(path, map[string]string{"en.msg": "Real translation"},
		func(v string) bool { return strings.HasPrefix(v, "TODO:") })
	if err != nil {
		t.Fatalf("UpsertEntries: %v", err)
	}
	if !changed {
		t.Error("changed = false when overwriting placeholder")
	}
	entries, _ := parseFile(path, dir)
	if !hasEntry(entries, "en.msg", "Real translation") {
		t.Error("placeholder was not overwritten")
	}
}

func TestUpsertEntries_SkipsNonPlaceholder(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "en.yml")
	WriteEntries(path, map[string]string{"en.msg": "Real value"})
	changed, err := UpsertEntries(path, map[string]string{"en.msg": "New value"},
		func(v string) bool { return strings.HasPrefix(v, "TODO:") })
	if err != nil {
		t.Fatalf("UpsertEntries: %v", err)
	}
	if changed {
		t.Error("changed = true when shouldOverwrite returned false")
	}
	entries, _ := parseFile(path, dir)
	if !hasEntry(entries, "en.msg", "Real value") {
		t.Error("non-placeholder value was overwritten")
	}
}

// ── setYAMLPath conflict detection ────────────────────────────────────────────

func TestSetYAMLPath_Conflict(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "conflict.yml")
	// Write "en.foo" as a scalar.
	WriteEntries(path, map[string]string{"en.foo": "scalar"})
	// Try to write "en.foo.bar" — should fail with conflict error.
	changed, err := WriteEntries(path, map[string]string{"en.foo.bar": "nested"})
	if err == nil {
		t.Error("expected conflict error when scalar exists where mapping needed, got nil")
	}
	if changed {
		t.Error("changed should be false on conflict")
	}
}

func TestSetYAMLPath_SameKeyNoOp(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "en.yml")
	WriteEntries(path, map[string]string{"en.foo": "value"})
	// Writing the same key again with WriteEntries is a skip (not conflict).
	changed, err := WriteEntries(path, map[string]string{"en.foo": "value"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if changed {
		t.Error("changed = true for same key/value no-op")
	}
}

// ── buildSkeletonDoc ──────────────────────────────────────────────────────────

func TestBuildSkeletonDoc_LangPrefixReplaced(t *testing.T) {
	entries := []Entry{
		{Key: "en.tools.foo", Value: "Foo"},
		{Key: "en.tools.bar", Value: "Bar"},
	}
	doc, err := buildSkeletonDoc("es", entries, "TODO: translate")
	if err != nil {
		t.Fatalf("buildSkeletonDoc: %v", err)
	}
	data, err := marshalDoc(doc)
	if err != nil {
		t.Fatalf("marshalDoc: %v", err)
	}
	content := string(data)
	if !strings.Contains(content, "es:") {
		t.Errorf("expected 'es:' in skeleton, got:\n%s", content)
	}
	if strings.Contains(content, "en:") {
		t.Errorf("skeleton should not contain 'en:', got:\n%s", content)
	}
	if !strings.Contains(content, "TODO: translate") {
		t.Errorf("expected placeholder in skeleton, got:\n%s", content)
	}
}

func TestBuildSkeletonDoc_ConflictReturnsError(t *testing.T) {
	// "en.foo" as leaf, then "en.foo.bar" creates conflict at build time.
	entries := []Entry{
		{Key: "en.foo", Value: "scalar"},
		{Key: "en.foo.bar", Value: "nested"},
	}
	_, err := buildSkeletonDoc("es", entries, "")
	if err == nil {
		t.Error("expected error for conflicting skeleton entries, got nil")
	}
}

func TestBuildSkeletonDoc_EmptyEntries(t *testing.T) {
	doc, err := buildSkeletonDoc("fr", nil, "")
	if err != nil {
		t.Fatalf("buildSkeletonDoc with nil entries: %v", err)
	}
	if doc == nil {
		t.Error("expected non-nil doc for empty entries")
	}
}

// ── Scan ──────────────────────────────────────────────────────────────────────

func TestScan_ReadsEntries(t *testing.T) {
	root := makeLocaleFixture(t, "en", "tools", `
en:
  tools:
    heading: "Hello"
    sub:
      label: "World"
`)
	entries, err := Scan(root, "en")
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	if !hasEntry(entries, "en.tools.heading", "Hello") {
		t.Errorf("missing en.tools.heading; entries: %v", entries)
	}
	if !hasEntry(entries, "en.tools.sub.label", "World") {
		t.Errorf("missing en.tools.sub.label; entries: %v", entries)
	}
}

func TestScan_MissingDirectoryReturnsEmpty(t *testing.T) {
	entries, err := Scan("/nonexistent/path", "en")
	// Should not error, just return empty.
	if err != nil {
		t.Fatalf("Scan on missing dir returned error: %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("expected 0 entries, got %d", len(entries))
	}
}

func TestScan_Deterministic(t *testing.T) {
	root := makeLocaleFixture(t, "en", "tools", `
en:
  tools:
    a: "A"
    b: "B"
    c: "C"
`)
	e1, _ := Scan(root, "en")
	e2, _ := Scan(root, "en")

	keys1 := extractKeys(e1)
	keys2 := extractKeys(e2)
	sort.Strings(keys1)
	sort.Strings(keys2)
	if strings.Join(keys1, ",") != strings.Join(keys2, ",") {
		t.Errorf("Scan not deterministic:\n  run1: %v\n  run2: %v", keys1, keys2)
	}
}

// ── FindDuplicates ────────────────────────────────────────────────────────────

func TestFindDuplicates_KeyDups(t *testing.T) {
	entries := []Entry{
		{Key: "en.a.error", KeyName: "error", Value: "Oops", ShortPath: "en/a.en.yml"},
		{Key: "en.b.error", KeyName: "error", Value: "Oops", ShortPath: "en/b.en.yml"},
	}
	dups := FindDuplicates(entries, 2)
	if len(dups.KeyDups) != 1 {
		t.Fatalf("expected 1 KeyDupGroup, got %d", len(dups.KeyDups))
	}
	g := dups.KeyDups[0]
	if g.KeyName != "error" {
		t.Errorf("KeyName = %q, want %q", g.KeyName, "error")
	}
	if !g.SameValue {
		t.Error("SameValue should be true when values are identical")
	}
}

func TestFindDuplicates_DifferentValues(t *testing.T) {
	entries := []Entry{
		{Key: "en.a.title", KeyName: "title", Value: "Hello", ShortPath: "en/a.en.yml"},
		{Key: "en.b.title", KeyName: "title", Value: "World", ShortPath: "en/b.en.yml"},
	}
	dups := FindDuplicates(entries, 2)
	if len(dups.KeyDups) != 1 {
		t.Fatalf("expected 1 KeyDupGroup, got %d", len(dups.KeyDups))
	}
	if dups.KeyDups[0].SameValue {
		t.Error("SameValue should be false when values differ")
	}
}

func TestFindDuplicates_BelowMinOccurrences(t *testing.T) {
	entries := []Entry{
		{Key: "en.a.msg", KeyName: "msg", Value: "Hi", ShortPath: "en/a.en.yml"},
	}
	dups := FindDuplicates(entries, 2)
	if len(dups.KeyDups) != 0 {
		t.Errorf("expected 0 KeyDupGroups below minOccurrences, got %d", len(dups.KeyDups))
	}
}

func TestFindDuplicates_ValueDups(t *testing.T) {
	entries := []Entry{
		{Key: "en.a.label", KeyName: "label", Value: "Submit", ShortPath: "en/a.en.yml"},
		{Key: "en.b.cta", KeyName: "cta", Value: "Submit", ShortPath: "en/b.en.yml"},
	}
	dups := FindDuplicates(entries, 2)
	if len(dups.ValueDups) != 1 {
		t.Fatalf("expected 1 ValueDupGroup, got %d", len(dups.ValueDups))
	}
	if dups.ValueDups[0].Value != "Submit" {
		t.Errorf("ValueDup.Value = %q, want %q", dups.ValueDups[0].Value, "Submit")
	}
}

func TestFindDuplicates_SharedFileExcluded(t *testing.T) {
	entries := []Entry{
		{Key: "en.shared.error", KeyName: "error", Value: "Oops", ShortPath: "en/shared.en.yml"},
		{Key: "en.a.error", KeyName: "error", Value: "Oops", ShortPath: "en/a.en.yml"},
	}
	dups := FindDuplicates(entries, 2)
	// shared.en.yml entries excluded → only 1 non-shared entry → no dup.
	if len(dups.KeyDups) != 0 {
		t.Errorf("shared file entries should be excluded from duplicates, got %d groups", len(dups.KeyDups))
	}
}

// ── helpers ───────────────────────────────────────────────────────────────────

func hasEntry(entries []Entry, key, value string) bool {
	for _, e := range entries {
		if e.Key == key && e.Value == value {
			return true
		}
	}
	return false
}

func extractKeys(entries []Entry) []string {
	keys := make([]string, len(entries))
	for i, e := range entries {
		keys[i] = e.Key
	}
	return keys
}

// makeLocaleFixture creates a temp Rails locale directory:
// {root}/config/locales/{lang}/{topic}.{lang}.yml
func makeLocaleFixture(t *testing.T, lang, topic, content string) string {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, "config", "locales", lang)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	path := filepath.Join(dir, topic+"."+lang+".yml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return root
}
