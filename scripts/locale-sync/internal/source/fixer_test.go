package source

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildFixPlanUsesI18nTForPlainRubyFiles(t *testing.T) {
	root := writeSourceFile(t, "app/models/private_deployment_request.rb", []string{
		`errors.add(:base, "must include at least one service")`,
	})

	plan := BuildFixPlan([]HardcodedString{{
		File:     "app/models/private_deployment_request.rb",
		Line:     1,
		Text:     "must include at least one service",
		Category: "ruby",
	}}, map[string]string{
		"must include at least one service": "en.private_deployment_request.must_include_at_least_one_service",
	}, "en", root)

	if got := onlyFix(t, plan).Patched; !strings.Contains(got, "I18n.t('private_deployment_request.must_include_at_least_one_service')") {
		t.Fatalf("patched model line should use I18n.t; got %q", got)
	}
}

func TestBuildFixPlanKeepsTForErbFiles(t *testing.T) {
	root := writeSourceFile(t, "app/views/apis/index.html.erb", []string{
		`<p>Welcome aboard</p>`,
	})

	plan := BuildFixPlan([]HardcodedString{{
		File:     "app/views/apis/index.html.erb",
		Line:     1,
		Text:     "Welcome aboard",
		Category: "tag_content",
	}}, map[string]string{
		"Welcome aboard": "en.apis.index.welcome_aboard",
	}, "en", root)

	if got := onlyFix(t, plan).Patched; got != `<p><%= t('apis.index.welcome_aboard') %></p>` {
		t.Fatalf("patched ERB line should use t helper; got %q", got)
	}
}

func TestBuildFixPlanKeepsTForControllerRubyFiles(t *testing.T) {
	root := writeSourceFile(t, "app/controllers/apis_controller.rb", []string{
		`redirect_to root_path, alert: "Something went wrong"`,
	})

	plan := BuildFixPlan([]HardcodedString{{
		File:     "app/controllers/apis_controller.rb",
		Line:     1,
		Text:     "Something went wrong",
		Category: "ruby",
	}}, map[string]string{
		"Something went wrong": "en.apis.flash.something_went_wrong",
	}, "en", root)

	if got := onlyFix(t, plan).Patched; !strings.Contains(got, "t('apis.flash.something_went_wrong')") {
		t.Fatalf("patched controller line should use t helper; got %q", got)
	}
	if got := onlyFix(t, plan).Patched; strings.Contains(got, "I18n.t") {
		t.Fatalf("patched controller line should not use I18n.t; got %q", got)
	}
}

func writeSourceFile(t *testing.T, relFile string, lines []string) string {
	t.Helper()

	root := t.TempDir()
	absPath := filepath.Join(root, relFile)
	if err := os.MkdirAll(filepath.Dir(absPath), 0o755); err != nil {
		t.Fatalf("mkdir source dir: %v", err)
	}
	if err := os.WriteFile(absPath, []byte(strings.Join(lines, "\n")+"\n"), 0o644); err != nil {
		t.Fatalf("write source file: %v", err)
	}
	return root
}

func onlyFix(t *testing.T, plan FixPlan) Fix {
	t.Helper()

	if len(plan.Fixes) != 1 {
		t.Fatalf("expected exactly one fix, got %d", len(plan.Fixes))
	}
	return plan.Fixes[0]
}

// ── ApplyFixes ────────────────────────────────────────────────────────────────

func TestApplyFixes_SingleLineReplacement(t *testing.T) {
	file := "app/views/show.html.erb"
	root := writeSourceFile(t, file, []string{`<h1>Hello World</h1>`})

	plan := FixPlan{
		YAMLAdds: map[string]string{},
		Fixes: []Fix{
			{
				File:    file,
				Line:    1,
				Patched: `<h1><%= t('show.heading') %></h1>`,
			},
		},
	}
	written, err := ApplyFixes(plan, root)
	if err != nil {
		t.Fatalf("ApplyFixes: %v", err)
	}
	if len(written) != 1 || written[0] != file {
		t.Errorf("written = %v, want [%s]", written, file)
	}
	content := readSourceFile(t, root, file)
	if !strings.Contains(content, "t('show.heading')") {
		t.Errorf("patched content not written:\n%s", content)
	}
	if strings.Contains(content, "Hello World") {
		t.Errorf("original hardcoded text still present:\n%s", content)
	}
}

func TestApplyFixes_TwoLinesInSameFile(t *testing.T) {
	file := "app/views/index.html.erb"
	root := writeSourceFile(t, file, []string{
		"<h1>Title Text Here</h1>",
		"<p>Body text goes here</p>",
	})

	plan := FixPlan{
		YAMLAdds: map[string]string{},
		Fixes: []Fix{
			{File: file, Line: 1, Patched: "<h1><%= t('index.title') %></h1>"},
			{File: file, Line: 2, Patched: "<p><%= t('index.body') %></p>"},
		},
	}
	if _, err := ApplyFixes(plan, root); err != nil {
		t.Fatalf("ApplyFixes: %v", err)
	}
	content := readSourceFile(t, root, file)
	if !strings.Contains(content, "t('index.title')") {
		t.Errorf("line 1 not patched:\n%s", content)
	}
	if !strings.Contains(content, "t('index.body')") {
		t.Errorf("line 2 not patched:\n%s", content)
	}
}

func TestApplyFixes_NonExistentFile_ReturnsError(t *testing.T) {
	root := t.TempDir()
	plan := FixPlan{
		YAMLAdds: map[string]string{},
		Fixes: []Fix{
			{File: "app/views/missing.html.erb", Line: 1, Patched: "anything"},
		},
	}
	_, err := ApplyFixes(plan, root)
	if err == nil {
		t.Error("expected error for non-existent file, got nil")
	}
}

func TestApplyFixes_EmptyPlan_NoFilesWritten(t *testing.T) {
	root := t.TempDir()
	written, err := ApplyFixes(FixPlan{YAMLAdds: map[string]string{}}, root)
	if err != nil {
		t.Fatalf("ApplyFixes(empty): %v", err)
	}
	if len(written) != 0 {
		t.Errorf("expected 0 written files, got %v", written)
	}
}

func TestApplyFixes_OutOfBoundsLineSkipped(t *testing.T) {
	file := "app/views/show.html.erb"
	root := writeSourceFile(t, file, []string{"only one line"})

	plan := FixPlan{
		YAMLAdds: map[string]string{},
		Fixes:    []Fix{{File: file, Line: 99, Patched: "should not appear"}},
	}
	if _, err := ApplyFixes(plan, root); err != nil {
		t.Fatalf("ApplyFixes: %v", err)
	}
	content := readSourceFile(t, root, file)
	if strings.Contains(content, "should not appear") {
		t.Error("out-of-bounds fix incorrectly applied")
	}
}

func TestApplyFixes_MultilineTagContent_InnerLinesBlank(t *testing.T) {
	file := "app/views/show.html.erb"
	root := writeSourceFile(t, file, []string{
		"<p>",
		"  Long text",
		"  spanning lines",
		"</p>",
	})

	plan := FixPlan{
		YAMLAdds: map[string]string{},
		Fixes: []Fix{
			{File: file, Line: 1, EndLine: 4, Patched: "<p><%= t('show.body') %></p>"},
		},
	}
	if _, err := ApplyFixes(plan, root); err != nil {
		t.Fatalf("ApplyFixes: %v", err)
	}
	content := readSourceFile(t, root, file)
	if !strings.Contains(content, "t('show.body')") {
		t.Errorf("multiline fix not applied:\n%s", content)
	}
	if strings.Contains(content, "Long text") || strings.Contains(content, "spanning lines") {
		t.Errorf("inner lines not blanked:\n%s", content)
	}
}

func readSourceFile(t *testing.T, root, relPath string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, relPath))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	return string(data)
}
