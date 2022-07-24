package f1

import (
	"context"
	"testing"
)

const (
	numOfCircuits = 22
)

func TestCircuitsService_ListAll(t *testing.T) {
	c := NewClient(nil)
	circuits, err := c.Circuits.ListAll(context.Background())
	if err != nil {
		t.Errorf("Circuits.ListAll returned error: %v", err)
	}

	if len(circuits) != numOfCircuits {
		t.Errorf("Circuits.ListAll returned %v circuits, want %v", len(circuits), numOfCircuits)
	}
}

func TestCircuitsService_GetOne(t *testing.T) {
	c := NewClient(nil)
	id := 11
	circuit, err := c.Circuits.GetOne(context.Background(), id)
	if err != nil {
		t.Errorf("Circuits.GetOne returned error: %v", err)
	}

	if circuit == nil {
		t.Errorf("Circuits.GetOne returned %v, want circuit with id %v", nil, id)
	}

	if circuit.GamePeriodID != id {
		t.Errorf("Circuits.GetOne returned circuit with GamePeriodId %v, want circuit with id %v", circuit.GamePeriodID, id)
	}

	if circuit.Circuit.ID != id {
		t.Errorf("Circuits.GetOne returned circuit with id %v, want circuit with id %v", circuit.Circuit.ID, id)
	}

	circuit2, err := c.Circuits.GetOne(context.Background(), 1000)
	if err != nil {
		t.Errorf("Circuits.GetOne returned error: %v", err)
	}

	if circuit2 != nil {
		t.Errorf("Circuits.GetOne returned circuit %v with ID %v, want %v", circuit2.Circuit.Name, circuit.GamePeriodID, nil)
	}
}
