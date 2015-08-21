package resources

import(
	"fmt"
)

type AbstractIntegrableFactory struct {}


func (this *RestErp) GetName() string {
	return this.name
}

func (this *RestErp) GetType() string {
	return this.erp_type
}

func (this *RestErp) SetConfigs() string {
	fmt.Println("configuring erp..")
	return ""
}

func (this *RestErp) UpdatePrice() string {
	fmt.Println("dispatching erp..")
	return ""
}

func (this *RestErp) UpdateStock() string {
	fmt.Println("resolving erp..")
	return ""
}
