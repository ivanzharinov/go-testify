package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerCorrectResponse(t *testing.T) { //Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, body)
}

func TestMainHandlerWhenCityNotMatch(t *testing.T) { //Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
	city := "moscow"
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moskek", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	reqCity := req.URL.Query().Get("city")
	require.Equal(t, reqCity, city, "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) { // Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.

	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	reqCount := req.URL.Query().Get("count")

	assert.Len(t, reqCount, len(list), body)
}
