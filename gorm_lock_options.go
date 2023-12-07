package gormlock

import (
	"context"
	"time"
)

type LockOption func(*gormLocker)

func WithJobIdentifier(f func(ctx context.Context, key string) string) LockOption {
	return func(l *gormLocker) {
		l.jobIdentifier = f
	}
}

func WithDefaultJobIdentifier(precision time.Duration) LockOption {
	return func(l *gormLocker) {
		l.jobIdentifier = defaultJobIdentifier(precision)
	}
}

// WithTTL when the locker records in the database exceeds the ttl, it is cleaned up.
// to avoid excessive data in the database.
func WithTTL(ttl time.Duration) LockOption {
	return func(l *gormLocker) {
		l.ttl = ttl
	}
}

// WithCleanInterval the time interval to run clean operation.
func WithCleanInterval(interval time.Duration) LockOption {
	return func(l *gormLocker) {
		l.interval = interval
	}
}
