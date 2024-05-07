package auth

import "github.com/golodash/galidator"

var (
	g = galidator.G()
)

func GetValidator(requestType interface{}) galidator.Validator {
	return g.Validator(requestType)
}
