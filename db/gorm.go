package db

import (
	"WeChatPusher/config"
	"WeChatPusher/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var DB *gorm.DB

// InitDb gorm初始化函数
func InitDb() error {
	var err error
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.HOST + ")/" + config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	DB.AutoMigrate(&model.DbData{})
	return nil
}
func GetDay() (error, string) {
	data := model.DbData{}
	err := DB.Model(&model.DbData{}).Last(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err, ""
	}
	return nil, data.AcquaintanceDay
}

func GetCount() (string, error) {
	data := model.DbData{}
	err := DB.Model(&model.DbData{}).Last(&data).Error
	if err != nil {
		return "0", err
	}
	count := int(data.Model.ID)
	return strconv.Itoa(count), nil
}
func SaveData(dbData *model.DbData) error {
	err := DB.Model(&model.DbData{}).Create(dbData).Error
	if err != nil {
		return err
	}
	return nil
}
