package factory

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"fabric/factory/builder"
	"fabric/util/request"
	"fabric/util/response"
	//"fabric/factory/builder/types"
)

type SoapErp struct {
	name             string
	erp_type         string
	erp_content_type string
	Response         map[string]interface{}
}

func (this *SoapErp) GetName() string {
	return this.name
}

func (this *SoapErp) GetType() string {
	return this.erp_type
}

func (this *SoapErp) GetContentType() string {
	return this.erp_content_type
}

func (this *SoapErp) SetConfigs(config io.Reader) {

	var rc request.RequestContainer

	decoder := json.NewDecoder(config)
	err := decoder.Decode(&rc)

	if err != nil {
		panic("Not Supported JSON data")
	}

	this.name = rc.Config["name"]
	this.erp_content_type = rc.Config["type"]

	this.Response = builder.SetMap(rc.Type, rc.Map, rc.Context)
}

func (this *SoapErp) GetResponse() []byte {

	r, err := response.To(this.GetContentType(), this.Response)

	if err != nil {
		panic(err)
	}

	return r
}

func (this *SoapErp) UpdatePrice(price float64) string {
	return "updating price to " + strconv.FormatFloat(price, 'f', 2, 64)
}

func (this *SoapErp) UpdateStock(quantity int64) string {
	return "update stock to.." + strconv.FormatInt(quantity, 10)
}

func (this *SoapErp) Dispatch() string {
	fmt.Println("dispatched to " + this.GetName() + " server")

	return string(this.GetResponse())
}

func (this *SoapErp) Resolve() string {
	fmt.Println(this.UpdatePrice(5.54))
	fmt.Println(this.UpdateStock(6))

	return "resolved the erp content to database.."
}
