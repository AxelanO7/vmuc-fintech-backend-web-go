package main

import (
	"log"
	"sync"

	"assyarif-backend-web-go/assyarif/delivery"
	"assyarif-backend-web-go/assyarif/repository"
	"assyarif-backend-web-go/assyarif/usecase"
	"assyarif-backend-web-go/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	Init()
	// initEnv()
	listenPort := ":4000"
	// appName := os.Getenv("APP_NAME")

	usrRepo := repository.NewPostgreUser(db.GormClient.DB)
	inRepo := repository.NewPostgreIn(db.GormClient.DB)
	outRepo := repository.NewPostgreOut(db.GormClient.DB)

	timeoutContext := fiber.Config{}.ReadTimeout

	userUseCase := usecase.NewUserUseCase(usrRepo, timeoutContext)
	inUseCase := usecase.NewInUseCase(inRepo, timeoutContext)
	outUseCase := usecase.NewOutUseCase(outRepo, timeoutContext)

	app := fiber.New(fiber.Config{})
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${green} ${status} ${white} | ${latency} | ${ip} | ${green} ${method} ${white} | ${path} | ${yellow} ${body} ${reset} | ${magenta} ${resBody} ${reset}\n",
		TimeFormat: "02 January 2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
	app.Use(cors.New())

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {

		//call delivery http here
		delivery.NewUserHandler(app, userUseCase)
		delivery.NewInHandler(app, inUseCase)
		delivery.NewOutHandler(app, outUseCase)
		log.Fatal(app.Listen(listenPort))
		wg.Done()
	}()
	wg.Wait()

}

func Init() {
	InitEnv()
	InitDB()
}

func InitEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}
}

func InitDB() {
	db.NewGormClient()
}
