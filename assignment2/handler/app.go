package handler

import (
	"assignment-2/database"
	"assignment-2/docs"
	"assignment-2/handler/http_handler"
	"assignment-2/repository/item_repository/item_pg"
	"assignment-2/repository/order_repository/order_pg"
	"assignment-2/service"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	docs.SwaggerInfo.Title = "Assigment 2"
	docs.SwaggerInfo.Description = "uhuy"
	docs.SwaggerInfo.Version = "v0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func StartApp() {
	db := database.GetPostgresInstance()

	itemRepo := item_pg.NewItemPG(db)
	orderRepo := order_pg.NewOrderPG(db, itemRepo)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := http_handler.NewOrderHandler(orderService)

	r := gin.Default()

	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders", orderHandler.GetAllOrders)
	r.GET("/orders/:orderID", orderHandler.GetOrderByID)
	r.PATCH("/orders/:orderID", orderHandler.UpdateOrderByID)
	r.DELETE("/orders/:orderID", orderHandler.DeleteOrderByID)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err.Error())
	}
}
