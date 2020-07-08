package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"not null"`
	Picture        string
	FirstName      string
	Surname        string
	DateOfBirth    string
	Gender         string
	PhoneNumber    string
	Nationality    string
	Occupation     string
	Address        string
	Country        string
	Region         string
	City           string
	AccName        string
	AccNumber      string
	AccBankName    string
	NkSurname      string
	NkFirstName    string
	NkRelationship string
	NkEmail        string
	NkPhoneNumber  string
	NkAddress      string
}

type Opportunity struct {
	gorm.Model
	Name        string  `gorm:"unique"`
	Amount      float64 `gorm:"not null"`
	Industry    string
	Description string
	UserId      uint64 `gorm:"not null"`
	Picture     string
	Returns     float64
	Duration    float32
	Location    string
}

type Investment struct {
	gorm.Model
	UserId        uint64  `gorm:"not null"`
	OpportunityId uint64  `gorm:"not null"`
	AmountBought  float64 `gorm:"not null"`
}

type Chat struct {
	gorm.Model
	SenderId     uint64 `gorm:"not null"`
	ReceiverId   uint64 `gorm:"not null"`
	Message      string
	SentAt       time.Time `gorm:"not null"`
	DelieveredAt time.Time
	ReadAt       time.Time
}

type ChangePassword struct {
	gorm.Model
	Active bool
	UserId uint64
	Code   string `gorm:"unique;not null"`
}

type Post struct {
	gorm.Model
	Picture       string
	Description   string
	OpportunityId uint64
}

type Auth struct {
	gorm.Model
	UserId uint64
	Token  string `gorm:"unique;not null"`
	Active bool
	Role   string `gorm:"not null"`
}
