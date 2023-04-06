package main

import (
	"net/http"

	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ApiResponse struct {
	Result string      `json:"result"`
	Ret    interface{} `json:"ret"`
}

type Account struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Balance      float32 `json:"balance"`
	ModifiedDate string  `json:"modified_date" gorm:"column:modify_date"`
	CreateDate   string  `json:"created_date" gorm:"column:create_date"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	db := connectDB()
	// var user = Account{ID: 1, Name: "test", Balance: 100.0, ModifiedDate: "2021-01-01", CreateDate: "2021-01-01"}
	var user Account
	result := db.Table("Account").Where("id = ?", id).First(&user)
	// result := db.Table("Account").Find(&users)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, ApiResponse{
				Result: "ok",
				Ret:    []Account{},
			})
		}
		panic(result.Error)
	}

	response := ApiResponse{
		Result: "ok",
		Ret:    user,
	}

	return c.JSON(http.StatusOK, response)
}

func connectDB() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/laravel_test?charset=utf8mb4&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
