package envisionAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ProductNumbers Array of SKUS associated to products
type ProductNumbers struct {
	Number string `json:"strPartNumber"`
}

// VehicleProductResponse The response that is returned the from web service
type VehicleProductResponse struct {
	RequestResult
	Numbers []ProductNumbers `json:"PartNumbers"`
}

// NoFitmentResponse The response from a NoFitment request
type NoFitmentResponse struct {
	RequestResult
	Numbers []ProductNumbers `json:"PartNumbers"`
}

// GetVehicleProducts Retruns the product numbers that are matched to the given
// vehcicle identifier
func GetVehicleProducts(config Config, vehicleID int) (*VehicleProductResponse, error) {

	vals := url.Values{}
	vals.Add("l", config.Login)
	vals.Add("p", config.Password)
	vals.Add("id", strconv.Itoa(vehicleID))

	resp, err := http.Get(
		fmt.Sprintf(
			"%s?%s",
			config.Domain,
			vals.Encode(),
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var vp VehicleProductResponse
	err = json.NewDecoder(resp.Body).Decode(&vp)
	if err != nil {
		return nil, err
	}

	return &vp, vp.Verify()
}

// NoFitment Returns a list of products that have no fitment information.
func NoFitment(config Config) (*NoFitmentResponse, error) {

	vals := url.Values{}
	vals.Add("l", config.Login)
	vals.Add("p", config.Password)
	vals.Add("fnct", "PNM")

	resp, err := http.Get(
		fmt.Sprintf(
			"%s?%s",
			config.Domain,
			vals.Encode(),
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var nf NoFitmentResponse
	err = json.NewDecoder(resp.Body).Decode(&nf)
	if err != nil {
		return nil, err
	}

	return &nf, nf.Verify()
}
