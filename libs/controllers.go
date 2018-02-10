package libs

import (
	"bytes"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckKeysExist(urlQuery url.Values, keys ...string) bool {
	for _, key := range keys {
		if _, exist := urlQuery[key]; !exist {
			return false
		}
	}
	return true
}

// APIAssertRequest represents the request and expected response data
type APIAssertRequest struct {
	Method string
	URL    string
	Values url.Values
	Body   []byte
	Status int
	Schema map[string]interface{}
}

// AssertResponseWithoutData asserts everything but the data is returned
func (r *APIAssertRequest) AssertResponse(t *testing.T, handler http.HandlerFunc) []byte {
	params := r.Values.Encode()
	reqURL := r.URL
	if params != "" {
		reqURL = fmt.Sprintf("%s?%s", reqURL, params)
	}
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(r.Method, reqURL, bytes.NewReader(r.Body))
	assert.NoError(t, err)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, r.Status, rr.Code)
	if r.Status >= 200 && r.Status <= 300 {
		loader := gojsonschema.NewGoLoader(r.Schema)
		result, err := gojsonschema.Validate(loader, gojsonschema.NewStringLoader(rr.Body.String()))
		assert.NoError(t, err)
		assert.True(t, result.Valid(), fmt.Sprintf("%v", result.Errors()))
	}
	return rr.Body.Bytes()
}
