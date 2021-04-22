package user

import "goyave.dev/goyave/v3/validation"

var (
	Register validation.RuleSet = validation.RuleSet{
		"email":      {"required", "string"},
		"password":   {"required", "string"},
		"first_name": {"required", "string"},
		"last_name":  {"required", "string"},
	}
)

var (
	Login validation.RuleSet = validation.RuleSet{
		"email":    {"required", "string"},
		"password": {"required", "string"},
	}
)
