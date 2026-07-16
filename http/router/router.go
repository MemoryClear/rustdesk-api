package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lejianwen/rustdesk-api/v2/global"
	"github.com/lejianwen/rustdesk-api/v2/http/controller/web"
)

// lostPage is served instead of the web admin interface when it is disabled.
const lostPage = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>404</title>
  <style>
    html,body{height:100%;margin:0}
    body{display:flex;align-items:center;justify-content:center;
      background:#0f172a;color:#e2e8f0;
      font-family:system-ui,-apple-system,"Segoe UI",Roboto,sans-serif}
    .box{text-align:center}
    .code{font-size:96px;font-weight:700;line-height:1;color:#38bdf8}
    .msg{margin-top:12px;font-size:20px;opacity:.85}
  </style>
</head>
<body>
  <div class="box">
    <div class="code">404</div>
    <div class="msg">You lost your way</div>
  </div>
</body>
</html>`

func WebInit(g *gin.Engine) {
	// web admin interface disabled: block the entry and admin static files,
	// and show a static "you lost your way" page instead.
	if global.Config.App.AdminWeb != 1 {
		g.GET("/", func(c *gin.Context) {
			c.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte(lostPage))
		})
		g.GET("/_admin/*any", func(c *gin.Context) {
			c.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte(lostPage))
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
