package factory

import(
	"fmt"
	"strconv"
	"encoding/json"
	"io"
	
	"fabric/factory/builder"
	//"fabric/factory/builder/types"
)

type RestErp struct {
	name 	 string
	erp_type string
	Response map[string]interface{}
}

func (this *RestErp) GetName() string {
	return this.name
}

func (this *RestErp) GetType() string {
	return this.erp_type
}

func (this *RestErp) SetConfigs(config io.Reader) {

	var rc builder.RequestContainer

	decoder := json.NewDecoder(config)
	err := decoder.Decode(&rc)

	if err != nil {
		panic("Not Supported data")
	}

	builder.SetMap(rc.Type, rc.Map, rc.Context)
}

func (this *RestErp) GetResponse() string {
	//fmt.Println(builder.GetMap())
	return "oi"
}

func (this *RestErp) UpdatePrice(price float64) string {
	return "updating price to " + strconv.FormatFloat(price, 'f', 2, 64)
}

func (this *RestErp) UpdateStock(quantity int64) string {
	return "update stock to.." + strconv.FormatInt(quantity, 10)
}

func (this *RestErp) Dispatch() string {
	s := this.GetResponse()
	return s
}

func (this *RestErp) Resolve() string {
	fmt.Println(this.UpdatePrice(5.54))
	fmt.Println(this.UpdateStock(6))

	return "resolved the erp content to database.."
}


