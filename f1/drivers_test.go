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
			t.Errorf("Drivers.ListAll returned team %v with ID %v. No teams should be returned.", d.FirstName, d.ID)
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
		t.Errorf("Drivers.GetOne returned nil, want driver with id %v", d1Id)
	}

	if d1.ID != d1Id {
		t.Errorf("Drivers.GetOne returned driver with id %v, want driver with id %v", d1.ID, d1Id)
	}

	if d1.IsConstructor {
		t.Errorf("Drivers.GetOne returned team %v with ID %v. want driver with id %v.", d1.FirstName, d1.ID, d1Id)
	}

	d2, err := c.Drivers.GetOne(context.Background(), 5)
	if err != nil {
		t.Errorf("Drivers.GetOne returned error: %v", err)
	}

	if d2 != nil {
		t.Errorf("Drivers.GetOne returned driver/team %v with ID %v, want nil", d2.FirstName, d2.ID)
	}

	d3, err := c.Drivers.GetOne(context.Background(), 1000)
	if err != nil {
		t.Errorf("Drivers.GetOne returned error: %v", err)
	}

	if d3 != nil {
		t.Errorf("Drivers.GetOne returned driver/team %v with ID %v, want nil", d3.FirstName, d3.ID)
	}
}