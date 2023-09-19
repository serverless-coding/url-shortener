package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/serverless-coding/url-shortener/service"
)

type ShortUrlRes struct {
	Code    int    `json:"code,omitempty" description:"code"`
	Message string `json:"message,omitempty" description:"message"`
	Data    string `json:"data" description:"断链"`
}

// func ShortUrl(w http.ResponseWriter, r *http.Request) {
// 	res := ShortUrlRes{}
// 	if r.URL.Query().Has("url") {
// 		short, err := service.NewUrlShortener().Short(r.URL.Query().Get("url"))
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		res.Data = r.Host + "/short/" + short
// 	}

// 	jv, _ := json.Marshal(res)
// 	fmt.Fprint(w, string(jv))
// }

var engine *gin.Engine

func init() {
	r := gin.Default()

	r.GET("/api/short", func(c *gin.Context) {
		res := ShortUrlRes{}
		if c.Query("url") != "" {
			short, err := service.NewUrlShortener().Short(c.Query("url"))
			if err != nil {
				fmt.Println(err)
			}
			res.Data = c.Request.Host + "/api/url/" + short
		}

		c.JSON(http.StatusOK, res)
	})

	r.GET("/api/url", func(ctx *gin.Context) {
		if ctx.Query("url") != "" {
			long, err := service.NewUrlShortener().UrlFromShort(ctx.Query("url"))
			if err != nil {
				ctx.AbortWithError(400, err)
				return
			}
			ctx.Redirect(301, long)
		}
		ctx.AbortWithError(400, errors.New("url required"))
	})

	engine = r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	engine.ServeHTTP(w, r)
	engine.Run()
}
