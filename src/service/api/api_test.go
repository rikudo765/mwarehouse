package api_test

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/Foundation-13/mwarehouse/src/service/api"
	"github.com/Foundation-13/mwarehouse/src/service/api/apimocks"
)

func TestStatus(t *testing.T) {
	a := newApi()

	a.r.GET("/media/123/status").
		SetDebug(true).
		Run(a.e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			data := []byte(r.Body.String())
			id, _ := jsonparser.GetString(data, "id")
			assert.Equal(t, "123", id)
		})
}

func TestUpload(t *testing.T) {
	a := newApi()

	a.r.POST("/media").
		SetDebug(true).
		SetFileFromPath([]gofight.UploadFile{
			{
				Path:    "/images/media.png",
				Name:    "file",
				Content: []byte("123"),
			},
		}).
		Run(a.e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

// helpers

type mocks struct {
	r *gofight.RequestConfig
	e *echo.Echo
	m *apimocks.Manager
}

func newApi() mocks {
	r := gofight.New()

	e := echo.New()
	m := &apimocks.Manager{}

	api.Assemble(e, m)

	return mocks{
		r: r,
		e: e,
		m: m,
	}
}
