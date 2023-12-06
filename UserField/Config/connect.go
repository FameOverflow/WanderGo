package Config

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
	if (!GLOBAL_DB.Migrator().HasTable(&Star{})) {
		err := GLOBAL_DB.AutoMigrate(&Star{})
		if err != nil {
			log.Println(err)
			return
		}
	}
	// input := `图书馆 115.79941，28.656973-115.800155，28.655495
	// 基础实验大楼 115.79684，28.658394-115.798664，28.656304
	// 天健19-23栋 115.795928，28.658503
	// 天健24-25栋 115.794871，28.656027
	// 理生楼 115.796851,28.659439-115.799088，28.658474
	// 材环楼 115.797044，28.660536-115.798997，28.659482
	// 信工楼 115.797977，28.661699-115.799694，28.660564
	// 机电楼 115.798058,28.661722-115.801668,28.662683
	// 建工楼 115.798964,28.662739-115.802559,28.66329
	// 休闲1-3栋 115.801743,28.66393-115.803578,28.665229
	// 一食堂 115.803739,28.66482-115.804527,28.664344
	// 润溪湖（1） 115.805922,28.66449-115.804366,28.665149
	// 润溪湖（2） 115.807692，28.662518-115.805268,28.663883
	// 润溪湖（3） 115.80722,28.662584-115.810235,28.661492
	// 艺术楼 115.807285,28.661303-115.808915,28.659976
	// 外经楼 115.805396,28.661501-115.80677,28.66265
	// 慧源楼 115.805257，28.663874-115.803658，28.662621
	// 文法楼 115.803744,28.661209-115.805332,28.660023
	// 正气广场 115.805461,28.659948-115.807285,28.658969
	// 白帆 115.810444,28.663709-115.813459,28.65967
	// 体育馆 115.810755,28.659557-115.812515,28.658682
	// 游泳馆 115.810187,28.663709-115.812043,28.664829
	// 休闲运动场 115.809919,28.664886-115.811689,28.667747
	// 休闲广场 115.808256,28.665747-115.806823,28.665968
	// 休闲6-8栋 115.807558,28.665681-115.806335，28.664645
	// 休闲9-12.16-17栋 115.808422,28.664664-115.809661,28.666942
	// 休闲13栋 115.806453,28.665992-115.807451,28.666476
	// 休闲14-18栋 115.806534,28.666519-115.809575,28.667507
	// 树人广场 115.802779,28.657034-115.803723,28.656224
	// 龙腾湖（1） 115.800552,28.656078-115.802151,28.654516
	// 龙腾湖（2） 115.802655,28.656088-115.803996,28.655382
	// 龙腾湖（3） 115.805434,28.657453-115.806378,28.657058
	// 龙腾湖（4） 115.806217,28.65863-115.809768,28.657952
	// 天健运动场 115.793815,28.653009-115.796679,28.654817
	// 医学实验大楼 115.796829,28.656088-115.798213,28.653216
	// 研究生 115.793461,28.65299-115.79625,28.651258`
	// parseInput(input)
}
