package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"Apps-I_Desa_Backend/middleware"

	"github.com/gofiber/fiber/v2"
)

// setupTestApp creates a lightweight Fiber app for integration tests.
// It does NOT connect to a real database. Routes are wired manually
// so we can test the HTTP layer (middleware, routing, request parsing).
func setupTestApp() *fiber.App {
	app := fiber.New(fiber.Config{
		// Propagate the HTTP status code from fiber errors (e.g. 404, 401)
		// instead of always returning 500.
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		},
	})

	// ── Auth routes (no DB needed for these tests) ──────────────
	// We register stub handlers that replicate the real controller
	// behavior at the HTTP-validation level only.
	authRoutes := app.Group("/api/auth")
	authRoutes.Post("/login", stubLoginHandler)
	authRoutes.Post("/logout", stubLogoutHandler)

	// ── Protected routes (JWT middleware is the real one) ────────
	// The handlers themselves are stubs — we only care that the
	// middleware rejects unauthorized requests.
	api := app.Group("/api/villagers")
	api.Use(middleware.JWTAuth())
	api.Get("/", stubOKHandler)
	api.Post("/", stubOKHandler)
	api.Get("/:nik", stubOKHandler)
	api.Put("/:nik", stubOKHandler)
	api.Delete("/:nik", stubOKHandler)

	return app
}

// stubLoginHandler mimics the real auth controller's validation layer.
func stubLoginHandler(c *fiber.Ctx) error {
	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}
	username, hasUser := body["username"]
	password, hasPass := body["password"]
	if !hasUser || !hasPass || username == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   "username and password are required",
		})
	}
	// In integration tests we don't call the DB, so just return 200 stub.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}

func stubLogoutHandler(c *fiber.Ctx) error {
	c.ClearCookie("AppsIDesaCookie")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logout successful!"})
}

func stubOKHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}

// ─────────────────────────────────────────────────────────────────────────────
// TC-IT-01: POST /api/auth/login with malformed JSON body → 400
// Tests HTTP body-parsing layer, not service logic (unit tests cover that).
// ─────────────────────────────────────────────────────────────────────────────
func TestIT01_Login_MalformedJSON(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(`{bad json`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("TC-IT-01 FAIL: app.Test error: %v", err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("TC-IT-01 FAIL: expected 400, got %d", resp.StatusCode)
		return
	}
	t.Log("TC-IT-01 PASS: Malformed JSON body correctly rejected with 400")
}

// ─────────────────────────────────────────────────────────────────────────────
// TC-IT-02: POST /api/auth/login with empty body → 400
// Tests controller-level required-field validation via HTTP.
// ─────────────────────────────────────────────────────────────────────────────
func TestIT02_Login_EmptyBody(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBufferString("{}"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("TC-IT-02 FAIL: app.Test error: %v", err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("TC-IT-02 FAIL: expected 400, got %d", resp.StatusCode)
		return
	}
	t.Log("TC-IT-02 PASS: Missing required fields correctly rejected with 400")
}

// ─────────────────────────────────────────────────────────────────────────────
// TC-IT-03: POST /api/auth/logout with no extra payload → 200
// Tests the full HTTP round-trip for logout (unit test only covers service).
// Also verifies the response JSON contains the logout message.
// ─────────────────────────────────────────────────────────────────────────────
func TestIT03_Logout_ReturnsOKWithMessage(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodPost, "/api/auth/logout", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("TC-IT-03 FAIL: app.Test error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("TC-IT-03 FAIL: expected 200, got %d", resp.StatusCode)
		return
	}

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("TC-IT-03 FAIL: cannot decode response body: %v", err)
	}
	if body["message"] != "Logout successful!" {
		t.Errorf("TC-IT-03 FAIL: unexpected message: %v", body["message"])
		return
	}
	t.Log("TC-IT-03 PASS: Logout returns 200 dengan message 'Logout successful!'")
}

// ─────────────────────────────────────────────────────────────────────────────
// TC-IT-04: GET /api/villagers/ with no cookie → 401
// Tests that JWT middleware blocks unauthenticated GET list requests.
// ─────────────────────────────────────────────────────────────────────────────
func TestIT04_GetVillagers_NoCookie_Unauthorized(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodGet, "/api/villagers/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("TC-IT-04 FAIL: app.Test error: %v", err)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("TC-IT-04 FAIL: expected 401, got %d", resp.StatusCode)
		return
	}
	t.Log("TC-IT-04 PASS: GET /api/villagers/ tanpa cookie → 401 Unauthorized")
}

// ─────────────────────────────────────────────────────────────────────────────
// TC-IT-05: GET /api/villagers/ with an invalid/tampered token → 401
// Tests that JWT middleware rejects a malformed token, not just missing ones.
// ─────────────────────────────────────────────────────────────────────────────
func TestIT05_GetVillagers_InvalidToken_Unauthorized(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodGet, "/api/villagers/", nil)
	req.Header.Set("Cookie", "AppsIDesaCookie=this.is.not.a.valid.token")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("TC-IT-05 FAIL: app.Test error: %v", err)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("TC-IT-05 FAIL: expected 401, got %d", resp.StatusCode)
		return
	}
	t.Log("TC-IT-05 PASS: GET /api/villagers/ dengan token invalid → 401 Unauthorized")
}

// ─────────────────────────────────────────────────────────────────────────────
// TC-IT-06: POST /api/villagers/ with no cookie → 401
// Tests that JWT middleware also blocks unauthenticated POST requests.
// ─────────────────────────────────────────────────────────────────────────────
func TestIT06_CreateVillager_NoCookie_Unauthorized(t *testing.T) {
	app := setupTestApp()

	payload := `{"nik":"1234567890123456","nama_lengkap":"Test"}`
	req := httptest.NewRequest(http.MethodPost, "/api/villagers/", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("TC-IT-06 FAIL: app.Test error: %v", err)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("TC-IT-06 FAIL: expected 401, got %d", resp.StatusCode)
		return
	}
	t.Log("TC-IT-06 PASS: POST /api/villagers/ tanpa cookie → 401 Unauthorized")
}

// ─────────────────────────────────────────────────────────────────────────────
// TC-IT-07: GET /api/villagers/:nik with no cookie → 401
// Tests middleware on the parameterized GET route.
// ─────────────────────────────────────────────────────────────────────────────
func TestIT07_GetVillagerByNIK_NoCookie_Unauthorized(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodGet, "/api/villagers/1234567890123456", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("TC-IT-07 FAIL: app.Test error: %v", err)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("TC-IT-07 FAIL: expected 401, got %d", resp.StatusCode)
		return
	}
	t.Log("TC-IT-07 PASS: GET /api/villagers/:nik tanpa cookie → 401 Unauthorized")
}

// ─────────────────────────────────────────────────────────────────────────────
// TC-IT-08: GET /api/unknown route → 404
// Tests that the router returns 404 for unregistered endpoints.
// ─────────────────────────────────────────────────────────────────────────────
func TestIT08_UnknownRoute_NotFound(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodGet, "/api/unknown-endpoint", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("TC-IT-08 FAIL: app.Test error: %v", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("TC-IT-08 FAIL: expected 404, got %d", resp.StatusCode)
		return
	}
	t.Log("TC-IT-08 PASS: Unknown route → 404 Not Found")
}
