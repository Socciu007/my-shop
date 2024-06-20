package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"my_shop/internal/routers"

	"github.com/gin-gonic/gin"
)

func TestPing(t *testing.T) {
	// Sử dụng Gin không trong chế độ debug để tránh các cảnh báo không cần thiết trong output kiểm tra
	gin.SetMode(gin.TestMode)

	// Thiết lập router cho test
	r := routers.SetupRouter()

	// Tạo một request HTTP test
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Tạo một ResponseRecorder để ghi lại response
	rr := httptest.NewRecorder()

	// Phục vụ request HTTP test
	r.ServeHTTP(rr, req)

	// Kiểm tra mã trạng thái
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler trả về mã trạng thái không đúng: nhận %v muốn %v", status, http.StatusOK)
	}

	// Kiểm tra body của response
	expected := "{\"message\":\"pong\"}"
	if rr.Body.String() != expected {
		t.Errorf("Handler trả về body không đúng: nhận %v muốn %v", rr.Body.String(), expected)
	}
}

func TestCreateUser(t *testing.T) {
    // Define the JSON payload
    payload := map[string]string{
        "name": "John Doe",
    }

    // Convert the payload to JSON
    jsonData, err := json.Marshal(payload)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new HTTP request
    req, err := http.NewRequest("POST", "/v1/users/create", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatal(err)
    }

    // Set the request content type to application/json
    req.Header.Set("Content-Type", "application/json")

    // Use httptest to create a ResponseRecorder
    rr := httptest.NewRecorder()

    // Create an instance of your handler
    handler := http.HandlerFunc(CreateUserHandler)

    // Serve the HTTP request
    handler.ServeHTTP(rr, req)

    // Check the status code is what you expect
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler trả về mã trạng thái không đúng: nhận %v muốn %v",
            status, http.StatusOK)
    }

    // Define the expected response
    expectedResponse := map[string]interface{}{
        "message": "User created",
        "user":    map[string]string{"name": "John Doe"},
    }
    expected, err := json.Marshal(expectedResponse)
    if err != nil {
        t.Fatal(err)
    }

    // Check the response body is what you expect
    if rr.Body.String() != string(expected) {
        t.Errorf("Handler trả về body không đúng: nhận %v muốn %v",
            rr.Body.String(), string(expected))
    }
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    // Your handler logic here
    var payload map[string]string
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&payload); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Respond with a success message
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    response := map[string]interface{}{
        "message": "User created",
        "user":    payload,
    }
    json.NewEncoder(w).Encode(response)
}

