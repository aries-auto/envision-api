package envisionAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Vehicle Stores the iConMedia Vehicle
type Vehicle struct {
	ID       string `json:"intVehicleID"`
	Year     string `json:"intYear"`
	Make     string `json:"strMake"`
	Model    string `json:"strModel"`
	BodyType string `json:"strBodyType"`
}

// ProductVehicleResponse Contains the response data from a Product Vehicle match
type ProductVehicleResponse struct {
	Result   int       `json:"Result"`
	Error    int       `json:"Error"`
	Message  string    `json:"Message"`
	Vehicles []Vehicle `json:"Vehicles"`
}

// Fitment Determines if a product (Number) matches a queried Vehicle. If
// Mapped is `1`, the product fits the provided vehicle
type Fitment struct {
	Number string `json:"strPartNumber"`
	Mapped string `json:"bitMapped"`
}

// FitmentResponse The response returned from a vehicle match (MatchFitment) query
type FitmentResponse struct {
	Resut    int       `json:"Result"`
	Error    int       `json:"Error"`
	Message  string    `json:"Message"`
	Fitments []Fitment `json:"PartNumbers"`
}

// GetVehiclesByProduct Returns the vehicles that match a given product
func GetVehiclesByProduct(c Config, productID string) (*ProductVehicleResponse, error) {

	vals := url.Values{}
	vals.Add("l", c.Login)
	vals.Add("p", c.Password)
	vals.Add("fnct", "VP")
	vals.Add("part", productID)

	resp, err := http.Get(
		fmt.Sprintf(
			"%s/ap-ar-vehicle-parts.cfm?%s",
			c.Domain,
			vals.Encode(),
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var pv ProductVehicleResponse
	err = json.NewDecoder(resp.Body).Decode(&pv)

	return &pv, err
}

// MatchFitment Returns a mapped response of Products that do or do not fit the
// provided vehicleID
func MatchFitment(c Config, vehicleID int, productIDs ...string) (*FitmentResponse, error) {

	vals := url.Values{}
	vals.Add("l", c.Login)
	vals.Add("p", c.Password)
	vals.Add("fnct", "VMP")
	vals.Add("id", strconv.Itoa(vehicleID))
	vals.Add("part", strings.Join(productIDs, ","))

	resp, err := http.Get(
		fmt.Sprintf(
			"%s/ap-ar-vehicle-parts.cfm?%s",
			c.Domain,
			vals.Encode(),
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var f FitmentResponse
	err = json.NewDecoder(resp.Body).Decode(&f)

	return &f, err
}
