package envisionAPI

import "testing"

func TestGetVehicleProducts(t *testing.T) {

	c := testConfigEnv()

	tmp := c.Domain
	c.Domain = ""
	_, err := GetVehicleProducts(c, 5460)
	if err == nil {
		t.Fatal(err)
	}

	c.Domain = tmp

	vp, err := GetVehicleProducts(c, 5460)
	if err != nil {
		t.Fatal(err)
	} else if vp == nil {
		t.Fatal("VehicleProductResponse should not be nil")
	}
}

func TestNoFitment(t *testing.T) {

	c := testConfigEnv()

	tmp := c.Domain
	c.Domain = ""
	_, err := NoFitment(c)
	if err == nil {
		t.Fatal(err)
	}

	c.Domain = tmp

	nf, err := NoFitment(c)
	if err != nil {
		t.Fatal(err)
	} else if nf == nil {
		t.Fatal("NoFitmentResponse should not be nil")
	}
}
