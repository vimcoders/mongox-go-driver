package mongox

type Identify interface {
	Identify() string
}

type Document interface {
	DocumentName() string
}
