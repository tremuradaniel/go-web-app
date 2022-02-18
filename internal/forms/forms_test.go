package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData

	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "smth")

	r := httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData

	form := New(r.PostForm)
	form.Has("a")

	if !form.Valid() {
		t.Error("form shows it does not have fields that are not empty when it does")
	}

	r, _ = http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	postedData.Add("a", "")
	r.PostForm = postedData

	form = New(r.PostForm)
	form.Has("a")

	if form.Valid() {
		t.Error("form shows it has field and is not empty, when it is present but empty")
	}

	r, _ = http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	r.PostForm = postedData

	form = New(r.PostForm)
	form.Has("a")

	if form.Valid() {
		t.Error("form shows it has field and is not empty, when it is not present")
	}
}

func TestForm_MinLength(t *testing.T) {
	r, _ := http.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("a", "smth")

	r.PostForm = postedData

	form := New(r.PostForm)

	form.MinLength("a", 5)

	if form.Valid() {
		t.Error("value bigger than min length")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r, _ := http.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("email", "real@mail.com")

	r.PostForm = postedData

	form := New(r.PostForm)

	form.IsEmail("email")

	if !form.Valid() {
		t.Error("did not determine real mail")
	}

	isError := form.Errors.Get("email")
	if isError != "" {
		t.Error("should not have an error, but did get one")
	}

	r, _ = http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	postedData.Add("email", "fakemail.com")

	r.PostForm = postedData

	form = New(r.PostForm)

	form.IsEmail("email")

	if form.Valid() {
		t.Error("did not determine fake mail")
	}
}
