package api

import (
	_ "add/api/docs"
	"add/api/handler"
	"add/storage"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandler(storage)

	r.POST("/user", handlerV1.CreateUser)
	r.GET("/user/:id", handlerV1.GetByIDUser)
	r.GET("/user", handlerV1.GetListUser)
	r.PUT("/user/:id", handlerV1.UpdateUser)
	r.DELETE("/user/:id", handlerV1.DeleteUser)

	r.POST("/subscription", handlerV1.CreateSubscription)
	r.GET("/subscription/:id", handlerV1.GetByIdSubscription)
	r.GET("/subscription", handlerV1.GetListSubscription)
	r.DELETE("/subscription/:id", handlerV1.DeleteSubscription)
	r.PUT("/subscription/:id", handlerV1.UpdateSubscription)

	r.POST("/subscription-user", handlerV1.CreateSubscriptionUser)
	r.GET("/subscription-get/:id", handlerV1.GetUserSubscription)
	r.GET("/subscription-user", handlerV1.GetByIdSubscriptionUser)
	r.PUT("/subscription-user/:id", handlerV1.UpdateSubscriptionUser)

	r.GET("/subscription-user/:id", handlerV1.GetByIdSubscriptionUser)
	r.GET("Check-access/:id",handlerV1.CheckAccess)

	r.GET("make-active/:id",handlerV1.MakeSubscriptionActive)


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
