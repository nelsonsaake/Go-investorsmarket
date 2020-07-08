package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"projects/investorsmarket/models"
	"time"
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
	db.AutoMigrate(&models.User{})
}

func GetUserA(email, password string) (user models.User, err error) {

	db.Where(&models.User{Email: email, Password: password}).First(&user)

	if user.Email != email {
		err = fmt.Errorf("Expecting %s, and got %s", email, user.Email)
		return
	}

	if user.Password != password {
		err = fmt.Errorf("Expecting %s, and got %s", password, user.Password)
		return
	}

	return
}

func CreateUser(email, password string) (user models.User, err error) {

	user = models.User{
		Email:    email,
		Password: password,
	}
	db.Create(&user)

	// new record means have no id
	// new record means not created
	if db.NewRecord(user) {
		err = fmt.Errorf("Failed to create user: %v", user)
		return
	}

	return
}

func GetAllUsers()(users []models.User){
	
	db.Find(&users)
	return
}

func GetUser(id uint64) (user models.User, err error) {

	err = db.First(&user, id).Error
	
	if err != nil {
		err = fmt.Errorf("Failed to get user: %v", err)
		return
	}
	
	if user.ID != uint(id) {
		err = fmt.Errorf("Failed to get user, id: %d", id)
		return
	}
	
	return
}

func UpdateUser(id uint64, email, picture, firstName, surname, dateOfBirth, gender, phoneNumber, nationality, occupation, address, country, region, city, accName, accNumber, accBankName, nkSurname, nkFirstName, nkRelationship, nkEmail, nkPhoneNumber, nkAddress string) (user models.User, err error) {

	err = db.First(&user, id).Error
	
	if err != nil {
		err = fmt.Errorf("Failed to find user: %v", err)
		return
	}
	
	if user.ID != uint(id) {
		err = fmt.Errorf("Failed to get user: %d", id)
		return
	}
	
	user.Email = email
	user.Picture = picture
	user.FirstName = firstName
	user.Surname = surname
	user.Gender = gender
	user.PhoneNumber = phoneNumber
	user.Nationality = nationality
	user.Occupation = occupation
	user.Address = address
	user.Country = country
	user.Region = region
	user.City = city
	user.AccName = accName
	user.AccNumber = accNumber
	user.AccBankName = accBankName
	user.NkSurname = nkSurname
	user.NkFirstName = nkFirstName
	user.NkRelationship = nkRelationship
	user.NkEmail = nkEmail
	user.NkPhoneNumber = nkPhoneNumber
	user.NkAddress = nkAddress
	user.DateOfBirth = dateOfBirth
	
	err = db.Save(&user).Error
	if err != nil {
		err = fmt.Errorf("failed to update user with id = %d", id)
		return
	}
	
	return
}