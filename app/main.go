package main

import (
	"log"
	"sync"

	"vmuc-fintech-backend-web-go/db"
	"vmuc-fintech-backend-web-go/vmuc/delivery"
	"vmuc-fintech-backend-web-go/vmuc/repository"
	"vmuc-fintech-backend-web-go/vmuc/usecase"

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
	refRepo := repository.NewPostgreRef(db.GormClient.DB)
	employeeRepo := repository.NewPostgreEmployee(db.GormClient.DB)
	payrollRepo := repository.NewPostgrePayroll(db.GormClient.DB)

	timeoutContext := fiber.Config{}.ReadTimeout

	userUseCase := usecase.NewUserUseCase(usrRepo, timeoutContext)
	refUseCase := usecase.NewRefUseCase(refRepo, timeoutContext)
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepo, timeoutContext)
	payrollUseCase := usecase.NewPayrollUseCase(payrollRepo, timeoutContext)

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
		delivery.NewUserHandler(app, userUseCase)
		delivery.NewRefHandler(app, refUseCase)
		delivery.NewEmployeeHandler(app, employeeUseCase)
		delivery.NewPayrollHandler(app, payrollUseCase)
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
