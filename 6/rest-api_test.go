package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	initDatabase()

	m.Run()
}

func TestSetProduct(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte("{\"kodeProduk\":\"a\", \"kuantitas\":10}")
	req, err = http.NewRequest("POST", "/product/set", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(setProduct)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "{\"message\":\"Product data has set to database\"}"
	assert.Equal(t, expected, rr.Body.String())

}

func TestReadProduct(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	req, err = http.NewRequest("GET", "/product/read?kodeProduk=a", nil)
	if err != nil {
		t.Error(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(readProduct)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "{\"message\":\"Product data available in database\",\"kodeProduk\":\"a\",\"kuantitas\":10}"
	assert.Equal(t, expected, rr.Body.String())
}

func TestDeleteProduct(t *testing.T) {
	var rr *httptest.ResponseRecorder
	var req *http.Request
	var handler http.HandlerFunc
	var err error

	jsonStr := []byte("{\"kodeProduk\":\"p\"}")
	req, err = http.NewRequest("POST", "/product/delete", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(deleteProduct)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	expected := "{\"message\":\"Data not available\"}"
	assert.Equal(t, expected, rr.Body.String())
}
