package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JohnnyOhms/BlogPost-api/pkg/models"
	"github.com/JohnnyOhms/BlogPost-api/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	blog := models.GetBlogs()
	res, _ := json.Marshal(blog)
	w.Header().Set("Content-Type:", "pkglicaton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	ID := param["id"]
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		panic("failed to convert to string")
	}
	blog, _ := models.GetBlog(id)
	res, _ := json.Marshal(blog)
	w.Header().Set("Content-Type:", "pkglicaton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	b := &models.Blogs{}
	utils.ParseBody(r, b)
	blog := b.CreateBlog()
	res, _ := json.Marshal(blog)
	w.Header().Set("Content-Type:", "pkglicaton/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdatedBlog(w http.ResponseWriter, r *http.Request) {
	b := &models.Blogs{}
	utils.ParseBody(r, b)
	param := mux.Vars(r)
	ID := param["id"]
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		panic("failed to convert to int64")
	}
	blog := b.UpdateBlog(id,updateData)


}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {

}
