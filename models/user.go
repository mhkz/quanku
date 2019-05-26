package models

import "time"

type User struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `sql:"index" json:"deletedAt"`
	ActivatedAt  *time.Time `json:"activatedAt"`
	Name         string     `json:"name"`
	Pass         string     `json:"-"`
	Email        string     `json:"-"`
	Sex          uint       `json:"sex"`
	Location     string     `json:"location"`
	Introduce    string     `json:"introduce"`
	Phone        string     `json:"-"`
	Score        uint       `json:"score"`
	ArticleCount uint       `json:"articleCount"`
	CommentCount uint       `json:"commentCount"`
	CollectCount uint       `json:"collectCount"`
	Signature    string     `json:"signature"` //个人签名
	Role         int        `json:"role"`      //角色
	AvatarURL    string     `json:"avatarURL"` //头像
	CoverURL     string     `json:"coverURL"`  //个人主页背景图片URL
	Status       int        `json:"status"`
	Schools      []School   `json:"schools"` //教育经历
	Careers      []Career   `json:"careers"` //职业经历
}