package envisionAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// IconVehicle Stores the iConMedia Vehicle
type IconVehicle struct {
	ID       string `json:"intVehicleID"`
	Year     string `json:"intYear"`
	Make     string `json:"strMake"`
	Model    string `json:"strModel"`
	BodyType string `json:"strBodyType"`
	ImageURL string `json:"strImageURL"`
}

type Vehicle struct {
	Identifier int    `json:"id"`
	Year       string `json:"year"`
	Make       string `json:"make"`
	Model      string `json:"model"`
}

// IconImage represents the iConMedia image object definition.
type IconImage struct {
	Source           string   `json:"src"`
	ColorIdentifiers []int    `json:"colorID"`
	ColorNames       []string `json:"colorName"`
	ColorSwatches    []string `json:"colorImage"`
}

// ColorOption defines the color swatch schema for a vehicle.
type ColorOption struct {
	Identifier int      `json:"id"`
	Name       string   `json:"name"`
	Image      *url.URL `json:"swatch"`
}

// Image defines the returned image response after some re-organizing of
// the iConMedia response.
type Image struct {
	Source       *url.URL      `json:"source"`
	ColorOptions []ColorOption `json:"colors"`
}

type ImageResponse struct {
	RequestResult
	Vehicle      Vehicle  `json:"vehicle"`
	Image        Image    `json:"image"`
	MappableSKUS []string `json:"mappable"`
}

// ProductVehicleResponse Contains the response data from a Product Vehicle match
type ProductVehicleResponse struct {
	RequestResult
	Vehicles []IconVehicle `json:"Vehicles"`
}

// VehicleImageResponse represents the visual representation of a provided
// vehicle, color, and SKU listing.
type VehicleImageResponse struct {
	RequestResult
	IconImages []IconImage `json:"img,omitempty"`
	Images     []Image     `json:"images,omitempty"`
}

// Fitment Determines if a product (Number) matches a queried Vehicle. If
// Mapped is `1`, the product fits the provided vehicle
type Fitment struct {
	Number string `json:"strPartNumber"`
	Mapped string `json:"bitMapped"`
}

// FitmentResponse The response returned from a vehicle match (MatchFitment) query
type FitmentResponse struct {
	RequestResult
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
			"%s?%s",
			c.VehicleDomain,
			vals.Encode(),
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var pv ProductVehicleResponse
	err = json.NewDecoder(resp.Body).Decode(&pv)
	if err != nil {
		return nil, err
	}

	return &pv, pv.Verify()
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
			"%s?%s",
			c.VehicleDomain,
			vals.Encode(),
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var f FitmentResponse
	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		return nil, err
	}

	return &f, f.Verify()
}

func GetImage(c Config, yearStr string, makeStr string, modelStr string, colorID int, skus []string) (*ImageResponse, error) {
	var vehicle *IconVehicle
	var imgResponse *VehicleImageResponse
	var allSKUS []string
	var err error

	vehicle, allSKUS, err = getFullestVehicle(c, yearStr, makeStr, modelStr, colorID, skus)
	if err != nil {
		return nil, err
	}

	imgResponse, err = GetVehicleImage(c, vehicle.ID, colorID, skus)
	if err != nil {
		// try again with no color
		imgResponse, err = GetVehicleImage(c, vehicle.ID, 0, skus)
		if err != nil {
			return nil, err
		}
	}

	if len(imgResponse.Images) == 0 {
		return nil, fmt.Errorf("no image available for %s %s %s in color %d with parts %+v", yearStr, makeStr, modelStr, colorID, skus)
	}

	var resp ImageResponse
	resp.MappableSKUS = allSKUS
	sort.Strings(resp.MappableSKUS)
	resp.Image = imgResponse.Images[0]
	resp.Vehicle = Vehicle{
		Year:  yearStr,
		Make:  makeStr,
		Model: modelStr,
	}

	resp.Vehicle.Identifier, _ = strconv.Atoi(vehicle.ID)

	return &resp, nil
}

func getFullestVehicle(c Config, yearStr string, makeStr, modelStr string, colorID int, skus []string) (*IconVehicle, []string, error) {
	var err error
	var vResponse *ProductVehicleResponse

	// get the trim options for a year/make/model
	vResponse, err = GetVehicleByYearMakeModel(c, yearStr, makeStr, modelStr)
	if err != nil {
		return nil, nil, err
	}

	var max = 0
	var maxProductVehicle *IconVehicle
	var allSKUS []string
	for _, v := range vResponse.Vehicles {
		var prodResponse *VehicleProductResponse
		var intVehicleID int
		intVehicleID, err = strconv.Atoi(v.ID)
		if err != nil {
			continue
		}
		prodResponse, err = GetVehicleProducts(c, intVehicleID)
		if err != nil {
			continue
		}

		for _, pn := range prodResponse.Numbers {
			if !contains(allSKUS, pn.Number) {
				allSKUS = append(allSKUS, pn.Number)
			}
		}

		if len(skus) > 0 {
			var matched = 0
			for _, pn := range prodResponse.Numbers {
				if contains(skus, pn.Number) {
					matched++
				}
			}

			if matched != len(skus) {
				continue
			}
		}

		if len(prodResponse.Numbers) > max {
			max = len(prodResponse.Numbers)
			maxProductVehicle = &v
		}
	}

	return maxProductVehicle, allSKUS, nil
}

// GetVehicleByYearMakeModel Returns the Vehicle(s) that match a given year, make, and model
func GetVehicleByYearMakeModel(c Config, yearStr, makeStr, modelStr string) (*ProductVehicleResponse, error) {
	vals := url.Values{}
	vals.Add("l", c.Login)
	vals.Add("p", c.Password)
	vals.Add("fnct", "VI")
	vals.Add("year", yearStr)
	vals.Add("make", makeStr)
	vals.Add("model", modelStr)

	resp, err := http.Get(
		fmt.Sprintf(
			"%s?%s",
			c.VehicleDomain,
			vals.Encode(),
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var f ProductVehicleResponse
	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		return nil, err
	}

	return &f, f.Verify()
}

// GetVehicleImage returns the an image that works with the provided
// vehicle, color, and SKU mapping.
func GetVehicleImage(c Config, vehicleID string, colorID int, skus []string) (*VehicleImageResponse, error) {
	vals := url.Values{}
	vals.Add("usejson", "1")
	vals.Add("uid", c.UserID)
	vals.Add("vehicle", vehicleID)
	vals.Add("part", strings.Join(skus, ","))
	vals.Add("color", strconv.Itoa(colorID))

	resp, err := http.Get(
		fmt.Sprintf(
			"%s?%s",
			c.ImageDomain,
			vals.Encode(),
		),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var i VehicleImageResponse
	err = json.NewDecoder(resp.Body).Decode(&i)
	if err != nil {
		return nil, err
	}

	for _, iconImg := range i.IconImages {
		var img *Image
		img, err = iconImg.transform(c)
		if err != nil || img == nil || len(img.ColorOptions) == 0 {
			continue
		}

		i.Images = append(i.Images, *img)
	}

	return &i, i.Verify()
}

func (i IconImage) transform(c Config) (*Image, error) {
	var img Image
	var err error

	img.Source, err = url.Parse(i.Source)
	if err != nil {
		return nil, err
	}

	for iter := range i.ColorIdentifiers {
		ci := ColorOption{
			Identifier: i.ColorIdentifiers[iter],
			Name:       i.ColorNames[iter],
		}

		ci.Image, err = url.Parse(
			fmt.Sprintf(
				"%s/%s",
				c.SwatchDomain,
				i.ColorSwatches[iter],
			),
		)

		if err != nil {
			continue
		}

		img.ColorOptions = append(img.ColorOptions, ci)
	}

	return &img, nil
}

func contains(skus []string, sku string) bool {
	for _, s := range skus {
		if strings.Compare(s, sku) == 0 {
			return true
		}
	}

	return false
}
