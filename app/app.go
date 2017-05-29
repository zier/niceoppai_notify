package app

import (
	"fmt"
	"regexp"
	"time"

	"github.com/zier/niceoppai_notify/entity"
)

// Service ...
type Service struct {
	TokenStore     TokenStore
	CartoonDict    map[string]*entity.Cartoon
	SourceCartoons SourceCartoons
	LineNotify     LineNotify
}

// New ...
func New(ts TokenStore, s SourceCartoons, ln LineNotify) *Service {
	service := &Service{
		TokenStore:     ts,
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
		cartoonNewChapters, err := s.FetchCartoonNewChapter()
		if err != nil {
			fmt.Println(err)
		}

		for cartoonName, newCartoon := range cartoonNewChapters {
			s.CartoonDict[cartoonName] = newCartoon
			err = s.SendAllPush(newCartoon)
			if err != nil {
				fmt.Println(err)
			}
		}

		time.Sleep(time.Minute * 5)
	}
}

// FetchCartoonNewChapter ...
func (s *Service) FetchCartoonNewChapter() (map[string]*entity.Cartoon, error) {
	newCartoonDict, err := s.SourceCartoons.GetAllCartoonDetail()
	if err != nil {
		return nil, err
	}

	cartoonNewChapters := map[string]*entity.Cartoon{}
	for cartoonName, newCartoon := range newCartoonDict {
		cartoon, exist := s.CartoonDict[cartoonName]
		if exist && cartoon.ChapterTitle == newCartoon.ChapterTitle {
			continue
		}

		cartoonNewChapters[cartoonName] = newCartoon
	}

	return cartoonNewChapters, nil
}

// SendAllPush ...
func (s *Service) SendAllPush(cartoon *entity.Cartoon) error {
	tokens, err := s.TokenStore.All()
	if err != nil {
		return err
	}

	reg, err := regexp.Compile("_[0-9]+")
	if err != nil {
		return err
	}

	for _, token := range tokens {
		text := fmt.Sprintf("%s : %s -> %s", cartoon.Name, cartoon.ChapterTitle, cartoon.GetURL())
		thumbnail := reg.ReplaceAllString(cartoon.Thumbnail, "_200")
		if err := s.LineNotify.SendPush(token, text, thumbnail); err != nil {
			if err.Error() == "invalid token" {
				s.TokenStore.Remove(token)
			}
		}
	}

	return nil
}
