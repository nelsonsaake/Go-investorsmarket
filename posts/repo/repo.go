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
	db.AutoMigrate(&models.Post{})
	db.Model(&models.Post{}).AddForeignKey("opportunity_id", "opportunities(id)", "RESTRICT", "RESTRICT")
}

func isValidOpportunity(id uint64) bool {

	type temp struct {
		Id uint64
	}

	var opportunity temp
	err := db.Table("opportunities").Select("id").Where("id=?", fmt.Sprint(id)).Scan(&opportunity).Error
	if err != nil {
		fmt.Printf("Error: opportunity record with %d not found: %v \n", id, err)
		return false
	}

	if opportunity.Id != id {
		fmt.Printf("Error: finding opportunity with id %d", id)
		return false
	}

	return true
}

func CreatePost(picture, description string, opportunityId uint64)(post models.Post, err error) {
	
	if !isValidOpportunity(opportunityId) {
		err = fmt.Errorf("invalid opportunity id = %v", opportunityId)
		return
	}
	
	post = models.Post{
		Picture: picture,
		Description: description,
		OpportunityId: opportunityId,
	}
	db.Create(&post)
	if db.NewRecord(post) {
		err = errors.New("failed to create new opportunity record")
	}
	return
}

func GetOpportunityPosts(id uint64)(posts []models.Post, err error) {
	
	err = db.Where(&models.Post{OpportunityId: id}).Find(&posts).Error
	if err != nil {
		err = fmt.Errorf("something went wrong reading from the database: %v\n", err)
	}
	return
}