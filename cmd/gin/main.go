package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"iot.api/internal/middlewares"
	"iot.api/internal/routes"
)

func main() {
	port := ":" + os.Getenv("APP_PORT")
	
	router := gin.Default()
	
	router.Use(cors.Default())
	router.Use(middlewares.BodyParser())
	
	gpsRouter := router.Group("/gps")
	{
		gpsRouter.GET("/count", routes.FindAll)
		gpsRouter.POST("/test", routes.TestWrite)
		gpsRouter.POST("/test/routine", routes.TestWriteRoutine)
		gpsRouter.GET("/test/random", routes.TestRandomWrite)
		gpsRouter.GET("/test/random/routine", routes.TestRandomWriteRoutine)
		gpsRouter.GET("/test/batch", routes.TestWriteBatch)
	}

	router.GET("/", routes.Home)
	router.GET("/hi", routes.Hi)
	
	router.Run(port)
	// //原本是用router.Run()，要使用net/http套件的shutdown的話，需要使用原生的ListenAndServe
	// srv := &http.Server{
	// 	Addr:    ":8787",
	// 	Handler: router,
	// }
	// //新增一個channel，type是os.Signal
	// ch := make(chan os.Signal, 1)
	// //call goroutine啟動http server
	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		fmt.Println("SERVER GG惹:", err)
	// 	}
	// }()
	// //Notify：將系統訊號轉發至channel
	// signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	// //阻塞channel
	// <-ch
	// fmt.Println("Graceful Shutdown start")
}
