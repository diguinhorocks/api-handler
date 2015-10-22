package response

import (
	"encoding/json"
	"github.com/clbanning/mxj"
)

type Map map[string]interface{}

func To(t string, content map[string]interface{}) ([]byte, error) {

	if t == "xml" {

		m, err := mxj.AnyXml(content)

		if err != nil {
			panic(err)
		}

		return m, err

	} else if t == "json" {
		return json.Marshal(content)
	}

	return nil, nil

}
