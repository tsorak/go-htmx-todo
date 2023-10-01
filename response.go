package main

import "net/http"

type Responder struct {
	_w http.ResponseWriter
}

func (r Responder) SendBytes(bytes []byte) Responder {
	r._w.Write(bytes)
	return r
}

func (r Responder) SendString(str string) Responder {
	r._w.Write([]byte(str))
	return r
}

func (r Responder) SetHeader(key string, value string) Responder {
	r._w.Header().Set(key, value)
	return r
}

func (r Responder) SetStatus(n int) Responder {
	r._w.WriteHeader(n)
	return r
}
