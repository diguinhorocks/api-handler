package request

type RequestContainer struct {
	Context map[string]interface{}
	Config  map[string]string
	Type    string
	Subject string
	Map     map[string]string
}
