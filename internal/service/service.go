package service

import (
	"context"
	"log"
	"strings"

	pb "github.com/hi20160616/fetchnews-api/proto/v1"
	"github.com/hi20160616/ms-kabar/internal/fetcher"
)

type Server struct {
	pb.UnimplementedFetchNewsServer
}

func (s *Server) ListArticles(ctx context.Context, in *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in ListArticles:\n%v\n", r)
		}
	}()
	// log.Printf("Received: %v", in.GetPageSize())
	a := fetcher.NewArticle()
	as, err := a.List()
	if err != nil {
		return nil, err
	}
	resp := &pb.ListArticlesResponse{}
	for _, a := range as {
		resp.Articles = append(resp.Articles, &pb.Article{
			Id:            a.Id,
			Title:         a.Title,
			Content:       a.Content,
			WebsiteId:     a.WebsiteId,
			WebsiteTitle:  a.WebsiteTitle,
			WebsiteDomain: a.WebsiteDomain,
			UpdateTime:    a.UpdateTime,
		})
	}
	return resp, nil
}

func (s *Server) GetArticle(ctx context.Context, in *pb.GetArticleRequest) (*pb.Article, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in GetArticle: %s\n%v\n", in.Id, r)
		}
	}()
	// log.Printf("Id: %v", in.Id)
	// Got article via json reading
	a := fetcher.NewArticle()
	a, err := a.Get(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Article{
		Id:            a.Id,
		Title:         a.Title,
		Content:       a.Content,
		WebsiteId:     a.WebsiteId,
		WebsiteTitle:  a.WebsiteTitle,
		WebsiteDomain: a.WebsiteDomain,
		UpdateTime:    a.UpdateTime,
	}, nil
}

func (s *Server) SearchArticles(ctx context.Context, in *pb.SearchArticlesRequest) (*pb.SearchArticlesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in SearchArticles:\n%v\n", r)
		}
	}()
	a := fetcher.NewArticle()
	as, err := a.Search(strings.Split(in.Keyword, ",")...)
	if err != nil {
		return nil, err
	}
	as2 := []*pb.Article{}
	for _, a := range as {
		as2 = append(as2, &pb.Article{
			Id:            a.Id,
			Title:         a.Title,
			Content:       a.Content,
			WebsiteId:     a.WebsiteId,
			WebsiteTitle:  a.WebsiteTitle,
			WebsiteDomain: a.WebsiteDomain,
			UpdateTime:    a.UpdateTime,
		})
	}
	return &pb.SearchArticlesResponse{Articles: as2}, nil
}
