package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(255)" json:"namaProduct"`
	Deskripsi   string `gorm:"type:text" json:"deskripssi"`
}
