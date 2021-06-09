package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/hi20160616/fetchnews-api/proto/v1"
	"github.com/hi20160616/ms-kabar/configs"
	"google.golang.org/grpc"
)

var address = "localhost" + configs.Data.MS.Addr

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFetchNewsClient(conn)

	// Contact the server and print out its response.
	articleId := "f8a3801262aea95b78403158c0b13758"
	if len(os.Args) > 1 {
		articleId = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ListArticles(ctx, &pb.ListArticlesRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetArticles())
	// r, err = c.GetArticle(ctx, &pb.GetArticleRequest{Id: name})
	article, err := c.GetArticle(ctx, &pb.GetArticleRequest{Id: articleId})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", article.Title)
	articles, err := c.SearchArticles(ctx, &pb.SearchArticlesRequest{Keyword: "戴安娜, 民主党"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("SearchArticles: %d", len(articles.Articles))
	for _, a := range articles.Articles {
		log.Println(a.Title)
	}
}
