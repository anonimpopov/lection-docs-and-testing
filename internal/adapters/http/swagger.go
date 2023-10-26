package http

// Dependencies:
// go get -u -t github.com/swaggo/swag/cmd/swag

// @Title Example for lecture
// @Version 1.0.0
// @host      localhost:8080
// @BasePath /auth/v1
// @Schemes http
// @securityDefinitions.basic Auth
// @authorizationurl /validate
// @name token
// @description Signed token protects our admin endpoints
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
