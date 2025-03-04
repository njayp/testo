package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/njayp/testo/pkg/manager"
)

type service struct {
	manager *manager.Manager
}

func newService() *service {
	return &service{
		manager: manager.NewManager(),
	}
}

type AddRequest struct {
	Num int32 `json:"num"`
}

type AddResponse struct {
	Count int32 `json:"count"`
}

func (s *service) HandleAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.HandleAddPost(w, r)
	case http.MethodGet:
		s.HandleAddGet(w, r)
	}
}

func (s *service) HandleAddPost(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// TODO
	}

	count := s.manager.Add(req.Num)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&AddResponse{
		Count: count,
	})
	if err != nil {
		// TODO
	}
}

func (s *service) HandleAddGet(w http.ResponseWriter, r *http.Request) {
	raw := r.URL.Query().Get("num")
	num, err := strconv.Atoi(raw)
	if err != nil {
		// TODO
	}

	count := s.manager.Add(int32(num))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&AddResponse{
		Count: count,
	})
	if err != nil {
		// TODO
	}
}
