package storage

import (
	"GroupExpenseTracker/pkg/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func New() error {
	databaseConfig := config.GetConfig().DatabaseConfig
	user := databaseConfig.User
	password := databaseConfig.Password
	address := databaseConfig.Address
	databaseName := databaseConfig.DatabaseName

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, address, databaseName)
	tmp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db = tmp
	return nil
}

func NewUser(telegram_user_id int64) error {
	result := db.Create(User{
		TelegramUserID: telegram_user_id,
	})
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func NewGroup(telegram_chat_id int64) error {
	result := db.Create(Group{
		TelegramChatID: telegram_chat_id,
	})
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func AddUserToGroup(telegram_user_id int64, telegram_chat_id int64) error {
	var user User
	var group Group
	result := db.Select("user_id").Where("telegram_user_id = ?", telegram_user_id).Find(&user)
	if result.Error != nil {
		return result.Error
	}
	result = db.Select("group_id").Where("telegram_chat_id = ?", telegram_chat_id).Find((&group))
	if result.Error != nil {
		return result.Error
	}
	result = db.Create(UsersInGroups{
		UserID: user.UserID,
		GroupID: group.GroupID,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddExpense()