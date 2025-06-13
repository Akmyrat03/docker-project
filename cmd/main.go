package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis-server:6379", // adjust if your Redis is somewhere else
	})

	// Use a background context for Redis commands
	ctx := context.Background()

	app := fiber.New()

	app.Get("/set", func(c *fiber.Ctx) error {
		err := rdb.Set(ctx, "greeting", "Hello, Mekan!", 10*time.Minute).Err()
		if err != nil {
			return c.Status(500).SendString("Failed to set value in Redis")
		}
		return c.SendString("Value set in Redis")
	})

	app.Get("/get", func(c *fiber.Ctx) error {
		val, err := rdb.Get(ctx, "greeting").Result()
		if err == redis.Nil {
			return c.SendString("Key does not exist")
		} else if err != nil {
			return c.Status(500).SendString("Failed to get value from Redis")
		}
		return c.SendString(val)
	})

	log.Fatal(app.Listen(":8081"))
}
