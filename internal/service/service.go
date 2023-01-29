package service

import (
	"context"
	"github.com/leon37/blog-server/global"
	"github.com/leon37/blog-server/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
