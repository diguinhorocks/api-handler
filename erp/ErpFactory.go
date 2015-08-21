package erp

import(
	"fmt"

	_ "labs/factory/lib"
)

type Erp struct {
	ais AbstractIntegrableFactory
}

func NewErpFactory (ais AbstractIntegrableFactory) *Erp {
	obj := new(Erp)
	obj.ais = ais
	return obj
}

func (this *Erp) Dispatch() string {
	fmt.Println("dispatching erp..")
	return ""
}

func (this *Erp) Resolve() string {
	fmt.Println("resolving erp..")
	return ""
}
