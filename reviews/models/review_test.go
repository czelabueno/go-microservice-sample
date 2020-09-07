package models

import (
	"fmt"
	"testing"
)

func NewReview(starts int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Starts:  starts,
		Comment: comment,
	}
}

func Test_withCorrectParams(t *testing.T) {
	var starts = 4
	var comment = "The iphone x"
	f := NewReview(starts, comment)
	err := f.validate()
	if err != nil { // With correct params result expected is no error
		t.Error("The validation did not pass with params: " + fmt.Sprint(starts) + " and " + comment)
		t.Fail() // So if error is not null the test must be fail
	}
}

func Test_withOutParamsAccepted(t *testing.T) {
	var starts = 8
	var comment = "The iphone x"
	f := NewReview(starts, comment)
	err := f.validate()
	if err == nil { // with incorrect params result expected is an error object
		t.Error("The validation did not pass with params: " + fmt.Sprint(starts) + " and " + comment)
		t.Fail() // So if error is null the test must be fail
	}
}
