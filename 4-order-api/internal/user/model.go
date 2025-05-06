package user

import "gorm.io/gorm"

type User struct {
	gorm.Model `gorm:"index"`
	Number     string `json:"number"`
	SessionID  string `json:"sessionId"`
	Code       string `json:"code"`
}
