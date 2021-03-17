package route

import (
	"github.com/ethicnology/uqac-851-software-engineering-api/database/model"
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
	UserRoutes(router, authenticator)
}

func UserRoutes(parent *goyave.Router, authenticator goyave.Middleware) {
	authRouter := parent.Subrouter("/auth")

	jwtController := auth.NewJWTController(&model.User{})
	jwtController.UsernameField = "email"
	authRouter.Post("/register", user.Store).Validate(user.Structure)
	authRouter.Post("/login", jwtController.Login).Validate(user.Structure)

	userRouter := parent.Subrouter("/users")
	userRouter.Middleware(authenticator)
	userRouter.Get("/{email}", user.Show).Middleware(middleware.Owner)
	userRouter.Patch("/{email}", user.Update).Middleware(middleware.Owner)
	userRouter.Delete("/{email}", user.Destroy).Middleware(middleware.Owner)
}
