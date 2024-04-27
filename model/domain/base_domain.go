package domain

import "time"

type BaseDomain struct {
	CreatedAt     time.Time
	CreatedBy     int64
	CreatedByName string
	UpdatedAt     time.Time
	UpdatedBy     int64
	UpdatedByName string
	IsDeleted     bool
}
