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

// ── tCallRe ───────────────────────────────────────────────────────────────────
// These tests guard against regressions where tCallRe fails to detect a valid
// t() call form, which would cause the Audit to report the key as orphaned
// (defined but never used) — a false positive that can cause unintended pruning.

func TestTCallRe_MatchesI18nT(t *testing.T) {
	m := tCallRe.FindStringSubmatch(`I18n.t('admin.users.index.title')`)
	if len(m) < 2 || m[1] != "admin.users.index.title" {
		t.Errorf("expected key %q, got %v", "admin.users.index.title", m)
	}
}

func TestTCallRe_MatchesI18nT_DoubleQuotes(t *testing.T) {
	m := tCallRe.FindStringSubmatch(`I18n.t("admin.sidebar.analytics")`)
	if len(m) < 2 || m[1] != "admin.sidebar.analytics" {
		t.Errorf("expected key %q, got %v", "admin.sidebar.analytics", m)
	}
}

func TestTCallRe_MatchesBareT(t *testing.T) {
	m := tCallRe.FindStringSubmatch(`t('some.key')`)
	if len(m) < 2 || m[1] != "some.key" {
		t.Errorf("expected key %q, got %v", "some.key", m)
	}
}

func TestTCallRe_MatchesTranslate(t *testing.T) {
	m := tCallRe.FindStringSubmatch(`translate('shared.nav.title')`)
	if len(m) < 2 || m[1] != "shared.nav.title" {
		t.Errorf("expected key %q, got %v", "shared.nav.title", m)
	}
}

func TestTCallRe_RelativeKey(t *testing.T) {
	m := tCallRe.FindStringSubmatch(`t('.title')`)
	if len(m) < 2 || m[1] != ".title" {
		t.Errorf("expected relative key %q, got %v", ".title", m)
	}
}

func TestTCallRe_WithoutParens(t *testing.T) {
	// t 'key' (space, no parens) — older Rails style used in some helpers.
	m := tCallRe.FindStringSubmatch(`t 'admin.key'`)
	if len(m) < 2 || m[1] != "admin.key" {
		t.Errorf("expected key %q (no parens form), got %v", "admin.key", m)
	}
}

func TestTCallRe_NoMatchOnStringLiteral(t *testing.T) {
	// A plain string literal is not a t() call — must not be flagged as a key usage.
	if tCallRe.MatchString(`"some hardcoded string"`) {
		t.Error("tCallRe must not match plain string literals")
	}
}

func TestTCallRe_NoMatchOnVariableAssignment(t *testing.T) {
	// Variables whose names start with 't' (e.g. title_text) must not trigger.
	if tCallRe.MatchString(`title_text = "My Title"`) {
		t.Error("tCallRe must not match variable assignments")
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

// ── CheckBareTCalls ───────────────────────────────────────────────────────────
// bare t() is valid in controllers, mailers, helpers (ActionController/ActionMailer
// inject the helper). In models, jobs, services, workers, lib — it raises
// NoMethodError at runtime. These tests lock down both sides of the boundary.

func writeRubyFileForAudit(t *testing.T, relPath, content string) string {
	t.Helper()
	root := t.TempDir()
	full := filepath.Join(root, filepath.FromSlash(relPath))
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	return root
}

func TestCheckBareTCalls_ControllerIsValid(t *testing.T) {
	root := writeRubyFileForAudit(t, "app/controllers/foo_controller.rb",
		"redirect_to root_path, notice: t('flash.done')\n")
	issues, err := CheckBareTCalls(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 0 {
		t.Errorf("controller: expected 0 issues, got %d: %v", len(issues), issues)
	}
}

func TestCheckBareTCalls_MailerIsValid(t *testing.T) {
	root := writeRubyFileForAudit(t, "app/mailers/foo_mailer.rb",
		"subject t('mailers.foo.subject')\n")
	issues, err := CheckBareTCalls(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 0 {
		t.Errorf("mailer: expected 0 issues, got %d: %v", len(issues), issues)
	}
}

func TestCheckBareTCalls_HelperIsValid(t *testing.T) {
	root := writeRubyFileForAudit(t, "app/helpers/application_helper.rb",
		"def page_title; t('shared.title'); end\n")
	issues, err := CheckBareTCalls(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 0 {
		t.Errorf("helper: expected 0 issues, got %d: %v", len(issues), issues)
	}
}

func TestCheckBareTCalls_ModelIsInvalid(t *testing.T) {
	root := writeRubyFileForAudit(t, "app/models/user.rb",
		"def label; t('models.user.label'); end\n")
	issues, err := CheckBareTCalls(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 1 {
		t.Fatalf("model: expected 1 issue, got %d: %v", len(issues), issues)
	}
	if issues[0].Key != "models.user.label" {
		t.Errorf("Key = %q, want %q", issues[0].Key, "models.user.label")
	}
}

func TestCheckBareTCalls_JobIsInvalid(t *testing.T) {
	root := writeRubyFileForAudit(t, "app/jobs/notify_job.rb",
		"message = t('jobs.notify.body')\n")
	issues, err := CheckBareTCalls(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 1 {
		t.Fatalf("job: expected 1 issue, got %d: %v", len(issues), issues)
	}
}

func TestCheckBareTCalls_LibIsInvalid(t *testing.T) {
	root := writeRubyFileForAudit(t, "lib/utility.rb",
		"puts t('lib.util.message')\n")
	issues, err := CheckBareTCalls(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 1 {
		t.Fatalf("lib: expected 1 issue, got %d: %v", len(issues), issues)
	}
}

func TestCheckBareTCalls_I18nTInModelIsOK(t *testing.T) {
	root := writeRubyFileForAudit(t, "app/models/user.rb",
		"def label; I18n.t('models.user.label'); end\n")
	issues, err := CheckBareTCalls(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 0 {
		t.Errorf("I18n.t in model: expected 0 issues, got %d: %v", len(issues), issues)
	}
}

func TestCheckBareTCalls_CommentSkipped(t *testing.T) {
	root := writeRubyFileForAudit(t, "app/models/user.rb",
		"# t('models.user.label') — example usage\n")
	issues, err := CheckBareTCalls(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 0 {
		t.Errorf("comment line: expected 0 issues, got %d: %v", len(issues), issues)
	}
}
