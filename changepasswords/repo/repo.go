package repo

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"projects/investorsmarket/models"
)

var db *gorm.DB

func init() {
	var err error
	// ? set password
	// set password as env 
	// removed password before commit
	db, err = gorm.Open("postgres", "user=postgres dbname=investorsmarket password= sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err)
	}
	db.AutoMigrate(&models.ChangePassword{})
	db.Model(&models.ChangePassword{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
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

func CreateChangePassword(userId uint64, code string) (cp models.ChangePassword, err error) {

	// check if userId is valid
	if !isUserValid(userId) {
		err = errors.New("invalid user id")
		return
	}
		
	// deactivate all old resquest for change password by the same user
	err = db.Table("change_passwords").Where("user_id = ?", fmt.Sprint(userId)).Updates(map[string]string{"active": "false"}).Error
	if err != nil {
		err = fmt.Errorf("error deactivating all old request to change password requested by the same user:\n %v\n", err)
		return
	}

	//
	cp = models.ChangePassword{
		Active: true,
		UserId:   userId,
		Code: code,
	}
	db.Create(&cp)
	if db.NewRecord(cp) {
		err = fmt.Errorf("failed to create a new change password record: %v", err)
		return
	}

	return
}

func GetChangePassword(userId uint64, code string) (cp models.ChangePassword, err error) {

	err = db.Where(models.ChangePassword{UserId: userId, Code: code}).First(&cp).Error
	if err != nil {
		err = fmt.Errorf("failed to get change password record: %v", err)
		return
	}

	return
}

func GetAllChangePasswords()(cps []models.ChangePassword, err error) {

	err = db.Find(&cps).Error
	if err != nil {
		err = fmt.Errorf("failed to get all change password records: %v", err)
		return
	}

	return
}