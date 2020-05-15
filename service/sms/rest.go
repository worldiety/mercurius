package sms

import (
	"context"
	"fmt"
	"github.com/golangee/uuid"
	"net/http"
)

// @ee.http.Controller
// @ee.http.Route("/api/v1/sms")
type RestController struct {
	sms Repository
}

func NewRestController(sms Repository) *RestController {
	return &RestController{sms}
}

// @ee.http.QueryParam("limit")
// @ee.http.Method("GET")
func (s *RestController) List(ctx context.Context, limit int) ([]SMS, error) {
	return s.sms.FindAll(ctx, limit)
}

// @ee.http.HeaderParam("value":"userAgent","alias":"User-Agent")
// @ee.http.Route("/a/:id")
// @ee.http.Method("GET")
func (s *RestController) Get(ctx context.Context, id uuid.UUID, userAgent string) (SMS, error) {
	fmt.Println(userAgent)
	return s.sms.FindById(ctx, id)
}

// @ee.http.Route("/asdf")
// @ee.http.Method("GET")
func (s *RestController) Get2(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello world"))
}
