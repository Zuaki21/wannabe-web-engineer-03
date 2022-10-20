package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jsonData struct {
	Number int    `json:"number,omitempty"`
	String string `json:"string,omitempty"`
	Bool   bool   `json:"bool,omitempty"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/post", postHandler)

	e.POST("/add", addHandler)

	e.POST("/hello/:name", helloHandler)

	e.POST("/fizzbuzz", fizzbuzzHandler)

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello,world.\n")
	})
	e.GET("/json", jsonHandler)

	e.GET("/ping", pingHandler)

	e.GET("/incremental", incrementalHandler)

	e.Start(":8080")
}

func jsonHandler(c echo.Context) error {
	res := jsonData{
		Number: 10,
		String: "hoge",
		Bool:   false,
	}

	return c.JSON(http.StatusOK, &res)
}

func postHandler(c echo.Context) error {
	var data jsonData
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}
	return c.JSON(http.StatusOK, data)
}

func helloHandler(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name+".\n")
}

func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

var i int

func incrementalHandler(c echo.Context) error {
	i++
	return c.String(http.StatusOK, fmt.Sprint(i))
}

func fizzbuzzHandler(c echo.Context) error {
	count := c.QueryParam("count")
	i, _ = strconv.Atoi(count)

	if !NumCheck(count) || i < 1 {
		return c.String(http.StatusBadRequest, "BadRequest")
	}

	var str string
	for x := 1; x <= i; x++ {
		if x%15 == 0 {
			str += "FizzBuzz\n"
		} else if x%5 == 0 {
			str += "Buzz\n"
		} else if x%3 == 0 {
			str += "Fizz\n"
		} else {
			str += strconv.Itoa(x) + "\n"
		}
	}
	return c.String(http.StatusOK, str)
}

type addJsonData struct {
	Right int `json:"right,omitempty"`
	Left  int `json:"left,omitempty"`
}
type answerJsonData struct {
	Answer int `json:"answer",omitempty`
}

func addHandler(c echo.Context) error {
	var data addJsonData
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	var answerData answerJsonData
	answerData.Answer = data.Left + data.Right

	return c.JSON(http.StatusOK, answerData)
}

func NumCheck(str string) bool {
	for _, r := range str {
		if '0' <= r && r <= '9' {
			return true
		}
	}
	return false
}
