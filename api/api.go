package api

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"online-print/api/controllers"
	"online-print/api/database"
	"online-print/api/models"
	"online-print/api/repository"
	"online-print/api/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	port        = flag.Int("p", 5000, "set port")
	resetTables = flag.Bool("rt", false, "reset tables")
)

func Run() {

	flag.Parse()

	db := database.Connect()
	if db != nil {
		defer db.Close()
	}

	fmt.Println("Database connected...")

	if *port != 5000 && *resetTables {
		createSuperTestTables()
	}

	productsRepository := repository.NewProductsRepository(db)
	productsorderRepository := repository.NewProductsorderRepository(db)
	usersRepository := repository.NewUsersRepository(db)
	locationsRepository := repository.NewLocationsRepository(db)
	locationusersRepository := repository.NewLocationusersRepository(db)

	productsController := controllers.NewProductsRepository(productsRepository)
	productsorderController := controllers.NewProductsorderRepository(productsorderRepository)
	usersController := controllers.NewUsersRepository(usersRepository)
	locationsController := controllers.NewLocationsRepository(locationsRepository)
	locationusersController := controllers.NewLocationusersRepository(locationusersRepository)

	productRoutes := routes.NewProductRoutes(productsController)
	productorderRoutes := routes.NewProductorderRoutes(productsorderController)
	userRoutes := routes.NewUserRoutes(usersController)
	locationRoutes := routes.NewLocationRoutes(locationsController)
	locationuserRoutes := routes.NewLocationuserRoutes(locationusersController)

	router := mux.NewRouter().StrictSlash(true)

	routes.Install(router, productRoutes, productorderRoutes,userRoutes, locationRoutes, locationuserRoutes)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Request", "Location", "Entity", "Accept"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("API Listening on", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.CORS(headers, methods, origins)(router)))
}

func createSuperTestTables() {
	db := database.Connect()
	if db != nil {
		defer db.Close()
	}

	tx := db.Begin()
	err := tx.Debug().DropTableIfExists(&models.Productorder{}, &models.Product{}).Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Debug().CreateTable(&models.Product{}).Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Debug().CreateTable(&models.Productorder{}).Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Debug().Model(&models.Productorder{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE").Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
}
