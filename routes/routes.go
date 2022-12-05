package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"query/api/v1"
	"query/middleware"
	"query/utils"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin","web/admin/dist/index.html")
	p.AddFromFiles("guest","web/guest/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()

	r.SetTrustedProxies(nil)
	r.HTMLRender = createMyRender()

	//将http请求重定向到https
	//r.Use(middleware.Tls())
	//打印日志
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	// CORS跨域请求
	r.Use(middleware.Cors())

	r.Static("/static","./web/guest/dist/static")
	r.Static("/admin","./web/admin/dist")
	r.StaticFile("/favicon.ico","/web/guest/dist/favicon.ico")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200,"guest",nil)
	})
	r.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(200,"admin",nil)
	})

	// 后台路由
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块路由接口
		auth.POST("user/add", v1.AddUser)
		auth.DELETE("user/:id", v1.DeleteUserById)
		auth.PUT("user/:id", v1.EditUserById)
		auth.GET("user/:id", v1.FindUserById)
		auth.GET("user/list",v1.GetUserListByUsername)

		// 修改密码
		auth.PUT("repassword/:id",v1.RePassword)

		// Swindler模块路由接口
		auth.POST("swindler/add",v1.AddSwindler)
		auth.DELETE("swindler/:id",v1.DeleteSwindlerById)
		auth.PUT("swindler/:id",v1.EditSwindlerById)
		auth.GET("swindler/:id",v1.FindSwindlerById)
	}
	unauth := r.Group("api/v1")
	{
		// 登录模块路由接口
		unauth.POST("login", v1.Login)

		// Swindler模块路由接口
		unauth.GET("swindler/list",v1.GetSwindlerListByTieba)
	}

	if utils.UseTls {
		r.Use(middleware.Tls())
		r.RunTLS(utils.HttpPort,utils.CrtPath,utils.KeyPath)
	} else {
		r.Run(utils.HttpPort)
	}
}