package cpfCnpj

type Repository interface {
	Save(cpfCnpj CpfCnpj) error
	GetAllDocuments() ([]CpfCnpj, error)
}