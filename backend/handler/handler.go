package handler

import (
	"net/http"

	"github.com/4hm3d92/community-app/backend/db"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var dbInstance db.Database

var sm *scs.SessionManager

func NewHandler(db db.Database, sessionManager scs.SessionManager) http.Handler {
	router := chi.NewRouter()
	dbInstance = db

	sm = &sessionManager
	/*
		router.Use(cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			// ExposedHeaders:   []string{"Link"},
			// AllowCredentials: false,
			// MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))
	*/

	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	router.Route("/api/members", members)
	router.Route("/api/users", users)
	router.Route("/api/payments", payments)

	//workDir, _ := os.Getwd()
	//filesDir := http.Dir(filepath.Join(workDir, "data"))
	//filesDir := http.Dir(filepath.Join(workDir, "public"))
	//FileServer(router, "/", filesDir)

	router.Handle("/*", http.FileServer(http.Dir("/var/www/community/public")))

	//router.

	return sm.LoadAndSave(router)
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
