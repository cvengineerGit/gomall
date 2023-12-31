package adminmodel

import "time"

type User struct {
	ID        int       `xorm:"not null pk autoincr BIGINT(20) id"`
	LoginName string    `xorm:"not null unique VARCHAR(50) login_name"`
	NickName  string    `xorm:"null VARCHAR(50) nick_name"`
	Passwd    string    `xorm:"not null VARCHAR(255) passwd"`
	Email     string    `xorm:"null unique VARCHAR(100) email"`
	Locked    int       `xorm:"not null default 0 TINYINT locked"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt time.Time `xorm:"deleted TIMESTAMP deleted_at"`
}

type UserInfo struct {
	ID        int       `json:"id"`
	LoginName string    `json:"login_name"`
	NickName  string    `json:"nick_name"`
	Email     string    `json:"email"`
	Locked    int       `json:"locked"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRegisterRequest struct {
	LoginName string `json:"login_name" validate:"required|min_len:5|max_len:10" message:"required:login_name 登录名不能为空|min_len:login_name 登录名应为5-10个字符|max_len:login_name 登录名应为5-10个字符"`
	Passwd    string `json:"passwd" validate:"required|min_len:8|max_len:20" message:"required:passwd 密码不能为空|min_len:passwd 密码应为8-20个字符|max_len:passwd 密码应为8-20个字符"`
}

type  UserLoginRequest struct {
	LoginName string `validate:"required" message:"required:账号不能为空" json:"login_name"`
	Passwd    string `validate:"required" message:"required:密码不能为空" json:"passwd"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}

type UpdateUserInfoRequest struct {
	ID        int    `json:"id" validate:"required" message:"required:用户ID不能为空"`
	LoginName string `json:"login_name" validate:"required|min_len:5|max_len:10" message:"required:login_name 登录名不能为空|min_len:login_name 登录名应为5-10个字符|max_len:login_name 登录名应为5-10个字符"`
	NickName  string `json:"nick_name" validate:"required" message:"required:昵称不能为空"`
	Email     string `json:"email" validate:"required|email" message:"required:邮箱不能为空|email:邮箱格式错误"`
}

func (u User) TableName() string {
	return "t_admin_users"
}
