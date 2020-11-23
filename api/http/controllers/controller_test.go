package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// controllerTest represents the suite of testing methods for controllers.
type controllerTest struct {
	testing  *testing.T
	recorder *httptest.ResponseRecorder
	gin      *gin.Context
	engine   *gin.Engine
}

// newResponseRecorder - New recorder for testing
// controllers, initalises gin & sets gin mode.
func newResponseRecorder(t *testing.T) *controllerTest {
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	gin, engine := gin.CreateTestContext(rr)

	return &controllerTest{
		testing:  t,
		recorder: rr,
		gin:      gin,
		engine:   engine,
	}
}

// Message gets the response message from the body
func (c *controllerTest) Message() string {
	b, ok := c.Body()["message"]
	if !ok {
		c.testing.Error("no message within the response")
	}
	return b.(string)
}

// Body gets the response body from the request
func (c *controllerTest) Body() map[string]interface{} {
	body := map[string]interface{}{}
	if err := json.Unmarshal(c.recorder.Body.Bytes(), &body); err != nil {
		fmt.Println(c.recorder.Body.String())
		c.testing.Fatal(fmt.Sprintf("error unmarshalling body %v", err))
	}
	return body
}

// getResponseData gets the response body & checks if the data key
// exists and then marshalls the data key to form a string.
//
// Returns a string of the marshalled data
func (c *controllerTest) Data() string {

	b, ok := c.Body()["data"]
	if !ok {
		c.testing.Fatal("no data within the response")
	}

	got, err := json.Marshal(b)
	if err != nil {
		c.testing.Fatal(fmt.Sprintf("error marshalling data %v", err))
	}

	return string(got)
}

// NewRequest makes a new http.Request and assigns the gin testing
// the request.
func (c *controllerTest) NewRequest(method string, url string, body io.Reader) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		c.testing.Fatal(err)
	}
	c.gin.Request = req
}

func (c *controllerTest) RequestAndServe(method string, url string, engineUrl string, body io.Reader, handler func(ctx *gin.Context)) {
	switch method {
	case "GET":
		{
			c.engine.GET(engineUrl, handler)
		}
	case "POST":
		{
			c.engine.POST(engineUrl, handler)
		}
	case "PUT":
		{
			c.engine.PUT(engineUrl, handler)
		}
	case "DELETE":
		{
			c.engine.DELETE(engineUrl, handler)
		}
	}
	c.NewRequest(method, url, body)
	c.engine.ServeHTTP(c.recorder, c.gin.Request)
}
