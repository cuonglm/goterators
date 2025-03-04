package goterators

import (
	"testing"
)

func TestMap(t *testing.T) {
	type Order struct {
		Quantity int
		Price    float64
	}
	testSource := []Order{
		{Quantity: 5, Price: 20.5},
		{Quantity: 2, Price: 3},
		{Quantity: 1, Price: 11.3},
	}
	expectedItems := []float64{102.5, 6, 11.3}

	mappedItems := Map(testSource, func(item Order) float64 {
		return float64(item.Quantity) * item.Price
	})
	if len(expectedItems) != len(mappedItems) {
		t.Fatalf("Expected length = %v, actual length = %v", len(expectedItems), len(mappedItems))
	}
	for index := range expectedItems {
		if expectedItems[index] != mappedItems[index] {
			t.Errorf("Index %v, expected = %v, actual = %v", index, expectedItems[index], mappedItems[index])
		}
	}
}
