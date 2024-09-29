package generator

type Generator interface {
	GenerateId(title string) string
}
