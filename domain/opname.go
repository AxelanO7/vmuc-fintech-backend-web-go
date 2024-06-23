package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Opname struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	IDOpaname string         `json:"id_opaname"`
	Name      string         `json:"name"`
	StartDate string         `json:"start_date"`
	EndDate   string         `json:"end_date"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ResByDate struct {
	In  []In  `json:"in"`
	Out []Out `json:"out"`
	Rtr []Rtr `json:"rtr"`
}

type ReqByDate struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// type StockOpname struct {
// 	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
// 	IdStuff   uint           `json:"id_stuff"`
// 	Name      string         `json:"name"`
// 	Type      string         `json:"type"`
// 	Quantity  float64        `json:"quantity"`
// 	Unit      string         `json:"unit"`
// 	Price     float64        `json:"price"`
// 	CreatedAt *time.Time     `json:"created_at"`
// 	UpdatedAt *time.Time     `json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
// }

type OpnameRepository interface {
	RetrieveAllOpname() ([]Opname, error)
	RetrieveOpnameByID(id uint) (*Opname, error)
	CreateOpname(opname *Opname) error
	RetriveByStartDateEndDate(startDate, endDate string) ([]In, []Out, []Rtr, error)
}

type OpnameUseCase interface {
	AddOpname(ctx context.Context, opname *Opname) error
	FetchOpnames(ctx context.Context) ([]Opname, error)
	FetchOpnameByID(ctx context.Context, id uint) (*Opname, error)
	FetchOpnameByDate(ctx context.Context, startDate, endDate string) (*ResByDate, error)
}
