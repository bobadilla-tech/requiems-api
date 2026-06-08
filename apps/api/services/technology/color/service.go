package color //nolint:revive // package name matches the service domain it implements

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"requiems-api/platform/svcerr"
)

// Formats holds the color expressed in every supported format.
type Formats struct {
	Hex  string `json:"hex"`
	RGB  string `json:"rgb"`
	HSL  string `json:"hsl"`
	CMYK string `json:"cmyk"`
}

// Response is the payload returned by GET /v1/convert/color.
type Response struct {
	Input   string  `json:"input"`
	Result  string  `json:"result"`
	Formats Formats `json:"formats"`
}

// Service provides color format conversion operations.
type Service struct{}

// NewService creates a new color Service.
func NewService() *Service { return &Service{} }

// rgb holds an sRGB color with components in [0, 255].
type rgb struct {
	r, g, b uint8
}

// Convert parses value in the from format and converts it to every supported
// format. Returns an error if the value cannot be parsed.
func (s *Service) Convert(from, to, value string) (Response, error) {
	c, err := parse(from, value)
	if err != nil {
		return Response{}, err
	}

	formats := allFormats(c)
	result := formatString(to, c)

	return Response{
		Input:   value,
		Result:  result,
		Formats: formats,
	}, nil
}

// BatchColorItem is the result for a single item in a batch color conversion request.
type BatchColorItem struct {
	From   string    `json:"from"`
	To     string    `json:"to"`
	Input  string    `json:"input"`
	Result *Response `json:"result,omitempty"`
	Error  string    `json:"error,omitempty"`
}

// ConvertBatch converts each item and returns results in input order.
// Items that fail to parse return an in-band error.
func (s *Service) ConvertBatch(items []ConvertQuery) []BatchColorItem {
	results := make([]BatchColorItem, len(items))
	for i, item := range items {
		r, err := s.Convert(item.From, item.To, item.Value)
		if err != nil {
			results[i] = BatchColorItem{From: item.From, To: item.To, Input: item.Value, Error: err.Error()}
		} else {
			results[i] = BatchColorItem{From: item.From, To: item.To, Input: item.Value, Result: &r}
		}
	}
	return results
}

// allFormats converts an rgb color into every supported string representation.
func allFormats(c rgb) Formats {
	return Formats{
		Hex:  toHex(c),
		RGB:  toRGB(c),
		HSL:  toHSL(c),
		CMYK: toCMYK(c),
	}
}

// formatString returns the string representation for the given format.
func formatString(format string, c rgb) string {
	switch format {
	case "hex":
		return toHex(c)
	case "rgb":
		return toRGB(c)
	case "hsl":
		return toHSL(c)
	case "cmyk":
		return toCMYK(c)
	}
	return toHex(c)
}

// parse reads a color string in the given format and returns an rgb value.
func parse(format, value string) (rgb, error) {
	switch format {
	case "hex":
		return parseHex(value)
	case "rgb":
		return parseRGB(value)
	case "hsl":
		return parseHSL(value)
	case "cmyk":
		return parseCMYK(value)
	}
	return rgb{}, invalidColor(value)
}

func invalidColor(value string) error {
	return svcerr.Invalid("invalid_color", fmt.Sprintf("cannot parse color value %q", value))
}

// parseHex parses a "#rrggbb" or "#rgb" hex color string.
func parseHex(value string) (rgb, error) {
	s := strings.TrimPrefix(value, "#")
	if len(s) == 3 {
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2]})
	}
	if len(s) != 6 {
		return rgb{}, invalidColor(value)
	}

	r, err1 := strconv.ParseUint(s[0:2], 16, 8)
	g, err2 := strconv.ParseUint(s[2:4], 16, 8)
	b, err3 := strconv.ParseUint(s[4:6], 16, 8)
	if err1 != nil || err2 != nil || err3 != nil {
		return rgb{}, invalidColor(value)
	}

	return rgb{uint8(r), uint8(g), uint8(b)}, nil
}

// parseRGB parses "rgb(r, g, b)" or "rgb(r,g,b)".
func parseRGB(value string) (rgb, error) {
	s := value
	s = strings.TrimPrefix(s, "rgb(")
	s = strings.TrimSuffix(s, ")")
	if s == value {
		return rgb{}, invalidColor(value)
	}

	parts := splitTrim(s, ",")
	if len(parts) != 3 {
		return rgb{}, invalidColor(value)
	}

	r, err1 := strconv.ParseUint(parts[0], 10, 8)
	g, err2 := strconv.ParseUint(parts[1], 10, 8)
	b, err3 := strconv.ParseUint(parts[2], 10, 8)
	if err1 != nil || err2 != nil || err3 != nil {
		return rgb{}, invalidColor(value)
	}

	return rgb{uint8(r), uint8(g), uint8(b)}, nil
}

// parseHSL parses "hsl(h, s%, l%)" or "hsl(h,s%,l%)".
func parseHSL(value string) (rgb, error) {
	s := value
	s = strings.TrimPrefix(s, "hsl(")
	s = strings.TrimSuffix(s, ")")
	if s == value {
		return rgb{}, invalidColor(value)
	}

	parts := splitTrim(s, ",")
	if len(parts) != 3 {
		return rgb{}, invalidColor(value)
	}

	h, err1 := strconv.ParseFloat(strings.TrimSuffix(parts[0], "%"), 64)
	sl, err2 := strconv.ParseFloat(strings.TrimSuffix(parts[1], "%"), 64)
	l, err3 := strconv.ParseFloat(strings.TrimSuffix(parts[2], "%"), 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return rgb{}, invalidColor(value)
	}

	return hslToRGB(h, sl/100, l/100), nil
}

// parseCMYK parses "cmyk(c%, m%, y%, k%)" or "cmyk(c,m,y,k)".
func parseCMYK(value string) (rgb, error) {
	s := value
	s = strings.TrimPrefix(s, "cmyk(")
	s = strings.TrimSuffix(s, ")")
	if s == value {
		return rgb{}, invalidColor(value)
	}

	parts := splitTrim(s, ",")
	if len(parts) != 4 {
		return rgb{}, invalidColor(value)
	}

	c, err1 := strconv.ParseFloat(strings.TrimSuffix(parts[0], "%"), 64)
	m, err2 := strconv.ParseFloat(strings.TrimSuffix(parts[1], "%"), 64)
	y, err3 := strconv.ParseFloat(strings.TrimSuffix(parts[2], "%"), 64)
	k, err4 := strconv.ParseFloat(strings.TrimSuffix(parts[3], "%"), 64)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return rgb{}, invalidColor(value)
	}

	return cmykToRGB(c/100, m/100, y/100, k/100), nil
}

// toHex formats an rgb value as "#rrggbb".
func toHex(c rgb) string {
	return fmt.Sprintf("#%02x%02x%02x", c.r, c.g, c.b)
}

// toRGB formats an rgb value as "rgb(r, g, b)".
func toRGB(c rgb) string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.r, c.g, c.b)
}

// toHSL formats an rgb value as "hsl(h, s%, l%)".
func toHSL(c rgb) string {
	r := float64(c.r) / 255
	g := float64(c.g) / 255
	b := float64(c.b) / 255

	maxVal := math.Max(r, math.Max(g, b))
	minVal := math.Min(r, math.Min(g, b))
	delta := maxVal - minVal

	l := (maxVal + minVal) / 2

	var s float64
	if delta != 0 {
		if l < 0.5 {
			s = delta / (maxVal + minVal)
		} else {
			s = delta / (2 - maxVal - minVal)
		}
	}

	var h float64
	if delta != 0 {
		switch maxVal {
		case r:
			h = math.Mod((g-b)/delta, 6)
		case g:
			h = (b-r)/delta + 2
		default: // b
			h = (r-g)/delta + 4
		}
		h *= 60
		if h < 0 {
			h += 360
		}
	}

	return fmt.Sprintf("hsl(%d, %d%%, %d%%)",
		int(math.Round(h)),
		int(math.Round(s*100)),
		int(math.Round(l*100)),
	)
}

// toCMYK formats an rgb value as "cmyk(c%, m%, y%, k%)".
func toCMYK(c rgb) string {
	r := float64(c.r) / 255
	g := float64(c.g) / 255
	b := float64(c.b) / 255

	k := 1 - math.Max(r, math.Max(g, b))
	if k == 1 {
		return "cmyk(0%, 0%, 0%, 100%)"
	}

	cy := (1 - r - k) / (1 - k)
	m := (1 - g - k) / (1 - k)
	y := (1 - b - k) / (1 - k)

	return fmt.Sprintf("cmyk(%d%%, %d%%, %d%%, %d%%)",
		int(math.Round(cy*100)),
		int(math.Round(m*100)),
		int(math.Round(y*100)),
		int(math.Round(k*100)),
	)
}

// hslToRGB converts HSL (h in [0,360), s and l in [0,1]) to RGB.
func hslToRGB(h, s, l float64) rgb {
	if s == 0 {
		v := uint8(math.Round(l * 255))
		return rgb{v, v, v}
	}

	var q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q

	r := hueToRGB(p, q, h/360+1.0/3)
	g := hueToRGB(p, q, h/360)
	b := hueToRGB(p, q, h/360-1.0/3)

	return rgb{
		uint8(math.Round(r * 255)),
		uint8(math.Round(g * 255)),
		uint8(math.Round(b * 255)),
	}
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	switch {
	case t < 1.0/6:
		return p + (q-p)*6*t
	case t < 0.5:
		return q
	case t < 2.0/3:
		return p + (q-p)*(2.0/3-t)*6
	}
	return p
}

// cmykToRGB converts CMYK (values in [0,1]) to RGB.
func cmykToRGB(c, m, y, k float64) rgb {
	r := 255 * (1 - c) * (1 - k)
	g := 255 * (1 - m) * (1 - k)
	b := 255 * (1 - y) * (1 - k)

	return rgb{
		uint8(math.Round(r)),
		uint8(math.Round(g)),
		uint8(math.Round(b)),
	}
}

// splitTrim splits s by sep and trims whitespace from each part.
func splitTrim(s, sep string) []string {
	parts := strings.Split(s, sep)
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}
	return parts
}
