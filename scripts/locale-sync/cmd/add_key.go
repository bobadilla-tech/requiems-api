package cmd

import (
	"fmt"
	"strings"

	"github.com/bobadilla-tech/locale-sync/internal/locale"
	"github.com/spf13/cobra"
)

var (
	addKeyLang string
	addKeyDry  bool
)

var addKeyCmd = &cobra.Command{
	Use:   "add-key",
	Short: "Add a single locale key-value pair to the appropriate YAML file",
	Long: `add-key writes a key-value pair into the locale YAML file that owns
the key's namespace, creating the file if needed.

Examples:
  locale-sync add-key --key shared.danger_zone.cancel_btn --value "Cancel" \
    --root ../../apps/dashboard
  locale-sync add-key --key admin.users.show.page_title --value "User Details - %{email}" \
    --root ../../apps/dashboard --lang en --dry-run`,
	RunE: runAddKey,
}

var addKeyKey, addKeyValue string

func init() {
	addKeyCmd.Flags().StringVarP(&rootPath, "root", "r", ".", "Rails app root directory")
	addKeyCmd.Flags().StringVar(&addKeyLang, "lang", "en", "Target language code")
	addKeyCmd.Flags().BoolVar(&addKeyDry, "dry-run", false, "Preview without writing")
	addKeyCmd.Flags().StringVar(&addKeyKey, "key", "", "Dot-notation key path (required)")
	addKeyCmd.Flags().StringVar(&addKeyValue, "value", "", "String value to store (required)")
	if err := addKeyCmd.MarkFlagRequired("key"); err != nil {
		panic(err)
	}
	if err := addKeyCmd.MarkFlagRequired("value"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(addKeyCmd)
}

func runAddKey(_ *cobra.Command, _ []string) error {
	if err := validateLangCode(addKeyLang); err != nil {
		return err
	}

	// Strip leading lang prefix if caller included it.
	keyPath := stripLangPrefixStr(addKeyKey, addKeyLang)

	// Derive topic from first segment of key path.
	parts := strings.SplitN(keyPath, ".", 2)
	if len(parts) < 2 {
		return fmt.Errorf("key %q must have at least two segments (topic.name)", addKeyKey)
	}
	topic := parts[0]
	fullKey := addKeyLang + "." + keyPath

	candidate := locale.MergeCandidate{
		KeyName:      lastSegment(keyPath),
		Value:        addKeyValue,
		SuggestedKey: fullKey,
	}

	fmt.Printf("Key:   %s\n", fullKey)
	fmt.Printf("Value: %q\n", addKeyValue)
	fmt.Printf("Topic: %s (%s/%s.%s.yml)\n", topic, addKeyLang, topic, addKeyLang)

	if addKeyDry {
		fmt.Println("--dry-run: no files modified.")
		return nil
	}

	path, changed, err := locale.UpsertTopicFile(rootPath, addKeyLang, topic, []locale.MergeCandidate{candidate})
	if err != nil {
		return fmt.Errorf("write key: %w", err)
	}
	if changed {
		fmt.Printf("Written to %s\n", shortPath(path, rootPath))
	} else {
		fmt.Printf("Already present in %s (skipped)\n", shortPath(path, rootPath))
	}
	return nil
}
