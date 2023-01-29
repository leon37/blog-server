package service

import (
	"github.com/leon37/blog-server/internal/model"
	"github.com/leon37/blog-server/internal/protocol"
	"github.com/leon37/blog-server/pkg/app"
)

func (svc *Service) CountTag(param *protocol.CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *protocol.TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *protocol.CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *protocol.UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *protocol.DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
