package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"news-2/pkg/models"
	"news-2/pkg/utils"

	"github.com/gorilla/mux"
)

// CreateNews - Function to create a new news
func CreateNews(w http.ResponseWriter, r *http.Request) {
	CreateNews := &models.News{}
	utils.ParseBody(r, CreateNews)
	if CreateNews.Author == "" || CreateNews.Body == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		b := CreateNews.CreateNews()
		res, _ := json.Marshal(b)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// GetNews - Function to get all the news
func GetNews(w http.ResponseWriter, r *http.Request) {
	newNews := models.GetAllNews()
	res, _ := json.Marshal(newNews)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetNewsByID - Function to get a news by ID
func GetNewsByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	newsID := vars["newsId"]
	ID, err := strconv.ParseInt(newsID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	newsDetails, _ := models.GetNewsByID(ID)
	res, _ := json.Marshal(newsDetails)
	println(newsDetails.ID)
	if newsDetails.ID != 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

// UpdateNews - Function to update a news by ID
func UpdateNews(w http.ResponseWriter, r *http.Request) {
	var updateNews = &models.News{}
	utils.ParseBody(r, updateNews)
	if updateNews.Author == "" || updateNews.Body == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		vars := mux.Vars(r)
		newsID := vars["newsId"]
		ID, err := strconv.ParseInt(newsID, 0, 0)
		if err != nil {
			fmt.Println("Error while parsing")
		}
		newsDetails, db := models.GetNewsByID(ID)
		if updateNews.Author != "" {
			newsDetails.Author = updateNews.Author
		}
		if updateNews.Body != "" {
			newsDetails.Body = updateNews.Body
		}
		db.Save(&newsDetails)
		res, _ := json.Marshal(newsDetails)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// DeleteNews - Function to delete news by ID
func DeleteNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	newsID := vars["newsId"]
	ID, err := strconv.ParseInt(newsID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	news := models.DeleteNews(ID)
	res, _ := json.Marshal(news)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
