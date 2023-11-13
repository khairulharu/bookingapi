package component

import (
	"fmt"
	"log"

	"github.com/khairulharu/bookingapi/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%v "+"port=%v "+"user=%v "+"password=%v "+"dbname=%v "+"sslmode=%v",
		cnf.DB.Host, cnf.DB.Port, cnf.DB.User, cnf.DB.Pass, cnf.DB.Name, cnf.DB.SSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error When Connect To Database: %v", err.Error())
	}
	fmt.Println("Database Connect")
	return db
}
