package transfer

import "goyave.dev/goyave/v3/validation"

var (
	Structure validation.RuleSet = validation.RuleSet{
		"receiver_id": {"required", "numeric", "min:1"},
		"scheduled":   {"required", "bool"},
		"amount":      {"required", "numeric"},
		"date":        {"required", "date"},
		"question":    {"required", "string"},
		"answer":      {"required", "string"},
	}
)
