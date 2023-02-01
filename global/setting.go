package global

import (
	"github.com/leon37/blog-server/pkg/logger"
	"github.com/leon37/blog-server/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	JWTSetting      *setting.JWTSettings
	EmailSetting    *setting.EmailSettings
	Logger          *logger.Logger
)
