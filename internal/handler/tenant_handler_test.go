package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTenant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTenantUseCase := mock.NewMockITenantUseCase(ctrl)
	tenantHandler := NewTenantHandler(mockTenantUseCase)

	e := echo.New()
	reqBody := `{
    "name": "app-008",
      "company": {
        "legal_name": "Maria e Amanda Transportes Ltda",
        "trade_name": "MA Transportes",
        "document": "56.057.455/0001-06",
        "state_registration": "025.129.841.534",
        "adress": {
            "street": "Rua José Gomez Diaz",
            "number": "966",
            "complement": "Bloco A",
            "district": "Vila Suiça",
            "city": "Pindamonhangaba",
            "state": "SP",
            "zip_code": "12405-230"
        },
        "type": "ME"
      },
      "user":{
        "name": "Guilherme Silva",
        "cpf": "359.162.208-71",
        "email": "guilherme.silva@mariaeamandatransportesltda.com.br",
        "password": "123456"
      }
}`

	mockTenantUseCase.EXPECT().CreateTenantWithCompanyAndAdmin(gomock.Any()).Return(&domain.Tenant{}, nil)

	req := httptest.NewRequest(http.MethodPost, "/tenant", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := tenantHandler.CreateTenant(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedJSON := `{"message":"Created tenant with sucessfully"}`
	var expectedData, actualData map[string]interface{}
	assert.NoError(t, json.Unmarshal([]byte(expectedJSON), &expectedData))
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &actualData))
	assert.Equal(t, expectedData, actualData)
}
