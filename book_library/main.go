package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"

	_ "github.com/lib/pq"
)

var BookMap = make(map[string]Book)

func main() {
	db := getDb()
	defer db.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	bookPath := r.Group("/book")
	{
		bookPath.POST("/register", func(ctx *gin.Context) {
			body := ctx.Request.Body
			value, err := ioutil.ReadAll(body)
			if err != nil {
				fmt.Println(err.Error())
			}

			var bodyData map[string]string
			json.Unmarshal(value, &bodyData)
			book := NewBook(bodyData["name"], bodyData["writer"], bodyData["publisher"])
			result, err := db.Exec("INSERT INTO book(name, writer, publisher) VALUES ($1, $2, $3)", book.Name, book.Writer, book.Publisher)
			fmt.Println(result.LastInsertId())
			fmt.Println(err)

			// BookMap[book.Id] = *book
			// fmt.Println(book)
			ctx.JSON(http.StatusOK, book)
		})

		bookPath.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			res, find := BookMap[id]
			if !find {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "잘못입력하셨습니다.",
				})
				return
			}
			ctx.JSON(http.StatusOK, res)
		})

	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getDb() *sql.DB {
	db, err := sql.Open("postgres", "user=qwer password=qwer1234 host=postgres port=5432 dbname=book sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
