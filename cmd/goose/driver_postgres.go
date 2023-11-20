//go:build !no_sqlanywhere
// +build !no_sqlanywhere

package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)
