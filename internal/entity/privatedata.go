package entity

const (
	CreditCardDataType   = "credit_card"
	BinaryDataType       = "binary_data"
	TextDataType         = "text_data"
	LoginAndPassDataType = "login_password"
)

type PrivateData struct {
	ID       string
	UserID   string
	Name     string
	DataType string
	Data     []byte
	MetaInfo string
}
