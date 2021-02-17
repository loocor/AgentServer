package model

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

var (
	ErrNotFound             = sqlx.ErrNotFound
	ErrNotMatchDestination  = sqlx.ErrNotMatchDestination
	ErrNotReadableValue     = sqlx.ErrNotReadableValue
	ErrNotSettable          = sqlx.ErrNotSettable
	ErrUnsupportedValueType = sqlx.ErrUnsupportedValueType
)
