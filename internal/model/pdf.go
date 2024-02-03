/*
Copyright © 2024 nu12
*/

package model

import (
	"io"
)

type PDF struct{}

type CompressionProfile struct {
	Value int
}

type MarginProfile struct {
	X int
	Y int
}

type ModeProfile struct {
	Format      string
	Orientation string
	Width       int
	Height      int
}

type Imager interface {
	Compress(CompressionProfile) (io.Reader, error)
	GetType() string
	GetFilename() string
}
