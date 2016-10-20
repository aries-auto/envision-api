package envisionAPI

import (
	"fmt"
	"sort"
	"strconv"
)

type MappedVehicle struct {
	Vehicle IconVehicle
	Parts   []ProductNumbers
}

type MappedVehicles []MappedVehicle

func getFullestVehicle(c Config, yearStr string, makeStr, modelStr string, colorID int, skus []string) (*IconVehicle, []string, error) {
	var err error
	var vResponse *ProductVehicleResponse

	// get the trim options for a year/make/model
	vResponse, err = GetVehicleByYearMakeModel(c, yearStr, makeStr, modelStr)
	if err != nil {
		return nil, nil, err
	}

	var mv MappedVehicles
	var allSKUS []string
	mv, allSKUS = vResponse.getProducts(c)
	sort.Sort(mv)

	if len(skus) > 0 {
		var tmp []MappedVehicle
		for _, m := range mv {
			var matched = 0
			for _, pn := range m.Parts {
				if contains(skus, pn.Number) {
					matched++
				}
			}

			if matched == len(skus) {
				tmp = append(tmp, m)
			}
		}

		mv = tmp
	}

	for _, m := range mv {
		var vImage *VehicleImageResponse
		vImage, err = GetVehicleImage(c, m.Vehicle.ID, colorID, skus)
		if err != nil || vImage == nil {
			continue
		}

		return &m.Vehicle, allSKUS, nil
	}

	return nil, nil, fmt.Errorf("no image available for %s %s %s in color %d with parts %+v", yearStr, makeStr, modelStr, colorID, skus)
}

func (vr *ProductVehicleResponse) getProducts(c Config) ([]MappedVehicle, []string) {
	var mapped []MappedVehicle
	var err error
	var allSKUS []string

	for _, v := range vr.Vehicles {
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

		mv := MappedVehicle{
			Vehicle: v,
			Parts:   prodResponse.Numbers,
		}

		mapped = append(mapped, mv)
	}

	return mapped, allSKUS
}

func (mv MappedVehicles) Len() int {
	return len(mv)
}

func (mv MappedVehicles) Less(i, j int) bool {
	return len(mv[i].Parts) > len(mv[j].Parts)
}

func (mv MappedVehicles) Swap(i, j int) {
	mv[i], mv[j] = mv[j], mv[i]
}
