package services

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-default/models"
	"golang-default/ws"
)

type GPSService struct {
	db *sql.DB
}

func NewGPSService(db *sql.DB) *GPSService {
	return &GPSService{db: db}
}

func (s *GPSService) InsertGPS(data models.GPSData) (int64, error) {
	if data.Latitude == 0 && data.Longitude == 0 {
		return 0, errors.New("invalid gps data")
	}
	result, err := s.db.Exec(
		` INSERT INTO gps_logs 
    (unit_id, latitude, longitude, speed, bearing, accuracy, altitude) 
    VALUES (?, ?, ?, ?, ?, ?, ?)`,
		data.UnitId,
		data.Latitude,
		data.Longitude,
		data.Speed,
		data.Bearing,
		data.Accuracy,
		data.Altitude,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert gps: %w", err)
	}

	_, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return 0, nil
}

func (s *GPSService) InsertLastGPS(data models.GPSData) error {
	if data.Latitude == 0 && data.Longitude == 0 {
		return errors.New("invalid gps data")
	}

	_, err := s.db.Exec(`
		INSERT INTO last_gps (
			unit_id, latitude, longitude, speed, bearing, accuracy, altitude
		) VALUES (?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			latitude = VALUES(latitude),
			longitude = VALUES(longitude),
			speed = VALUES(speed),
			bearing = VALUES(bearing),
			accuracy = VALUES(accuracy),
			altitude = VALUES(altitude),
			updated_at = CURRENT_TIMESTAMP
	`,
		data.UnitId,
		data.Latitude,
		data.Longitude,
		data.Speed,
		data.Bearing,
		data.Accuracy,
		data.Altitude,
	)
	if err != nil {
		return fmt.Errorf("failed to upsert gps: %w", err)
	}

	// Kirim ke semua client via WebSocket
	ws.GetHub().Broadcast("gps", data)

	return nil
}
