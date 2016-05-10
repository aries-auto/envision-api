package envisionAPI

import "fmt"

// RequestResult Stores the global status variables returned by
// the iConMedia API
type RequestResult struct {
	Result  int    `json:"Result"`
	Error   int    `json:"Error"`
	Message string `json:"Message"`
}

// Verify Checks the Error property of the RequestResult and
// returns an error if necessary
func (r RequestResult) Verify() error {
	if r.Error == 1 {
		return fmt.Errorf("request error: %s", r.Message)
	}

	return nil
}
