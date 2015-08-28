package builder

import (
	"fmt"
	"regexp"
	"strings"
	"reflect"

	//"github.com/fatih/structs"
)

type ConcreteMappable interface {
	GetName() 	 string
	GetContext() interface{}
}

type RequestContainer struct {
	Context map[string]interface{}
	Type 	string
	Subject string
	Map 	map[string]string
}

func SetMap (t string, m map[string]string, context map[string]interface{}) {
	fmt.Println("configuring map " + t)
	
	buildMap(m, context)
}

// func GetMap () (map[string]interface{}) {
// 	return Map
// }

func buildMap (m map[string]string, c map[string]interface{}) {

	buildedContext := make(map[string]interface{})

	for i, _ := range c {

		var key string

		key = m[i]

		if isDotted(m[i]) {
			key = strings.Split(m[i], ".")[0]
		}

		buildedContext[key] = reflect.ValueOf(parse(buildedContext[i], i, m[i], c[i])).Interface().(map[string]interface{})

	}

	fmt.Println("OPA", buildedContext)

}

func parse (buildedContext interface{}, i string, m string, c interface{}) interface{} {

	if buildedContext == nil {

		buildedContext := make(map[string]interface{})

		externalKey := reflect.ValueOf(m).Interface().(string)

		if isDotted(externalKey) {

			s := strings.Split(externalKey, ".")

			//result := make(map[string]interface{})

			for x, z := range s {
				//primeiro infice não pega
				if (x > 0) {

					var currentIndexContextValue interface{}
					var currentIndexMapValue interface{}

					switch t := reflect.ValueOf(c).Interface().(type) {

						case string, int64, float64:
							currentIndexContextValue = t
						default:
							panic("Unknown value type")

					}

					currentIndexMapValue = reflect.ValueOf(m).Interface().(string)
					
					buildedContext = buildAttribute(buildedContext, z, currentIndexContextValue, currentIndexMapValue)

					fmt.Println("123", buildedContext)
				}

			}

		} else {
			buildedContext = buildAttribute(buildedContext, i, c, m)
		}

		return buildedContext

	} 


	return buildedContext
	

}

func isDotted (value string) bool {
	match, _ := regexp.MatchString("\\.", value)
	return match
}

func buildAttribute (container map[string]interface{}, key string, context interface{}, m interface{}) map[string]interface{} {

	t := reflect.TypeOf(context)

	//k := reflect.ValueOf(m).Interface().(string)

	if t.String() == "map[string]interface {}" || t.String() == "map[string]string" {

		//convertendo para o tipo pertinente
		test := reflect.ValueOf(context).Interface().(map[string]interface{})

		result := make(map[string]interface{})

		for i, v := range test {
			
			result[i] = v
		}

		fmt.Println("TESTE", result)

		container[key] = result


	} else {

		fmt.Println(key)

		container[key] = context
	}

	return container

}

func buildComplexKey (key string) {

}