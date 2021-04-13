package operation

import "goyave.dev/goyave/v3/validation"

var (
	Structure validation.RuleSet = validation.RuleSet{
		"amount":      {"required", "numeric"},
		"receiver_id": {"required", "numeric", "min:1"},
	}
)
