package services

import (
	"database/sql"
	"fmt"
	"golang-default/models"
)

type UnitService struct {
	db *sql.DB
}

func NewUnitService(db *sql.DB) *UnitService {
	return &UnitService{db: db}
}

func (s *UnitService) CreateUnit(unit models.UnitData) (int64, error) {
	// ✅ 1. Validasi UnitCode
	var exists bool
	err := s.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM units WHERE unit_code = ?)`, unit.UnitCode).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("failed to check UnitCode: %w", err)
	}
	if exists {
		return 0, fmt.Errorf("UnitCode already registered")
	}

	result, err := s.db.Exec(
		`INSERT INTO units (unit_code, unit_type,name, description) VALUES (?,?,?,?)`,
		unit.UnitCode,
		unit.UnitType,
		unit.Name,
		unit.Description,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert unit: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

// Get Unit by ID
func (s *UnitService) GetUnitByID(id int64) (models.UnitData, error) {
	var unit models.UnitData
	err := s.db.QueryRow(`SELECT id,unit_code,unit_type,name,description FROM units WHERE id = ?`, id).
		Scan(&unit.ID, &unit.UnitCode, &unit.UnitType, &unit.Name, &unit.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return unit, fmt.Errorf("unit not found")
		}
		return unit, fmt.Errorf("failed to get unit: %w", err)
	}
	return unit, nil
}

// Update unit
func (s *UnitService) UpdateUnit(unit models.UnitData) error {
	_, err := s.db.Exec(`UPDATE units SET unit_code = ?, unit_type = ?, name = ?, description = ? WHERE id = ?`, unit)
	if err != nil {
		return fmt.Errorf("failed to update unit: %w", err)
	}
	return nil
}

// Delete unit
func (s *UnitService) DeleteUnit(id int64) error {
	_, err := s.db.Exec(`DELETE FROM units WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete unit: %w", err)
	}
	return nil
}
