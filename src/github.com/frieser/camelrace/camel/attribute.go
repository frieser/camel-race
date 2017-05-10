package camel

type Attribute struct {
	Name     string
	Modifier func(n float32) interface{}
}
