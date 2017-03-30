package main

import (
	"github.com/zier/niceoppai_notify/app"
	"github.com/zier/niceoppai_notify/config"
	"github.com/zier/niceoppai_notify/linenotify"
	"github.com/zier/niceoppai_notify/niceoppai"
)

func main() {
	c := config.New()
	if err := c.ReadCLIParams(); err != nil {
		panic(err)
	}

	s := app.New(c, &niceoppai.Service{}, &linenotify.Service{})
	if err := s.InitCartoonDic(); err != nil {
		panic(err)
	}

	s.Start()
}
