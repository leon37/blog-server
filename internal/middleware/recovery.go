package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leon37/blog-server/global"
	"github.com/leon37/blog-server/pkg/app"
	"github.com/leon37/blog-server/pkg/email"
	"github.com/leon37/blog-server/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != any(nil) {
				s := "panic recover err: %v"
				global.Logger.WithCallerFrames().Errorf(s, err)

				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf("mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToResponse(errcode.ServerError)
				c.Abort()
			}
		}()

		c.Next()
	}
}
