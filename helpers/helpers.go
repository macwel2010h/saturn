package helpers

import (
	"errors"
	"net/http"
	"serv-test/config"

	"github.com/go-playground/form/v4"
)

func DecodeForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = config.App.FormDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError
		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}

	return nil
}

func IsAuthenticated(r *http.Request) bool {
	return config.App.SessionManager.Exists(r.Context(), "authenticatedUserID")
}
