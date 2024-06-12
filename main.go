package main

import (
	"context"
	"log"
	"net/http"

	"bhajann/pkg/config"
	"bhajann/pkg/controller"
	"bhajann/pkg/repository"
	"bhajann/pkg/services"

	"github.com/gorilla/mux"
)

func InitializeHomeController() *controller.HomeController {
	categoryModelService := repository.CategoryModelService{}
	categoryCollectionClient := categoryModelService.GetCollection()
	categoryServiceModel := repository.NewCategoryModelService(categoryCollectionClient)
	categoryService := services.NewCategoryService(categoryServiceModel)

	contentModelService := repository.ContentModelService{}
	contentCollectionClient := contentModelService.GetCollection()
	contentServiceModel := repository.NewContentModelService(contentCollectionClient)
	contentService := services.NewContentService(contentServiceModel)
	return controller.NewHomeController(
		categoryService,
		contentService,
	)
}

func main() {
	// Connect to MongoDB
	config.ConnectMongo("mongodb://localhost:27017")

	// Ensure MongoDB connection is closed on exit
	defer func() {
		if err := config.Client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	router := mux.NewRouter()

	homeController := InitializeHomeController()
	router.HandleFunc("/categories/list", homeController.ListCategories).Methods("GET")
	router.HandleFunc("/categories", homeController.AddCategories).Methods("POST")

	router.HandleFunc("/content/list", homeController.ListContent).Methods("GET")
	router.HandleFunc("/content", homeController.AddContent).Methods("POST")

	// router.HandleFunc("/content/list", homeController.GetContentList).Methods("GET")
	// router.HandleFunc("/content/{id}/meta", homeController.AddContentList).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
