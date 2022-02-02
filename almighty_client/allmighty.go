package almighty

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

type Allmighty struct{}

func (a *Allmighty) AddXY(num1, num2 int) (string, error) {
	res := num1 + num2
	return a.hashify(fmt.Sprint(res))
}

func (a *Allmighty) SubXY(num1, num2 int) (string, error) {
	res := num1 - num2
	return a.hashify(fmt.Sprint(res))
}

func (a *Allmighty) SumXY(num1, num2 int) (string, error) {
	res := num1 + num2
	return a.hashify(fmt.Sprint(res))
}

func (a *Allmighty) DivXY(num1, num2 int) (string, error) {
	if num1 == 0 {
		return "", errors.New("invalid operation")
	}

	res := num1 / num2
	return a.hashify(fmt.Sprint(res))
}

func (a *Allmighty) hashify(num string) (string, error) {
	i, err := strconv.Atoi(num)
	if err != nil {
		return "", err
	}

	is2Devident := i%2 == 0
	is5Devident := i%5 == 0
	is10Devident := i%10 == 0

	if is2Devident {
		return a.toMD5(i), nil
	}

	if is5Devident {
		return a.toSHA1(i), nil
	}

	if is10Devident {
		return a.toSHA256(i), nil
	}

	return "", errors.New("not devident number")
}

func (a *Allmighty) toMD5(num int) string {
	strNum := strconv.Itoa(num)

	h := md5.New()
	h.Write([]byte(strNum))
	return hex.EncodeToString(h.Sum(nil))
}

func (a *Allmighty) toSHA1(num int) string {
	strNum := fmt.Sprint(num)

	h := sha1.New()
	h.Write([]byte(strNum))
	return hex.EncodeToString(h.Sum(nil))
}

func (a *Allmighty) toSHA256(num int) string {
	strNum := fmt.Sprint(num)

	h := sha256.New()
	h.Write([]byte(strNum))
	return hex.EncodeToString(h.Sum(nil))
}
