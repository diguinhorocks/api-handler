package factory

import (
	"io"
)

type AbstractIntegrableFactory interface {
	GetName() string
	GetType() string
	SetConfigs(c io.Reader)
	Dispatch() map[string]interface{}
	Resolve() string
}

type Integrable struct {
	ais AbstractIntegrableFactory
}

func Factory(ais AbstractIntegrableFactory) *Integrable {
	obj := new(Integrable)
	obj.ais = ais
	return obj
}

func (this *Integrable) SetConfigs(c io.Reader) {
	this.ais.SetConfigs(c)
}

func (this *Integrable) Dispatch() map[string]interface{} {
	return this.ais.Dispatch()
}

func (this *Integrable) Resolve() string {
	return this.ais.Resolve()
}
