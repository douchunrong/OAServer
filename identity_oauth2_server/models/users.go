package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"identity_openid_server/models/base"
	"time"
)

type User struct {
	ID                  int64     `bson:"_id"          				json:"_id"`
	Name                string    `bson:"name"        				json:"name"`
	Email               string    `bson:"email"       				json:"email"`
	Password            string    `bson:"password"     				json:"password"`
	ResetPasswordToken  string    `bson:"reset_password_token"  	json:"reset_password_token"`
	ResetPasswordSentAt time.Time `bson:"reset_assword_sent_at"     json:"reset_assword_sent_at"`
	RememberCreatedAt   time.Time `bson:"remember_created_at"     	json:"remember_created_at"`
	SignInCount         int64     `bson:"sign_in_count"     		json:"sign_in_count"`
	CurrentSignInAt     time.Time `bson:"current_sign_in_at"     	json:"current_sign_in_at"`
	LastSignInAt        time.Time `bson:"last_sign_in_at"     		json:"last_sign_in_at"`
	CurrentSignInIp     string    `bson:"current_sign_in_ip"     	json:"current_sign_in_ip"`
	LastSignInIp        string    `bson:"last_sign_in_ip"     		json:"last_sign_in_ip"`
	IdentityUrl         string    `bson:"identity_url"     			json:"identity_url"`
	CreatedAt           time.Time `bson:"created_at"     			json:"created_at"`
	UpdatedAt           time.Time `bson:"updated_at"     			json:"updated_at"`
}

func NewUser(name, email, password string) *User {
	user := User{
		Name:        name,
		Email:       email,
		Password:    password,
		SignInCount: 1,
		IdentityUrl: beego.AppConfig.String("rootPath") + name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return &user
}

func FindUserByName(username string) *User {
	beego.Debug("用户名:" + username)
	user := User{}

	mConn := base.Conn()
	defer mConn.Close()

	tb := mConn.DB("").C("users")
	err := tb.Find(bson.M{"name": username}).One(user)

	if err != nil {
		return nil
	}

	return &user
}

func (u *User) InsertUser() (code int, err error) {
	mConn := base.Conn()
	defer mConn.Close()

	tb := mConn.DB("").C("users")

	err = tb.Insert(u)
	if err != nil {
		if err == mgo.ErrNotFound {
			code = 100
		} else {
			code = -1
		}
	} else {
		code = 0
	}
	return
}
