package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetMovie(t *testing.T) {
	testCases := []struct {
		target    string
		expOutput MoviesData
		expError  error
	}{
		//{
		//	target: "/movie/$",
		//	expOutput: MoviesData{
		//		Error: "No movie found with the id"},
		//},
		{
			target: "/movie",
			expOutput: MoviesData{
				Code:   200,
				Status: "SUCCESS",
				Data: &Data{
					Movie: &Movies{
						Id:       1,
						Name:     "Silicon Valley",
						Genre:    "Comedy",
						Rating:   4.5,
						Plot:     "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.",
						Released: true,
					},
				},
				Error: "",
			},
		},
	}

	for _, tt := range testCases {
		r := httptest.NewRequest(http.MethodGet, tt.target, nil)
		w := httptest.NewRecorder()
		params := map[string]string{
			"id": "1",
		}
		r = mux.SetURLVars(r, params)
		getOneMovie(w, r)

		var movieResponse MoviesData
		res := w.Result()
		data, err := io.ReadAll(res.Body)

		if err != tt.expError {
			t.Errorf("Wrong Output.")
			return
		}
		err = json.Unmarshal(data, &movieResponse)
		if err != nil {
			t.Errorf("Wrong Output nil.")
			return
		}

		if !reflect.DeepEqual(movieResponse, tt.expOutput) {
			t.Fatalf("Wrong Output, Expected: %v, Got: %v", tt.expOutput, movieResponse)
		}
		fmt.Printf("output, Expected: %v, Got: %v\n", tt.expOutput, movieResponse)
	}
}
