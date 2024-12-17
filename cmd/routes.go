package app

//    app.Use(func(c *fiber.Ctx) error {
//         return c.SendStatus(404) // => 404 "Not Found"
//     })


// CORS AND CSRF 
// app.Use(helmet.New())
// app.Use(idempotency.New(idempotency.Config{
// 	Lifetime: 42 * time.Minute,
// 	// ...
// }))
// app.Use(healthcheck.New())
// app.Use(limiter.New(limiter.Config{
//     Next: func(c *fiber.Ctx) bool {
//         return c.IP() == "127.0.0.1"
//     },
//     Max:          20,
//     Expiration:     30 * time.Second,
//     KeyGenerator:          func(c *fiber.Ctx) string {
//         return c.Get("x-forwarded-for")
//     },
//     LimitReached: func(c *fiber.Ctx) error {
//         return c.SendFile("./toofast.html")
//     },
//     Storage: myCustomStorage{},
// }))
// app.Get("/metrics", monitor.New())
// app.Use(recover.New())

// app.Use(requestid.New())

// // Or extend your config for customization
// app.Use(requestid.New(requestid.Config{
//     Header:    "X-Custom-Header",
//     Generator: func() string {
//         return "static-id"
//     },
// }))

// app := fiber.New(fiber.Config{
// 		IdleTimeout: idleTimeout,
// 	})

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello world!")
// 	})

// 	// Listen from a different goroutine
// 	go func() {
// 		if err := app.Listen(":3000"); err != nil {
// 			log.Panic(err)
// 		}
// 	}()

// 	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

// 	_ = <-c // This blocks the main thread until an interrupt is received
// 	fmt.Println("Gracefully shutting down...")
// 	_ = app.Shutdown()

// 	fmt.Println("Running cleanup tasks...")

// 	// Your cleanup tasks go here
// 	// db.Close()
// 	// redisConn.Close()
// 	fmt.Println("Fiber was successful shutdown.")