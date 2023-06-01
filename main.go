package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	// Configuração do cliente MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping no servidor MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Selecionando o banco de dados e a coleção
	collection := client.Database("mydatabase").Collection("posts")

	// Fazendo a requisição HTTP para obter os posts
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Decodificando a resposta JSON em uma slice de objetos Post
	var posts []Post
	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		log.Fatal(err)
	}

	// Iterar sobre os objetos Post e inserir no banco de dados
	for _, post := range posts {
		_, err := collection.InsertOne(context.Background(), post)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Post inserido com sucesso!")
	}
}
