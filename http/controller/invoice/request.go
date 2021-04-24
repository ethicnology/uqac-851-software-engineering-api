package invoice

import "goyave.dev/goyave/v3/validation"

var (
	Post validation.RuleSet = validation.RuleSet{
		"to":       {"required", "string"},
		"amount":   {"required", "numeric"},
		"due_date": {"required", "date"},
	}
)
var (
	Patch validation.RuleSet = validation.RuleSet{
		"acquitted": {"required", "bool"},
	}
)
