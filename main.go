/**
Golang upper v1.11
with echo, gorm modules with mysql database sample
by Yoshi Sakai
 */
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "strconv"
)
/*
** Config for Database table
 */
type Product struct {
    gorm.Model
    Code string
    Price uint
}
/*
** Config for Json data
 */
type Item struct {
    Code string `json:"code"`
    Price uint `json:"price"`
}
/*
** Config for MySQL setting
 */
func gormConnect() *gorm.DB {
    DBMS     := "mysql"
    USER     := "docker"
    PASS     := "docker"
    PROTOCOL := "tcp(mysql_host:3306)"
    DBNAME   := "test_database"
    // add parseTime option
    CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=true"
    db,err := gorm.Open(DBMS, CONNECT)

    if err != nil {
        panic(err.Error())
    }
    return db
}
/*
** Common return of Json
 */
func returnJson(c echo.Context, code string) error {
    return c.JSON(
        http.StatusOK,
        struct {
            Code int    `json:"code"`
            Text string `json:"body"`
        }{
            Code: http.StatusOK,
            Text: code,
        },
    )
}
/*
** HELLO WORLD!
 */
func helloWorld(c echo.Context) error {
    u := new(Item)
    if err := c.Bind(u); err != nil {
        return err
    }
    return c.JSON(http.StatusCreated,u)
}
/*
** Insert record by Json data
 */
func insert(c echo.Context) error {
    db := gormConnect()
    db.AutoMigrate(&Product{})
    u := new(Item)
    if err := c.Bind(u); err != nil {
        return err
    }
    db.Create(&Product{Code: u.Code, Price: u.Price})
    defer db.Close()
    return returnJson(c, u.Code)
}
/*
** Find record by Json data
 */
func find(c echo.Context) error {
    db := gormConnect()
    u := new(Item)
    if err := c.Bind(u); err != nil {
        return err
    }
    var product Product
    db.First(&product, "code = ?", u.Code)
    defer db.Close()
    return returnJson(c, strconv.FormatUint(uint64(product.Price),10))
}
/*
** Update record by Json data
 */
func update(c echo.Context) error {
    db := gormConnect()
    db.AutoMigrate(&Product{})
    u := new(Item)
    if err := c.Bind(u); err != nil {
        return err
    }
    var product Product
    db.First(&product, "code = ?", u.Code)
    db.Model(&product).Update("Price", u.Price)
    defer db.Close()
    return returnJson(c, u.Code)
}
/*
** Delete record by Json data
 */
func delete(c echo.Context) error {
    db := gormConnect()
    db.AutoMigrate(&Product{})
    u := new(Item)
    if err := c.Bind(u); err != nil {
        return err
    }
    var product Product
    db.First(&product, "code = ?", u.Code)
    db.Delete(&product)
    defer db.Close()
    return returnJson(c, u.Code)
}
/*
** main loop
 */
func main() {
    e := echo.New()
    // Routes
    e.POST("/helloWorld", helloWorld)
    e.POST("/insert", insert)
    e.POST("/find", find)
    e.POST("/update", update)
    e.POST("/delete", delete)
    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}

