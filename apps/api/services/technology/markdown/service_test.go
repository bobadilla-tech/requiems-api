package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Convert(t *testing.T) {
	t.Parallel()
	svc := NewService()

	tests := []struct {
		name     string
		input    string
		sanitize bool
		wantHTML string
	}{
		{
			name:     "heading and bold",
			input:    "# Hello\n\nThis is **bold** text.",
			sanitize: false,
			wantHTML: "<h1>Hello</h1>\n<p>This is <strong>bold</strong> text.</p>",
		},
		{
			name:     "italic text",
			input:    "*italic*",
			sanitize: false,
			wantHTML: "<p><em>italic</em></p>",
		},
		{
			name:     "link",
			input:    "[link](https://example.com)",
			sanitize: false,
			wantHTML: `<p><a href="https://example.com">link</a></p>`,
		},
		{
			name:     "unordered list",
			input:    "- item one\n- item two",
			sanitize: false,
			wantHTML: "<ul>\n<li>item one</li>\n<li>item two</li>\n</ul>",
		},
		{
			name:     "code block",
			input:    "```\nhello world\n```",
			sanitize: false,
			wantHTML: "<pre><code>hello world\n</code></pre>",
		},
		{
			name:     "inline html passes through without sanitize",
			input:    "Hello <strong>world</strong>",
			sanitize: false,
			wantHTML: "<p>Hello <strong>world</strong></p>",
		},
		{
			name:     "script tag stripped when sanitize true",
			input:    "Hello <script>alert('xss')</script> world",
			sanitize: true,
			wantHTML: "<p>Hello <!-- raw HTML omitted -->alert('xss')<!-- raw HTML omitted --> world</p>",
		},
		{
			name:     "empty markdown returns empty string",
			input:    "",
			sanitize: false,
			wantHTML: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := svc.Convert(tt.input, tt.sanitize)
			require.NoError(t, err)
			assert.Equal(t, tt.wantHTML, got.HTML)
		})
	}
}

func TestService_ConvertBatch(t *testing.T) {
	t.Parallel()
	svc := NewService()

	tests := []struct {
		name      string
		inputs    []string
		sanitize  bool
		wantBatch []Response
	}{
		{
			name: "multiple markdown documents",
			inputs: []string{
				"# Hello",
				"**bold**",
				"- item",
			},
			sanitize: false,
			wantBatch: []Response{
				{HTML: "<h1>Hello</h1>"},
				{HTML: "<p><strong>bold</strong></p>"},
				{HTML: "<ul>\n<li>item</li>\n</ul>"},
			},
		},
		{
			name: "sanitize html in batch",
			inputs: []string{
				"Hello <script>alert('xss')</script>",
				"<strong>safe?</strong>",
			},
			sanitize: true,
			wantBatch: []Response{
				{
					HTML: "<p>Hello <!-- raw HTML omitted -->alert('xss')<!-- raw HTML omitted --></p>",
				},
				{
					HTML: "<p><!-- raw HTML omitted -->safe?<!-- raw HTML omitted --></p>",
				},
			},
		},
		{
			name:      "empty batch",
			inputs:    []string{},
			sanitize:  false,
			wantBatch: []Response{},
		},
		{
			name: "batch with empty markdown",
			inputs: []string{
				"",
				"# Title",
			},
			sanitize: false,
			wantBatch: []Response{
				{HTML: ""},
				{HTML: "<h1>Title</h1>"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := svc.ConvertBatch(tt.inputs, tt.sanitize)
			require.NoError(t, err)
			require.Len(t, got, len(tt.wantBatch))
			for i := range got {
				assert.Equal(t, tt.wantBatch[i].HTML, got[i].HTML)
			}
		})
	}
}
