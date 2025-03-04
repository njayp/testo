package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddPost(t *testing.T) {
	for _, test := range []struct {
		name string
		req  *AddRequest
		resp *AddResponse
	}{{
		name: "num is 2",
		req: &AddRequest{
			Num: 2,
		},
		resp: &AddResponse{
			Count: 2,
		}},
		{
			name: "num is 4",
			req: &AddRequest{
				Num: 4,
			},
			resp: &AddResponse{
				Count: 4,
			}},
	} {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.req)
			reader := bytes.NewReader(data)
			req, err := http.NewRequest("POST", "/add", reader)
			if err != nil {
				t.Fatal(err)
			}

			s := newService()
			handler := http.HandlerFunc(s.HandleAdd)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			var resp AddResponse
			err = json.NewDecoder(w.Body).Decode(&resp)
			if err != nil {
				t.Fatal(err)
			}

			if resp.Count != test.resp.Count {
				t.Errorf("expected %v, got %v", test.resp.Count, resp.Count)
			}
		})
	}
}

func TestAddGet(t *testing.T) {
	for _, test := range []struct {
		name string
		num  int
		resp *AddResponse
	}{{
		name: "num is 2",
		num:  2,
		resp: &AddResponse{
			Count: 2,
		}},
		{
			name: "num is 4",
			num:  4,
			resp: &AddResponse{
				Count: 4,
			}},
	} {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("/add?num=%v&", test.num), nil)
			if err != nil {
				t.Fatal(err)
			}

			s := newService()
			handler := http.HandlerFunc(s.HandleAdd)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			var resp AddResponse
			err = json.NewDecoder(w.Body).Decode(&resp)
			if err != nil {
				t.Fatal(err)
			}

			if resp.Count != test.resp.Count {
				t.Errorf("expected %v, got %v", test.resp.Count, resp.Count)
			}
		})
	}
}
