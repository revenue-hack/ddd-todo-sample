package vo

import (
	"reflect"
	"regexp"

	"golang.org/x/xerrors"
)

type URL string

var (
	urlFormat = `^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`
	urlRegExp = regexp.MustCompile(urlFormat)
)

const urlMaxLength = 2048

func NewURL(url string) (URL, error) {
	if len(url) == 0 {
		return URL(""), xerrors.New("url must not be empty")
	}

	if len(url) > urlMaxLength {
		return URL(""), xerrors.Errorf("url must less than %d: %s", urlMaxLength, url)
	}

	if ok := urlRegExp.MatchString(url); !ok {
		return URL(""), xerrors.Errorf("invalid url format. url is %s", url)
	}

	return URL(url), nil
}

func (e URL) Value() string {
	return string(e)
}

func (e URL) Equals(e2 URL) bool {
	return reflect.DeepEqual(e, e2)
}
