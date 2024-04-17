package router

import (
	"net/http"
	"time"

	"ticket/internal/api"
	"ticket/internal/middleware"
	"ticket/internal/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

type Router struct {
	router *gin.Engine
	store  storage.Store
}

func New(s storage.Store, secretKey string) *Router {
	r := &Router{
		router: gin.Default(),
		store:  s,
	}

	r.initSessions(secretKey)
	r.setCors()
	r.initRoutes()
	return r
}

func (r *Router) initSessions(secretKey string) {
	sessionStore := cookie.NewStore([]byte(secretKey))
	r.router.Use(sessions.Sessions("session", sessionStore))
}

func (r *Router) setCors() {
	config := cors.DefaultConfig()
	config.AllowOriginFunc = func(origin string) bool { return true }
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	r.router.Use(cors.New(config))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r *Router) initRoutes() {
	// API Router Group
	apiGroup := r.router.Group("/api")

	unauthorizedApiGroup := apiGroup.Group("")

	authorizedApiGroup := apiGroup.Group("")
	authorizedApiGroup.Use(middleware.AuthRequired)

	// Auth controllers
	{
		unauthorizedApiGroup.Handle(http.MethodPost, "/auth/register", api.RegisterUser(r.store))
		unauthorizedApiGroup.Handle(http.MethodPost, "/auth/login", api.LoginUser(r.store))
		authorizedApiGroup.Handle(http.MethodPost, "/auth/logout", api.LogoutUser())
	}

	// User controllers
	{
		authorizedApiGroup.Handle(http.MethodGet, "/user/me", api.GetUserSelfInfo(r.store))
		authorizedApiGroup.Handle(http.MethodGet, "/user/tickets", api.GetUserSelfTickets(r.store))
	}

	// Ticket controllers
	{
		authorizedApiGroup.Handle(http.MethodGet, "/ticket/:id", api.FindTicketByID(r.store))
		authorizedApiGroup.Handle(http.MethodPost, "/ticket/buy", api.BuyTicket(r.store))
		authorizedApiGroup.Handle(http.MethodPost, "/ticket/eat", api.EatTicket(r.store))
	}

	// Interview controllers
	{
		authorizedApiGroup.Handle(http.MethodPost, "/interview", api.CheckLuckForInterview(r.store))
	}
}
