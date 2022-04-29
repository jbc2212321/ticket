package database

type UserDao interface {
	ExistUser(phone, category int64) bool
	AddUser(user *User) error
	CheckUser(phone, category int64, password string) int64
}

type VatDao interface {
	AddVat(vat *VatInvoice) error
	GetVatByUserId(idList ...int64) ([]*VatInvoice, error)
	DelVatByTicketId(idList ...int64) error
}

type ImageDao interface {
	AddImage(img *Image) error
	GetImageByTicketId(idList ...int64) ([]*Image, error)
	ListImages() ([]*Image, error)
	DelImageByTicketId(idList ...int64) error
}
