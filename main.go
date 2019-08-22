package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func sample(c echo.Context) error {
    return c.JSON(
        http.StatusOK,
        struct {
            Code int    `json:"code"`
            Text string `json:"body"`
        }{
            Code: http.StatusOK,
            Text: http.StatusText(http.StatusOK),
        },
    )
}

func main() {
    e := echo.New()
    // Routes
    e.GET("/sample", sample)
    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}

