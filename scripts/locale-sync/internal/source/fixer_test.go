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
