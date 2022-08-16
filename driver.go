package mongox

type Identify interface {
	Id() string
}

type Document interface {
	DocumentName() string
}
