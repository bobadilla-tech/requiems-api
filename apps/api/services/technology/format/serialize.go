package convformat

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
	yaml "gopkg.in/yaml.v3"

	"requiems-api/platform/svcerr"
)

// serializeOutput converts a generic Go value to the given format string.
func serializeOutput(format string, v any) (string, error) {
	switch format {
	case "json":
		return toJSON(v)
	case "yaml":
		return toYAML(v)
	case "csv":
		return toCSV(v)
	case "xml":
		return toXML(v)
	case "toml":
		return toTOML(v)
	default:
		return "", svcerr.Invalid("unsupported_format", fmt.Sprintf("unsupported output format: %s", format))
	}
}

// --- JSON ---

func toJSON(v any) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to serialize to JSON: %w", err)
	}
	return string(b), nil
}

// --- YAML ---

func toYAML(v any) (string, error) {
	b, err := yaml.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("failed to serialize to YAML: %w", err)
	}
	return string(b), nil
}

// --- CSV ---

// toCSV serializes a generic value to CSV. The value must be a slice of maps.
// Headers are derived from the union of keys across all rows so that fields
// present in later rows are not silently dropped.
func toCSV(v any) (string, error) {
	rows, ok := v.([]any)
	if !ok {
		return "", svcerr.Unknown("conversion_error", "CSV output requires a JSON array of objects")
	}
	if len(rows) == 0 {
		return "", nil
	}

	// Validate all rows and collect the union of keys for complete headers.
	maps := make([]map[string]any, len(rows))
	seenKeys := make(map[string]bool)
	var headers []string
	for i, row := range rows {
		m, ok := row.(map[string]any)
		if !ok {
			return "", svcerr.Unknown("conversion_error", "CSV output requires all array elements to be objects")
		}
		maps[i] = m
		for k := range m {
			if !seenKeys[k] {
				seenKeys[k] = true
				headers = append(headers, k)
			}
		}
	}
	sort.Strings(headers)

	var buf strings.Builder
	w := csv.NewWriter(&buf)

	if err := w.Write(headers); err != nil {
		return "", fmt.Errorf("failed to write CSV headers: %w", err)
	}

	for _, m := range maps {
		record := make([]string, len(headers))
		for i, h := range headers {
			if val, exists := m[h]; exists {
				record[i] = fmt.Sprintf("%v", val)
			}
		}
		if err := w.Write(record); err != nil {
			return "", fmt.Errorf("failed to write CSV row: %w", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return "", fmt.Errorf("failed to flush CSV writer: %w", err)
	}

	return buf.String(), nil
}

// --- XML ---

// toXML serializes a generic value to XML. Maps are converted to elements;
// the root element is named <root>. The "#text" key (produced by the XML
// parser for text-content nodes) is emitted as character data rather than
// a child element, preserving round-trip fidelity.
func toXML(v any) (string, error) {
	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")

	root := xml.StartElement{Name: xml.Name{Local: "root"}}
	if err := enc.EncodeToken(root); err != nil {
		return "", fmt.Errorf("failed to serialize to XML: %w", err)
	}
	if err := encodeXMLValue(enc, v); err != nil {
		return "", fmt.Errorf("failed to serialize to XML: %w", err)
	}
	if err := enc.EncodeToken(root.End()); err != nil {
		return "", fmt.Errorf("failed to serialize to XML: %w", err)
	}
	if err := enc.Flush(); err != nil {
		return "", fmt.Errorf("failed to serialize to XML: %w", err)
	}

	return buf.String(), nil
}

func encodeXMLValue(enc *xml.Encoder, v any) error {
	switch val := v.(type) {
	case map[string]any:
		return encodeXMLMap(enc, val)
	case []any:
		for _, item := range val {
			elem := xml.StartElement{Name: xml.Name{Local: "item"}}
			if err := enc.EncodeToken(elem); err != nil {
				return err
			}
			if err := encodeXMLValue(enc, item); err != nil {
				return err
			}
			if err := enc.EncodeToken(elem.End()); err != nil {
				return err
			}
		}
	case nil:
		// emit nothing
	default:
		if err := enc.EncodeToken(xml.CharData(fmt.Sprintf("%v", val))); err != nil {
			return err
		}
	}
	return nil
}

// encodeXMLMap encodes a map[string]any as XML child elements. The "#text" key
// is treated as character data rather than a child element, preserving
// round-trip fidelity with the XML parser output.
func encodeXMLMap(enc *xml.Encoder, val map[string]any) error {
	keys := make([]string, 0, len(val))
	for k := range val {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if k == "#text" {
			// Emit text content directly rather than wrapping in a child element.
			if err := enc.EncodeToken(xml.CharData(fmt.Sprintf("%v", val[k]))); err != nil {
				return err
			}
			continue
		}
		elem := xml.StartElement{Name: xml.Name{Local: sanitizeXMLName(k)}}
		if err := enc.EncodeToken(elem); err != nil {
			return err
		}
		if err := encodeXMLValue(enc, val[k]); err != nil {
			return err
		}
		if err := enc.EncodeToken(elem.End()); err != nil {
			return err
		}
	}
	return nil
}

// sanitizeXMLName replaces characters that are invalid in XML element names
// with underscores. The result is always non-empty.
func sanitizeXMLName(name string) string {
	if name == "" {
		return "_"
	}
	var b strings.Builder
	for i, r := range name {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r == '_':
			b.WriteRune(r)
		case i > 0 && (r >= '0' && r <= '9' || r == '-' || r == '.'):
			b.WriteRune(r)
		default:
			b.WriteRune('_')
		}
	}
	return b.String()
}

// --- TOML ---

func toTOML(v any) (string, error) {
	m, ok := v.(map[string]any)
	if !ok {
		return "", svcerr.Unknown("conversion_error", "TOML output requires a JSON object (not an array or scalar)")
	}
	var buf strings.Builder
	enc := toml.NewEncoder(&buf)
	if err := enc.Encode(m); err != nil {
		return "", fmt.Errorf("failed to serialize to TOML: %w", err)
	}
	return buf.String(), nil
}
