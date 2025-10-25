package errors

import "errors"

var (
	ErrCarsNotFound          = errors.New("cars not found")
	ErrCarNotFound           = errors.New("cars not found")
	ErrAvailableCarsNotFound = errors.New("available cars not found")
	ErrCarNotAvailable       = errors.New("car is not available")
)
