// Test the main function

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHomePage tests the homePage handler
func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/home", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homePage)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("homePage returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("homePage returned unexpected content type: got %v want %v", contentType, expectedContentType)
	}

	expectedBodySnippet := "<h1>Welcome to My Portfolio</h1>"
	if !strings.Contains(rr.Body.String(), expectedBodySnippet) {
		t.Errorf("homePage returned unexpected body content: missing %v", expectedBodySnippet)
	}
}

// TestSkillsPage tests the skillsPage handler
func TestSkillsPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/skills", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(skillsPage)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("skillsPage returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("skillsPage returned unexpected content type: got %v want %v", contentType, expectedContentType)
	}

	expectedBodySnippet := "<h1>My Skills</h1>"
	if !strings.Contains(rr.Body.String(), expectedBodySnippet) {
		t.Errorf("skillsPage returned unexpected body content: missing %v", expectedBodySnippet)
	}
}

// TestAboutPage tests the aboutPage handler
func TestAboutPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(aboutPage)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("aboutPage returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("aboutPage returned unexpected content type: got %v want %v", contentType, expectedContentType)
	}

	expectedBodySnippet := "<h1>About Me</h1>"
	if !strings.Contains(rr.Body.String(), expectedBodySnippet) {
		t.Errorf("aboutPage returned unexpected body content: missing %v", expectedBodySnippet)
	}
}

// TestContactPage tests the contactPage handler
func TestContactPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/contact", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contactPage)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("contactPage returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("contactPage returned unexpected content type: got %v want %v", contentType, expectedContentType)
	}

	expectedBodySnippet := "<h1>Contact Me</h1>"
	if !strings.Contains(rr.Body.String(), expectedBodySnippet) {
		t.Errorf("contactPage returned unexpected body content: missing %v", expectedBodySnippet)
	}
}
