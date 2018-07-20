package utils

// Response struct
type Response struct {
	Payload map[string](interface{})
}

// NewResponse : create a new response
func NewResponse(k string, v interface{}) Response {
	resp := Response{(map[string](interface{}){})}
	resp.Payload[k] = v

	return resp
}

// Set : add a new key/value to the reponses payload
func (r *Response) Set(k string, v interface{}) {

	r.Payload[k] = v

}
