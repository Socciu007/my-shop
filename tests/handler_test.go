package tests

import (
	// "bytes"
	// "encoding/json"
	// "net/http"
	// "net/http/httptest"
	// "testing"

	// "my_shop/internal/models"
	// "my_shop/internal/routers"

	// "github.com/gin-gonic/gin"
)

// func TestPing(t *testing.T) {
//     // Set up a temporary database for testing
// 	mysqlService := config.New()
//     db := mysqlService.GetDB()

// 	// Sử dụng Gin không trong chế độ debug để tránh các cảnh báo không cần thiết trong output kiểm tra
// 	gin.SetMode(gin.TestMode)

// 	// Thiết lập router cho test
// 	r := routers.SetupRouter(db)

// 	// Tạo một request HTTP test
// 	req, err := http.NewRequest("GET", "/ping", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Tạo một ResponseRecorder để ghi lại response
// 	rr := httptest.NewRecorder()

// 	// Phục vụ request HTTP test
// 	r.ServeHTTP(rr, req)

// 	// Kiểm tra mã trạng thái
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Handler trả về mã trạng thái không đúng: nhận %v muốn %v", status, http.StatusOK)
// 	}

// 	// Kiểm tra body của response
// 	expected := "{\"message\":\"pong\"}"
// 	if rr.Body.String() != expected {
// 		t.Errorf("Handler trả về body không đúng: nhận %v muốn %v", rr.Body.String(), expected)
// 	}
// }

// func TestCreateUser(t *testing.T) {
//     // Set up a temporary database for testing
// 	mysqlService := config.New()
//     db := mysqlService.GetDB()

//     // Auto migrate the User model
//     db.AutoMigrate(&models.Users{})

//     // Create a new instance of the UserService
//     userService := us.NewUserService(db)

//     // Create a new instance of the UserController
//     userController := uc.NewUserController(&userService)

//     // Set up Gin router
//     gin.SetMode(gin.TestMode)
//     r := routers.SetupRouter(db)
//     r.POST("/api/create-user", userController.CreateUser)

//     // Define the JSON payload
//     payload := map[string]string{
//         "username": "user001",
//         "email": "user001@gmail.com",
//         "password": "12345678",
//     }

//     // Convert the payload to JSON
//     jsonData, err := json.Marshal(payload)
//     if err != nil {
//         t.Fatal(err)
//     }

//     // Create a new HTTP request
//     req, err := http.NewRequest("POST", "/api/create-user", bytes.NewBuffer(jsonData))
//     if err != nil {
//         t.Fatal(err)
//     }

//     // Set the request content type to application/json
//     req.Header.Set("Content-Type", "application/json")

//     // Use httptest to create a ResponseRecorder
//     rr := httptest.NewRecorder()

//     // Serve the HTTP request
//     r.ServeHTTP(rr, req)

//     // Check the status code is what you expect
//     if status := rr.Code; status != http.StatusCreated {
//         t.Errorf("Handler returned wrong status code: got %v want %v",
//             status, http.StatusCreated)
//     }

//     // Define the expected response
//     expectedResponse := map[string]interface{}{
//         "status":  1,
//         "message": "User created successfully",
//     }
//     expected, err := json.Marshal(expectedResponse)
//     if err != nil {
//         t.Fatal(err)
//     }

//     // Check the response body is what you expect
//     if rr.Body.String() != string(expected) {
//         t.Errorf("Handler returned unexpected body: got %v want %v",
//             rr.Body.String(), string(expected))
//     }
// }