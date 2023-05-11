//@Author: Dennis
//@Date: 2023/5/11

package helloworld

import (
	"crypto/rand"
	"encoding/base64"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("hello empty name", func(t *testing.T) {
		got := Hello("")
		want := "hello world"
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})
	t.Run("hello not empty name", func(t *testing.T) {
		nameBytes := make([]byte, 5)
		_, _ = rand.Read(nameBytes)
		name := base64.StdEncoding.EncodeToString(nameBytes)
		got := Hello(name)
		want := "hello " + name
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})
}
