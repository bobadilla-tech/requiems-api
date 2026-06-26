package translate

import (
	"encoding/json"
	"os"
)

// Cache stores translations keyed by source language and source text.
// On-disk format: {"en\x00Hello": {"es": "Hola", "fr": "Bonjour"}}
// The null-byte separator makes accidental collisions between lang and text impossible.
type Cache struct {
	path    string
	Entries map[string]map[string]string // "{srcLang}\x00{srcText}" → {targetLang: translation}
}

// LoadCache reads the cache file from disk. Returns an empty cache if the file
// doesn't exist yet.
func LoadCache(path string) (*Cache, error) {
	c := &Cache{
		path:    path,
		Entries: make(map[string]map[string]string),
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return c, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &c.Entries); err != nil {
		return nil, err
	}
	return c, nil
}

func cacheKey(srcLang, src string) string { return srcLang + "\x00" + src }

// Get returns a cached translation for the given source language, source text,
// and target language, plus whether it was found.
func (c *Cache) Get(srcLang, src, targetLang string) (string, bool) {
	langs, ok := c.Entries[cacheKey(srcLang, src)]
	if !ok {
		return "", false
	}
	v, ok := langs[targetLang]
	return v, ok
}

// Set stores a translation in memory (call Save to persist).
func (c *Cache) Set(srcLang, src, targetLang, translation string) {
	k := cacheKey(srcLang, src)
	if c.Entries[k] == nil {
		c.Entries[k] = make(map[string]string)
	}
	c.Entries[k][targetLang] = translation
}

// Save writes the cache to disk as pretty JSON.
func (c *Cache) Save() error {
	data, err := json.MarshalIndent(c.Entries, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(c.path, data, 0o644)
}
