package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Product2 struct {
	ID          int64     `json:"id"`
	KngID       int64     `json:"kng_id"`
	PlanetID    int64     `json:"planet_id"`
	ProgressID  int64     `json:"progress_id"`
	EsscoID     int64     `json:"essco_id"`
	Article     string    `json:"article"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product2) KngIDVal() string {
	return fmt.Sprintf("%06d", p.KngID)
}
func (p *Product2) PlanetIDVal() string {
	return fmt.Sprintf("%06d", p.PlanetID)
}
func (p *Product2) ProgressIDVal() string {
	return fmt.Sprintf("%06d", p.ProgressID)
}
func (p *Product2) EsscoIDVal() string {
	return fmt.Sprintf("%06d", p.EsscoID)
}

type Product struct {
	ID int64 `json:"id"`
	//KngID       int64     `json:"kng_id"`
	//PlanetID    int64     `json:"planet_id"`
	//ProgressID  int64     `json:"progress_id"`
	//EsscoID     int64     `json:"essco_id"`
	KngCode      *ProductCode
	PlanetCode   *ProductCode
	ProgressCode *ProductCode
	EsscoCode    *ProductCode
	Article      sql.NullString `json:"article"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type ProductCode struct {
	ID           int64          `json:"id"`
	ProductID    int64          `json:"product_id"`
	SupID        int64          `json:"sup_id"`
	SupProductID sql.NullInt64  `json:"sup_product_id"`
	SupCode      sql.NullString `json:"sup_code"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func (p *Product) KngIDVal() string {
	//return fmt.Sprintf("%06d", p.KngID)
	if p.KngCode != nil {
		return p.KngCode.SupCode.String
	}
	return ""
}
func (p *Product) PlanetIDVal() string {
	//return fmt.Sprintf("%06d", p.PlanetID)
	if p.PlanetCode != nil {
		return p.PlanetCode.SupCode.String
	}
	return ""
}
func (p *Product) ProgressIDVal() string {
	//return fmt.Sprintf("%06d", p.ProgressID)
	if p.ProgressCode != nil {
		return p.ProgressCode.SupCode.String
	}
	return ""
}
func (p *Product) EsscoIDVal() string {
	//return fmt.Sprintf("%06d", p.EsscoID)
	if p.EsscoCode != nil {
		return p.EsscoCode.SupCode.String
	}
	return ""
}
