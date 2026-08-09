package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	v1 "github.com/ComputerScienceHouse/marks/internal/api/v1"
	"github.com/ComputerScienceHouse/marks/internal/database"
	"github.com/ComputerScienceHouse/marks/internal/logging"
	"github.com/ComputerScienceHouse/marks/internal/models"
	"github.com/ComputerScienceHouse/marks/internal/redis"

	cshauth "github.com/computersciencehouse/csh-auth/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/ComputerScienceHouse/marks/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed web/dist
var distFS embed.FS

var indexBuffer []byte

func serveIcon(c *gin.Context, icon []byte) {
	c.Data(http.StatusOK, "image/png", icon)
}

func serveFrontend(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", indexBuffer)
}

func chooseRouter() *gin.Engine {
	if os.Getenv("DEV_MODE") == "true" {
		return gin.Default()
	}

	return gin.New()
}

func createFrontend(frontend *gin.RouterGroup) error {
	dist, err := fs.Sub(distFS, "web/dist")
	if err != nil {
		return fmt.Errorf("failed to initialize dist for the app %s", err)
	}

	assets, err := fs.Sub(dist, "assets")
	if err != nil {
		return fmt.Errorf("failed to initialize assets for the app %s", err)
	}

	index, err := fs.ReadFile(dist, "index.html")
	if err != nil {
		return fmt.Errorf("failed to load index %s", err)
	}
	indexBuffer = index

	icon, err := fs.ReadFile(dist, "favicon.png")
	if err != nil {
		return fmt.Errorf("failed to load favicon %s", err)
	}

	frontend.Use(cors.Default())
	frontend.StaticFS("/assets", http.FS(assets))

	frontend.GET("/favicon.ico", func(ctx *gin.Context) {
		serveIcon(ctx, icon)
	})

	frontend.GET("/favicon.svg", func(ctx *gin.Context) { // just do both and make my life easy
		serveIcon(ctx, icon)
	})

	frontend.GET("/favicon.png", func(ctx *gin.Context) { // theres a third now ig.
		serveIcon(ctx, icon)
	})

	frontend.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	frontend.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	return nil
}

// @title CSH Marks
// @version 1.0
// @description How Fast Can You Eat?
// @BasePath /api/v1
func main() {
	if err := redis.InitRedis(); err != nil {
		//its dead, but thats fine because we dont neeeeed it
		logging.Logger.Warnf("failed to initialize redis! continuing without %s", err)
	}

	if err := database.InitDatabase(); err != nil {
		//yeah we need ts gng :pray:
		// logging.Logger.Fatalf("failed to initialize database for the app %s", err) ignore for now
	}

	hostUrl := os.Getenv("SERVER_HOST")
	auth, err := cshauth.Init(
		os.Getenv("AUTH_OIDC_ID"),
		os.Getenv("AUTH_OIDC_SECRET"),
		hostUrl,
		hostUrl+"/auth/login",
		hostUrl+"/auth/callback",
		[]string{"profile", "email", "groups"},
	)
	if err != nil {
		logging.Logger.Fatalf("failed to initialize csh auth for the app %s", err)
	}

	router := chooseRouter()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNoContent, gin.H{})
	})

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{ //shut the f up
		SkipPaths: []string{"/health"},
	}))

	frontend := router.Group("")
	api := router.Group("/api")
	v1.SetRoutes(api)

	if os.Getenv("DEV_MODE") == "true" {
		logging.Logger.Warn("RUNNING IN DEV MODE!")
		router.NoRoute(createViteProxy())
	} else {
		if err := createFrontend(frontend); err != nil {
			logging.Logger.Fatalf("failed to create frontend %s", err)
		}

		router.NoRoute(auth.CookieMiddleware(), func(c *gin.Context) {
			reqURL := c.Request.URL.String()

			if strings.Contains(reqURL, "api/") {
				c.JSON(http.StatusNotFound, models.ErrorResponse{
					Message: "API not found!",
				})
				return
			}
			serveFrontend(c)
		})

	}

	router.Run(":8080")
}
