package math

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddPositive(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {
	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}
	_, err = Add(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}
	_, err = Add(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestAddZero(t *testing.T) {
	_, err := Add(0, 2)
	if err == nil {
		t.Error("first arg zero - expected error not be nil")
	}
	_, err = Add(1, 0)
	if err == nil {
		t.Error("second arg zero - expected error not be nil")
	}
	_, err = Add(0, 0)
	if err == nil {
		t.Error("all arg zero - expected error not be nil")
	}
}

func TestEstimateValue(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{
			name:  "test 1",
			value: 1,
			want:  "small",
		},
		{
			name:  "test 11",
			value: 11,
			want:  "medium",
		},
		{
			name:  "test 111",
			value: 111,
			want:  "big",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EstimateValue(tt.value); got != tt.want {
				t.Errorf("EstimateValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEstimateValueTestify(t *testing.T) {
	t.Run("Small", func(t *testing.T) {
		assert.Equal(t, "small", EstimateValue(9))
	})

	t.Run("Medium", func(t *testing.T) {
		assert.Equal(t, "medium", EstimateValue(99))
	})

	t.Run("Big", func(t *testing.T) {
		assert.Equal(t, "big", EstimateValue(100))
	})
}

