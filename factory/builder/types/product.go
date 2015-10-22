package types

//mudar para o tipo correto posteriormente
type ProductMap struct {
	Name  string
	Attrs map[string]interface{}
}

func (this *ProductMap) GetName() string {
	return this.Name
}

func (this *ProductMap) GetContext() interface{} {
	return this.Attrs
}
