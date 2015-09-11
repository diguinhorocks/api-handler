package builder

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	//"github.com/fatih/structs"
)

type ConcreteMappable interface {
	GetName() string
	GetContext() interface{}
}

type RequestContainer struct {
	Context map[string]interface{}
	Type    string
	Subject string
	Map     map[string]string
}

type M map[string]interface{}

func SetMap(t string, m map[string]string, context map[string]interface{}) map[string]interface{} {
	fmt.Println("configuring map " + t)

	mp := make(M)

	return buildMap(mp, m, context)
}

func (m *M) GetMap() M {
	return *m
}

func buildMap(mp M, m map[string]string, c map[string]interface{}) map[string]interface{} {

	//valor do caminho do recurso da base de dados
	var leftAssignment interface{}
	//recurso externo
	var rightAssignment map[string]interface{}

	for i, v := range m {

		if isDotted(i) {

			mainKey := strings.Split(i, ".")[0]
			list := strings.Split(i, ".")

			leftAssignment = findValue(list[1:], c[mainKey])

		} else {

			leftAssignment = c[i]

		}

		if len(rightAssignment) < 1 {
			rightAssignment = make(map[string]interface{})
		}

		if isDotted(v) {

			mainKey := strings.Split(v, ".")[0]
			list := strings.Split(v, ".")

			if _, ok := rightAssignment[mainKey]; ok == false {
				rightAssignment[mainKey] = make(map[string]interface{})
			}

			attachValue(leftAssignment, list[1:], rightAssignment[mainKey])

		} else {

			if _, ok := rightAssignment[v]; ok == false {
				rightAssignment[v] = make(map[string]interface{})
			}

			attachValue(leftAssignment, []string{v}, rightAssignment)
		}

	}

	return rightAssignment

}

func attachValue(value interface{}, list []string, context interface{}) {

	if len(list) > 1 {

		if _, ok := context.(map[string]interface{})[list[0]]; !ok {
			context.(map[string]interface{})[list[0]] = make(map[string]interface{})
		}

		attachValue(value, list[1:], context.(map[string]interface{})[list[0]])

	} else {

		if reflect.TypeOf(value).Kind().String() == "slice" || reflect.TypeOf(value).Kind().String() == "map" {

			if reflect.TypeOf(value).Kind().String() == "slice" {

				for _, v := range value.([]interface{}) {

					if reflect.TypeOf(v).Kind().String() == "string" {

						if _, ok := context.(map[string]interface{})[list[0]]; !ok {
							context.(map[string]interface{})[list[0]] = make([]string, 0)
						}

						context.(map[string]interface{})[list[0]] = append(context.(map[string]interface{})[list[0]].([]string), v.(string))

					} else {

						for x, z := range v.(map[string]interface{}) {

							if _, ok := context.(map[string]interface{})[list[0]]; !ok {
								context.(map[string]interface{})[list[0]] = make([]float64, 0)
							}

							if x == list[0] {
								context.(map[string]interface{})[list[0]] = append(context.(map[string]interface{})[list[0]].([]float64), z.(float64))
							}
						}
					}

				}

			} else if reflect.TypeOf(value).Kind().String() == "map" {

			}

		} else {
			context.(map[string]interface{})[list[0]] = value
		}

	}

}

func findValue(attachment []string, context interface{}) interface{} {

	c := convert(context)

	k := convert(attachment[0])

	fmt.Println(attachment)

	if len(attachment) > 1 {

		if attachment[1] == "$" {

			fmt.Println(attachment)

			for _, v := range context.(map[string][]map[string]string)[k.(string)] {
				if reflect.TypeOf(v).Kind().String() == "map" {
					for i, _ := range v {
						if i == attachment[2] {
							fmt.Println(v[i])
							return v[i]
						}
					}
				}
			}

			//attachment = append(attachment[:1], attachment[1-1])
		}

	} else {

		var v interface{}

		switch t := reflect.ValueOf(c).Interface().(type) {
		case string:
			v = t
		case int64:
			v = t
		case float64:
			v = t
		case map[string]string:
			v = t[k.(string)]
		case map[string]interface{}:
			v = t[k.(string)]
		case map[string][]string:
			v = t
		case map[string][]map[string]string:
			v = t[k.(string)]
		case map[string][]map[string]float64:
			v = t[k.(string)]
		case map[string][]map[string]int64:
			v = t[k.(string)]
		case map[string]map[string]map[string]string:
			v = t[k.(string)]
		case map[string]int64:
			v = t[k.(string)]
		case map[string]map[string]int64:
			v = t[k.(string)]
		case map[string]map[string]map[string]int64:
			v = t[k.(string)]
		case map[string]float64:
			v = t[k.(string)]
		case map[string]map[string]float64:
			v = t[k.(string)]
		case map[string]map[string]map[string]float64:
			v = t[k.(string)]
		default:
			panic(reflect.TypeOf(t))
		}

		if len(attachment) > 1 {
			return findValue(attachment[1:], v)

		} else {
			switch t := reflect.ValueOf(v).Interface().(type) {
			case map[string]float64:
				return t[attachment[0]]
			case map[string]int64:
				return t[attachment[0]]
			case map[string]string:
				return t[attachment[0]]
			case map[string][]string:
				return t[attachment[0]]
			case string, float64, int64:
				return t
			default:
				panic("Unknown values type")
			}
		}

	}

	return c

}

func convert(value interface{}) interface{} {

	switch t := reflect.ValueOf(value).Interface().(type) {
	case string:
		value = reflect.ValueOf(t).Interface().(string)
	case int64:
		value = reflect.ValueOf(t).Interface().(int64)
	case float64:
		value = reflect.ValueOf(t).Interface().(float64)
	case map[string]interface{}:
		value = reflect.ValueOf(t).Interface().(map[string]interface{})
	case []string:
		value = reflect.ValueOf(t).Interface().([]string)
	case []interface{}:
		value = reflect.ValueOf(t).Interface().([]interface{})
	case map[string][]map[string]string:
		value = reflect.ValueOf(t).Interface().(map[string][]map[string]string)
	case map[string]string:
		value = reflect.ValueOf(t).Interface().(map[string]string)
	case map[string][]string:
		value = reflect.ValueOf(t).Interface().(map[string][]string)
	case map[string]map[string]string:
		value = reflect.ValueOf(t).Interface().(map[string]map[string]string)
	case map[string]map[string]map[string]string:
		value = reflect.ValueOf(t).Interface().(map[string]map[string]map[string]string)
	case map[string]int64:
		value = reflect.ValueOf(t).Interface().(map[string]int64)
	case map[string]map[string]int64:
		value = reflect.ValueOf(t).Interface().(map[string]map[string]int64)
	case map[string]map[string]map[string]int64:
		value = reflect.ValueOf(t).Interface().(map[string]map[string]map[string]int64)
	case map[string]float64:
		value = reflect.ValueOf(t).Interface().(map[string]float64)
	case map[string]map[string]float64:
		value = reflect.ValueOf(t).Interface().(map[string]map[string]float64)
	case map[string]map[string]map[string]float64:
		value = reflect.ValueOf(t).Interface().(map[string]map[string]map[string]float64)
	default:
		panic(reflect.TypeOf(t))
	}

	return value

}

func parseDotted() {

}

func parseList() {

}

func parseComplex() {

}

func isDotted(value string) bool {
	match, _ := regexp.MatchString("\\.", value)
	return match
}

func isComplex(value string) bool {
	return isDotted(value) && isList(value)
}

func isList(value string) bool {
	match, _ := regexp.MatchString("\\$", value)
	return match
}
