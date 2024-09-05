package api

import (
	"encoding/json"
	"net/http"

	"github.com/Flake-chat/Flake-auth/auth"
	"github.com/Flake-chat/Flake-auth/model"
	"github.com/Flake-chat/Flake-auth/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
	auth   *auth.Auth
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.confireLog(); err != nil {
		return err
	}
	s.logger.Info("Start server")
	s.configureRouter()
	s.configureAuth()
	if err := s.configureStore(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.Addr, s.router)
}

func (s *ApiServer) confireLog() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/reg", s.register())
	s.router.HandleFunc("/login", s.login())
	s.router.HandleFunc("/token", s.token_Check())
}

func (s *ApiServer) configureAuth() {
	stl := auth.New(s.config.Auth)
	s.auth = stl
}

func (s *ApiServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}
	s.logger.Info("Database Conneted")
	s.store = st

	return nil
}

func (s *ApiServer) register() http.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Name:     req.Name,
			Fullname: req.Fullname,
			Email:    req.Email,
			Password: req.Password,
		}
		if _, err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			s.logger.Error(err, " ", "Name: ", u.Name, " Email: ", u.Email, " Password: ", u.Password)
			return
		}
		s.logger.Info("Create user ID: ", u.ID)
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *ApiServer) login() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		type token struct {
			Token string `json:"token"`
		}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Find(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		t, err := s.auth.Createtoken(u)

		tokn := &token{
			Token: t,
		}

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, tokn)
	}
}

func (s *ApiServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *ApiServer) token_Check() http.HandlerFunc {
	type request struct {
		Jwt string `json:"token"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		type claims struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			EXP  int64  `json:"time"`
		}
		user, err := s.auth.Get_User(req.Jwt)

		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		u := claims{
			ID:   user.ID,
			Name: user.Name,
			EXP:  user.Exp,
		}

		s.respond(w, r, http.StatusAccepted, u)
	}
}

func (s *ApiServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
