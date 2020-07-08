package repo

import (
	"errors"
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
	db.AutoMigrate(&models.Chat{})
	db.Model(&models.Chat{}).AddForeignKey("sender_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Chat{}).AddForeignKey("receiver_id", "users(id)", "RESTRICT", "RESTRICT")
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

func CreateChat(senderId, receiverId uint64, message string, sentAt time.Time) (chat models.Chat, err error) {

	// check if senderId is valid
	if !isUserValid(senderId) {
		err = errors.New("invalid sender id")
		return
	}

	// check if receiverId is valid
	if !isUserValid(receiverId) {
		err = errors.New("invalid receiver id")
		return
	}
	
	if sentAt.After(time.Now()) {
		err = errors.New("time is not acceptable, it appears to be in future!")
		return
	}

	chat = models.Chat{
		SenderId:   senderId,
		ReceiverId: receiverId,
		Message:    message,
		SentAt:     sentAt,
	}
	db.Create(&chat)
	if db.NewRecord(chat) {
		err = fmt.Errorf("Error creating a new chat: %v", err)
		return
	}

	return
}

type partner struct{
	SenderId uint64
	ReceiverId uint64
}

func filterCouterParts(id uint64, panters []partner) (counterparts []uint64, err error) {
	
	for _, partner := range panters {

		if partner.SenderId != id && partner.ReceiverId == id {
			counterparts = append(counterparts, partner.SenderId)
		} 

		if partner.SenderId == id && partner.ReceiverId != id {
			counterparts = append(counterparts, partner.ReceiverId)
		} 
	}

	return
}

func GetChatCounterPartsOf(id uint64) (counterparts []models.User, err error) {

	var partners []partner

	err = db.Raw("SELECT DISTINCT sender_Id, receiver_Id FROM chats WHERE sender_id=? or receiver_id=?", id, id).Scan(&partners).Error
	if err != nil {
		err = fmt.Errorf("failed to get all distinct partners: %v\n", err)
		return
	}

	counterpartsId, err := filterCouterParts(id, partners)
	if err != nil {
		err = fmt.Errorf("failed to filter out the counterparts id of id=%d, err: %v\n", id, err)
		return
	}

	err = db.Table("users").Where("id IN (?) ", counterpartsId).Scan(&counterparts).Error
	if err != nil {
		err = fmt.Errorf("failed to read counterparts into user struct, err: %v\n", err)
		return
	}

	return
}

func GetChats(u1, u2 uint64) (chats []models.Chat, err error) {

	err = db.Where("(sender_id=? and receiver_id=?) or (sender_id=? and receiver_id=?)", u1, u2, u2, u1).Find(&chats).Error
	if err != nil {
		err = fmt.Errorf("error getting chats from db")
		return
	}

	return
}
