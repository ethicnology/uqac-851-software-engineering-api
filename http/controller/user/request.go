package user

import "goyave.dev/goyave/v3/validation"

var (
	Structure validation.RuleSet = validation.RuleSet{
		"first_name": {"required", "string"},
		"last_name":  {"required", "string"},
		"email":      {"required", "string"},
		"password":   {"required", "string"},
	}
)
