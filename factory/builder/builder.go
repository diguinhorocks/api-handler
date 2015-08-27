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

func buildMap (m map[string]string, c map[string]interface{}){

	var buildedContext = make(map[string]interface{})

	for i, _ := range c {

		if isDotted(externalKey) {

			s := strings.Split(externalKey, ".")

			fmt.Println("composto", s[0])

		} else {
			buildAttribute(buildedContext, i, c[i], m[i])
		}
	}

	fmt.Println(buildedContext)

}

func isDotted (value string) bool {
	match, _ := regexp.MatchString("\\.", value)
	return match
}

func buildAttribute (container map[string]interface{}, key string, context interface{}, m interface{}) {

	externalKey := reflect.ValueOf(m).Interface().(string)

	t := reflect.TypeOf(context)

	if t.String() == "map[string]interface {}" || t.String() == "map[string]string" {

		//convertendo para o tipo pertinente
		test := reflect.ValueOf(context).Interface().(map[string]interface{})

		result := make(map[string]interface{})

		for i, v := range test {
			
			result[i] = v
		}

		container[key] = result


	} else {

		container[k] = context
	}

}

func buildComplexKey (key string) {

}