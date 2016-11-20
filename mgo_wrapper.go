package main

import (
	"net/http"
	"os"

	"github.com/gorilla/context"

	mgo "gopkg.in/mgo.v2"
)

type Server struct {
	dbSession *mgo.Session
}

func NewServer() (*Server, error) {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return &Server{dbSession: session}, nil
}
func (s *Server) Close() {
	s.dbSession.Close()
}

func (s *Server) WithData(ourFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionCopy := s.dbSession.Copy()
		defer sessionCopy.Close()

		// lets save some objects, shall we
		context.Set(r, "sessionCopy", sessionCopy)

		// now we call our handler :)
		ourFunction(w, r)
	}
}
