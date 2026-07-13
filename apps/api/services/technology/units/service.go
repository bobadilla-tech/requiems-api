package units

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
)

// unit represents a single unit with its conversion factor relative to the base unit.
// Temperature units have a zero factor; conversions are handled separately.
type unit struct {
	category string
	// factor is the number of base units in one of this unit.
	// Used to compute the conversion ratio and the human-readable formula.
	factor float64
}

var units = map[string]unit{
	// Length (base: meters)
	"mm":    {category: "length", factor: 0.001},
	"cm":    {category: "length", factor: 0.01},
	"m":     {category: "length", factor: 1},
	"km":    {category: "length", factor: 1000},
	"in":    {category: "length", factor: 0.0254},
	"ft":    {category: "length", factor: 0.3048},
	"yd":    {category: "length", factor: 0.9144},
	"miles": {category: "length", factor: 1609.344},
	"nmi":   {category: "length", factor: 1852},

	// Weight (base: grams)
	"mg":    {category: "weight", factor: 0.001},
	"g":     {category: "weight", factor: 1},
	"kg":    {category: "weight", factor: 1000},
	"t":     {category: "weight", factor: 1e6},
	"oz":    {category: "weight", factor: 28.3495},
	"lb":    {category: "weight", factor: 453.592},
	"stone": {category: "weight", factor: 6350.29},

	// Volume (base: milliliters)
	"ml":    {category: "volume", factor: 1},
	"l":     {category: "volume", factor: 1000},
	"tsp":   {category: "volume", factor: 4.92892},
	"tbsp":  {category: "volume", factor: 14.7868},
	"fl_oz": {category: "volume", factor: 29.5735},
	"cup":   {category: "volume", factor: 236.588},
	"pt":    {category: "volume", factor: 473.176},
	"qt":    {category: "volume", factor: 946.353},
	"gal":   {category: "volume", factor: 3785.41},

	// Temperature (base: celsius; special handling)
	"c": {category: "temperature"},
	"f": {category: "temperature"},
	"k": {category: "temperature"},

	// Area (base: square meters)
	"mm2":  {category: "area", factor: 1e-6},
	"cm2":  {category: "area", factor: 1e-4},
	"m2":   {category: "area", factor: 1},
	"km2":  {category: "area", factor: 1e6},
	"in2":  {category: "area", factor: 6.4516e-4},
	"ft2":  {category: "area", factor: 0.092903},
	"yd2":  {category: "area", factor: 0.836127},
	"acre": {category: "area", factor: 4046.86},
	"ha":   {category: "area", factor: 10000},

	// Speed (base: km/h)
	"m_s":   {category: "speed", factor: 3.6},
	"km_h":  {category: "speed", factor: 1},
	"mph":   {category: "speed", factor: 1.60934},
	"knots": {category: "speed", factor: 1.852},
}

// aliases maps common long-form / plural names onto canonical unit keys.
// Discovery (Units) still returns only canonical keys; Convert accepts both.
var aliases = map[string]string{
	"meter": "m", "meters": "m", "metre": "m", "metres": "m",
	"kilometer": "km", "kilometers": "km", "kilometre": "km", "kilometres": "km",
	"centimeter": "cm", "centimeters": "cm", "centimetre": "cm", "centimetres": "cm",
	"millimeter": "mm", "millimeters": "mm", "millimetre": "mm", "millimetres": "mm",
	"foot": "ft", "feet": "ft",
	"inch": "in", "inches": "in",
	"yard": "yd", "yards": "yd",
	"nautical_mile": "nmi", "nautical_miles": "nmi",
	"kilogram": "kg", "kilograms": "kg",
	"gram": "g", "grams": "g",
	"milligram": "mg", "milligrams": "mg",
	"pound": "lb", "pounds": "lb",
	"ounce": "oz", "ounces": "oz",
	"ton": "t", "tons": "t", "tonne": "t", "tonnes": "t",
	"celsius": "c", "fahrenheit": "f", "kelvin": "k",
	"liter": "l", "liters": "l", "litre": "l", "litres": "l",
	"milliliter": "ml", "milliliters": "ml", "millilitre": "ml", "millilitres": "ml",
	"gallon": "gal", "gallons": "gal",
	"quart": "qt", "quarts": "qt",
	"pint": "pt", "pints": "pt",
	"cups":        "cup",
	"fluid_ounce": "fl_oz", "fluid_ounces": "fl_oz",
	"teaspoon": "tsp", "teaspoons": "tsp",
	"tablespoon": "tbsp", "tablespoons": "tbsp",
	"square_meter": "m2", "square_meters": "m2", "square_metre": "m2", "square_metres": "m2",
	"square_kilometer": "km2", "square_kilometers": "km2", "square_kilometre": "km2", "square_kilometres": "km2",
	"square_centimeter": "cm2", "square_centimeters": "cm2",
	"square_millimeter": "mm2", "square_millimeters": "mm2",
	"square_foot": "ft2", "square_feet": "ft2",
	"square_inch": "in2", "square_inches": "in2",
	"square_yard": "yd2", "square_yards": "yd2",
	"acres":   "acre",
	"hectare": "ha", "hectares": "ha",
	"meters_per_second": "m_s", "metres_per_second": "m_s",
	"kilometers_per_hour": "km_h", "kilometres_per_hour": "km_h",
	"miles_per_hour": "mph",
}

// ErrUnknownUnit is returned when a unit is not recognised.
var ErrUnknownUnit = errors.New("unknown unit")

// ErrIncompatibleUnits is returned when from and to belong to different categories.
var ErrIncompatibleUnits = errors.New("incompatible units: cannot convert between different measurement types")

type Result struct {
	From    string  `json:"from"`
	To      string  `json:"to"`
	Input   float64 `json:"input"`
	Result  float64 `json:"result"`
	Formula string  `json:"formula"`
}

// Results maps each measurement category to its supported unit keys.
type Results struct {
	Length      []string `json:"length"`
	Weight      []string `json:"weight"`
	Volume      []string `json:"volume"`
	Temperature []string `json:"temperature"`
	Area        []string `json:"area"`
	Speed       []string `json:"speed"`
}

// BatchResponse represents the result of a single batch conversion operation.
type BatchResponse struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    Result `json:"data"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// Units returns all supported unit keys grouped by measurement category,
// sorted alphabetically within each group.
func (s *Service) Units() Results {
	var result Results
	for key, u := range units {
		switch u.category {
		case "length":
			result.Length = append(result.Length, key)
		case "weight":
			result.Weight = append(result.Weight, key)
		case "volume":
			result.Volume = append(result.Volume, key)
		case "temperature":
			result.Temperature = append(result.Temperature, key)
		case "area":
			result.Area = append(result.Area, key)
		case "speed":
			result.Speed = append(result.Speed, key)
		}
	}
	sort.Strings(result.Length)
	sort.Strings(result.Weight)
	sort.Strings(result.Volume)
	sort.Strings(result.Temperature)
	sort.Strings(result.Area)
	sort.Strings(result.Speed)
	return result
}

// resolveUnit maps a caller-supplied unit key (canonical or alias) to the
// canonical key and its definition.
func resolveUnit(key string) (canonical string, u unit, ok bool) {
	key = strings.ToLower(strings.TrimSpace(key))
	if u, ok = units[key]; ok {
		return key, u, true
	}
	if canonical, ok = aliases[key]; ok {
		u, ok = units[canonical]
		return canonical, u, ok
	}
	return "", unit{}, false
}

func (s *Service) Convert(from, to string, value float64) (Result, error) {
	fromKey, fromUnit, ok := resolveUnit(from)
	if !ok {
		return Result{}, fmt.Errorf("%w: %q", ErrUnknownUnit, from)
	}

	toKey, toUnit, ok := resolveUnit(to)
	if !ok {
		return Result{}, fmt.Errorf("%w: %q", ErrUnknownUnit, to)
	}

	if fromUnit.category != toUnit.category {
		return Result{}, ErrIncompatibleUnits
	}

	var result float64
	var formula string

	if fromUnit.category == "temperature" {
		result, formula = convertTemperature(fromKey, toKey, value)
	} else {
		factor := fromUnit.factor / toUnit.factor
		result = roundTo(value * factor)
		formula = fmt.Sprintf("%s × %s", fromKey, formatFactor(factor))
	}

	return Result{
		From:    fromKey,
		To:      toKey,
		Input:   value,
		Result:  result,
		Formula: formula,
	}, nil
}

func convertTemperature(from, to string, value float64) (result float64, formula string) {
	if from == to {
		return roundTo(value), fmt.Sprintf("%s (no conversion needed)", from)
	}

	celsius := toCelsius(from, value)
	result = fromCelsius(to, celsius)
	formula = getTemperatureFormula(from, to)

	return roundTo(result), formula
}

func toCelsius(from string, value float64) float64 {
	switch from {
	case "c":
		return value
	case "f":
		return (value - 32) * 5 / 9
	case "k":
		return value - 273.15
	default:
		return value
	}
}

func fromCelsius(to string, celsius float64) float64 {
	switch to {
	case "c":
		return celsius
	case "f":
		return celsius*9/5 + 32
	case "k":
		return celsius + 273.15
	default:
		return celsius
	}
}

func getTemperatureFormula(from, to string) string {
	formulas := map[string]string{
		"c-f": "°C × 9/5 + 32",
		"f-c": "(°F − 32) × 5/9",
		"c-k": "°C + 273.15",
		"k-c": "K − 273.15",
		"f-k": "(°F − 32) × 5/9 + 273.15",
		"k-f": "(K − 273.15) × 9/5 + 32",
	}
	return formulas[from+"-"+to]
}

func roundTo(v float64) float64 {
	const decimals = 6
	p := math.Pow(10, float64(decimals))
	return math.Round(v*p) / p
}

func formatFactor(f float64) string {
	// Round to 6 decimal places first to avoid floating-point artefacts.
	r := roundTo(f)

	if r == math.Trunc(r) {
		return fmt.Sprintf("%.0f", r)
	}

	s := fmt.Sprintf("%.6f", r)

	// Trim trailing zeros but keep at least one decimal place.
	i := len(s) - 1
	for i > 0 && s[i] == '0' {
		i--
	}

	if s[i] == '.' {
		i++
	}

	return s[:i+1]
}

// ConvertBatch processes multiple unit conversion operations and returns a result for each one.
// Each operation is converted independently — if one fails, the rest continue processing.
// The returned slice preserves the same order as the input operations.
func (s *Service) ConvertBatch(ctx context.Context, operations []BatchItem) []BatchResponse {
	results := make([]BatchResponse, len(operations))

	for i, op := range operations {
		if op.Value == nil {
			results[i] = BatchResponse{
				From:    op.From,
				To:      op.To,
				Success: false,
				Error:   "value is required",
			}
			continue
		}

		result, err := s.Convert(op.From, op.To, *op.Value)
		if err != nil {
			results[i] = BatchResponse{
				From:    op.From,
				To:      op.To,
				Success: false,
				Error:   err.Error(),
			}
			continue
		}

		results[i] = BatchResponse{
			From:    op.From,
			To:      op.To,
			Success: true,
			Data:    result,
		}

	}

	return results
}

// Ptr returns a pointer to the given float64 value.
//
//go:fix inline
func Ptr(v float64) *float64 { return new(v) }
