package apiserver

import (
	"encoding/json"
	"net/http"
	"testapi/internal/app/store"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/profile", s.authMiddleware(s.handleProfile())).Methods("GET")
}

func (s *server) authMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Api-key")

		keyExists := s.store.Auth().KeyExists(apiKey)

		if keyExists == false {
			s.error(w, r, http.StatusForbidden, store.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func (s *server) handleProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")

		udata, err := s.store.User().GetUsersList(username)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, store.ErrInternalServerError)
			return
		}
		s.respond(w, r, http.StatusOK, udata)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"Error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
