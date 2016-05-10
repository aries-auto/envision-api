package envisionAPI

import (
	"os"
	"testing"
)

func testConfigEnv() Config {
	return Config{
		Login:    os.Getenv("TEST_LOGIN"),
		Password: os.Getenv("TEST_PASSWORD"),
		Domain:   os.Getenv("TEST_DOMAIN"),
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

	c, err = NewConfig("", "", "")
	if err != nil {
		t.Fatalf("error should have been nil: %s\n", err)
	}
	if c == nil {
		t.Fatal("config should not have been nil")
	}

}
