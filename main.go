package main

import (
	_ "emtesttask/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Message string `json:"message"`
}

// @title     Subscriptions Service
// @version         1.0
// @description     Effective Mobile test task
func main() {
	/*sqlDB := createConnection()
	defer sqlDB.Close()
	applyMigrations(sqlDB)

	// GORM будет использовать тот же *sql.DB
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	subscriptions := router.Group("/subscriptions")
	subscriptions.GET("/test/:id", func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorInfo{Message: "Failed to parse id " + ctx.Param("id")})
			return
		}
		startTime := time.Date(2020, time.May, 11, 0, 0, 0, 0, time.Local)
		ctx.JSON(http.StatusOK, SubscriptionDTO{
			//ID: 1,
			ServiceName: "Service 1", Price: 1000, UserId: id, StartDate: startTime, EndDate: time.Now(),
		})
	})
	subscriptions.GET("/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorInfo{Message: "Failed to parse id " + ctx.Param("id")})
			return
		}
		c := context.Background()
		sub, err := gorm.G[entity.Subscription](db).Where("id = ?", id).First(c)
		ctx.JSON(http.StatusOK, SubscriptionDTO{
			//ID: 1,
			ServiceName: sub.ServiceName, Price: int64(sub.Price), UserId: sub.UserId, StartDate: sub.StartDate, EndDate: sub.EndDate,
		})
	})*/
	router, cleanup, _ := InitializeApp()
	defer cleanup()

	router.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080") // f5bb8a33-ab4a-4659-a8cd-39eb83cbdd24
}
