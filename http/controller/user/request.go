package user

import "goyave.dev/goyave/v3/validation"

var (
	Structure validation.RuleSet = validation.RuleSet{
		"email":    {"required", "string"},
		"password": {"required", "string"},
	}
)
