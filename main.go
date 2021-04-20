package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/ethicnology/uqac-851-software-engineering-api/http/route"

	"goyave.dev/goyave/v3"
	_ "goyave.dev/goyave/v3/database/dialect/mysql"
)

func main() {
	// Init seed
	rand.Seed(time.Now().UnixNano())
	// Application entry point.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
