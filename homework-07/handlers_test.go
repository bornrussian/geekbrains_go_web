package main

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"context"
)

func TestHandlers(t *testing.T) {
	lg := logrus.New()
	lg.SetLevel(0)
	serv := NewServer(context.Background(), lg).Start(":8080")

	// Генерируем случайный пост
	uuid, _ := uuid.NewRandom()
	testuuid := "test-" + uuid.String()
	post := &Post{
		Autor: testuuid,
		Date: testuuid,
		Header: testuuid,
		Content: testuuid,
	}

	// Пробуем вставить его в базу данных
	if err := post.Insert(context.TODO(), serv.db); err != nil {
		t.Error("Insert failed", err)
	} else {
		// Если получилось вставить запись в базу, то пробуем увидеть её в HTML-ответе от сервера
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serv.wwwIndex)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("status code isn't 200, got: %v", status)
		}
		if !strings.Contains(rr.Body.String(),testuuid) {
			t.Errorf("HTML Body is not contains testuuid %s",testuuid)
		}

		// Удаляем тестовый пост за собой, чтобы не мусорить
		post.Delete(context.TODO(), serv.db)

		serv.Stop()
	}
}