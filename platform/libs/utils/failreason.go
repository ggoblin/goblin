package utils

import ()

type FailReason struct {
	Success bool
	Message string
}

func NewFailReason(success bool, msg string) FailReason {
	return FailReason{Success: success, Message: msg}
}
