package server

import (
	"math"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	// Mean service time (in milliseconds)
	MeanServiceTime = 50.0
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	s.App.Get("/", s.HelloWorldHandler)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	// Generate exponentially distributed service time
	// Using -mean * ln(U) where U is uniform(0,1)
	u := rand.Float64()
	serviceTime := -MeanServiceTime * float64(time.Millisecond) * math.Log(u)

	// Sleep for the generated service time
	time.Sleep(time.Duration(serviceTime))

	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}
