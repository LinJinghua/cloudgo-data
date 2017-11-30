package entitiesorm

import (
	"log"
	"time"
)

const userInfoID = "user_id"

// UserInfo .
type UserInfo struct {
	UID        int64  `xorm:"pk autoincr" 'user_id'"`
	UserName   string `xorm:"notnull 'user_name'"`
	DepartName string `xorm:"'depart_name'"`
	CreateAt   *time.Time `xorm:"created 'create_time'"`
}

func init()  {
	// 同步结构体与数据表
	if err := orm.Sync(new(UserInfo)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	// if u.CreateAt == nil {
	// 	t := time.Now()
	// 	u.CreateAt = &t
	// }
	return &u
}
