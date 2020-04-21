package render

import (
	"net/http"

	"github.com/go-chi/chi/middleware"

	"go.uber.org/zap"

	"github.com/go-chi/render"
)

// Render util with logging
type Render struct {
	logger *zap.Logger
}

// New renderer
func New(logger *zap.Logger) *Render {
	return &Render{logger}
}

// Response for successful response
type Response struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	ErrorCode int         `json:"errorCode,omitempty"`
	Error     error       `json:"-"`
	Logger    *zap.Logger `json:"-"`
}

// Render hook for response
func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	if res.Logger != nil && (res.Error != nil || res.ErrorCode > 0) {
		var err string
		if res.Error != nil {
			err = res.Error.Error()
		}
		res.Logger.Error(res.Message,
			zap.String("req_id", middleware.GetReqID(r.Context())),
			zap.Int("error_code", res.ErrorCode),
			zap.String("error", err),
		)
	}
	return nil
}

// Message maps and status and message into JSON formatted string
func (rr *Render) Message(status bool, message string) render.Renderer {
	return &Response{
		Status:  status,
		Message: message,
	}
}

// ErrorMessage maps status and message into JSON formatted string
func (rr *Render) ErrorMessage(errorCode int, err error, message string) render.Renderer {
	return &Response{
		Status:    false,
		Message:   message,
		ErrorCode: errorCode,
		Error:     err,
		Logger:    rr.logger,
	}
}

// DataMessage maps data, status and message into JSON formatted string
func (rr *Render) DataMessage(data interface{}, status bool, message string) render.Renderer {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// Respond encodes a JSON response to a http request
func (rr *Render) Respond(w http.ResponseWriter, r *http.Request, renderer render.Renderer) {
	render.Render(w, r, renderer)
}

// RespondWithStatus encodes a JSON response to a http request and modifies response status code
func (rr *Render) RespondWithStatus(w http.ResponseWriter, r *http.Request, statusCode int, renderer render.Renderer) {
	render.Status(r, statusCode)
	rr.Respond(w, r, renderer)
}
