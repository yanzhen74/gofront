package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/recover"
	"github.com/yanzhen74/gofront/src/inits/parse"
	"github.com/yanzhen74/gofront/src/supports"
)

//所有路由
func Hub(app *iris.Application) {
	preSettring(app)
	var main = corsSetting(app)
	HomeHub(main)
	CCTC_Hub(main)
	Protocal_Hub(main)
}

func corsSetting(app *iris.Application) (main iris.Party) {
	var (
		crs context.Handler
	)

	crs = cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //允许通过的主机名称
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
		//AllowCredentials: true,
	})

	/* 定义路由 */
	main = app.Party("/", crs).AllowMethods(iris.MethodOptions)
	//main.Use(middleware.ServeHTTP)
	//main := app.Party("/")
	//main.Use(middleware.ServeHTTP)

	return main
}

func preSettring(app *iris.Application) {
	app.Logger().SetLevel(parse.AppConfig.LogLevel)
	logger, close := supports.NewRequestLogger()
	defer close()
	app.Use(
		recover.New(),
		logger, // 记录请求
		//middleware.ServeHTTP
	)

	// ---------------------- 定义错误处理 ------------------------
	app.OnErrorCode(iris.StatusNotFound, logger, func(ctx iris.Context) {
		supports.Error(ctx, iris.StatusNotFound, supports.NotFound, nil)
	})
	app.OnErrorCode(iris.StatusInternalServerError, logger, func(ctx iris.Context) {
		supports.Error(ctx, iris.StatusInternalServerError, supports.StatusInternalServerError, nil)
	})
	//app.OnErrorCode(iris.StatusForbidden, customLogger, func(ctx iris.Context) {
	//	ctx.JSON(utils.Error(iris.StatusForbidden, "权限不足", nil))
	//})
	//捕获所有http错误:
	//app.OnAnyErrorCode(customLogger, func(ctx iris.Context) {
	//	//这应该被添加到日志中，因为`logger.Config＃MessageContextKey`
	//	ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
	//	ctx.JSON(utils.Error(500, "服务器内部错误", nil))
	//})
}
