package envisionAPI

import "fmt"

// Config Stores the global access information
type Config struct {
	Login         string
	Password      string
	VehicleDomain string
	ImageDomain   string
	SwatchDomain  string
	UserID        string
}

// NewConfig Generates a Config based off the params
func NewConfig(params ...string) (*Config, error) {
	if len(params) < 2 {
		return nil, fmt.Errorf("you must provide a minimum of the login and password")
	}

	c := &Config{
		Login:    params[0],
		Password: params[1],
	}

	if len(params) > 2 {
		c.VehicleDomain = params[2]
	}
	if len(params) > 3 {
		c.ImageDomain = params[3]
	}
	if len(params) > 4 {
		c.SwatchDomain = params[4]
	}
	if len(params) > 5 {
		c.UserID = params[5]
	}

	return c, nil
}
