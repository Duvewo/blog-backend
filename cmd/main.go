package main

import (
	"blog/cache"
	"blog/endpoints/posts"
	"blog/storage"
	"github.com/labstack/echo/v4"
)

func main()  {
	e := echo.New()

	redis, err := cache.New()

	if err != nil {
		panic(err)
	}

	db, err := storage.New()

	if err != nil {
		panic(err)
	}

	/* posts */
	e.GET("/posts/:id", posts.Get(redis, db))
	e.POST("/posts", posts.Post)
	e.PATCH("/posts/:id", posts.Patch)

	e.Start(":8080")
}
