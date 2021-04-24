package transfer

import (
	"goyave.dev/goyave/v3/validation"
)

var (
	Post validation.RuleSet = validation.RuleSet{
		"to":        {"required", "string"},
		"scheduled": {"required", "bool"},
		"instant":   {"required", "bool"},
		"amount":    {"required", "numeric"},
		"date":      {"required", "date"},
		"question":  {"required", "string"},
		"answer":    {"required", "string"},
	}
)
