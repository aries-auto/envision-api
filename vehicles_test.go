package envisionAPI

import "testing"

func TestGetVehiclesByProduct(t *testing.T) {

	c := testConfigEnv()

	tmp := c.VehicleDomain
	c.VehicleDomain = ""
	_, err := GetVehiclesByProduct(c, "203001")
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = "http://www.google.com"
	_, err = GetVehiclesByProduct(c, "203001")
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = tmp

	pv, err := GetVehiclesByProduct(c, "203001")
	if err != nil {
		t.Fatal("failed")
	} else if pv == nil {
		t.Fatal("ProductVehicleResponse should not be nil")
	}
}

func TestMatchFitment(t *testing.T) {

	c := testConfigEnv()

	tmp := c.VehicleDomain
	c.VehicleDomain = ""
	_, err := MatchFitment(c, 5460, "203001", "1042")
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = "http://www.google.com"
	_, err = MatchFitment(c, 5460, "203001", "1042")
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = tmp

	f, err := MatchFitment(c, 5460, "203001", "1042")
	if err != nil {
		t.Fatal(err)
	} else if f == nil {
		t.Fatal("FitmentResponse should not be nil")
	}
}

func TestYearMakeModel(t *testing.T) {
	c := testConfigEnv()
	yearStr := "1996"
	makeStr := "JEEP"
	modelStr := "GRAND CHEROKEE"
	tmp := c.VehicleDomain
	c.VehicleDomain = ""
	_, err := GetVehicleByYearMakeModel(c, yearStr, makeStr, modelStr)
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = "http://www.google.com"
	_, err = GetVehicleByYearMakeModel(c, yearStr, makeStr, modelStr)
	if err == nil {
		t.Fatal(err)
	}

	c.VehicleDomain = tmp
	f, err := GetVehicleByYearMakeModel(c, yearStr, makeStr, modelStr)
	if err != nil {
		t.Fatal(err)
	} else if f == nil {
		t.Fatal("ProductReponse should not be nil")
	}
}

func TestVehicleImage(t *testing.T) {
	c := testConfigEnv()
	vehicleID := "0"
	colorID := "0"
	skus := []string{}
	tmp := c.ImageDomain

	c.ImageDomain = ""
	_, err := GetVehicleImage(c, vehicleID, colorID, skus)
	if err == nil {
		t.Fatal(err)
	}

	c.ImageDomain = "http://www.google.com"
	img, err := GetVehicleImage(c, vehicleID, colorID, skus)
	if err == nil {
		t.Fatal(err)
	}

	c.ImageDomain = tmp

	tmp = c.SwatchDomain
	c.SwatchDomain = ""
	img, err = GetVehicleImage(c, vehicleID, colorID, skus)
	if err != nil || (len(img.IconImages) > 0 && len(img.IconImages[0].ColorNames) > 0) {
		t.Fatal(err)
	}
	c.SwatchDomain = tmp

	img, err = GetVehicleImage(c, vehicleID, colorID, skus)
	if err != nil {
		t.Fatal(err)
	} else if img == nil {
		t.Fatal("VehicleImageResponse should not be nil")
	}
}
