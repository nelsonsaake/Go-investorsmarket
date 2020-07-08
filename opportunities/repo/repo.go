package repo

import (
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
	db.AutoMigrate(&models.Opportunity{})
	db.Model(&models.Opportunity{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

func CreateOpportunity(name string, amount float64, industry string, description string, userId uint64, picture string, returns float64, duration float32, location string) (opp models.Opportunity, err error) {

	// make sure the user exist
	type tempUser struct {
		Id uint64
	}

	var user tempUser
	err = db.Table("users").Select("id").Where("id = ?", fmt.Sprint(userId)).Scan(&user).Error
	if err != nil {
		err = fmt.Errorf("Error creating a new opportunity, a problem with the user id: %v", err)
		return
	}

	if user.Id != userId {
		err = fmt.Errorf("Error creating a new opportunity, a problem with the user id: %v", err)
		return
	}

	// make sure the no other opportunity with the same name
	searchOpportunity := models.Opportunity{}
	err = db.Where("name = ?", name).First(&searchOpportunity).Error
	if searchOpportunity.Name == name {
		err = fmt.Errorf("Error creating new opportunity, name already taken.")
		return
	}

	opp = models.Opportunity{
		Name:        name,
		Amount:      amount,
		Industry:    industry,
		Description: description,
		UserId:      userId,
		Picture:     picture,
		Returns:     returns,
		Duration:    duration,
		Location:    location,
	}

	err = db.Create(&opp).Error
	if err != nil {
		err = fmt.Errorf("Error creating a new opportunity: %v", err)
		return
	}

	if db.NewRecord(opp) {
		err = fmt.Errorf("Failed to create new opportunity!")
		return
	}

	return
}

func GetAllOpportunities() (opportunities []models.Opportunity, err error) {

	err = db.Find(&opportunities).Error
	if err != nil {
		err = fmt.Errorf("Error getting all opportunities: %v", err)
	}
	return
}

func GetOpportunityCreators() (creators []models.User, err error) {

	var creator_ids []uint64
	err = db.Raw("SELECT DISTINCT user_id from opportunities").Pluck("user_id", &creator_ids).Error
	if err != nil {
		err = fmt.Errorf("Error getting creator_ids from opportunities: %v", err)
		return
	}

	err = db.Where("id IN(?)", creator_ids).Find(&creators).Error
	if err != nil {
		err = fmt.Errorf("Error getting creators from users table using creator_ids: %v", err)
		return
	}
	return
}

func GetCreatorHistory(id uint64) (opportunities []models.Opportunity, err error) {

	err = db.Where(&models.Opportunity{UserId: id}).Find(&opportunities).Error
	if err != nil {
		err = fmt.Errorf("Error getting creators history from repo: %v", err)
		return
	}
	return
}

func GetOpportunity(id uint64) (opportunity models.Opportunity, err error) {

	err = db.Where("id=?",fmt.Sprint(id)).First(&opportunity).Error
	if err != nil {
		err = fmt.Errorf("Error getting opportunity from repo: %v", err)
		return
	}
	return
}