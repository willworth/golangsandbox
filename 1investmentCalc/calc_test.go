// package main

// func TestCalc(t *testing.T){

// }

package main

import (
	"errors"
	"io/ioutil"
	"math"
	"os"
	"testing"
)

// Helper function to compare floats with a tolerance
func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestCalculateFinancials(t *testing.T) {
	tests := []struct {
		revenue  float64
		expenses float64
		taxRate  float64
		ebt      float64
		profit   float64
		ratio    float64
		err      error
	}{
		{100, 50, 10, 50, 45, 1.111111, nil},
		{200, 150, 20, 50, 40, 1.25, nil},
		{100, 100, 10, 0, 0, 0, errors.New("Profit is zero, can't calculate ratio")},
	}

	tolerance := 0.0001 // Define a small tolerance for floating-point comparison

	for _, tt := range tests {
		ebt, profit, ratio, err := calculateFinancials(tt.revenue, tt.expenses, tt.taxRate)
		if !almostEqual(ebt, tt.ebt, tolerance) || !almostEqual(profit, tt.profit, tolerance) || !almostEqual(ratio, tt.ratio, tolerance) || ((err != nil) != (tt.err != nil)) {
			t.Errorf("calculateFinancials(%f, %f, %f): expected %f, %f, %f, error: %v, got %f, %f, %f, error: %v",
				tt.revenue, tt.expenses, tt.taxRate, tt.ebt, tt.profit, tt.ratio, tt.err, ebt, profit, ratio, err)
		}
	}
}

func TestStoreResults(t *testing.T) {
	ebt, profit, ratio := 100.0, 80.0, 1.25
	err := storeResults(ebt, profit, ratio)
	if err != nil {
		t.Errorf("storeResults returned an unexpected error: %v", err)
	}

	content, err := ioutil.ReadFile("results.txt")
	if err != nil {
		t.Errorf("Error reading results.txt: %v", err)
	}

	expected := "EBT: 100.0\nProfit: 80.0\nRatio: 1.250\n"
	if string(content) != expected {
		t.Errorf("Expected file content:\n%s\nGot:\n%s\n", expected, content)
	}

	// Cleanup
	os.Remove("results.txt")
}
