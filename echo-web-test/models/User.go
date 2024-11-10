package models

type User struct {
	Userid             string `gorm:"column:user_id;primary_key" json:"userid"`
	Nickname           string `gorm:"column:nick_name" json:"nickname"`
	Userpassword       string `gorm:"column:user_password" json:"userpassword"`
	Email              string `gorm:"column:email" json:"email"`
	Sex                string `gorm:"column:sex" json:"sex"`
	Birthday           string `gorm:"column:birthday" json:"birthday"`
	Personintroduction string `gorm:"column:person_introduction" json:"personintroduction"`
	Avatar             string `gorm:"column:avatar" json:"avatar"`
	Createtime         string `gorm:"column:create_time" json:"createtime"`
	Lastlogintime      string `gorm:"column:last_login_time" json:"lastlogintime"`
	Lastloginip        string `gorm:"column:last_login_ip" json:"lastloginip"`
	Status             string `gorm:"column:status" json:"status"`
	Experience         int    `gorm:"column:experience" json:"experience"`
	Threme             int    `gorm:"column:threme" json:"threme"`
	Role               string `gorm:"column:role" json:"role"`
}

func (User) TableName() string {
	return "user_info"
}
