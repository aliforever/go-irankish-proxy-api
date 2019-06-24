package responses

import "encoding/json"

type Response struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Result  interface{} `json:"object,omitempty"`
}

func (r *Response) JSON() string {
	j, _ := json.Marshal(&r)
	return string(j)
}
