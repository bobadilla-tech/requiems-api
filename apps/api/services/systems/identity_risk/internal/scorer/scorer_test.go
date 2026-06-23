package scorer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignalConfidence_NoSignals(t *testing.T) {
	s := Signals{}

	got := signalConfidence(s)

	assert.Equal(t, 0.0, got)
}

func TestSignalConfidence_SingleHighQualitySignal(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailDisposable: false,
		EmailNoMX:       false,
	}

	got := signalConfidence(s)

	assert.Equal(t, 1.0, got)
}

func TestSignalConfidence_EmailAllRisk(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailNoMX:       true,
		EmailDisposable: true,
	}

	got := signalConfidence(s)

	assert.Equal(t, 0.0, got)
}

func TestSignalConfidence_AllSignalsClean(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: false,

		PhonePresent: true,
		PhoneVoIP:    false,
		PhoneVirtual: false,

		IPPresent:  true,
		IsProxy:    false,
		IsTOR:      false,
		FraudScore: 0,
	}

	got := signalConfidence(s)

	assert.Equal(t, 1.0, got)
}

func TestSignalConfidence_Mixed(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: true,

		PhonePresent: true,
		PhoneVoIP:    false,
		PhoneVirtual: false,

		IPPresent:  true,
		IsProxy:    true,
		IsTOR:      false,
		FraudScore: 0,
	}

	got := signalConfidence(s)

	assert.Greater(t, got, 0.0)
	assert.Less(t, got, 1.0)
}

func TestCompute_ConfidenceMatchesSignalConfidence(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: false,
	}

	result := Compute(s)
	expected := signalConfidence(s)

	assert.Equal(t, expected, result.Confidence)
}

func TestCompute_IsSafe_LowScoreHighConfidence(t *testing.T) {
	s := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: false,
	}

	result := Compute(s)

	assert.True(t, result.IsSafe)
}

func TestCompute_IsSafe_FalseWhenTOROrProxy(t *testing.T) {
	tor := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: false,
		IPPresent:       true,
		IsTOR:           true,
		FraudScore:      0,
	}

	proxy := Signals{
		EmailPresent:    true,
		EmailNoMX:       false,
		EmailDisposable: false,
		IPPresent:       true,
		IsProxy:         true,
		FraudScore:      0,
	}

	assert.False(t, Compute(tor).IsSafe)
	assert.False(t, Compute(proxy).IsSafe)
}
