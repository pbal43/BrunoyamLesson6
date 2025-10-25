package domain

import "time"

const (
	ContextTimeout    = 5 * time.Second
	LeewayTimeout     = 60 * time.Second
	AccessTTL         = 15 * time.Minute
	RefreshTTL        = 24 * 7 * time.Hour
	ReadHeaderTimeout = 15 * time.Second
	RefreshAge        = 3600 * 24 * 7
)
