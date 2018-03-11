package main

import (
	"reflect"
	"testing"
)

func TestConvertTemp(t *testing.T) {
	var tests = []struct {
		input  string
		output map[string]float64
	}{
		{"100c", map[string]float64{"f": 212, "k": 100 + 273.15}},
		{"120k", map[string]float64{"c": -153.14999999999998, "f": -243.66999999999996}},
	}
	for _, test := range tests {
		if out := convertTemp(test.input); !reflect.DeepEqual(out, test.output) {
			t.Errorf("test result: %v\n", test.output)
			t.Errorf("actual result: %v\n", out)
		}
	}
}
