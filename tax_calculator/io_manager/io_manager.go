package io_manager

type IOManager interface {
	ReadLines() ([]string, error)
	Write(data interface{}) error
}