package main

import (
	"testing"
)

func TestAdd(t *testing.T) {

	//tcs := []struct {
	//	desc  string
	//	movie MovieData
	//}{
	//	{"1", MovieData{"Harry Potter", 1, "Fiction", 10, "Lord Voldemort Wins", true}},
	//	{"2", MovieData{"Welcome", 1, "Comedy", 1, "Lord Voldemort always Wins", true}},
	//}
	//
	//for _, tt := range tcs {
	//	t.Run(tt.desc, func(t *testing.T) {
	//		finalJson, _ := json.Marshal(tt.movie)
	//		req := httptest.NewRequest("POST", "/movie", bytes.NewBuffer(finalJson))
	//		w := httptest.NewRecorder()
	//
	//		createOneMovie(w, req)
	//
	//		if !reflect.DeepEqual(req.Body, tt.movie) {
	//			t.Error("Test Case didn't pass at description ", tt.desc)
	//		}
	//	})
	//
	//}

}
