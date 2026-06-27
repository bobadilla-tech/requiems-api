package source

import (
	"os"
	"path/filepath"
	"testing"
)

// ── ResolveRelativeKey ────────────────────────────────────────────────────────

func TestResolveRelativeKey(t *testing.T) {
	cases := []struct {
		name      string
		shortPath string
		leaf      string
		want      string
	}{
		{
			"erb view",
			"app/views/admin/users/index.html.erb",
			"title",
			"admin.users.index.title",
		},
		{
			"erb partial underscore stripped",
			"app/views/partials/apis_show/_endpoint_documentation.html.erb",
			"heading",
			"partials.apis_show.endpoint_documentation.heading",
		},
		{
			"haml view",
			"app/views/tools/email_validator/show.html.haml",
			"description",
			"tools.email_validator.show.description",
		},
		{
			"plain erb extension",
			"app/views/shared/_footer.erb",
			"links",
			"shared.footer.links",
		},
		{
			"ruby file",
			"app/views/sales_inquiries/new.rb",
			"heading",
			"sales_inquiries.new.heading",
		},
		{
			"leaf with leading dot stripped",
			"app/views/admin/dashboard/index.html.erb",
			".title",
			"admin.dashboard.index.title",
		},
		{
			"partials/ prefix",
			"partials/home/_hero.html.erb",
			"cta",
			"partials.home.hero.cta",
		},
		{
			"no views prefix fallback",
			"other/path/file.html.erb",
			"key",
			"other.path.file.key",
		},
		{
			"deeply nested",
			"app/views/admin/analytics/revenue/chart.html.erb",
			"label",
			"admin.analytics.revenue.chart.label",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ResolveRelativeKey(tc.shortPath, tc.leaf)
			if got != tc.want {
				t.Errorf("ResolveRelativeKey(%q, %q) = %q, want %q",
					tc.shortPath, tc.leaf, got, tc.want)
			}
		})
	}
}

// ── extractTCalls ─────────────────────────────────────────────────────────────

func TestExtractTCalls_Absolute(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "show.html.erb")
	os.WriteFile(path, []byte(`
<h1><%= t('tools.email_validator.show.heading') %></h1>
<p><%= t('tools.email_validator.show.description') %></p>
`), 0o644)

	res := extractTCalls(path, dir)
	if len(res.absolute) != 2 {
		t.Fatalf("expected 2 absolute keys, got %d: %v", len(res.absolute), res.absolute)
	}
	keys := map[string]bool{}
	for _, u := range res.absolute {
		keys[u.Key] = true
	}
	if !keys["tools.email_validator.show.heading"] {
		t.Error("missing tools.email_validator.show.heading")
	}
	if !keys["tools.email_validator.show.description"] {
		t.Error("missing tools.email_validator.show.description")
	}
}

func TestExtractTCalls_Relative(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "index.html.erb")
	os.WriteFile(path, []byte(`
<h1><%= t('.title') %></h1>
<p><%= t('.subtitle') %></p>
`), 0o644)

	res := extractTCalls(path, dir)
	if len(res.relative) != 2 {
		t.Fatalf("expected 2 relative keys, got %d: %v", len(res.relative), res.relative)
	}
}

func TestExtractTCalls_SkipsCommentLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "show.html.erb")
	os.WriteFile(path, []byte(`
<%# t('ignored.by.comment') %>
<%= t('real.key') %>
`), 0o644)

	res := extractTCalls(path, dir)
	keys := map[string]bool{}
	for _, u := range res.absolute {
		keys[u.Key] = true
	}
	if keys["ignored.by.comment"] {
		t.Error("comment line key should be skipped")
	}
	if !keys["real.key"] {
		t.Error("real key should be detected")
	}
}

func TestExtractTCalls_I18nT(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "helper.rb")
	os.WriteFile(path, []byte(`I18n.t('some.i18n.key')`), 0o644)

	res := extractTCalls(path, dir)
	if len(res.absolute) != 1 || res.absolute[0].Key != "some.i18n.key" {
		t.Errorf("I18n.t() not detected; got %v", res.absolute)
	}
}

func TestExtractTCalls_NonExistentFile(t *testing.T) {
	res := extractTCalls("/nonexistent/file.rb", "/")
	if len(res.absolute) != 0 || len(res.relative) != 0 {
		t.Errorf("expected empty result for nonexistent file, got %v", res)
	}
}

// ── Audit: missing key detection for resolved relative keys (fix 4a) ──────────

func TestAudit_RelativeKeyMissingDetected(t *testing.T) {
	// A relative key t('.title') resolved to a full key that has no YAML definition
	// should appear in MissingKeys (not silently excluded).
	root := t.TempDir()
	appViews := filepath.Join(root, "app", "views", "admin", "users")
	os.MkdirAll(appViews, 0o755)

	// View file with only a relative t() call whose target is undefined in YAML.
	viewFile := filepath.Join(appViews, "index.html.erb")
	os.WriteFile(viewFile, []byte(`<%= t('.title') %>`), 0o644)

	// definedKeys has nothing for admin.users.index.title.
	definedKeys := map[string]bool{
		"en.admin.users.index.other": true,
	}

	result, err := Audit(root, "en", definedKeys)
	if err != nil {
		t.Fatalf("Audit: %v", err)
	}

	found := false
	for _, m := range result.MissingKeys {
		if m.Key == "admin.users.index.title" {
			found = true
		}
	}
	if !found {
		t.Errorf("resolved relative key 'admin.users.index.title' should appear in MissingKeys; got: %v", result.MissingKeys)
	}
}

func TestAudit_OrphanedKey(t *testing.T) {
	root := t.TempDir()
	appViews := filepath.Join(root, "app", "views")
	os.MkdirAll(appViews, 0o755)

	// No view files use 'tools.unused.key'.
	definedKeys := map[string]bool{
		"en.tools.unused.key": true,
	}

	result, err := Audit(root, "en", definedKeys)
	if err != nil {
		t.Fatalf("Audit: %v", err)
	}

	found := false
	for _, k := range result.OrphanedKeys {
		if k == "en.tools.unused.key" {
			found = true
		}
	}
	if !found {
		t.Errorf("'en.tools.unused.key' should be orphaned; got: %v", result.OrphanedKeys)
	}
}

func TestAudit_UsedKeyNotOrphaned(t *testing.T) {
	root := t.TempDir()
	appViews := filepath.Join(root, "app", "views")
	os.MkdirAll(appViews, 0o755)

	viewFile := filepath.Join(appViews, "show.html.erb")
	os.WriteFile(viewFile, []byte(`<%= t('tools.heading') %>`), 0o644)

	definedKeys := map[string]bool{
		"en.tools.heading": true,
	}

	result, err := Audit(root, "en", definedKeys)
	if err != nil {
		t.Fatalf("Audit: %v", err)
	}

	for _, k := range result.OrphanedKeys {
		if k == "en.tools.heading" || k == "tools.heading" {
			t.Errorf("used key should not be orphaned: %q", k)
		}
	}
}

func TestAudit_MissingKeyDetected(t *testing.T) {
	root := t.TempDir()
	appViews := filepath.Join(root, "app", "views")
	os.MkdirAll(appViews, 0o755)

	viewFile := filepath.Join(appViews, "show.html.erb")
	os.WriteFile(viewFile, []byte(`<%= t('tools.undefined.key') %>`), 0o644)

	// definedKeys does NOT contain tools.undefined.key.
	definedKeys := map[string]bool{
		"en.tools.other.key": true,
	}

	result, err := Audit(root, "en", definedKeys)
	if err != nil {
		t.Fatalf("Audit: %v", err)
	}

	found := false
	for _, m := range result.MissingKeys {
		if m.Key == "tools.undefined.key" {
			found = true
		}
	}
	if !found {
		t.Errorf("'tools.undefined.key' should be in MissingKeys; got: %v", result.MissingKeys)
	}
}

// ── stripLangPrefix ───────────────────────────────────────────────────────────

func TestStripLangPrefix(t *testing.T) {
	cases := []struct {
		key, lang, want string
	}{
		{"en.tools.heading", "en", "tools.heading"},
		{"es.foo.bar", "es", "foo.bar"},
		{"fr.foo", "en", "fr.foo"}, // wrong lang → unchanged
		{"tools.heading", "en", "tools.heading"}, // no prefix → unchanged
	}
	for _, tc := range cases {
		got := stripLangPrefix(tc.key, tc.lang)
		if got != tc.want {
			t.Errorf("stripLangPrefix(%q, %q) = %q, want %q", tc.key, tc.lang, got, tc.want)
		}
	}
}
