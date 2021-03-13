package main

import (
	"os"

	"github.com/ethicnology/uqac-851-software-engineering-api/http/route"
	_ "github.com/ethicnology/uqac-851-software-engineering-api/http/validation"

	"goyave.dev/goyave/v3"
	_ "goyave.dev/goyave/v3/database/dialect/mysql"
)

func main() {
	// Application entry point.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
