package migration

import (
	"fmt"
	"log"
	databases "tes/database"
	"tes/model/entity"
)

func RunMigration() {
	err := databases.DB.AutoMigrate(
		entity.User{},
	)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Databases Migrated")
}
