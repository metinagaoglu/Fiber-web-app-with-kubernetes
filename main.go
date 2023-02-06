package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"log"
	"net"
	"time"
)

func connectRedis() *redis.Client {
	// new redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
	})

	// test connection
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	// return pong if server is online

	return client
}

func main() {
	app := fiber.New()
	rdb := connectRedis()

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("[GET] - /")
		err := rdb.Set("app", "up", 2*time.Second).Err()
		if err != nil {
			panic(err)
		}
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/redis-status", func(c *fiber.Ctx) error {
		status, _ := rdb.Get("app").Result()
		return c.SendString("Hello, World 👋! Redis is " + status)
	})

	app.Get("/flush", func(c *fiber.Ctx) error {
		rdb.FlushAll()
		return c.SendString("Flushed")
	})

	app.Get("/server-info", func(c *fiber.Ctx) error {
		log.Println("[GET] - /server-info")

		addrs, err := net.InterfaceAddrs()
		if err != nil {
			fmt.Println(err)
		}

		var ip string
		for _, address := range addrs {
			// check the address type and if it is not a loopback the display it
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
				}
			}
		}

		return c.SendString("server ip adress " + ip)
	})

	log.Println("Listening 3000....")
	app.Listen(":3000")
}
