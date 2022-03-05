package binding

import "net/http"

type cookieBinding struct{}

func (cookieBinding) Name() string {
	return "cookie"
}

func (b cookieBinding) Bind(req *http.Request, obj interface{}) error {
	cookies := req.Cookies()

	form := make(map[string][]string, len(cookies))
	for i := 0; i < len(cookies); i++ {
		form[cookies[i].Name] = []string{cookies[i].Value}
	}

	if err := mapFormByTag(obj, form, "cookie"); err != nil {
		return err
	}

	return validate(obj)
}
