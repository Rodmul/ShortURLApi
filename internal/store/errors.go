package store

import "ShortLink/internal/genErr"

var (
	ErrTransactionCreationFailed = genErr.New("transaction creation failed")
	ErrTransactionFailed         = genErr.New("transaction failed, rollback")
	ErrTransactionRollbackFailed = genErr.New("transaction rollback failed")
	ErrTransactionCommitFailed   = genErr.New("transaction commit failed")
	ErrScanStructFailed          = genErr.New("failed to Scan structure")
)
