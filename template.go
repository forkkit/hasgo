// Code generated by go generate; DO NOT EDIT.
package main

var hasgoTemplates = map[string]string{
	"Sum.go": `
func (s SliceType) Sum() ElementType {
	var sum ElementType
	for _, v := range s {
		sum += v
	}
	return sum
}
`,
}

const (
	ForNumbers = "ForNumbers"
	ForStrings = "ForStrings"
)

var funcDomains = map[string][]string{
	"Sum.go": []string{ForNumbers, ForStrings},
}