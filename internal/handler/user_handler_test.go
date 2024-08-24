package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/br4tech/auth-nex/internal/handler"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert" // or your preferred assertion library
	"go.uber.org/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock.NewMockIUserUseCase(ctrl)
	userHandler := handler.NewUserHandler(mockUserUseCase)

	e := echo.New()
	reqBody := `{"username": "testuser", "email": "test@example.com", "password": "password"}`

	mockUserUseCase.EXPECT().Create(gomock.Any()).Return(nil, nil)

	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := userHandler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedJSON := `{"message":"Created user with successfully!!"}`

	var expectedData, actualData map[string]interface{}
	assert.NoError(t, json.Unmarshal([]byte(expectedJSON), &expectedData))
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &actualData))
	assert.Equal(t, expectedData, actualData)
}

func TestValidateAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock.NewMockIUserUseCase(ctrl)
	userHandler := handler.NewUserHandler(mockUserUseCase)

	e := echo.New()
	validToken := "your_valid_test_token"

	mockUserUseCase.EXPECT().ValidateAccessToken(validToken).Return(nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", validToken)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := userHandler.ValidateAccessToken(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedJSON := `{"message":"token valido"}`
	var expectedData, actualData map[string]interface{}
	assert.NoError(t, json.Unmarshal([]byte(expectedJSON), &expectedData))
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &actualData))
	assert.Equal(t, expectedData, actualData)
}
