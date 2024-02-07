package main

import (
	"net/http"

	"github.com/joaomarcosbc/api-fc/configs"
	"github.com/joaomarcosbc/api-fc/internal/entity"
	"github.com/joaomarcosbc/api-fc/internal/infra/database"
	"github.com/joaomarcosbc/api-fc/internal/webservice/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if _, err := configs.LoadConfig("."); err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)
}
