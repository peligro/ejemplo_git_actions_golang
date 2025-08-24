package modelos

import (
	"proyecto-golang/database"
)

type Categoria struct {
    Id     uint   `json:"id" gorm:"primaryKey"`  // ← Agregar primaryKey
    Nombre string `gorm:"type:varchar(100)" json:"nombre"`
    Slug   string `gorm:"type:varchar(100)" json:"slug"`
}

type Categorias []Categoria



func Migraciones() {
	database.Database.AutoMigrate(&Categoria{})

}