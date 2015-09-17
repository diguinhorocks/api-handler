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

			if isComplex(i) {
				attachValue(leftAssignment, list[1:], rightAssignment[mainKey], list[len(list)-1], m, strings.Split(i, ".")[0])
			} else {
				attachValue(leftAssignment, list[1:], rightAssignment[mainKey], list[len(list)-1], m, "")
			}

		} else {

			if _, ok := rightAssignment[v]; ok == false {
				rightAssignment[v] = make(map[string]interface{})
			}

			attachValue(leftAssignment, []string{v}, rightAssignment, "", m, "")
		}

	}

	return rightAssignment

}

func attachValue(value interface{}, list []string, context interface{}, key string, mp map[string]string, mk string) {

	if len(list) > 1 {

		if reflect.TypeOf(context).Kind().String() == "map" {
			// item seguinte é um array
			if list[1] == "$" {

				if _, ok := context.(map[string]interface{})[list[0]]; !ok {
					context.(map[string]interface{})[list[0]] = make(map[string]interface{})
				}

				context.(map[string]interface{})[list[0]] = make([]interface{}, 0)

				if reflect.TypeOf(value).Kind().String() == "slice" {

					for i, v := range value.([]interface{}) {

						var m = make(map[string]interface{})

						keys := getKeyCandidates(mp, mk)

						for n, k := range keys {
							m[n] = v.(map[string]interface{})[k]
						}

						if len(context.(map[string]interface{})[list[0]].([]interface{})) == i {
							context.(map[string]interface{})[list[0]] = append(context.(map[string]interface{})[list[0]].([]interface{}), m)
						}

					}
				} else {

					for i, _ := range value.(map[string]interface{}) {

						attachValue(value.(map[string]interface{})[i], list, context, key, mp, mk)

					}

				}

			} else {

				if _, ok := context.(map[string]interface{})[list[0]]; !ok {
					context.(map[string]interface{})[list[0]] = make(map[string]interface{})
				}

				attachValue(value, list[1:], context.(map[string]interface{})[list[0]], list[len(list)-1], mp, mk)
			}

		} else {

			if reflect.TypeOf(context).Kind().String() == "slice" {

				for i, _ := range context.([]interface{}) {
					attachValue(value, list[1:], context.([]interface{})[i], list[len(list)-1], mp, mk)
				}
			}

		}

	} else {

		context.(map[string]interface{})[list[0]] = value

	}

}

func getKeyCandidates(m map[string]string, mainKey string) map[string]string {

	res := make(map[string]string)

	for i, x := range m {

		if isComplex(i) {

			str := strings.Split(i, ".")

			if str[0] == mainKey {
				for _, v := range str[1:] {
					if v != "$" {
						outerKey := strings.Split(x, ".")
						res[outerKey[len(outerKey)-1]] = v
					}
				}

			}
		}
	}

	return res

}

func findValue(attachment []string, context interface{}) interface{} {

	k := reflect.ValueOf(attachment[0]).Interface()

	var v interface{}

	switch t := reflect.ValueOf(context).Interface().(type) {
	case string, float64, int64, []interface{}, map[string]interface{}:

		if reflect.TypeOf(t).Kind().String() == "map" {
			v = t.(map[string]interface{})[k.(string)]
		} else {
			v = t
		}

	default:
		panic("Unknown type for key " + k.(string) + " on value " + reflect.TypeOf(t).String())
	}

	switch t := reflect.ValueOf(v).Interface().(type) {
	case string, float64, int64, []interface{}, map[string]interface{}:
		return t
	default:
		fmt.Println("unknown value", t)
	}

	return context

}

func getMapValue(m interface{}, key string) interface{} {

	if reflect.TypeOf(m).Kind().String() == "slice" {

		for i, _ := range m.([]interface{}) {

			return getMapValue(m.([]interface{})[i], key)

		}

	} else if reflect.TypeOf(m).Kind().String() == "map" {

		for i, v := range m.(map[string]interface{}) {

			if i != key {

				return getMapValue(v, key)

			}
		}

	}

	return m

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
