package translate

import (
	"encoding/json"
	"os"
)

// Cache stores source text → {langCode: translatedText} mappings on disk.
type Cache struct {
	path    string
	Entries map[string]map[string]string
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

// Get returns a cached translation and whether it was found.
func (c *Cache) Get(src, lang string) (string, bool) {
	langs, ok := c.Entries[src]
	if !ok {
		return "", false
	}
	v, ok := langs[lang]
	return v, ok
}

// Set stores a translation in memory (call Save to persist).
func (c *Cache) Set(src, lang, translation string) {
	if c.Entries[src] == nil {
		c.Entries[src] = make(map[string]string)
	}
	c.Entries[src][lang] = translation
}

// Save writes the cache to disk as pretty JSON.
func (c *Cache) Save() error {
	data, err := json.MarshalIndent(c.Entries, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(c.path, data, 0o644)
}
