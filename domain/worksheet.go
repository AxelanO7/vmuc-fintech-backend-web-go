package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Worksheet struct {
	ID            uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	NameWorksheet string         `json:"name_worksheet"`
	Date          string         `json:"date"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type WorksheetRepository interface {
	RetrieveAllWorksheet() ([]Worksheet, error)
	RetrieveWorksheetByID(id uint) (*Worksheet, error)
	GetWorksheetByWorksheetPeriodeId(id uint) ([]Worksheet, error)
	CreateWorksheet(req *Worksheet) (*Worksheet, error)
	CreateBulkWorksheet(req []*Worksheet) ([]*Worksheet, error)
	UpdateWorksheet(req *Worksheet) (*Worksheet, error)
	UpdateBulkWorksheet(req []*Worksheet) ([]*Worksheet, error)
	DeleteWorksheet(id uint) error
}

type WorksheetUseCase interface {
	FetchWorksheets(ctx context.Context) ([]Worksheet, error)
	FetchWorksheetByID(ctx context.Context, id uint) (*Worksheet, error)
	AddWorksheet(ctx context.Context, req *Worksheet) (*Worksheet, error)
	AddBulkWorksheet(ctx context.Context, req []*Worksheet) ([]*Worksheet, error)
	EditWorksheet(ctx context.Context, req *Worksheet) (*Worksheet, error)
	EditBulkWorksheet(ctx context.Context, req []*Worksheet) ([]*Worksheet, error)
	DeleteWorksheet(ctx context.Context, id uint) error
}
