package telegram

import (
	"encoding/json"
	"fmt"
)

const requestAddress = "https://api.telegram.org"

type RequestHandler func(methodName string, req interface{}) (json.RawMessage, error)

type Response struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
	// Parameters  *ResponseParameters `json:"parameters"`
}

type ErrorResponse Response

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("resp not ok, descr=%v, code=%v", e.Description, e.ErrorCode)
}

// func (b *Bot) executeRequest(methodName string, req interface{}) (json.RawMessage, error) {
// 	url := fmt.Sprintf("%s/bot%s/%s", requestAddress, b.token, methodName)
//
// 	var httpReq *http.Request
//
// 	if upload, ok := isFileUpload(req); ok {
// 		if upload.err != nil {
// 			return nil, upload.err
// 		}
//
// 		ms := multipartstreamer.New()
// 		ms.WriteFields(upload.params)
//
// 		r, err := upload.file.Reader()
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		ms.WriteReader(upload.fieldname, upload.file.Name(), upload.file.Size(), r)
// 		if rc, ok := r.(io.ReadCloser); ok {
// 			defer rc.Close()
// 		}
//
// 		httpReq, err = http.NewRequest("POST", url, nil)
// 		if err != nil {
// 			return nil, err
// 		}
// 		ms.SetupRequest(httpReq)
// 	} else {
// 		body, err := json.Marshal(req)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		httpReq, err = http.NewRequest("POST", url, bytes.NewReader(body))
// 		if err != nil {
// 			return nil, err
// 		}
// 		httpReq.Header.Set("Content-Type", "application/json")
// 	}
//
// 	resp, err := b.client.Do(httpReq)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
//
// 	var respObj Response
// 	err = json.NewDecoder(resp.Body).Decode(&respObj)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if !respObj.Ok {
// 		return nil, ErrorResponse(respObj)
// 	}
//
// 	return respObj.Result, nil
// }
