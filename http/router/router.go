package router

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lejianwen/rustdesk-api/v2/global"
	"github.com/lejianwen/rustdesk-api/v2/http/controller/web"
)

// forbiddenHTML is the static page served instead of the web admin interface
// when it is disabled (RUSTDESK_API_APP_ADMIN_WEB != 1).
//
//go:embed forbidden.html
var forbiddenHTML []byte

func WebInit(g *gin.Engine) {
	// web admin interface disabled: block the entry and admin static files,
	// and show the static forbidden.html page instead.
	if global.Config.App.AdminWeb != 1 {
		g.GET("/", func(c *gin.Context) {
			c.Data(http.StatusForbidden, "text/html; charset=utf-8", forbiddenHTML)
		})
		g.GET("/_admin/*any", func(c *gin.Context) {
			c.Data(http.StatusForbidden, "text/html; charset=utf-8", forbiddenHTML)
		})
		return
	}

	i := &web.Index{}
	g.GET("/", i.Index)

	if global.Config.App.WebClient == 1 {
		g.GET("/webclient-config/index.js", i.ConfigJs)
	}

	if global.Config.App.WebClient == 1 {
		g.StaticFS("/webclient", http.Dir(global.Config.Gin.ResourcesPath+"/web"))
		g.StaticFS("/webclient2", http.Dir(global.Config.Gin.ResourcesPath+"/web2"))
	}
	g.StaticFS("/_admin", http.Dir(global.Config.Gin.ResourcesPath+"/admin"))
}
