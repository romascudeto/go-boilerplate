package helpers

type Response struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewResponse(meta interface{}, data interface{}) (response Response) {
	newResponse := Response{
		Meta: meta,
		Data: data,
	}
	return newResponse
}
