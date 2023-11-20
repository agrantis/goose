//go:build !no_sqlanywhere && windows
// +build !no_sqlanywhere,windows

package main

import (
	_ "github.com/a-palchikov/sqlago"
)
