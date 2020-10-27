package config

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"
)

// These debugging codes are copeid and modified from https://github.com/motemen/go-loghttp
// Copyright (c) 2014 motemen

var _ http.RoundTripper = &DebugTransport{}

// DebugTransport implements http.RoundTripper and showing request/response contents.
type DebugTransport struct {
	Transport http.RoundTripper
}

var DefaultDebugTransport = &DebugTransport{
	Transport: http.DefaultTransport,
}

type contextKey struct{}

var contextKeyRequestStart = &contextKey{}

func (t *DebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := context.WithValue(req.Context(), contextKeyRequestStart, time.Now())
	req = req.WithContext(ctx)

	num := rand.Int63() // #nosec G404 | this number is used only for debugging and readability
	t.showDumpRequest(req, num)

	resp, err := t.transport().RoundTrip(req)
	if err != nil {
		return resp, err
	}

	t.showDumpResponse(resp, num)
	return resp, err
}

// showDumpRequest shows dumps of requests.
func (t *DebugTransport) showDumpRequest(req *http.Request, num int64) {
	if req == nil || req.Body == nil {
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	dump, err := httputil.DumpRequest(req, false)
	if err != nil {
		return
	}
	fmt.Printf("---- [[AWS Request (#%d)]] ----\n[[Request Header]]\n%s[[Request Body]]\n%s\n\n", num, string(dump), body)
}

func (t *DebugTransport) showDumpResponse(resp *http.Response, num int64) {
	if resp == nil || resp.Body == nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	resp.Body = ioutil.NopCloser(bytes.NewReader(body))
	dump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		return
	}
	fmt.Printf("---- [[AWS Response (#%d)]] ----\n[[Response Header]]\n%s[[Response Body]]\n%s\n\n\n", num, string(dump), body)
}

func (t *DebugTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}

	return http.DefaultTransport
}
