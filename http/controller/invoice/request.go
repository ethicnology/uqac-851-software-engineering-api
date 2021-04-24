package invoice

import "goyave.dev/goyave/v3/validation"

var (
	Post validation.RuleSet = validation.RuleSet{
		"to":        {"required", "string"},
		"amount":    {"required", "numeric"},
		"acquitted": {"required", "bool"},
		"due_date":  {"required", "date"},
	}
)
