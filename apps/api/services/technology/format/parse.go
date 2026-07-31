package convformat

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	yaml "gopkg.in/yaml.v3"

	"requiems-api/platform/svcerr"
)

// parseInput parses the content string from the given format into a generic
// Go value (map, slice, or scalar).
func parseInput(format, content string) (any, error) {
	switch format {
	case "json":
		return parseJSON(content)
	case "yaml":
		return parseYAML(content)
	case "csv":
		return parseCSV(content)
	case "xml":
		return parseXML(content)
	case "toml":
		return parseTOML(content)
	default:
		return nil, svcerr.Invalid("unsupported_format", fmt.Sprintf("unsupported input format: %s", format))
	}
}

// --- JSON ---

func parseJSON(content string) (any, error) {
	dec := json.NewDecoder(strings.NewReader(content))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return nil, svcerr.Unknown("invalid_json", fmt.Sprintf("invalid JSON: %s", err.Error()))
	}
	return normalizeNumbers(v), nil
}

// --- YAML ---

func parseYAML(content string) (any, error) {
	var v any
	if err := yaml.Unmarshal([]byte(content), &v); err != nil {
		return nil, svcerr.Unknown("invalid_yaml", fmt.Sprintf("invalid YAML: %s", err.Error()))
	}
	// yaml.v3 unmarshals maps as map[string]any, which is what we want.
	return v, nil
}

// --- CSV ---

// parseCSV parses CSV content into a slice of maps ([]map[string]string).
// The first row is treated as the header row. Duplicate header names are rejected.
func parseCSV(content string) (any, error) {
	r := csv.NewReader(strings.NewReader(content))
	records, err := r.ReadAll()
	if err != nil {
		return nil, svcerr.Unknown("invalid_csv", fmt.Sprintf("invalid CSV: %s", err.Error()))
	}
	if len(records) == 0 {
		return []any{}, nil
	}

	headers := records[0]
	seen := make(map[string]bool, len(headers))
	for _, h := range headers {
		if seen[h] {
			return nil, svcerr.Unknown("invalid_csv", fmt.Sprintf("duplicate header column: %q", h))
		}
		seen[h] = true
	}

	rows := make([]any, 0, len(records)-1)
	for rowIdx, record := range records[1:] {
		if len(record) > len(headers) {
			return nil, svcerr.Unknown("invalid_csv", fmt.Sprintf("row %d has %d columns but header defines %d", rowIdx+2, len(record), len(headers)))
		}
		row := make(map[string]any, len(headers))
		for i, h := range headers {
			if i < len(record) {
				row[h] = record[i]
			} else {
				row[h] = ""
			}
		}
		rows = append(rows, row)
	}
	return rows, nil
}

// --- XML ---

// parseXML parses XML content into a generic map structure.
// XML attributes are captured with an "@" prefix (e.g. id="1" → "@id": "1").
// Text content of leaf elements is stored under the "#text" key.
func parseXML(content string) (any, error) {
	dec := xml.NewDecoder(strings.NewReader(content))
	result, err := xmlDecodeElement(dec)
	if err != nil {
		return nil, svcerr.Unknown("invalid_xml", fmt.Sprintf("invalid XML: %s", err.Error()))
	}
	return result, nil
}

// xmlDecodeElement reads the next XML element from the decoder and converts it
// to a map[string]any, recursively.
func xmlDecodeElement(dec *xml.Decoder) (any, error) {
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			return xmlReadElement(dec, t)
		case xml.CharData:
			s := strings.TrimSpace(string(t))
			if s != "" {
				return s, nil
			}
		}
	}
}

// xmlReadElement reads the content of an already-opened XML element (start token
// already consumed) and returns it as a map. Attributes are stored with an "@"
// prefix; duplicate child tags become slices; text content goes under "#text".
func xmlReadElement(dec *xml.Decoder, start xml.StartElement) (map[string]any, error) {
	result := make(map[string]any)

	// Capture XML attributes with "@" prefix.
	for _, attr := range start.Attr {
		result["@"+attr.Name.Local] = attr.Value
	}

	var textContent strings.Builder

	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			child, err := xmlReadElement(dec, t)
			if err != nil {
				return nil, err
			}
			key := t.Name.Local
			if existing, ok := result[key]; ok {
				// Multiple children with same tag → convert to slice
				switch ev := existing.(type) {
				case []any:
					result[key] = append(ev, child)
				default:
					result[key] = []any{existing, child}
				}
			} else {
				result[key] = child
			}
		case xml.CharData:
			textContent.Write(t)
		case xml.EndElement:
			text := strings.TrimSpace(textContent.String())
			if len(result) == 0 && text != "" {
				return map[string]any{"#text": text}, nil
			}
			if text != "" {
				result["#text"] = text
			}
			return result, nil
		}
	}
}

// --- TOML ---

func parseTOML(content string) (any, error) {
	var v map[string]any
	if _, err := toml.Decode(content, &v); err != nil {
		return nil, svcerr.Unknown("invalid_toml", fmt.Sprintf("invalid TOML: %s", err.Error()))
	}
	return v, nil
}

// --- helpers ---

// normalizeNumbers converts json.Number values (produced by json.Decoder with
// UseNumber()) to int64 (when the number has no fractional part) or float64,
// so that downstream serializers (YAML, TOML, XML) emit proper numeric values
// rather than quoted strings.
//
// Numbers with no fractional part are converted to int64. Numbers that exceed
// int64 range fall through to float64, which may lose precision for very large
// integers (e.g. values > 2^53). In the rare case both conversions fail the
// original string representation is preserved.
func normalizeNumbers(v any) any {
	switch val := v.(type) {
	case json.Number:
		if i, err := val.Int64(); err == nil {
			return i
		}
		if f, err := val.Float64(); err == nil {
			return f
		}
		return val.String()
	case map[string]any:
		out := make(map[string]any, len(val))
		for k, child := range val {
			out[k] = normalizeNumbers(child)
		}
		return out
	case []any:
		out := make([]any, len(val))
		for i, item := range val {
			out[i] = normalizeNumbers(item)
		}
		return out
	default:
		return val
	}
}
