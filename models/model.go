package models

import "time"

type Paper struct {
	ID            uint      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Judul         string    `json:"judul"`
	Deskripsi     string    `json:"deskripsi"`
	Abstrak       string    `json:"abstrak"`
	Link          string    `json:"link"`
	FilePaper     string    `json:"file_paper"`
	Author        string    `json:"author"`
	TanggalTerbit time.Time `json:"tanggal_terbit"`
}

type File struct {
	ID       int    `gorm:"primaryKey;autoIncrement:true"`
	Filename string `gorm:"not null"`
	UUID     string `gorm:"unique;not null"`
}
