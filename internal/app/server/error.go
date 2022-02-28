package server

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	_validator "wager-be/pkg/validator"
)

type errResponse struct {
	Description string `json:"error"`
}

const (
	validationErrorFmt = "field %s - error: %s"
)

func (s *server) responseErr(ctx *gin.Context, code int, err error) {
	var ve validator.ValidationErrors
	if ok := errors.As(err, &ve); ok {
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = fmt.Sprintf(validationErrorFmt, fe.Field(), _validator.GetErrorMsg(fe))
		}
		ctx.JSON(code, &errResponse{
			Description: strings.Join(out, "\n"),
		})
		return
	}

	// other error types
	ctx.JSON(code, &errResponse{
		Description: err.Error(),
	})
}
