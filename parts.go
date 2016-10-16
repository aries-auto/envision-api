package envisionAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ProductNumbers Array of SKUS associated to products
type ProductNumbers struct {
	Number string `json:"strPartNumber"`
}

// VehicleProductResponse Matches the assigned product numbers to a given vehicle
type VehicleProductResponse struct {
	RequestResult
	Numbers []ProductNumbers `json:"PartNumbers"`
}

// NoFitmentResponse Shows all products that have no vehice fitment
// information
type NoFitmentResponse struct {
	RequestResult
	Numbers []ProductNumbers `json:"PartNumbers"`
}

// LayersResponse Shows the Layer associated to a given product
type LayersResponse struct {
	RequestResult
	Layers []ProductLayer `json:"PartNumbers"`
}

// ProductLayer Describes the Layer information associated to a product
type ProductLayer struct {
	Name          string `json:"strName"`
	ProductNumber string `json:"strPartNumber"`
	Layer         string `json:"strLayer"`
	LayerID       string `json:"intLayerID"`
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
			config.VehicleDomain,
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
			config.VehicleDomain,
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

// GetLayers Returns a list of products that have no fitment information.
func GetLayers(config Config, vehicleID string, parts ...string) (*LayersResponse, error) {

	vals := url.Values{}
	vals.Add("l", config.Login)
	vals.Add("p", config.Password)
	vals.Add("fnct", "PL")
	vals.Add("id", vehicleID)
	vals.Add("part", strings.Join(parts, ","))

	resp, err := http.Get(
		fmt.Sprintf(
			"%s?%s",
			config.VehicleDomain,
			vals.Encode(),
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var lr LayersResponse
	err = json.NewDecoder(resp.Body).Decode(&lr)
	if err != nil {
		return nil, err
	}

	return &lr, lr.Verify()
}
