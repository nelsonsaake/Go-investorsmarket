package repo

import (
	"errors"
	"fmt"
	"projects/investorsmarket/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func init() {
	var err error
	// ? set password
	// set password as env 
	// removed password before commit
	db, err = gorm.Open("postgres", "user=postgres dbname=investorsmarket password= sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to the database!")
		panic(err)
	}
	db.AutoMigrate(&models.Auth{})
	db.Model(&models.Auth{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

func getUserId(email, password string) (uint64, error) {

	type GetUserId struct {
		Id    uint64
		Email string
	}

	var user GetUserId
	err := db.Table("users").Select("id, email").Where("email=? and password=?", email, password).Scan(&user).Error
	if err != nil {
		return 0, err
	}

	if user.Id == 0 {
		return 0, errors.New("user id is zero")
	}

	if user.Email != email {
		return 0, errors.New("user email doesn't match")
	}

	return user.Id, nil
}

func checkToken(token string) error {

	if len(token) < 32 {
		return errors.New("token is too short")
	}
	return nil
}

func checkRole(role string) error {

	if !(role == "admin" || role == "user") {
		return fmt.Errorf("role is invalid: should be admin or user")
	}

	return nil
}

func CreateAuth(email, password, token, role string) (auth models.Auth, err error) {

	// get user Id
	userId, err := getUserId(email, password)
	if err != nil {
		err = fmt.Errorf("user doesn't exist: %v %s", err, "\n")
		return
	}

	// check/verify token
	err = checkToken(token)
	if err != nil {
		err = fmt.Errorf("invalid token: \nerr: %v \ntoken: ", err, token)
		return
	}

	err = db.Table("auths").Where("user_id=?", fmt.Sprint(userId)).Updates(map[string]string{"active": "false"}).Error
	if err != nil {
		err = fmt.Errorf("Failed to deactivate all old tokens: %v %s", err, "\n")
		return
	}

	// check/verify role
	err = checkRole(role)
	if err != nil {
		err = fmt.Errorf("invalid role: \nerr: %v \nrole: ", err, role)
		return
	}

	// create auth
	auth = models.Auth{
		UserId: userId,
		Token:  token,
		Active: true,
		Role:   role,
	}
	db.Create(&auth)
	if db.NewRecord(auth) {
		err = fmt.Errorf("failed to create auth, \nauth: %v\n", auth)
		return
	}

	return
}

func GetAuthGivenEP(email, password string) (auth models.Auth, err error) {

	// find user
	userId, err := getUserId(email, password)
	if err != nil {
		err = fmt.Errorf("user doesn't exist: %v %s", err, "\n")
		return
	}

	// get code
	err = db.Where(models.Auth{UserId: userId, Active: true}).First(&auth).Error
	if err != nil {
		err = fmt.Errorf("failed to get active auth token for user: %v", err)
		return
	}

	if !auth.Active {
		err = errors.New("no active token available for this user")
		return
	}

	return
}

func GetAuthGivenToken(token string) (auth models.Auth, err error) {

	// get code
	err = db.Where(models.Auth{Token: token, Active: true}).First(&auth).Error
	if err != nil {
		err = fmt.Errorf("failed to get active auth token for user: %v", err)
		return
	}

	if !auth.Active {
		err = errors.New("no active token available for this user")
		return
	}

	return
}

func isUserValid(id uint64) bool {

	type userTemp struct {
		Id uint64
	}

	var user userTemp
	err := db.Table("users").Select("id").Where("id=?", fmt.Sprint(id)).Scan(&user).Error
	if err != nil {
		fmt.Printf("Error finding user with id %d: %v \n", id, err)
		return false
	}

	if user.Id != id {
		fmt.Printf("Error finding user with id %d", id)
		return false
	}

	return true
}

func GetAuthGivenUserAndToken(userId uint64, token string) (auth models.Auth, err error) {
	// check if userId is valid
	if !isUserValid(userId) {
		err = errors.New("invalid user id")
		return
	}

	// get code
	err = db.Where(models.Auth{UserId: userId, Token: token, Active: true}).First(&auth).Error
	if err != nil {
		err = fmt.Errorf("failed to get active auth token for user: %v", err)
		return
	}

	if !auth.Active {
		err = errors.New("no active token available for this user")
		return
	}

	if auth.UserId != userId {
		err = errors.New("record not found")
		return
	}

	return
}
