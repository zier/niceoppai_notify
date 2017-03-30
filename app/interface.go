package app

import "github.com/zier/niceoppai_notify/entity"

// SourceCartoons ...
type SourceCartoons interface {
	GetAllCartoonDetail() (map[string]*entity.Cartoon, error)
}

// LineNotify ...
type LineNotify interface {
	SendPush(token, text string) error
}
