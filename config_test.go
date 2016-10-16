package envisionAPI

import (
	"os"
	"testing"
)

func testConfigEnv() Config {
	return Config{
		Login:         os.Getenv("TEST_LOGIN"),
		Password:      os.Getenv("TEST_PASSWORD"),
		VehicleDomain: os.Getenv("TEST_VEHICLE_DOMAIN"),
		ImageDomain:   os.Getenv("TEST_IMAGE_DOMAIN"),
		SwatchDomain:  os.Getenv("TEST_SWATCH_DOMAIN"),
		UserID:        os.Getenv("TEST_USER_ID"),
	}
}

func TestNewConfig(t *testing.T) {

	c, err := NewConfig()
	if c != nil {
		t.Fatal("config should be nil")
	}
	if err == nil {
		t.Fatal("should have returned an error on too few params")
	}

	c, err = NewConfig("", "", "", "", "", "")
	if err != nil {
		t.Fatalf("error should have been nil: %s\n", err)
	}
	if c == nil {
		t.Fatal("config should not have been nil")
	}

}
