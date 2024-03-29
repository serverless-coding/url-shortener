package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/serverless-coding/url-shortener/service"
)

type ShortUrl struct {
	Url    string `json:"url,omitempty" description:"原始"`
	Short  string `json:"short,omitempty" description:"短链"`
	Target string `json:"target,omitempty" description:"目标"`
}

type ShortUrlRes struct {
	Code    int      `json:"code,omitempty" description:"code"`
	Message string   `json:"message,omitempty" description:"message"`
	Data    ShortUrl `json:"data" description:"结果"`
}

var engine *gin.Engine

func init() {
	r := gin.New()
	g := r.Group("/api")
	g.GET("/short", func(c *gin.Context) {
		res := ShortUrlRes{
			Code: 200,
		}
		if c.Query("url") != "" {
			short, err := service.NewUrlShortener().Short(c.Query("url"))
			if err != nil {
				fmt.Println(err)
			}
			schema := "https://"
			if c.Request.Host == "" || strings.Contains(c.Request.Host, "localhost") || strings.Contains(c.Request.Host, "127.0.0.1") {
				schema = "http://"
			}
			res.Data.Target = schema +
				c.Request.Host + "/api/url/redirect?short=" + short
			res.Data.Short = short
			res.Data.Url = c.Query("url")
		}

		c.JSON(http.StatusOK, res)
	})

	g.GET("/url", func(ctx *gin.Context) {
		short := ctx.Query("short")
		if short != "" {
			long, _ := service.NewUrlShortener().UrlFromShort(short)
			res := ShortUrlRes{
				Code:    200,
				Message: "success",
				Data: ShortUrl{
					Short:  short,
					Url:    long,
					Target: long,
				},
			}
			ctx.JSON(200, res)
			return
		}
		ctx.JSON(200, "not found")
	})

	g.GET("/url/redirect", func(ctx *gin.Context) {
		short := ctx.Query("short")
		if short != "" {
			long, _ := service.NewUrlShortener().UrlFromShort(short)

			ctx.Redirect(http.StatusTemporaryRedirect, long)
			return
		}
		ctx.JSON(200, "not found")
	})

	engine = r
}

func Handler(w http.ResponseWriter, r *http.Request) {

	engine.ServeHTTP(w, r)
}
