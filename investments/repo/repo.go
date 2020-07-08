package repo

import (
	"fmt"
	"errors"
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
	db.AutoMigrate(&models.Investment{})
	db.Model(&models.Investment{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Investment{}).AddForeignKey("opportunity_id", "opportunities(id)", "RESTRICT", "RESTRICT")
}


func isUserIdValid(userId uint64) bool{

	type userTemp struct{
		Id uint64
		Name string
	}
	var user userTemp
	
	err := db.Table("users").Select("id, name").Where("id = ?", fmt.Sprint(userId)).Scan(&user).Error
	if err != nil {
		fmt.Println("Error validating user id: ", err)
		return false
	}
	
	if user.Id != userId {
		return false
	}

	return true
}

func isOpportunityIdValid(opportunityId uint64) bool{
	
	type opportunityTemp struct {
		Id uint64
	}
	
	var res opportunityTemp
	err := db.Table("opportunities").Select("id").Where("id=?", fmt.Sprint(opportunityId)).Scan(&res).Error
	if err != nil {
		fmt.Println("Error validating opportunity id: ", err)
		return false
	}
	
	if res.Id != opportunityId {
		fmt.Println("Error validating opportunity id!")
		return false
	}

	return true
}

func CreateInvestment(userId, opportunityId uint64, amountBought float64)(investment models.Investment, err error) {
	
	// validate userId
	if !isUserIdValid(userId) {
		err = errors.New("User id is invalid\n")
		return
	}
	
	// validate opportunityId
	if !isOpportunityIdValid(opportunityId) {
		err = errors.New("Opportunity id is invalid\n")
		return
	}	
	
	investment = models.Investment{
		UserId: userId,
		OpportunityId: opportunityId,
		AmountBought: amountBought,
	}
	db.Create(&investment)
	if db.NewRecord(investment) {
		err = errors.New("failed to create investment\n")
		return
	}
	return
}

func GetInvestmentsByInvestor(id uint64)(investments []models.Investment, err error){
	
	err = db.Where(&models.Investment{UserId: id}).Find(&investments).Error
	if err != nil {
		err = fmt.Errorf("failed to retrieve investments made by an investor: %v", err)
		return
	}
	return
}

func GetInvestment(id uint64)(investment models.Investment, err error) {
	
	err = db.Where("id=?",fmt.Sprint(id)).First(&investment).Error
	if err != nil {
		err = fmt.Errorf("failed to retrieve investment: %v", err)
		return
	}
	return
}

func GetAllInvestments()(investments []models.Investment, err error){
	
	err = db.Find(&investments).Error
	if err != nil {
		err = fmt.Errorf("failed to retrieve all investments: %v", err)
		return
	}
	return
}