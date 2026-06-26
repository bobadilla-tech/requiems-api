package translate

import (
	"github.com/bobadilla-tech/locale-sync/internal/locale"
)

// TranslatedEntry pairs a MissingEntry with its translated text.
type TranslatedEntry struct {
	MissingEntry
	TranslatedText string
	FromCache      bool
}

// WriteTranslations writes translated entries into their respective target YAML files.
// Entries are grouped by TargetFile; locale.WriteEntries handles the actual YAML merge.
// Returns the list of files that were modified.
func WriteTranslations(translated []TranslatedEntry) ([]string, error) {
	// Group by target file.
	byFile := make(map[string]map[string]string)
	for _, e := range translated {
		if byFile[e.TargetFile] == nil {
			byFile[e.TargetFile] = make(map[string]string)
		}
		byFile[e.TargetFile][e.TargetKey] = e.TranslatedText
	}

	var written []string
	for filePath, entries := range byFile {
		changed, err := locale.WriteEntries(filePath, entries)
		if err != nil {
			return written, err
		}
		if changed {
			written = append(written, filePath)
		}
	}
	return written, nil
}
