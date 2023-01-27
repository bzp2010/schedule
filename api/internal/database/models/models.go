package models

import (
	"fmt"
	"io"
	"strconv"
)

// Status indicates the current item status
type Status uint8

const (
	// StatusDisabled means this item is disabled
	StatusDisabled Status = iota

	// StatusEnabled means this item is enabled
	StatusEnabled

	// StatusDisabledString means this item is disabled (used for graphql)
	StatusDisabledString string = "DISABLED"

	// StatusEnabledString means this item is enabled (used for graphql)
	StatusEnabledString string = "ENABLED"
)

// MarshalGQL implements GQLGEN dependent encoding function
func (s Status) MarshalGQL(w io.Writer) {
	if s == StatusEnabled {
		fmt.Fprint(w, strconv.Quote(StatusEnabledString))
	} else {
		fmt.Fprint(w, strconv.Quote(StatusDisabledString))
	}
}

// UnmarshalGQL implements GQLGEN dependent decoding function
func (s *Status) UnmarshalGQL(v interface{}) error {
	statusStr, ok := v.(string)
	if !ok {
		return fmt.Errorf("status must be string")
	}

	if statusStr == StatusEnabledString {
		*s = StatusEnabled
	} else {
		*s = StatusDisabled
	}

	return nil
}
