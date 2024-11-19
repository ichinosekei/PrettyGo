package generator

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestFirstGenerator_GenerateId(t *testing.T) {
	gen := FirstGenerator{}
	id := gen.GenerateId("Test Book")
	hash := sha1.New()
	hash.Write([]byte("Test Book"))
	expected := fmt.Sprintf("%x", hash.Sum(nil))
	// для sha1 проверка
	if id != expected {
		t.Errorf("Expected ID %s, but got %s", expected, id)
	}
}

func TestSecondGenerator_GenerateId(t *testing.T) {
	gen := SecondGenerator{}
	id := gen.GenerateId("Test Book")
	hash := md5.New()
	hash.Write([]byte("Test Book"))
	expected := fmt.Sprintf("%x", hash.Sum(nil))
	// для md5 проверка
	if id != expected {
		t.Errorf("Expected ID %s, but got %s", expected, id)
	}
}
