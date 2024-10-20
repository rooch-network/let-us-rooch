package routes

import (
	"gohub/internal/controller/app"
	"gohub/internal/routes/middlewares"
	"gohub/pkg/fileP"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler(router)

	// 静态资源路由
	router.Static("/static", fileP.GetRootPath())
}

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	appGroup := r.Group("/app")
	appRoutes(appGroup)

	adminGroup := r.Group("/admin")
	adminRoutes(adminGroup)
}

func appRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/order")
	ocl := new(app.OrderController)
	authGroup.POST("/save", ocl.Save)
	authGroup.POST("/execute", ocl.Execute)
	authGroup.GET("/list", ocl.List)

	seedGroup := r.Group("/seed")
	scl := new(app.SeedController)
	seedGroup.GET("/randomUsableSeed", scl.RandomUsableSeed)
	seedGroup.GET("/usedTempSeed", scl.UsedTempSeed)
	seedGroup.GET("/address", scl.GetSeedsByAddress)
	seedGroup.GET("/seedHTML/:hSeed", scl.SeedHtml)
}

func adminRoutes(r *gin.RouterGroup) {}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.Cors(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
