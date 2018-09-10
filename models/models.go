package models

import (
    "errors"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3" // 驱动
    "os"
    "path"
    "time"
)

const (
    _DB_NAME         = "data/beeblog.db"
    _SQL_ITE3_DRIVER = "sqlite3"
)

// 用于分类
type Category struct {
    Id              int64  // 默认int64的Id字段被orm视为主键
    Title           string // 默认在orm里长度为255
    Created         time.Time `orm:"index"`
    Views           int64     `orm:"index"` // 浏览次数
    TopicTime       time.Time `orm:"index"` // 最后一篇文章的创建时间
    TopicCount      int64                   // 文章数
    TopicLastUserId int64     `orm:"index"` // 最后一个操作的用户id
}

// 文章结构
type Topic struct {
    Id              int64                        // orm 主键
    Uid             int64                        // 用户id
    Title           string                       // 标题
    Content         string    `orm:"size(5000)"` // 设置大小为5000字节
    Attachment      string                       // 附件
    Created         time.Time `orm:"index"`
    Updated         time.Time `orm:"index"`
    Views           int64     `orm:"index"` // 浏览次数
    Auchor          string                  // 作者
    ReplyTime       time.Time               // 回复时间
    ReplyCount      int64                   // 回复次数
    ReplyLastUserId int64                   // 最后回复的主播id
}

func RegisterDB() error {
    _, err := os.Stat(_DB_NAME)
    if err != nil {
        if os.IsNotExist(err) {
            // 不存在则创建
            os.MkdirAll(path.Dir(_DB_NAME), os.ModeDir)
            os.Create(_DB_NAME)
        } else {
            // 非不存在的错误
            return errors.New("check db file error, err: " + err.Error())
        }
    }

    orm.RegisterModel(new(Category), new(Topic))
    orm.RegisterDriver(_SQL_ITE3_DRIVER, orm.DRSqlite) // 引入sqlite后初始化注册，这里可以不用加
    orm.RegisterDataBase("default", _SQL_ITE3_DRIVER, _DB_NAME, 10)

    return nil
}
