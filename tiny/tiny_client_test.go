package tiny

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ziskandyigal/stringculator/mocks"
)

func Test_tiny_AddXY_Sucess(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	calculatorMock := mocks.NewMockCalculator(ctl)
	presenterMock := mocks.NewMockPresenter(ctl)
	tinyCilent := &tiny{
		calc: calculatorMock,
		pres: presenterMock,
	}

	expectedHashRes := "xoxoxoxo123"
	expectedCalcRes := "10"

	calculatorMock.EXPECT().Add(5, 5).Return(expectedCalcRes, nil).Times(1)
	presenterMock.EXPECT().Hashify(expectedCalcRes).Return(expectedHashRes, nil).Times(1)

	res, err := tinyCilent.AddXY(5, 5)
	assert.NoError(t, err)
	assert.Equal(t, res, expectedHashRes)
}

func Test_tiny_Add_Return_Error_1(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	calculatorMock := mocks.NewMockCalculator(ctl)
	presenterMock := mocks.NewMockPresenter(ctl)
	tinyCilent := &tiny{
		calc: calculatorMock,
		pres: presenterMock,
	}

	expectedHashRes := ""
	expectedCalcRes := ""
	expectedError := errors.New("some error")

	calculatorMock.EXPECT().Add(5, 5).Return("", expectedError).Times(1)
	presenterMock.EXPECT().Hashify(expectedCalcRes).Return(expectedHashRes, nil).Times(0)

	res, err := tinyCilent.AddXY(5, 5)
	assert.ErrorIs(t, err, expectedError)
	assert.Equal(t, res, expectedHashRes)
}

func Test_tiny_Add_Return_Error_2(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	calculatorMock := mocks.NewMockCalculator(ctl)
	presenterMock := mocks.NewMockPresenter(ctl)
	tinyCilent := &tiny{
		calc: calculatorMock,
		pres: presenterMock,
	}

	expectedHashRes := ""
	expectedCalcRes := "10"
	expectedError := errors.New("some error")

	calculatorMock.EXPECT().Add(5, 5).Return(expectedCalcRes, nil).Times(1)
	presenterMock.EXPECT().Hashify(expectedCalcRes).Return(expectedHashRes, expectedError).Times(1)

	res, err := tinyCilent.AddXY(5, 5)
	assert.ErrorIs(t, err, expectedError)
	assert.Equal(t, res, expectedHashRes)
}
