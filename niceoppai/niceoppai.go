package niceoppai

import (
	"net/http"

	"github.com/yhat/scrape"
	"github.com/zier/niceoppai_notify/entity"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Service ...
type Service struct{}

func (n *Service) matchRow(node *html.Node) bool {
	if node.DataAtom == atom.Div && node.Attr[0].Val == "row" && node.Parent.Attr[0].Val == "wpm_pag mng_lts_chp grp" {
		return true
	}

	return false
}

func (n *Service) matchName(node *html.Node) bool {
	if node.DataAtom == atom.A && node.Attr[0].Val == "ttl" {
		return true
	}

	return false
}

func (n *Service) matchChapterTitle(node *html.Node) bool {
	if node.DataAtom == atom.B && node.Attr[0].Val == "val lng_" {
		return true
	}

	return false
}

func (n *Service) matchURL(node *html.Node) bool {
	if node.DataAtom == atom.A && node.Attr[0].Val == "lst" {
		return true
	}

	return false
}

func (n *Service) matchThumbnail(node *html.Node) bool {
	if node.DataAtom == atom.Img {
		return true
	}

	return false
}

// GetAllCartoonDetail ...
func (n *Service) GetAllCartoonDetail() (map[string]*entity.Cartoon, error) {
	cartoonDict := map[string]*entity.Cartoon{}

	resp, err := http.Get("http://www.niceoppai.net")
	if err != nil {
		return nil, err
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	elements := scrape.FindAll(root, n.matchRow)
	for _, element := range elements {
		cartoon := entity.NewCartoon()

		name, ok := scrape.Find(element, n.matchName)
		if !ok {
			continue
		}

		chapterTitle, ok := scrape.Find(element, n.matchChapterTitle)
		if !ok {
			continue
		}

		url, ok := scrape.Find(element, n.matchURL)
		if !ok {
			continue
		}

		thumbnail, ok := scrape.Find(element, n.matchThumbnail)
		if !ok {
			continue
		}

		cartoon.Name = scrape.Text(name)
		cartoon.ChapterTitle = scrape.Text(chapterTitle)
		cartoon.URL = scrape.Attr(url, "href")
		cartoon.Thumbnail = scrape.Attr(thumbnail, "src")

		cartoonDict[cartoon.Name] = cartoon
	}

	return cartoonDict, nil
}
