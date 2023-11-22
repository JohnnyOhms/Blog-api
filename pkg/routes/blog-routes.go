package routes

import (
	"github.com/JohnnyOhms/BlogPost-api/pkg/controllers"
	"github.com/gorilla/mux"
)

func BlogRoutes(router *mux.Router) {
	router.HandleFunc("/getblogs", controllers.GetBlog).Methods("GET")
	router.HandleFunc("/create", controllers.CreateBlog).Methods("POST")
	router.HandleFunc("/getblog/{id}", controllers.GetBlog).Methods("GET")
	router.HandleFunc("/updatedblog/{id}", controllers.UpdatedBlog).Methods("PUT")
	router.HandleFunc("/deleteblog/{id}", controllers.DeleteBlog).Methods("DELETE")
}
