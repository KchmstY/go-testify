package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrectRequest(t *testing.T) {
	//сервис возвращает код ответа 200
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверяем статус 200
	require.Equal(t, responseRecorder.Code, http.StatusOK)

	// проверяем не пустое ли тело
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerCityIsMissing(t *testing.T) {
	// Сервис возвращает код ответа 400 и ошибку
	req := httptest.NewRequest("GET", "/cafe?count=23&city=piter", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверяем ответ 400
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	// проверяем тело
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	// формируем запрос с count > 4
	req := httptest.NewRequest("GET", "/cafe?count=12&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// проверяем статус 200
	require.Equal(t, responseRecorder.Code, http.StatusOK)

	// проверяем что вернулось totalCount
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)
}
