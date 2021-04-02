package bank

import "goyave.dev/goyave/v3/validation"

var (
	Structure validation.RuleSet = validation.RuleSet{
		"balance": {"required", "numeric"},
	}
)
