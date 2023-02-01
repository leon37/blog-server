package service

import (
	"errors"
	"github.com/leon37/blog-server/internal/protocol"
)

func (svc *Service) CheckAuth(param *protocol.AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
