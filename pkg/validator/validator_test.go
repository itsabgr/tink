package validator

import (
	"testing"
)

func TestValidator(t *testing.T) {
	result := ValidateURI([]byte("https://example.com"))
	if result != nil {
		t.FailNow()
	}
	result = ValidateURI([]byte("http://google.com"))
	if result != ErrNonOKResp {
		t.FailNow()
	}
}
