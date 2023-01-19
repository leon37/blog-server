package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/leon37/blog-server/global"
	"github.com/leon37/blog-server/pkg/setting"
)

type Model struct {
	ID         int32  `json:"id"`          // id
	CreatedOn  int32  `json:"created_on"`  // 创建时间
	CreatedBy  string `json:"created_by"`  // 创建人
	ModifiedOn int32  `json:"modified_on"` // 修改时间
	ModifiedBy string `json:"modified_by"` // 修改人
	DeletedOn  int32  `json:"deleted_on"`  // 删除时间
	IsDel      int8   `json:"is_del"`      // 是否删除 0为未删除、1为已删除
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
