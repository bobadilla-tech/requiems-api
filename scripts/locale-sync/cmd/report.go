package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/bobadilla-tech/locale-sync/internal/source"
	"github.com/spf13/cobra"
)

var (
	reportFormat    string
	reportErbOnly   bool
	reportFailOnAny bool
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Scan source files and report hardcoded strings that should be i18n'd",
	Long: `report scans Ruby/ERB source files for hardcoded user-facing strings and
prints a lint report. Useful in CI to track i18n coverage over time.

Exit codes:
  0  — no hardcoded strings found
  1  — hardcoded strings found (only when --fail-on-found is set)`,
	RunE: runReport,
}

func init() {
	reportCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	reportCmd.Flags().StringVar(&reportFormat, "format", "text", "Output format: text or json")
	reportCmd.Flags().BoolVar(&reportErbOnly, "erb-only", false,
		"Only scan .erb/.haml/.slim files, skip .rb files")
	reportCmd.Flags().BoolVar(&reportFailOnAny, "fail-on-found", false,
		"Exit with code 1 if any hardcoded strings are found (useful in CI)")
	rootCmd.AddCommand(reportCmd)
}

func runReport(_ *cobra.Command, _ []string) error {
	found, err := source.Extract(rootPath)
	if err != nil {
		return fmt.Errorf("scan: %w", err)
	}

	if reportErbOnly {
		filtered := found[:0]
		for _, h := range found {
			if strings.HasSuffix(h.File, ".erb") ||
				strings.HasSuffix(h.File, ".haml") ||
				strings.HasSuffix(h.File, ".slim") {
				filtered = append(filtered, h)
			}
		}
		found = filtered
	}

	switch reportFormat {
	case "json":
		return printReportJSON(found)
	case "text":
		return printReportText(found)
	default:
		return fmt.Errorf("unknown --format %q: want \"text\" or \"json\"", reportFormat)
	}
}

func printReportText(found []source.HardcodedString) error {
	if len(found) == 0 {
		fmt.Println("✓ No hardcoded strings found.")
		return nil
	}

	// Group by file.
	byFile := make(map[string][]source.HardcodedString)
	var fileOrder []string
	seen := map[string]bool{}
	for _, h := range found {
		if !seen[h.File] {
			fileOrder = append(fileOrder, h.File)
			seen[h.File] = true
		}
		byFile[h.File] = append(byFile[h.File], h)
	}

	fmt.Printf("=== Hardcoded String Report (%d strings in %d files) ===\n\n",
		len(found), len(fileOrder))

	for _, file := range fileOrder {
		items := byFile[file]
		fmt.Printf("  %s (%d)\n", file, len(items))
		for _, h := range items {
			loc := fmt.Sprintf("%d", h.Line)
			if h.EndLine > h.Line {
				loc = fmt.Sprintf("%d-%d", h.Line, h.EndLine)
			}
			fmt.Printf("    [%s] line %-6s %q\n", h.Category, loc, h.Text)
		}
		fmt.Println()
	}

	if reportFailOnAny {
		os.Exit(1)
	}
	return nil
}

type jsonReportEntry struct {
	File     string `json:"file"`
	Line     int    `json:"line"`
	EndLine  int    `json:"end_line,omitempty"`
	Category string `json:"category"`
	Text     string `json:"text"`
}

func printReportJSON(found []source.HardcodedString) error {
	entries := make([]jsonReportEntry, 0, len(found))
	for _, h := range found {
		entries = append(entries, jsonReportEntry{
			File:     h.File,
			Line:     h.Line,
			EndLine:  h.EndLine,
			Category: h.Category,
			Text:     h.Text,
		})
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(entries); err != nil {
		return err
	}
	if reportFailOnAny && len(found) > 0 {
		os.Exit(1)
	}
	return nil
}
