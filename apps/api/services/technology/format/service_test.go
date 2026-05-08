package convformat

import (
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
