package database

type UserDao interface {
	ExistUser(phone, category int64) bool
	AddUser(user *User) error
	CheckUser(phone, category int64, password string) int64
}

type VatDao interface {
	AddVat(vat *VatInvoice) error
}

type ImageDao interface {
	AddImage(img *Image) error
	GetImageByTicketId(idList ...int64) ([]*Image, error)
}
