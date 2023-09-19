package service

import (
	"github.com/pkg/errors"
	"github.com/serverless-coding/url-shortener/util/base62"
	"github.com/serverless-coding/url-shortener/util/id"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	_workIdKey        = "url-shortener:workId"
	_urlContentPrefix = "url-shortener:kv:"
	_defaultExpire    = 60 * 60 * 24 * 7 // 7 days
)

type UrlShortener interface {
	Short(uri string) (string, error)
	UrlFromShort(short string) (string, error)
}

type _defaultShortener struct {
	snowflakeId *id.Snowflake
	r           *redis.Redis
}

func (d *_defaultShortener) init() error {
	// TODO: serverless 这种,如何确定workid,加心跳? 默认
	r, err := id.NewRedis()
	if err != nil {
		return err
	}
	d.r = r
	wordkid, _ := r.Incr(_workIdKey)
	s, err := id.NewSnowflake(0, wordkid, false)
	if err != nil {
		return err
	}
	d.snowflakeId = s
	return nil
}

func (d *_defaultShortener) Short(uri string) (string, error) {
	if d.snowflakeId == nil {
		return "", errors.Errorf("not initialized")
	}
	id := d.snowflakeId.Id()
	short, err := base62.Encode(id)
	if err != nil {
		return "", err
	}
	_, _ = d.r.SetnxEx(_urlContentPrefix+short, uri, _defaultExpire)
	return short, nil
}

func (d *_defaultShortener) UrlFromShort(short string) (string, error) {
	if d.r == nil {
		return "", errors.Errorf("not initialized")
	}
	v, err := d.r.Get(_urlContentPrefix + short)
	if err != nil {
		return "", errors.New("not found")
	}
	return v, nil
}

func NewUrlShortener() UrlShortener {
	s := &_defaultShortener{}
	s.init()
	return s
}
