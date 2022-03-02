package forms

import (
	"net/http"
	"net/url"
	"strings"
)

//Form Creates a custom Form struct embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

//Valid Returns true if There Are No Errors otherwise False
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//Initializes a new Form Structre
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {

		}
	}
}

//Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This Field Cannot Be Blank")
		return false
	}
	return true

}
