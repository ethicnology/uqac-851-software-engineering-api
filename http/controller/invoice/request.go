package invoice

import "goyave.dev/goyave/v3/validation"

var (
	Structure validation.RuleSet = validation.RuleSet{
		"amount":      {"required", "numeric"},
		"receiver_id": {"required", "numeric", "min:1"},
		"acquitted":   {"required", "bool"},
		"due_date":    {"required", "date"},
	}
)
