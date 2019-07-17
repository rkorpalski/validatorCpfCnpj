package cpfCnpj

import (
	"github.com/pkg/errors"
	"github.com/rkorpalski/validatorCpfCnpj/backend/pkg/messages"
	"regexp"
	"strconv"
	"strings"
)

const (
	CpfRegex = `^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`
	CnpjRegex = `^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`
)

var (
	invalidsCpfs = []string{
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

	invalidsCnpjs = []string{
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
)

type Service interface {
	Validate(documentNumber string) (bool, error)
	Save(cpfCnpj *CpfCnpj) error
}

type CpfCnpjService struct {
	repo Repository
}

func NewCpfCnpjService(repo Repository) * CpfCnpjService {
	return &CpfCnpjService{
		repo: repo,
	}
}

func (s *CpfCnpjService) Validate(documentNumber string) (bool, error){

	regex, err := regexp.Compile(CpfRegex)
	if err != nil {
		return false, errors.Wrap(err, messages.RegexCompileError)
	}

	if match := regex.MatchString(documentNumber); match {
		return ValidateCpf(documentNumber), nil
	}

	regex, err = regexp.Compile(CnpjRegex)
	if err != nil {
		return false, errors.Wrap(err, messages.RegexCompileError)
	}

	if match := regex.MatchString(documentNumber); match {
		return ValidateCnpj(documentNumber), nil
	}

	return false, errors.New(messages.DocumentInvalidError)
}

func (s *CpfCnpjService) Save(cpfCnpj CpfCnpj) error {
	isValid, err := s.Validate(cpfCnpj.Number)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New(messages.DocumentInvalidError)
	}
	return s.repo.Save(cpfCnpj)
}

func (s *CpfCnpjService) GetAllDocuments(isBlacklsit bool) ([]CpfCnpj, error) {
	return s.repo.GetDocuments(isBlacklsit)
}

func (s *CpfCnpjService) MoveToBlacklist(documentId string) error {
	return s.repo.MoveToBlacklist(documentId)
}

func (s *CpfCnpjService) DeleteDocument(documentId string) error {
	return s.repo.DeleteDocument(documentId)
}

func (s *CpfCnpjService) RemoveFromBlacklist(documentId string) error {
	return s.repo.RemoveFromBlacklist(documentId)
}

func (s *CpfCnpjService) FindByDocument(document string) ([]CpfCnpj, error) {
	return s.repo.FindByDocument(document)
}
/*
	Valida um CPF calculando os dois últimos dígitos.
	Para entender a regra de validação de CPF ver:
	https://souforce.cloud/regra-de-validacao-para-cpf-e-cnpj/
*/
func ValidateCpf(cpf string) bool {

	if isInvalid := CheckKnowInvalidCpf(cpf); isInvalid {
		return false
	}

	digits := CleanCpf(cpf)

	firstDigit := CalculateCpfDigit(digits[:9])
	secondDigit := CalculateCpfDigit(digits[:10])

	if firstDigit == digits[9] && secondDigit == digits[10] {
		return true
	}

	return false
}

/*
	Existem CPFs que passam pela validação, mas são reconhecidos como inválidos.
	Este método verifica se o CPF é um destes.
*/
func CheckKnowInvalidCpf(cpf string) bool {
	for _, invalidCpf := range invalidsCpfs {
		if cpf == invalidCpf {
			return true
		}
	}
	return false
}

func CleanCpf(cpf string) []int{
	unmaskCpf := strings.ReplaceAll(cpf, ".", "")
	unmaskCpf = strings.ReplaceAll(unmaskCpf, "-", "")
	strDigits := strings.Split(unmaskCpf, "")

	digits := make([]int, 0)

	for _, strDigit := range strDigits {
		digit, _ := strconv.Atoi(strDigit)
		digits = append(digits, digit)
	}

	return digits
}

func CalculateCpfDigit(digits []int) int{

	size := len(digits) + 1
	sum := 0
	for i := 0; size >= 2; i++ {
		sum += digits[i] * size
		size--
	}

	result := (sum * 10) % 11

	if result == 10 {
		return 0
	}

	return result
}

/*
	Valida um CNPJ calculando os dois últimos dígitos.
	Para entender a regra de validação de CNPJ ver:
	https://souforce.cloud/regra-de-validacao-para-cpf-e-cnpj/
*/
func ValidateCnpj(cnpj string) bool {

	if isInvalid := CheckKnowInvalidCnpj(cnpj); isInvalid {
		return false
	}

	digits := CleanCnpj(cnpj)

	firstDigit := CalculateCnpjDigit(digits[:12], 5)
	secondDigit := CalculateCnpjDigit(digits[:13], 6)

	if firstDigit == digits[12] && secondDigit == digits[13] {
		return true
	}

	return false
}

/*
	Existem CNPJs que passam pela validação, mas são reconhecidos como inválidos.
	Este método verifica se o CNPJ é um destes.
*/
func CheckKnowInvalidCnpj(cnpj string) bool {
	for _, invalidCnpj := range invalidsCnpjs {
		if cnpj == invalidCnpj {
			return true
		}
	}
	return false
}

func CleanCnpj(cpf string) []int{
	unmaskCnpj := strings.ReplaceAll(cpf, ".", "")
	unmaskCnpj = strings.ReplaceAll(unmaskCnpj, "-", "")
	unmaskCnpj = strings.ReplaceAll(unmaskCnpj, "/", "")
	strDigits := strings.Split(unmaskCnpj, "")

	digits := make([]int, 0)

	for _, strDigit := range strDigits {
		digit, _ := strconv.Atoi(strDigit)
		digits = append(digits, digit)
	}

	return digits
}

func CalculateCnpjDigit(digits []int, factor int) int {

	sum := 0

	for i:= 0; i < len(digits); i++ {
		sum += digits[i] * factor
		factor--
		if factor == 1 {
			factor = 9
		}
	}

	result := sum % 11

	if result == 0 || result == 1 {
		return 0
	}

	return 11 - result
}