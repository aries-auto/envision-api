package envisionAPI

import "fmt"

// Config Stores the global access information
type Config struct {
	Login    string
	Password string
	Domain   string
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

	if len(params) == 3 {
		c.Domain = params[2]
	}

	return c, nil
}
