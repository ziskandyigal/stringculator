package calculator

type Calculator interface {
	Add(num1, num2 int) (string, error)
	Sub(num1, num2 int) (string, error)
	Sum(num1, num2 int) (string, error)
	Div(num1, num2 int) (string, error)
}
