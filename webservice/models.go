package webservice

type httpResponse struct {
	Meta httpResponseMeta `json:"meta"`
	Data interface{}      `json:"data"`
}

type httpResponseMeta struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
