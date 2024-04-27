package domain

import "time"

type Country struct {
	Id            int64
	Name          string
	CreatedAt     time.Time
	CreatedBy     int64
	CreatedByName string
	UpdatedAt     time.Time
	UpdatedBy     int64
	UpdatedByName string
	IsDeleted     bool
}
