package apiserver

import (
	"io/ioutil"
	"net/http"
	"time"

	"modality/internal/app/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *gin.Engine
	logger *logrus.Logger
	store  store.Store
}

// func newServer
func newServer(store store.Store, config *Config) *server {

	// after finish debug
	if config.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &server{
		router: gin.New(),
		logger: logrus.New(),
		store:  store,
	}

	// disable logs to console
	if !config.ConsoleLog {
		s.logger.SetOutput(ioutil.Discard)
	}
	s.logger.SetReportCaller(false)

	s.configureRouter()
	s.logger.Info("start server")

	return s
}

func (s *server) configureRouter() {

	// for get real IP by nginx
	s.router.ForwardedByClientIP = false

	// add parameter requestid for logging current request
	s.router.Use(requestid.New())
	s.router.Use(cors.Default())
	s.router.Use(static.Serve("/", static.LocalFile("./static", false)))
	s.router.Use(s.logRequest())
	s.router.PUT("/text", s.handleTextAdd())
	s.router.GET("/types", s.handleTypesGet())
	s.router.GET("/langs", s.handleLangsGet())
	s.router.POST("/texts", s.handlePageTextsGet())
	s.router.GET("/text", s.handleCurTextGet())
	s.router.DELETE("/text", s.handleCurTextDelete())
	s.router.PATCH("/text", s.handleCurTextUpdate())
	s.router.PUT("/modality", s.handleModalityAdd())
	s.router.GET("/modality", s.handleModalityGet())
	s.router.DELETE("/modality", s.handleCurModalityDelete())
	s.router.PATCH("/modality", s.handleModalityUpdate())
	s.router.GET("/modalities", s.handleModalitiesGet())
	s.router.POST("/statistic", s.handleStatisticLanguagesGet())

}

// logRequest logging requests
func (s *server) logRequest() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": ctx.Request.Header.Get("X-Real-IP"),
			"request_id":  requestid.Get(ctx),
		})
		logger.Infof("started %s %s", ctx.Request.Method, ctx.Request.RequestURI)

		start := time.Now()

		// before request
		ctx.Next()
		// after request

		code := ctx.Writer.Status()

		var level logrus.Level
		switch {
		case code >= 500:
			level = logrus.ErrorLevel
		case code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}

		logger.Logf(
			level,
			"completed with %d %s in %v",
			code,
			http.StatusText(code),
			time.Now().Sub(start),
		)
	}
}

// respondWithError send error
func (s *server) respondWithError(ctx *gin.Context, code int, err error) {
	ctx.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
}
