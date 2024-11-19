package generator

import (
	"crypto/md5"
	"fmt"
)

type SecondGenerator struct{}

func (SecondGenerator) GenerateId(title string) string {
	hash := md5.New()
	hash.Write([]byte(title))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
