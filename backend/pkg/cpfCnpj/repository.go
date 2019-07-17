package cpfCnpj

type Repository interface {
	Save(cpfCnpj CpfCnpj) error
	GetDocuments(isBlacklist bool) ([]CpfCnpj, error)
	MoveToBlacklist(documentId string) error
	DeleteDocument(documentId string) error
	RemoveFromBlacklist(documentId string) error
	FindByDocument(document string) ([]CpfCnpj, error)
}