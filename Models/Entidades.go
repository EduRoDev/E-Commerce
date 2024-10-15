package models

import "time"

type Cliente struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Nombre    string `gorm:"size:100;not null" json:"nombre"`
	Apellido  string `gorm:"size:100;not null" json:"apellido"`
	Email     string `gorm:"size:100;not null" json:"email"`
	Telefono  string `gorm:"size:100;not null" json:"telefono"`
	Direccion string `gorm:"size:100;not null" json:"direccion"`
}

type Clientes []Cliente

func (Cliente) TableName() string {
	return "clientes"
}

type Producto struct {
	Id          int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre      string  `gorm:"size:100;not null" json:"nombre"`
	Precio      float64 `gorm:"type:decimal(10,2);not null" json:"precio"`
	Cantidad    int     `gorm:"type:int;not null" json:"cantidad"`
	Categoria   string  `gorm:"type:enum('Electronicos','Alimentacion','Otros');not null" json:"categoria"`
	Descripcion string  `gorm:"type:text;not null" json:"descripcion"`
}

type Productos []Producto

func (Producto) TableName() string {
	return "productos"
}

type Orden struct {
	Id         int       `gorm:"primaryKey; autoIncrement"`
	Fecha      time.Time `gorm:"type:datetime;not null" json:"fecha"`
	IdCliente  int       `gorm:"type:int;not null" json:"id_cliente"`
	IdProducto int       `gorm:"type:int;not null" json:"id_producto"`
	Cliente    Cliente   `gorm:"foreignKey:IdCliente" json:"cliente"`
	Producto   Producto  `gorm:"foreignKey:IdProducto" json:"producto"`
	Cantidad   int       `gorm:"type:int;not null" json:"cantidad"`
	Total      float64   `gorm:"type:decimal(10,2);not null" json:"total"`
}

type Ordenes []Orden

func (Orden) TableName() string {
	return "ordenes"
}
