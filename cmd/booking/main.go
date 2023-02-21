package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/e346m/bsg/di"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	conn := initDB()

	defer func() {
		conn.Close()
	}()

	bookingHandler := di.NewbookingHTTP(conn)

	e := initEcho()
	api := e.Group("/api")
	{
		api.GET("/booking", bookingHandler.GetBooking)
	}
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func initDB() *sql.DB {
	conn, err := sql.Open("postgresql", "url")

	if err != nil {
		panic(err)
	}

	if err = conn.Ping(); err != nil {
		panic(err)
	}

	// handle with config
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(5 * time.Minute)

	return conn
}

func initEcho() *echo.Echo {
	e := echo.New()
	e.Use(echoRecover())
	e.Use(echoSecureHeaderConfig())
	return e
}

func echoRecover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10,
		DisableStackAll:   false,
		DisablePrintStack: false,
	})
}

func echoSecureHeaderConfig() echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	})
}
