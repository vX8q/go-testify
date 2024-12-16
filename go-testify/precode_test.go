package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	expectedResponse := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedResponse, responseRecorder.Body.String())

	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenCityIsInvalid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=2&city=spb", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	assert.Equal(t, "wrong city value", responseRecorder.Body.String())

	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	expectedResponse := "Мир кофе,Сладкоежка"
	assert.Equal(t, expectedResponse, responseRecorder.Body.String())

	assert.NotEmpty(t, responseRecorder.Body.String())

	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), 2)
}
