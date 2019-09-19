package services

type Type int

type Service interface {
	Validate() error
	Assert() error
}
