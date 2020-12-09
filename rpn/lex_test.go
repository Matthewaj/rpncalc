package rpn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name       string
		expression []byte
		want       float64
	}{
		{"Should be == 2", []byte("1 1 +"), 2.00},
		{"Should be == 1", []byte("(1)"), 1.00},
		{"Should be == 5", []byte("5"), 5.00},
		{"Should be == 4", []byte("2 2 *"), 4.00},
		{"Should be == 0.841", []byte("1 sin"), 0.841},
		{"Should be == 0.154", []byte("30 cos"), 0.154},
		{"Should be == 7.244", []byte("14 tan"), 7.244},
		{"Should be == 15.708", []byte("(3.141 (2 3 +) (1.571 sin) *)"), 15.708},
		{"Should be == 200", []byte("((5 2 * (2 2 +)) (10 2 /))"), 200.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.expression)

			assert.NoError(t, err)
			assert.InDelta(t, tt.want, got, 0.01)
		})
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		name       string
		expression []byte
	}{
		{"sin_5", []byte("sin 5")},
		{"tan_30", []byte("tan 30")},
		{"cos_30", []byte("cos 30")},
		{"invalid brackets", []byte(")(")},
		{"empty brackets", []byte("()")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.expression)

			assert.NotNil(t, err)
		})
	}
}
