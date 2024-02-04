package main

import (
	"database/sql"
	"fmt"
	"net/http"

	datab "github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/database"
	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/service"
	"github.com/Math-O5/Go/projects/imersao-full-cycle/goapi/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	print("Hello API WOLRD")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")

	if err != nil {
		print("ERROR DATABASE")
		panic(err.Error())
	}
	print("Hello API WOLRD")

	defer db.Close()

	categoryDB := datab.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := datab.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	WebCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/category/{id}", WebCategoryHandler.GetCategory)
	c.Get("/category", WebCategoryHandler.GetCategories)
	c.Post("/category", WebCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
