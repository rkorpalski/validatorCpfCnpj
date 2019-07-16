package tests

import (
	"github.com/rkorpalski/validatorCpfCnpj/pkg/cpfCnpj"
	"github.com/rkorpalski/validatorCpfCnpj/pkg/messages"
	"github.com/rkorpalski/validatorCpfCnpj/pkg/tests/mock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func TestValidCpf(t *testing.T) {
	validCpfs := []string{
		"953.016.910-82",
		"962.567.980-40",
		"137.756.220-47",
	}

	for _, cpf:= range validCpfs {
		isValid := cpfCnpj.ValidateCpf(cpf)
		assert.True(t, isValid)
	}
}

func TestInvalidCpf(t *testing.T) {
	validCpfs := []string{
		"111.111.111-11",
		"777.777.777-77",
		"137.756.220-56",
		"962.567.980-85",
	}

	for _, cpf:= range validCpfs {
		isValid := cpfCnpj.ValidateCpf(cpf)
		assert.False(t, isValid)
	}
}

func TestCheckInvalidCpf(t *testing.T) {
	invalidsCpfs := []string{
		"000.000.000-00",
		"111.111.111-11",
		"222.222.222-22",
		"333.333.333-33",
		"444.444.444-44",
		"555.555.555-55",
		"666.666.666-66",
		"777.777.777-77",
		"888.888.888-88",
		"999.999.999-99",
	}

	validCpfs := []string{
		"953.016.910-82",
		"962.567.980-40",
		"137.756.220-47",
	}

	for _, cpf:= range invalidsCpfs {
		isValid := cpfCnpj.CheckInvalidCpf(cpf)
		assert.True(t, isValid)
	}

	for _, cpf:= range validCpfs {
		isValid := cpfCnpj.CheckInvalidCpf(cpf)
		assert.False(t, isValid)
	}
}

func TestCleanCpf(t *testing.T) {
	cpf := "953.016.910-82"
	unmaskCpf := []int{9, 5, 3, 0, 1, 6, 9, 1, 0, 8, 2}

	result := cpfCnpj.CleanCpf(cpf)

	assert.Equal(t, unmaskCpf, result)
}

func TestCalculateCpfDigit(t *testing.T) {
	cpf := []int{9, 5, 3, 0, 1, 6, 9, 1, 0, 8, 2}
	firstDigitExpect := 8
	secondDigitExpect := 2

	firstDigitResult := cpfCnpj.CalculateCpfDigit(cpf[:9])
	secondDigitResult := cpfCnpj.CalculateCpfDigit(cpf[:10])

	assert.Equal(t, firstDigitExpect, firstDigitResult)
	assert.Equal(t, secondDigitExpect, secondDigitResult)
}

func TestValidate(t *testing.T) {
	validCpf := "953.016.910-82"
	invalidCpf := "95.01.91-82"

	validCnpj := "78.421.417/0001-15"
	invalidCnpj := "7.42.47/001-15"

	service := cpfCnpj.CpfCnpjService{}

	result, err := service.Validate(validCpf)
	assert.NoError(t, err)
	assert.True(t, result)

	result, err = service.Validate(validCnpj)
	assert.NoError(t, err)
	assert.True(t, result)

	result, err = service.Validate(invalidCpf)
	assert.EqualError(t, err, messages.DocumentInvalidError)
	assert.False(t, result)

	result, err = service.Validate(invalidCnpj)
	assert.EqualError(t, err, messages.DocumentInvalidError)
	assert.False(t, result)
}

func TestValidCnpj(t *testing.T) {
	validCnpjs := []string{
		"78.421.417/0001-15",
		"01.693.748/0001-80",
		"95.023.793/0001-70",
	}

	for _, cnpj:= range validCnpjs {
		isValid := cpfCnpj.ValidateCnpj(cnpj)
		assert.True(t, isValid)
	}
}

func TestInvalidCnpj(t *testing.T) {
	validCnpjs := []string{
		"78.421.417/0001-56",
		"01.693.748/0001-74",
		"95.023.793/0001-96",
		"99.999.999/9999-99",
	}

	for _, cnpj:= range validCnpjs {
		isValid := cpfCnpj.ValidateCnpj(cnpj)
		assert.False(t, isValid)
	}
}

func TestCalculateCnpjDigit(t *testing.T) {
	cnpj := []int{7, 8, 4, 2, 1, 4, 1, 7, 0, 0, 0, 1, 1, 5}
	firstDigitExpect := 1
	secondDigitExpect := 5

	firstDigitResult := cpfCnpj.CalculateCnpjDigit(cnpj[:12], 5)
	secondDigitResult := cpfCnpj.CalculateCnpjDigit(cnpj[:13], 6)

	assert.Equal(t, firstDigitExpect, firstDigitResult)
	assert.Equal(t, secondDigitExpect, secondDigitResult)
}

func TestCleanCnpj(t *testing.T) {
	cnpj := "78.421.417/0001-15"
	expect := []int{7, 8, 4, 2, 1, 4, 1, 7, 0, 0, 0, 1, 1, 5}

	result := cpfCnpj.CleanCnpj(cnpj)

	assert.Equal(t, expect, result)
}

func TestCheckInvalidCnpj(t *testing.T) {
	invalidsCnpjs := []string{
		"00.000.000/0000-00",
		"11.111.111/1111-11",
		"22.222.222/2222-22",
		"33.333.333/3333-33",
		"44.444.444/4444-44",
		"55.555.555/5555-55",
		"66.666.666/6666-66",
		"77.777.777/7777-77",
		"88.888.888/8888-88",
		"99.999.999/9999-99",
	}

	validCnpjs := []string{
		"78.421.417/0001-56",
		"01.693.748/0001-74",
		"95.023.793/0001-96",
	}

	for _, cnpj:= range invalidsCnpjs {
		isInvalid := cpfCnpj.CheckInvalidCnpj(cnpj)
		assert.True(t, isInvalid)
	}

	for _, cnpj:= range validCnpjs {
		isInvalid := cpfCnpj.CheckInvalidCnpj(cnpj)
		assert.False(t, isInvalid)
	}
}

func TestSaveCpfCnpj(t *testing.T) {
	cpfCnpjMock := cpfCnpj.CpfCnpj{
		Number: "892.385.660-62",
		Type: "CPF",
	}
	dbMock := mocks.Repository{}
	dbMock.On("Save", cpfCnpjMock).Return(nil)

	service := cpfCnpj.NewCpfCnpjService(&dbMock)
	err := service.Save(cpfCnpjMock)
	assert.NoError(t, err)
}

func TestSaveInvalidCnpj(t *testing.T) {
	cpfCnpjMock := cpfCnpj.CpfCnpj{
		Number: "03.689.262/0001-78",
		Type: "CNPJ",
	}
	dbMock := mocks.Repository{}
	dbMock.On("Save", cpfCnpjMock).Return(nil)

	service := cpfCnpj.NewCpfCnpjService(&dbMock)
	err := service.Save(cpfCnpjMock)
	assert.Error(t, err, messages.DocumentInvalidError)
}

func TestGetAllDocuments(t *testing.T) {
	listMock := []cpfCnpj.CpfCnpj{
		{
			Number: "03.689.262/0001-78",
			Type: "CNPJ",
			BlackList: false,
			CreateDate: time.Now().String(),
		},
		{
			Number: "892.385.660-62",
			Type: "CPF",
			BlackList: false,
			CreateDate: time.Now().String(),
		},
	}
	dbMock := mocks.Repository{}
	dbMock.On("GetAllDocuments").Return(listMock, nil)

	service := cpfCnpj.NewCpfCnpjService(&dbMock)
	results, err := service.GetAllDocuments()
	assert.NoError(t, err)
	assert.Len(t, results, 2)
}