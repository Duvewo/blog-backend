package posts

import (
	"blog/cache"
	"blog/storage"
	"blog/types"
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func Get(redis *cache.Cache, db *storage.Connection) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")

		if !redis.KeyExists(id) {

			count := 0
			rows, err := db.Conn.Query("SELECT COUNT(*) FROM posts WHERE id = " + id)

			if err != nil {
				panic(err)
			}

			for rows.Next() {
				rows.Scan(&count)
			}

			if count == 0 {
				var response types.Error
				response.Error.ErrorMsg = types.PostNotFoundError
				response.Error.ErrorCode = types.PostNotFoundCode

				return c.JSON(200, response)
			}

			rows, err = db.Conn.Query("SELECT * FROM posts WHERE id = " + id)

			if err != nil {
				panic(err)
			}

			var response types.Response
			for rows.Next() {
				rows.Scan(&response.Response.ID, &response.Response.Title, &response.Response.Body, &response.Response.Date)
			}

			redis.AddValue(id, response.Response.ID, time.Minute*10)
			redis.AddValue(id+"title", response.Response.Title, time.Minute*10)
			redis.AddValue(id+"body", response.Response.Body, time.Minute*10)
			redis.AddValue(id+"date", response.Response.Date, time.Minute*10)

			return c.JSON(200, response)

		} else {
			postID, err := strconv.Atoi(redis.GetValue(id))

			if err != nil {
				fmt.Println(err)
			}

			body := redis.GetValue(id + "body")
			date := redis.GetValue(id + "date")
			title := redis.GetValue(id + "title")

			var response types.Response
			response.Response.ID = postID
			response.Response.Body = body
			response.Response.Date = date
			response.Response.Title = title

			return c.JSON(200, response)

		}
	}
}