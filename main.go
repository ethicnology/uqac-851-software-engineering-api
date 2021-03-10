package main

import (
	"os"

	"github.com/851-software-engineering/PrixBanque-api/http/route"
	_ "github.com/851-software-engineering/PrixBanque-api/http/validation"

	"goyave.dev/goyave/v3"
	_ "goyave.dev/goyave/v3/database/dialect/mysql"
)

func main() {
	// Application entry point.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
