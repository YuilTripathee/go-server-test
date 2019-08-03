package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// 	fmt.Printf("\n" + time.ANSIC)
	// }))
	// e.Use(middleware.Logger())
	// t := time.Now()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Format: "[${method}] | uri=${uri} [${status}]\n",
		Format: "\n[API] | ${host} | ${time_rfc3339} | ${status} | ${latency_human} | ${remote_ip} | ${method} ${uri} ",
		Output: os.Stdout,
	}))
	e.GET("/", func(c echo.Context) (err error) {
		return c.JSON(200, []string{
			"Jack",
			"Jill",
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
