package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/s-larionov/telegram-api/models"
)

const apiURL = "https://api.telegram.org"

var fileType = reflect.TypeOf(models.InputFile(""))

type Response struct {
	Ok          bool                       `json:"ok"`
	Result      json.RawMessage            `json:"result"`
	ErrorCode   int                        `json:"error_code"`
	Description string                     `json:"description"`
	Parameters  *models.ResponseParameters `json:"parameters"`
}

type Requester struct {
	token  string
	client *http.Client
}

func NewRequester(token string) *Requester {
	return NewRequesterWithClient(token, http.DefaultClient)
}

func NewRequesterWithClient(token string, client *http.Client) *Requester {
	return &Requester{
		token:  token,
		client: client,
	}
}

func (r *Requester) JSONRequest(method string, request interface{}) (json.RawMessage, error) {
	url := fmt.Sprintf("%s/bot%s/%s", apiURL, r.token, method)

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	log.WithField("body", string(body)).Trace("request")

	response, err := r.jsonRequest(url, body)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"status":      response.Ok,
		"error_code":  response.ErrorCode,
		"description": response.Description,
		"parameters":  response.Parameters,
		"body":        string(response.Result),
	}).Trace("response")

	return response.Result, nil
}

func (r *Requester) jsonRequest(url string, body []byte) (*Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return r.execute(req)
}

func (r *Requester) MultipartRequest(method string, request interface{}) (json.RawMessage, error) {
	url := fmt.Sprintf("%s/bot%s/%s", apiURL, r.token, method)

	params, files, err := r.prepareMultipartRequestData(request)
	if err != nil {
		return nil, err
	}

	response, err := r.multipartRequest(url, params, files)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"status":      response.Ok,
		"error_code":  response.ErrorCode,
		"description": response.Description,
		"parameters":  response.Parameters,
		"body":        string(response.Result),
	}).Debug("response")

	return response.Result, nil
}

func (r Requester) prepareMultipartRequestData(
	request interface{},
) (params map[string]string, files map[string]string, err error) {
	params = make(map[string]string)
	files = make(map[string]string)

	v := reflect.ValueOf(request)
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	// FIXME: Support t.Kind() == reflect.Map

	if t.Kind() != reflect.Struct {
		return nil, nil, errors.New("incorrect type of request: must be struct")
	}

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		ft := fv.Type()

		name, omitempty := parseJSONTag(t.Field(i))
		if name == "" {
			name = ft.Name()
		}

		isNil := ft.Kind() == reflect.Ptr && fv.IsNil()
		isZero := ft.Kind() != reflect.Ptr && fv.IsZero()
		isEmpty := (ft.Kind() == reflect.Slice || ft.Kind() == reflect.Array) && fv.Len() == 0

		if omitempty && (isNil || isZero || isEmpty) {
			continue
		}

		if ft == fileType {
			file := fv.Interface().(models.InputFile)
			files[name] = string(file)
			continue
		}

		if ft.Kind() == reflect.String {
			params[name] = fv.String()
			continue
		}

		data, err := json.Marshal(fv.Interface())
		if err != nil {
			return nil, nil, err
		}

		params[name] = string(data)
	}

	return params, files, err
}

func (r *Requester) multipartRequest(url string, params, files map[string]string) (*Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for name, file := range files {
		part, err := writer.CreateFormFile(name, file)
		if err != nil {
			return nil, err
		}

		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, f)
		_ = f.Close()
		if err != nil {
			return nil, err
		}
	}

	for name, value := range params {
		err := writer.WriteField(name, value)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return r.execute(req)
}

func (r *Requester) execute(req *http.Request) (*Response, error) {
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if !response.Ok {
		return nil, fmt.Errorf("[%d] %s", response.ErrorCode, response.Description)
	}

	return &response, nil
}

func parseJSONTag(field reflect.StructField) (string, bool) {
	tag := field.Tag.Get("json")

	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], contains(tag[idx+1:], "omitempty")
	}

	return tag, false
}

func contains(options, option string) bool {
	for options != "" {
		var next string
		i := strings.Index(options, ",")

		if i >= 0 {
			options, next = options[:i], options[i+1:]
		}

		if options == option {
			return true
		}

		options = next
	}

	return false
}
