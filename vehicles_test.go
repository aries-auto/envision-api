package envisionAPI

import "testing"

func TestGetVehiclesByProduct(t *testing.T) {

	c := testConfigEnv()

	tmp := c.Domain
	c.Domain = ""
	_, err := GetVehiclesByProduct(c, "203001")
	if err == nil {
		t.Fatal(err)
	}

	c.Domain = tmp

	pv, err := GetVehiclesByProduct(c, "203001")
	if err != nil {
		t.Fatal("failed")
	} else if pv == nil {
		t.Fatal("ProductVehicleResponse should not be nil")
	}
}

func TestMatchFitment(t *testing.T) {

	c := testConfigEnv()

	tmp := c.Domain
	c.Domain = ""
	_, err := MatchFitment(c, 5460, "203001", "1042")
	if err == nil {
		t.Fatal(err)
	}

	c.Domain = tmp

	f, err := MatchFitment(c, 5460, "203001", "1042")
	if err != nil {
		t.Fatal(err)
	} else if f == nil {
		t.Fatal("FitmentResponse should not be nil")
	}
}
