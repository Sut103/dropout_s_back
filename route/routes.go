package route

import (
	"dropout_s_back/config"
	"dropout_s_back/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Init ルーティング
func Init(conn *gorm.DB) *gin.Engine {
	r := gin.Default()
	ctrler := controller.NewController(conn)

	r.GET("/ble/get", ctrler.GetBle)
	r.GET("/ble/getall", ctrler.GetBleAll)
	r.GET("/message/get", ctrler.GetMessage)
	r.POST("/user/signup", ctrler.SignUp)
	r.POST("/message/post", ctrler.PostMessage)
	r.GET("/user/get", ctrler.GetUsers)

	return r
}
