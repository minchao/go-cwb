package cwb

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the CWB client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// CWB client configured to use test server
	client = NewClient("token", nil)
	u, _ := url.Parse(server.URL)
	client.BaseURL = u
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method is %v, want %v", got, want)
	}
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters is %v, want %v", got, want)
	}
}

func areEqualJSON(j1, j2 []byte) (bool, error) {
	var v1 interface{}
	var v2 interface{}

	var err error
	err = json.Unmarshal(j1, &v1)
	if err != nil {
		return false, fmt.Errorf("Unmarshal JSON 1 error: %v", err)
	}
	err = json.Unmarshal(j2, &v2)
	if err != nil {
		return false, fmt.Errorf("Unmarshal JSON 2 error: %v", err)
	}

	return reflect.DeepEqual(v1, v2), nil
}

func TestNewClient(t *testing.T) {
	c := NewClient("token", http.DefaultClient)

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
}

func TestClient_NewRequest(t *testing.T) {
	c := NewClient("token", http.DefaultClient)

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := "Hello, 世界", "Hello, 世界"
	req, _ := c.NewRequest("GET", inURL, strings.NewReader(inBody))

	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}

	if got, want := req.Header.Get("Authorization"), c.token; got != want {
		t.Errorf("NewRequest() Authorization is %v, want %v", got, want)
	}

	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}
}

func TestClient_Do(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		Hello string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method is %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, `{"Hello":"世界"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(foo)
	client.Do(context.Background(), req, body)

	want := &foo{"世界"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body is %s, want %s", body, want)
	}
}

func TestClient_Do_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(context.Background(), req, nil)
	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}
}

func TestClient_Do_noContent(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var body json.RawMessage

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(context.Background(), req, &body)
	if err == nil {
		t.Error("Expected empty body error.")
	}
}

func TestClient_Do_contextWithCancel(t *testing.T) {
	var body json.RawMessage

	req, _ := client.NewRequest("GET", "/", nil)
	req.URL = nil

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := client.Do(ctx, req, &body)
	if err != context.Canceled {
		t.Logf("Expected context canceled error, got %v", err)
	}
}

func Test_checkResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(strings.NewReader("data not found")),
	}
	err := checkResponse(res).(*ErrorResponse)

	if err == nil {
		t.Error("Expected error response.")
	}

	want := &ErrorResponse{
		Response: res,
		Message:  "data not found",
	}
	if !reflect.DeepEqual(err, want) {
		t.Errorf("Error is %v, want %v", err, want)
	}
}
