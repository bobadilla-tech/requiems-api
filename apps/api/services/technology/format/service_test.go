package convformat

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Convert_SameFormat(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "json", To: "json", Content: `{"name":"Alice"}`}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	assert.Equal(t, req.Content, resp.Result)
}

func TestService_Convert_ContentTooLarge(t *testing.T) {
	t.Parallel()
	svc := NewService()
	big := strings.Repeat("a", maxContentSize+1)
	req := Request{From: "json", To: "yaml", Content: big}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

// --- JSON ↔ YAML ---

func TestService_JSONToYAML(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{
		From:    "json",
		To:      "yaml",
		Content: `{"name":"Alice","age":30}`,
	}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	assert.True(t, strings.Contains(resp.Result, "name: Alice"), "expected YAML with 'name: Alice', got %q", resp.Result)
	assert.True(t, strings.Contains(resp.Result, "age: 30"), "expected YAML with 'age: 30', got %q", resp.Result)
}

func TestService_YAMLToJSON(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{
		From:    "yaml",
		To:      "json",
		Content: "name: Alice\nage: 30\n",
	}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	assert.True(t, strings.Contains(resp.Result, `"name"`), "expected JSON with 'name' key, got %q", resp.Result)
	assert.True(t, strings.Contains(resp.Result, "Alice"), "expected JSON with 'Alice', got %q", resp.Result)
}

func TestService_InvalidJSON(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "json", To: "yaml", Content: `{invalid`}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

func TestService_InvalidYAML(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "yaml", To: "json", Content: ":\t:bad yaml\n"}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

// --- JSON ↔ CSV ---

func TestService_JSONToCSV(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{
		From:    "json",
		To:      "csv",
		Content: `[{"name":"Alice","age":"30"},{"name":"Bob","age":"25"}]`,
	}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	assert.True(t, strings.Contains(resp.Result, "name"), "expected CSV with 'name' header, got %q", resp.Result)
	assert.True(t, strings.Contains(resp.Result, "Alice"), "expected CSV with 'Alice', got %q", resp.Result)
}

func TestService_CSVToJSON(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{
		From:    "csv",
		To:      "json",
		Content: "name,age\nAlice,30\nBob,25\n",
	}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	assert.True(t, strings.Contains(resp.Result, `"name"`), "expected JSON with 'name' key, got %q", resp.Result)
	assert.True(t, strings.Contains(resp.Result, "Alice"), "expected JSON with 'Alice', got %q", resp.Result)
}

func TestService_JSONToCSV_NonArray(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "json", To: "csv", Content: `{"name":"Alice"}`}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

// --- JSON ↔ XML ---

func TestService_JSONToXML(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{
		From:    "json",
		To:      "xml",
		Content: `{"name":"Alice","age":30}`,
	}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	assert.True(t, strings.Contains(resp.Result, "<root>"), "expected XML with <root>, got %q", resp.Result)
	assert.True(t, strings.Contains(resp.Result, "Alice"), "expected XML with 'Alice', got %q", resp.Result)
}

func TestService_XMLToJSON(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{
		From:    "xml",
		To:      "json",
		Content: `<root><name>Alice</name><age>30</age></root>`,
	}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	assert.True(t, strings.Contains(resp.Result, "Alice"), "expected JSON with 'Alice', got %q", resp.Result)
}

func TestService_InvalidXML(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "xml", To: "json", Content: "<unclosed>"}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

// --- JSON ↔ TOML ---

func TestService_JSONToTOML(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{
		From:    "json",
		To:      "toml",
		Content: `{"name":"Alice","age":30}`,
	}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	assert.True(t, strings.Contains(resp.Result, "Alice"), "expected TOML with 'Alice', got %q", resp.Result)
}

func TestService_TOMLToJSON(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{
		From:    "toml",
		To:      "json",
		Content: "name = \"Alice\"\nage = 30\n",
	}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	assert.True(t, strings.Contains(resp.Result, `"name"`), "expected JSON with 'name' key, got %q", resp.Result)
	assert.True(t, strings.Contains(resp.Result, "Alice"), "expected JSON with 'Alice', got %q", resp.Result)
}

func TestService_InvalidTOML(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "toml", To: "json", Content: "= invalid toml"}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

func TestService_JSONToTOML_Array(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "json", To: "toml", Content: `[1,2,3]`}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

// --- Unsupported formats ---

func TestService_UnsupportedInputFormat(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "msgpack", To: "json", Content: "data"}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

func TestService_UnsupportedOutputFormat(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "json", To: "msgpack", Content: `{"x":1}`}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

// --- CSV edge cases ---

func TestService_CSVToJSON_EmptyContent(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "csv", To: "json", Content: ""}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	assert.Equal(t, "[]", strings.TrimSpace(resp.Result))
}

func TestService_CSV_RowWithFewerColumnsThanHeaders(t *testing.T) {
	t.Parallel()
	svc := NewService()
	// Go's csv.Reader enforces consistent field counts by default,
	// so a row with fewer columns than the header returns an error.
	req := Request{From: "csv", To: "json", Content: "a,b,c\n1,2\n"}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

func TestService_CSV_RowWithMoreColumnsThanHeaders(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "csv", To: "json", Content: "a,b\n1,2,3\n"}
	_, err := svc.Convert(req)
	require.Error(t, err, "expected error when a row has more columns than headers")
}

func TestService_JSONToCSV_NonMapFirstRow(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "json", To: "csv", Content: `["string-not-a-map"]`}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

func TestService_JSONToCSV_NonMapSubsequentRow(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{From: "json", To: "csv", Content: `[{"a":"1"},"bad"]`}
	_, err := svc.Convert(req)
	require.Error(t, err)
}

// --- XML edge cases ---

func TestService_XML_DuplicateTags(t *testing.T) {
	t.Parallel()
	svc := NewService()
	req := Request{
		From:    "xml",
		To:      "json",
		Content: `<root><item>A</item><item>B</item></root>`,
	}
	resp, err := svc.Convert(req)
	require.NoError(t, err)
	// Both items must appear in output.
	assert.Contains(t, resp.Result, "A")
	assert.Contains(t, resp.Result, "B")
}

// --- sanitizeXMLName ---

func TestSanitizeXMLName_EmptyString(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "_", sanitizeXMLName(""))
}

func TestSanitizeXMLName_StartsWithDigit(t *testing.T) {
	t.Parallel()
	// Leading digit is invalid in XML element names → replaced with _
	result := sanitizeXMLName("1abc")
	assert.Equal(t, "_abc", result)
}

func TestSanitizeXMLName_ValidName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "valid_Name", sanitizeXMLName("valid_Name"))
}

// --- normalizeNumbers ---

func TestNormalizeNumbers_Float64(t *testing.T) {
	t.Parallel()
	result := normalizeNumbers(jsonNumber("1.5"))
	assert.Equal(t, float64(1.5), result)
}

func TestNormalizeNumbers_Int64(t *testing.T) {
	t.Parallel()
	result := normalizeNumbers(jsonNumber("42"))
	assert.Equal(t, int64(42), result)
}

func TestNormalizeNumbers_NestedMap(t *testing.T) {
	t.Parallel()
	input := map[string]any{"n": jsonNumber("3")}
	result := normalizeNumbers(input)
	m := result.(map[string]any)
	assert.Equal(t, int64(3), m["n"])
}

func TestNormalizeNumbers_Slice(t *testing.T) {
	t.Parallel()
	input := []any{jsonNumber("7")}
	result := normalizeNumbers(input)
	s := result.([]any)
	assert.Equal(t, int64(7), s[0])
}

func TestNormalizeNumbers_NonNumber(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "hello", normalizeNumbers("hello"))
}

// jsonNumber returns a json.Number value for use in normalizeNumbers tests.
func jsonNumber(s string) any {
	return json.Number(s)
}
