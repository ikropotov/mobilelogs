package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"encoding/json"
)




func TestProcessLogs(t *testing.T) {

	r := CreateRouter()

	mapD := map[string]string{"id": "id1", "password": "123", "log": "Test Log"}
	mapB, _ := json.Marshal(mapD)

	// Create a request to send to the above route
	req, _ := http.NewRequest("POST", "/logs", strings.NewReader(string(mapB)))

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)


}