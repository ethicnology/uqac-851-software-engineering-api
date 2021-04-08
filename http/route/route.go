package route

import (
	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/controller/bank"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/controller/user"
	"github.com/ethicnology/uqac-851-software-engineering-api/http/middleware"

	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/auth"
	"goyave.dev/goyave/v3/cors"
	"goyave.dev/goyave/v3/middleware/ratelimiter"
)

// Register all the application routes. This is the main route registrer.
func Register(router *goyave.Router) {
	// Applying default CORS settings (allow all methods and all origins)
	router.CORS(cors.Default())
	router.Middleware(ratelimiter.New(model.RateLimiterFunc))

	authenticator := auth.Middleware(&model.User{}, &auth.JWTAuthenticator{})
	myRoutes(router, authenticator)
}

func myRoutes(parent *goyave.Router, authenticator goyave.Middleware) {
	authRouter := parent.Subrouter("/auth")

	jwtController := auth.NewJWTController(&model.User{})
	jwtController.UsernameField = "email"
	authRouter.Post("/register", user.Store).Validate(user.Structure)
	authRouter.Post("/login", jwtController.Login).Validate(user.Structure)

	userRouter := parent.Subrouter("/users/{email}")
	userRouter.Middleware(authenticator)
	userRouter.Middleware(middleware.Owner)
	userRouter.Get("/", user.Show)
	userRouter.Patch("/", user.Update)
	userRouter.Delete("/", user.Destroy)

	bankRouter := userRouter.Subrouter("/banks")
	bankRouter.Get("/", bank.Index)
	bankRouter.Post("/", bank.Store).Validate(bank.Structure)
	bankRouter.Get("/{id:[0-9]+}", bank.Show)
	bankRouter.Patch("/{id:[0-9]+}", bank.Update).Validate(bank.Structure)
	bankRouter.Delete("/{id:[0-9]+}", bank.Destroy)

	operationRouter := bankRouter.Subrouter("/operations")
	operationRouter.Get("/", operation.Index)
	operationRouter.Post("/", operation.Store).Validate(operation.Structure)
	operationRouter.Get("/{id:[0-9]+}", operation.Show)
	operationRouter.Patch("/{id:[0-9]+}", operation.Update).Validate(operation.Structure)
	operationRouter.Delete("/{id:[0-9]+}", operation.Destroy)
}
