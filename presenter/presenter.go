package presenter

type Presenter interface {
	Hashify(num string) (string, error)
}
