package main

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/go-redis/redis"
	"github.com/zier/niceoppai_notify/app"
	"github.com/zier/niceoppai_notify/linenotify"
	"github.com/zier/niceoppai_notify/niceoppai"
	"github.com/zier/niceoppai_notify/tokenstore"
)

func main() {
	r := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)

	ts, err := tokenstore.New(r)
	if err != nil {
		panic(err)
	}

	s := app.New(ts, &niceoppai.Service{}, &linenotify.Service{})
	if err := s.InitCartoonDic(); err != nil {
		panic(err)
	}

	go s.Start()

	gin.SetMode("release")
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.GET("/", s.Index)
	router.POST("/token", s.Token)
	router.Run(":8080")
}
