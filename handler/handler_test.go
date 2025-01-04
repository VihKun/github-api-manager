package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	Init()
}

func TestHelthCheck(t *testing.T) {
	r := gin.Default()
	r.GET("/health", HealthCheckHandler)

	req, _ := http.NewRequest("GET", "/health", nil)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	expected := `"message":"Service Running"`
	assert.Contains(t, rec.Body.String(), expected)
}

func TestListRepos(t *testing.T) {
	r := gin.Default()
	r.GET("/repos", ListReposHandler)

	req, _ := http.NewRequest("GET", "/repos", nil)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateRepo(t *testing.T) {
	r := gin.Default()
	r.POST("/repos", CreateRepoHandler)

	body := `{"name": "test-repo", "description": "test", "private": true}`

	req, _ := http.NewRequest("POST", "/repos", bytes.NewBufferString(body))

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestDeleteRepo(t *testing.T) {
	r := gin.Default()
	r.DELETE("/repos/:name", DeleteRepoHandler)

	req, _ := http.NewRequest("DELETE", "/repos/test-repo", nil)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
