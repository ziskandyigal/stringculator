//go:generate mockgen -destination ../mocks/mock_converter.go --package mocks -source converter.go
package converter

type Converter interface {
	ToMD5(num int) string
	ToSHA1(num int) string
	ToSHA256(num int) string
}
