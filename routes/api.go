package routes

import (
	"github.com/labstack/echo"
	"github.com/sqrmyy/youtube-manager-go/web/api"
)

func Init(e *echo.Echo) {

	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
	}

}
