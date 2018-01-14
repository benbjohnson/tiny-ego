package http

import (
	"net/http"
	"net/url"
)

func PostFormString(r *http.Request, key string) string {
	r.ParseForm()
	return URLString(r.PostForm, key)
}

func PostFormStringPtr(r *http.Request, key string) *string {
	r.ParseForm()
	return URLStringPtr(r.PostForm, key)
}

func PostFormInt(r *http.Request, key string) int {
	r.ParseForm()
	return URLInt(r.PostForm, key)
}

func PostFormIntPtr(r *http.Request, key string) *int {
	r.ParseForm()
	return URLIntPtr(r.PostForm, key)
}

func FormString(r *http.Request, key string) string {
	r.ParseForm()
	return URLString(r.Form, key)
}

func FormStringPtr(r *http.Request, key string) *string {
	r.ParseForm()
	return URLStringPtr(r.Form, key)
}

func FormInt(r *http.Request, key string) int {
	r.ParseForm()
	return URLInt(r.Form, key)
}

func FormIntPtr(r *http.Request, key string) *int {
	r.ParseForm()
	return URLIntPtr(r.Form, key)
}

func URLString(v url.Values, key string) string {
	return v.Get(key)
}

func URLStringPtr(v url.Values, key string) *string {
	if len(v[key]) == 0 {
		return nil
	}
	s := v.Get(key)
	return &s
}

func URLInt(v url.Values, key string) int {
	i, _ := strconv.Atoi(v.Get(key))
	return i
}

func URLIntPtr(v url.Values, key string) *int {
	if len(v[key]) == 0 {
		return nil
	}
	i, _ := strconv.Atoi(v.Get(key))
	return &i
}
