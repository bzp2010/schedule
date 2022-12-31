package models

// Status indicates the current task status
type Status uint8

const (
	// StatusDisabled means this item is disabled
	StatusDisabled Status = iota

	// StatusEnabled means this item is enabled
	StatusEnabled
)
