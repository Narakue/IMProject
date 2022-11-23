package main

import (
	"IMProject/common"
	"IMProject/router"
	"IMProject/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func main() {
	util.InitConfig()
	db := common.InitDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			println("error")
		}
	}(db)
	r := gin.Default()
	//r.Use(util.LoggerToFile())
	//r.Use(common.MiddleLog)
	r = router.CommonRouter(r)
	r = router.UserRouter(r)
	port := viper.GetString("server.port")
	err := r.Run(":" + port)
	if err != nil {
		return
	}
}
