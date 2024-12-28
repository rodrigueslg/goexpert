package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/rodrigueslg/fc-goexpert/api/configs"
	_ "github.com/rodrigueslg/fc-goexpert/api/docs"
	"github.com/rodrigueslg/fc-goexpert/api/internal/entity"
	"github.com/rodrigueslg/fc-goexpert/api/internal/infra/database"
	"github.com/rodrigueslg/fc-goexpert/api/internal/infra/webserver/handlers"
)

// @title API GoExpert
// @version 1.0
// @description API GoExpert
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://swagger.io/support
// @contact.email support@email.com

// @license.name MIT
// @license.url http://mitlicense.org

// @host localhost:8080
// @BasePath /
// @SecurityDefinitions.api_key ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUserDB(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.JWT))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))
	//r.Use(pkgMiddleware.RequestLogger)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.JWT))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/auth", userHandler.Auth)
	})

	docsUrl := fmt.Sprintf("http://localhost:%s/docs/doc.json", configs.WebServerPort)
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsUrl)))

	http.ListenAndServe(":"+configs.WebServerPort, r)
}
