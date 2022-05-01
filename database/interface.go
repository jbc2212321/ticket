package database

type UserDao interface {
	ExistUser(phone, category int64) bool
	AddUser(user *User) error
	CheckUser(phone, category int64, password string) int64
	GetUserById(idList ...int64) ([]*User, error)
	ListUser() ([]*User, error)
	DelUserById(idList ...int64) error
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

type SongsStoreDao interface {
	AddSong(song *SongsStore) error
}

type VerifyListDao interface {
	GetVerifyByID(idList ...int64) ([]*Verifylist, error)
	GetVerifyByStatus(statuList ...int64) ([]*Verifylist, error)
	AddVerifyList(verify *Verifylist) error
}

type TrainDao interface {
	AddTrain(train *Train) error
	GetTrainByUserId(idList ...int64) ([]*Train, error)
	DelTrainByTicketId(idList ...int64) error
}

type LogDao interface {
	AddLog(log *Log) error
	ListLog() ([]*Log, error)
	DelLogById(idList ...int64) error
}
