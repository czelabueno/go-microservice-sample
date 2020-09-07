package models

import (
	"errors"
	"time"
)

const maxLengthInComments = 400

// Review model struct for review bean
type Review struct {
	Id      int64
	Starts  int // range from 1-5
	Comment string
	Date    time.Time // created at timestamp
}

// CreateReviewCMD command to create new Review entity
type CreateReviewCMD struct {
	Starts  int    `json:"starts"`
	Comment string `json:"comment"`
}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Starts < 1 || cmd.Starts > 5 {
		return errors.New("starts must be between 1 - 5")
	}
	if len(cmd.Comment) > maxLengthInComments {
		return errors.New("comment must be less or equal that 400 chars")
	}
	return nil
}
