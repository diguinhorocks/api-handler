package lib

type AbstractIntegrableFactory interface {
	GetName() string
	GetType() string
	SetConfigs() string
	Dispatch() string
	Resolve() string
}