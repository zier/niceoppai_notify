package app

import (
	"fmt"
	"time"

	"github.com/zier/niceoppai_notify/config"
	"github.com/zier/niceoppai_notify/entity"
)

// Service ...
type Service struct {
	*config.Config
	CartoonDict    map[string]*entity.Cartoon
	SourceCartoons SourceCartoons
	LineNotify     LineNotify
}

// New ...
func New(c *config.Config, s SourceCartoons, ln LineNotify) *Service {
	service := &Service{
		Config:         c,
		SourceCartoons: s,
		LineNotify:     ln,
		CartoonDict:    map[string]*entity.Cartoon{},
	}

	return service
}

// InitCartoonDic ...
func (s *Service) InitCartoonDic() error {
	cartoonDic, err := s.SourceCartoons.GetAllCartoonDetail()
	if err != nil {
		return err
	}

	s.CartoonDict = cartoonDic
	return nil
}

// Start ...
func (s *Service) Start() {
	fmt.Println("Start Niceoppai Notify")

	for {
		s.FetchCartoon()
		time.Sleep(time.Minute * 5)
	}
}

// FetchCartoon ...
func (s *Service) FetchCartoon() error {
	newCartoonDict, err := s.SourceCartoons.GetAllCartoonDetail()
	if err != nil {
		return err
	}

	for cartoonName, newCartoon := range newCartoonDict {
		cartoon, exist := s.CartoonDict[cartoonName]
		if exist && cartoon.ChapterTitle == newCartoon.ChapterTitle {
			continue
		}
		s.CartoonDict[cartoonName] = newCartoon

		err = s.SendAllPush(newCartoon)
		if err != nil {
			return err
		}
	}

	return nil
}

// SendAllPush ...
func (s *Service) SendAllPush(cartoon *entity.Cartoon) error {
	for _, token := range s.AppConfig.Tokens {
		text := fmt.Sprintf("%s : %s -> %s", cartoon.Name, cartoon.ChapterTitle, cartoon.GetURL())
		if err := s.LineNotify.SendPush(token, text); err != nil {
			return err
		}
	}

	return nil
}
