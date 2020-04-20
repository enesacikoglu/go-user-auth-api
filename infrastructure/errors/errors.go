package errors

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

const (
	resourceNotFoundTitle    = "Not found"
	badRequestTitle          = "Bad Request"
	resourceNotFoundMessage  = "Resource not found"
	internalServerErrorTitle = "Internal Server Error"
)

// Echo Error handler
func CustomEchoHTTPErrorHandler(err error, c echo.Context) {

	var (
		customError = &CustomError{
			Detail: err.Error(),
		}
	)

	if ce, ok := err.(*CustomError); ok {
		customError = ce
	} else if he, ok := err.(*echo.HTTPError); ok {
		customError.Status = he.Code
	} else {
		customError.Status = http.StatusInternalServerError
		customError.Title = internalServerErrorTitle
		customError.Message = err.Error()
	}
	customError.RequestMethod = c.Request().Method
	customError.RequestUri = c.Request().RequestURI
	customError.Instant = time.Now()

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(customError.Status)
		} else {
			err = c.JSON(customError.Status, customError)
		}
		request := c.Request()
		logFormat := map[string]interface{}{
			"real-ip":           c.RealIP(),
			"host":              request.Host,
			"method":            request.Method,
			"request-uri":       request.RequestURI,
			"response-status":   customError.Status,
			"request-referer":   request.Referer(),
			"request-useragent": request.UserAgent(),
		}

		log.Error(logFormat, "%s", customError.Detail)
	}

}

type CustomError struct {
	Title         string    `json:"title"`
	Status        int       `json:"status"`
	Detail        string    `json:"detail"`
	RequestUri    string    `json:"requestUri"`
	RequestMethod string    `json:"requestMethod"`
	Instant       time.Time `json:"instant"`
	Message       string    `json:"message"`
}

func (err CustomError) Error() string {
	return err.Detail
}

func BadRequest(detail string) error {
	return makeError(http.StatusBadRequest, detail, badRequestTitle, "")
}

func InternalServerError(detail string) error {
	return makeError(http.StatusInternalServerError, detail, internalServerErrorTitle, "")
}

func NotFound(detail string) error {
	return makeError(http.StatusNotFound, detail, resourceNotFoundTitle, resourceNotFoundMessage)
}

func makeError(code int, detail string, title string, message string) error {
	return &CustomError{
		Title:   title,
		Message: message,
		Status:  code,
		Detail:  detail,
		Instant: time.Now(),
	}
}
