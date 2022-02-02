//go:generate mockgen -destination ../mocks/mock_presenter.go --package mocks -source presenter.go

package presenter

type Presenter interface {
	Hashify(num string) (string, error)
}
