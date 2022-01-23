package tiny

import (
	"github.com/ziskandyigal/stringculator/calculator"
	"github.com/ziskandyigal/stringculator/presenter"
)

type tiny struct {
	calc calculator.Calculator
	pres presenter.Presenter
}

func New(c calculator.Calculator, p presenter.Presenter) tiny {
	return tiny{
		calc: c,
		pres: p,
	}
}

func (t *tiny) Add(num1, num2 int) (string, error) {
	res, err := t.calc.Add(num1, num2)
	if err != nil {
		return "", err
	}

	return t.pres.Hashify(res)
}

func (t *tiny) Sub(num1, num2 int) (string, error) {
	res, err := t.calc.Sub(num1, num2)
	if err != nil {
		return "", err
	}

	return t.pres.Hashify(res)
}

func (t *tiny) Sum(num1, num2 int) (string, error) {
	res, err := t.calc.Sum(num1, num2)
	if err != nil {
		return "", err
	}

	return t.pres.Hashify(res)
}

func (t *tiny) Div(num1, num2 int) (string, error) {
	res, err := t.calc.Div(num1, num2)
	if err != nil {
		return "", err
	}

	return t.pres.Hashify(res)
}
