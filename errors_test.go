package envisionAPI

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	rr := RequestResult{
		Result:  1,
		Error:   1,
		Message: "random error",
	}

	err := rr.Verify()
	if err == nil {
		t.Fatal("error should not have been nil")
	}
	if err.Error() != fmt.Sprintf("request error: %s", rr.Message) {
		t.Fatal("error doesn't match")
	}

	rr.Error = 0
	rr.Message = ""

	err = rr.Verify()
	if err != nil {
		t.Fatal("error should have been nil")
	}
}
