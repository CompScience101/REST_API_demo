package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bookid uint
var router = setupRouter()

func TestHealthRoute(t *testing.T) { //PASSING
	//router := setupRouter()

	w := httptest.NewRecorder()
	//request
	req, _ := http.NewRequest("HEAD", "/api/health/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
	//assert.Equal(t, "pong", w.Body.String())
}

func TestCreateBookRoute(t *testing.T) { //Ensure createbook test occurs before all other CRUD tests
	//router := setupRouter()

	w := httptest.NewRecorder()
	//create book
	var myjson Book
	myjson.Isbn = "alkeoidsjfkldsjalkf"
	myjson.Author = "Orlando Pacer"
	myjson.Title = "NBA"
	data, err := json.Marshal(myjson)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data) //convert myjson to bytes then io.Reader
	//request
	req, _ := http.NewRequest("POST", "/api/createbook", reader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	var resp Book
	json.Unmarshal([]byte(w.Body.Bytes()), &resp)
	bookid = resp.ID
	fmt.Println("bookid:", bookid, "resp.ID:", resp.ID)
	assert.Equal(t, myjson.Isbn, resp.Isbn)
	assert.Equal(t, myjson.Author, resp.Author)
	assert.Equal(t, myjson.Title, resp.Title)
	//assert.Equal(t, "pong", w.Body.String())
}

func TestGetBookRoute(t *testing.T) {
	//router := setupRouter()

	w := httptest.NewRecorder()
	//request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/getbook/%v", bookid), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var resp Book
	json.Unmarshal([]byte(w.Body.Bytes()), &resp)
	assert.Equal(t, bookid, resp.ID)
	fmt.Println("bookid:", bookid, " resp id:", resp.ID) //global variable test works!
	//assert.Equal(t, "pong", w.Body.String())
}
func TestGetBooksRoute(t *testing.T) {
	//router := setupRouter()

	w := httptest.NewRecorder()
	//request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/getbooks/"), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var resp []Book
	json.Unmarshal([]byte(w.Body.Bytes()), &resp)
	for _, service := range resp {
		if service.ID == bookid {
			assert.Equal(t, bookid, service.ID)
			fmt.Println("bookid:", bookid, " resp id:", service.ID) //global variable test works!
			break
		}
	}
	//assert.Equal(t, "pong", w.Body.String())
}

func TestUpdatePUTRoute(t *testing.T) {
	//router := setupRouter()

	w := httptest.NewRecorder()
	//create book
	var myjson Book
	myjson.Isbn = "itoghwljdlkdsklfjkljsda"
	myjson.Title = "the football league2"
	myjson.Author = "Star Trek2"
	data, err := json.Marshal(myjson)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data) //convert myjson to bytes then io.Reader
	//request
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/updatebook/%v", bookid), reader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var resp Book
	json.Unmarshal([]byte(w.Body.Bytes()), &resp)
	assert.Equal(t, myjson.Isbn, resp.Isbn)
	assert.Equal(t, myjson.Author, resp.Author)
	assert.Equal(t, myjson.Title, resp.Title)
	assert.Equal(t, bookid, resp.ID) //uint(10)
	//assert.Equal(t, "pong", w.Body.String())
}

func TestUpdatePATCHRoute(t *testing.T) {
	//router := setupRouter()

	w := httptest.NewRecorder()
	//create book
	var myjson BookPatch
	myjson.Title = "the football league"
	myjson.Author = "Star Trek"
	data, err := json.Marshal(myjson)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data) //convert myjson to bytes then io.Reader
	//request
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/api/updatebook/%v", bookid), reader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var resp Book
	json.Unmarshal([]byte(w.Body.Bytes()), &resp)
	assert.Equal(t, myjson.Author, resp.Author)
	assert.Equal(t, myjson.Title, resp.Title)
	assert.Equal(t, bookid, resp.ID) //uint(10)
	//fmt.Println("bookid:", bookid, " resp id:", resp.ID)//global variable test works!
	//assert.Equal(t, "pong", w.Body.String())
}

func TestDeleteRoute(t *testing.T) {
	//router := setupRouter()

	w := httptest.NewRecorder()
	//request
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/deletebook/%v", bookid), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	type myResponse struct {
		Success string `json:"success" binding:"required"`
	}
	var resp myResponse
	json.Unmarshal([]byte(w.Body.Bytes()), &resp)
	assert.Equal(t, fmt.Sprintf("deleted book id:%v", bookid), resp.Success)
	//fmt.Println("bookid:", bookid, " resp id:", resp.ID)//global variable test works!
	//assert.Equal(t, "pong", w.Body.String())
}

func TestCalculateTransaction1(t *testing.T) { // test transaction with $120.00
	//var test int64 = 120
	//points := calculateTransaction(test)
	//fmt.Printf("in test, rewarded client %d  points \n", points )
	//if points != 90 {
	//	t.Errorf("calculateTransaction(120.00) = %d; want 90", points)
	//}
}

/* func TestFormatCheck1 (t *testing.T){
	bool := formatCheck(120.00)
	if bool != true {
		t.Errorf("formatCheck(120.00) = %t; want true", bool)
	}
} */
