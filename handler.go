package main

type HTTPHandler struct{}

// NewHTTPHandler :nodoc:
func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{}
}

func (h *HTTPHandler) FindUserByIDFromDB() {

}

func (h *HTTPHandler) FindUserByIDFromES() {

}
