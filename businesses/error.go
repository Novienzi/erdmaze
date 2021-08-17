package businesses

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrNewsIDResource = errors.New("(NewsID) not found or empty")

	ErrNewsTitleResource = errors.New("(NewsTitle) not found or empty")

	ErrActivityNotFound = errors.New("activity not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrUsernamePasswordNotFound = errors.New("(Username) or (Password) empty")

	ErrLocationNotFound = errors.New("location not found")

	ErrTourismsIDResource = errors.New("(TourismsID) not found or empty")

	ErrTourismsNameResource = errors.New("(TourismsName) not found or empty")
)
