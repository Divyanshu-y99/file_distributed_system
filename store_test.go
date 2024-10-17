package main

import (
	"bytes"
	"fmt"
	"io"

	"testing"
)

// func TestPathTransformFunc(t *testing.T) {
// 	key := "momsbestpicture"
// 	pathKey := CASPathTransformFun(key)
// 	expectedOriginalKey := "6804429f74181a63c50c3d81d733a12f14a353ff"
// 	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"

// 	if pathKey.PathName != expectedPathName {
// 		t.Error(t, "have %s want %s", pathKey.PathName, expectedPathName)
// 	}

// 	if pathKey.Filename != expectedPathName {
// 		t.Error(t, "have %s want %s", pathKey.Filename, expectedOriginalKey)
// 	}

// }

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFun,
	}
	s := NewStore(opts)
	key := "momsspecials"

	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)

	fmt.Println(string(b))

	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}
}
