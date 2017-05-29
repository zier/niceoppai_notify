package app

import "github.com/zier/niceoppai_notify/entity"

// SourceCartoons ...
type SourceCartoons interface {
	GetAllCartoonDetail() (map[string]*entity.Cartoon, error)
}

// LineNotify ...
type LineNotify interface {
	SendPush(token, text string, thumbnail string) error
}

// TokenStore ...
type TokenStore interface {
	Save(token string) error
	Remove(token string) error
	All() ([]string, error)
}
