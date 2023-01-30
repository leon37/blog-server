package service

import (
	"github.com/leon37/blog-server/internal/model"
	"github.com/leon37/blog-server/internal/protocol"
	"github.com/leon37/blog-server/pkg/app"
)

func (svc *Service) CountArticles(param *protocol.CountArticleRequest) (int, error) {
	return svc.dao.CountArticle(param.Title, param.State)
}

func (svc *Service) GetSingleArticle(param *protocol.ArticleSingleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}

func (svc *Service) GetArticleList(param *protocol.ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateArticle(param *protocol.CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.CreatedBy)
}

func (svc *Service) UpdateArticle(param *protocol.UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteArticle(param *protocol.DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}
