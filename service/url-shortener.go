package service

import "github.com/serverless-coding/url-shortener/util/base62"

const (
	_baseUrl = "https://"
)

type UrlShortener interface {
	Short(uri string) (string, error)
	UrlFromShort(short string) (string, error)
}

type _defaultShortener struct {
}

func (s *_defaultShortener) Short(uri string) (string, error) {
	return base62.Encode(0)
}

func (s *_defaultShortener) UrlFromShort(short string) (string, error) {
	return "", nil
}
