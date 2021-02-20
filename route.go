package main

import (
	"encoding/json"
	"net/http"
	"server/entity"
	"server/repository"
	"math/rand"



)

var (repo repository.PostRespository = repository.NewPostRespository())

// func init(){
// 	posts= []Post {Post{Id:1,Title: "Title 1",Text:"Text 1"}}
// }

func getPosts(resp http.ResponseWriter, req *http.Request){

	resp.Header().Set("Content-type","application/json")
	posts, err := repo.FindAll()
	if err !=nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error : "Error getting the post"}`))
	} 
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
	// resp.Write(posts)

}

func addPost(resp http.ResponseWriter, req *http.Request){
	resp.Header().Set("Content-type","application/json")

	var post  entity.Post
	err:=json.NewDecoder(req.Body).Decode(&post)
	if err!=nil{
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"{"error :"Error unmarchalling the request "}`))
		return
	}
	post.Id = rand.Int()
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	resp.Write(result)
}