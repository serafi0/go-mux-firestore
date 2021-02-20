package repository

import (
	"context"
	"server/entity"
	"cloud.google.com/go/firestore"
	"log"

)

type PostRespository interface {
	Save (post *entity.Post) (*entity.Post,error)
	FindAll() ([]entity.Post,error)
}

type repo struct {}

func NewPostRespository() PostRespository{
	return &repo{}
}

const (
	projectId string  ="go-api-3ecca"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post,error){
	ctx := context.Background();
	client,err :=firestore.NewClient(ctx,projectId)
	if err !=nil{
		log.Fatalf("couldn't connect to firestore : %v",err)
	}

	defer client.Close()

	_,_, err =client.Collection(collectionName).Add(ctx, map[string] interface{}{

		"ID" : post.Id,
		"Title" :post.Title,
		"Text":post.Text,


	})

	if err!=nil{
		log.Fatal("Failed to add a new Post : %v",err)
		return nil , err
	}
	return post,nil
}

func(*repo) FindAll() ([]entity.Post,error){
		ctx := context.Background();
	client,err :=firestore.NewClient(ctx,projectId)
	if err !=nil{
		log.Fatalf("couldn't connect to firestore : %v",err)
	}
	defer client.Close()

	var posts[ ]entity.Post
	iterator :=client.Collection(collectionName).Documents(ctx)
	for {
		doc, err:=iterator.Next()
		if err!=nil{
			log.Fatalf("Failed to iterate a list of post %v", err)

		}
		post :=entity.Post{
			Id : doc.Data()["ID"].(int),
			Title:doc.Data()["Title"].(string),
			Text:doc.Data()["Text"].(string),

		}
		posts = append(posts,post)
	}
	return posts, nil
}
