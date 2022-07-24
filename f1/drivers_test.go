package f1

import (
	"context"
	"testing"
)

const (
	numOfDrivers = 20
)

func TestDriversService_ListAll(t *testing.T) {
	c := NewClient(nil)
	drivers, err := c.Drivers.ListAll(context.Background())
	if err != nil {
		t.Errorf("Drivers.ListAll returned error: %v", err)
	}

	if len(drivers) != numOfDrivers {
		t.Errorf("Drivers.ListAll returned %v drivers, want %v", len(drivers), numOfDrivers)
	}

	for _, d := range drivers {
		if d.IsConstructor {
			t.Errorf("Drivers.ListAll returned team %v with ID %v. No teams should be returned.", d.DisplayName, d.ID)
		}
	}
}

func TestDriversService_GetOne(t *testing.T) {
	c := NewClient(nil)
	d1Id := 11
	d1, err := c.Drivers.GetOne(context.Background(), d1Id)
	if err != nil {
		t.Errorf("Drivers.GetOne returned error: %v", err)
	}

	if d1 == nil {
		t.Errorf("Drivers.GetOne returned %v, want driver with id %v", nil, d1Id)
	}

	if d1.ID != d1Id {
		t.Errorf("Drivers.GetOne returned driver with id %v, want driver with id %v", d1.ID, d1Id)
	}

	if d1.IsConstructor {
		t.Errorf("Drivers.GetOne returned team %v with ID %v. want driver with id %v.", d1.DisplayName, d1.ID, d1Id)
	}

	d2, err := c.Drivers.GetOne(context.Background(), 5)
	if err != nil {
		t.Errorf("Drivers.GetOne returned error: %v", err)
	}

	if d2 != nil {
		t.Errorf("Drivers.GetOne returned driver/team %v with ID %v, want %v", d2.DisplayName, d2.ID, nil)
	}

	d3, err := c.Drivers.GetOne(context.Background(), 1000)
	if err != nil {
		t.Errorf("Drivers.GetOne returned error: %v", err)
	}

	if d3 != nil {
		t.Errorf("Drivers.GetOne returned driver/team %v with ID %v, want %v", d3.DisplayName, d3.ID, nil)
	}
}
