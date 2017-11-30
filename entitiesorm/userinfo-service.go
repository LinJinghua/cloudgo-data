package entitiesorm

import (
	"fmt"
)

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

const mode = `SET GLOBAL sql_mode=''`
// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	if _, err := orm.Query(mode); err != nil {
		panic(err)
	}
	if affected, err := orm.Insert(u); err != nil {
		fmt.Println("Save Error:", affected, err)
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	users := make([]UserInfo, 0)
	if err := orm.Find(&users); err != nil {
		fmt.Println("FindAll Error:", err)
	}
	return users
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	fmt.Println("FindByID Enter:")
	// var u UserInfo
	// has, err := orm.Where(userInfoID + " = ?", id).Get(&u);
	u := &UserInfo{UID: int64(id)}
	has, err := orm.Get(u)
	if err != nil {
		fmt.Println("FindByID Error:", err)
	}
	if has {
		return u
		fmt.Println("FindByID:", u)
	}
	return nil
}
