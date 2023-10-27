package app

type Exporter interface {
	Export() ([]byte, error)
}
