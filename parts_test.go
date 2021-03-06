package envisionAPI

import "testing"

func TestGetVehicleProducts(t *testing.T) {

	c := testConfigEnv()

	tmp := c.VehicleDomain
	c.VehicleDomain = ""
	_, err := GetVehicleProducts(c, 5460)
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = "http://www.google.com"
	_, err = GetVehicleProducts(c, 5460)
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = tmp

	vp, err := GetVehicleProducts(c, 5460)
	if err != nil {
		t.Fatal(err)
	} else if vp == nil {
		t.Fatal("VehicleProductResponse should not be nil")
	}
}

func TestNoFitment(t *testing.T) {

	c := testConfigEnv()

	tmp := c.VehicleDomain
	c.VehicleDomain = ""
	_, err := NoFitment(c)
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = "http://www.google.com"
	_, err = NoFitment(c)
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = tmp

	nf, err := NoFitment(c)
	if err != nil {
		t.Fatal(err)
	} else if nf == nil {
		t.Fatal("NoFitmentResponse should not be nil")
	}
}

func TestGetLayers(t *testing.T) {

	c := testConfigEnv()

	tmp := c.VehicleDomain
	c.VehicleDomain = ""
	_, err := GetLayers(c, "", "")
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = "http://www.google.com"
	_, err = GetLayers(c, "", "")
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = tmp

	lr, err := GetLayers(c, "", "")
	if err != nil {
		t.Fatal(err)
	} else if lr == nil {
		t.Fatal("LayersResponse should not be nil")
	}
}
