package Database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var GLOBAL_DB *gorm.DB

func ConnectToDb() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:fakepwd@tcp(127.0.0.1:3306)/spark_forge?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Println(err)
		return
	}
	GLOBAL_DB = db
	if (!GLOBAL_DB.Migrator().HasTable(&User{})) {
		err := GLOBAL_DB.AutoMigrate(&User{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	if (!GLOBAL_DB.Migrator().HasTable(&Avatar{})) {
		err := GLOBAL_DB.AutoMigrate(&Avatar{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	if (!GLOBAL_DB.Migrator().HasTable(&Photo{})) {
		err := GLOBAL_DB.AutoMigrate(&Photo{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	if (!GLOBAL_DB.Migrator().HasTable(&Place{})) {
		err := GLOBAL_DB.AutoMigrate(&Place{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	if (!GLOBAL_DB.Migrator().HasTable(&Comment{})) {
		err := GLOBAL_DB.AutoMigrate(&Comment{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	// GLOBAL_DB.Model(&Place{}).Create(&Place{
	// 	PlaceName: "慧源楼",
	// 	PlaceUID:  0,
	// 	TopLeftPoint: Address{
	// 		X: -30,
	// 		Y: 30,
	// 	},
	// 	BottomRightPoint: Address{
	// 		X: 30,
	// 		Y: -30,
	// 	},
	// 	CenterPoint: Address{
	// 		X: 0,
	// 		Y: 0,
	// 	},
	// })
	// GLOBAL_DB.Model(&Place{}).Create(&Place{
	// 	PlaceName: "一食堂",
	// 	PlaceUID:  1,
	// 	TopLeftPoint: Address{
	// 		X: -50,
	// 		Y: 5,
	// 	},
	// 	BottomRightPoint: Address{
	// 		X: -40,
	// 		Y: -5,
	// 	},
	// 	CenterPoint: Address{
	// 		X: -45,
	// 		Y: 0,
	// 	},
	// })
}
