package main

import (
	"io"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

var entropy io.Reader

func generateULID() string {
	if entropy == nil {
		entropy = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	}
	return ulid.MustNew(ulid.Now(), entropy).String()
}
