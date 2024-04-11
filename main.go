package main

import (
	"fmt"
	"strconv"
	"strings"
	"think-intern-2023/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		logs.Error("Error loading env file")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func main() {
	// db, err := database.CreateDB()
	// if err != nil {
	// 	logs.Error(err)
	// }
	// database.AutoMigrate(db)
	// thinkRepo := thinkRepository.NewThinkRepository(db)
	// thinkService := thinkService.NewThinkService(thinkRepo)
	// thinkHandler := thinkHandler.NewThinkHandler(thinkService)
	app := fiber.New()
	fiberConfig(app)
	//authorized := app.Group("/api/v1/authorized", entityHandlers.JWTAuthen()) ไม่จำเป็นต้องใช้
	app.Post("/api/v1/think/:order", primeOrder)
	app.Listen(fmt.Sprintf(":%s", viper.GetString("APP_PORT")))
}
func primeOrder(c *fiber.Ctx) error {
	order := c.Params("order")
	orderInt, err := strconv.Atoi(order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	primeNum := 0
	count := 0
	i := 2
	for count < orderInt {
		if isPrime(i) {
			count++
			primeNum = i
		}
		i++
	}

	return c.JSON(fiber.Map{"prime_number": primeNum, "order": orderInt})
}
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func fiberConfig(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max",
	}))
	app.Use(logger.New(logger.Config{
		CustomTags: map[string]logger.LogFunc{
			"port": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(viper.GetString("APP_PORT"))
			},
			"msg": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				message := fmt.Sprintf("Response body: %s", c.Response().Body())
				return output.WriteString(message)
			},
		},
		Format:     "[${time}] Status: ${status} - Medthod: ${method} API: ${path} Port:${port}\n Message: ${msg}\n\n",
		TimeFormat: "2 Jan 2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))
}
