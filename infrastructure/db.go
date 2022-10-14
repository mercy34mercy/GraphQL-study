package infrastructure

import (
	"fmt"
	"graphql/infrastructure/dto"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)


func Init() *gorm.DB {
	err = godotenv.Load(".env")
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
	db, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		fmt.Println("db init error: ", err)
	}
	autoMigrate()
	fmt.Println("[INFO] db setup done!")
	return db
}

func GetDB() *gorm.DB {
	return db
}
 

func autoMigrate() {
	db.
		AutoMigrate(&dto.Todo{})
	db.
		AutoMigrate(&dto.User{})
}

