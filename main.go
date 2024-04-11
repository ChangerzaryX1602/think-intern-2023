package main

import (
	"fmt"
	"math"
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
	if err != nil || orderInt <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order parameter. Must be a positive integer.",
		})
	}

	// Estimate the upper bound using the prime number theorem
	upperBound := int(float64(orderInt) * (math.Log(float64(orderInt)) + math.Log(math.Log(float64(orderInt)))))
	if upperBound < 2 {
		upperBound = 2
	}

	// Generate primes up to the upper bound
	primes := generatePrimes(upperBound)

	// If the list of primes does not contain enough primes, extend the list
	if len(primes) < orderInt {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate enough primes",
		})
	}

	// Retrieve the `order`-th prime number
	primeNum := primes[orderInt-1]

	// Return the prime number in the response
	return c.JSON(fiber.Map{
		"order":        orderInt,
		"prime_number": primeNum,
	})
}

// Generate primes up to the specified upper bound using the Sieve of Eratosthenes
func generatePrimes(upperBound int) []int {
	// Create a boolean array for the sieve
	isPrime := make([]bool, upperBound+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false // 0 and 1 are not prime

	// Sieve of Eratosthenes
	for i := 2; i*i <= upperBound; i++ {
		if isPrime[i] {
			for j := i * i; j <= upperBound; j += i {
				isPrime[j] = false
			}
		}
	}

	// Collect the primes
	primes := make([]int, 0)
	for i := 2; i <= upperBound; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
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
