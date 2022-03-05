package binding

import (
	"net/http"
)

type requestBinding struct{}

func (requestBinding) Name() string {
	return "request"
}

func (b requestBinding) Bind(obj interface{}, req *http.Request, form map[string][]string) error {
	if err := Uri.BindUri(form, obj); err != nil {
		return err
	}

	values := req.URL.Query()
	if err := mapFormByTag(obj, values, "query"); err != nil {
		return err
	}

	binders := []Binding{Header, Cookie}
	for _, binder := range binders {
		if err := binder.Bind(req, obj); err != nil {
			return err
		}
	}

	// default json
	if req.Method == http.MethodPut || req.Method == http.MethodPost {
		contentType := req.Header.Get("Content-Type")
		if contentType == "" {
			contentType = MIMEJSON
		}
		bb := Default(req.Method, contentType)
		return bb.Bind(req, obj)
	}

	return validate(obj)
}
