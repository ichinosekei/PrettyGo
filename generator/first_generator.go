package generator

import (
	"crypto/sha1"
	"fmt"
)

type FirstGenerator struct{}

func (FirstGenerator) GenerateId(title string) string {
	hash := sha1.New()
	hash.Write([]byte(title))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
