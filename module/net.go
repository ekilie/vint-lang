package module

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ekilie/vint-lang/object"
)

var NetFunctions = map[string]object.ModuleFunction{}

func init() {
	NetFunctions["get"] = getRequest
	NetFunctions["post"] = postRequest
}

func getRequest(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		var url *object.String
		var headers, params *object.Dict
		for k, v := range defs {
			switch k {
			case "url":
				strUrl, ok := v.(*object.String)
				if !ok {
					return &object.Error{Message: "URL must be a string"}
				}
				url = strUrl
			case "headers":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Headers must be a dictionary"}
				}
				headers = dictHead
			case "body":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Body must be a dictionary"}
				}
				params = dictHead
			default:
				return &object.Error{Message: "Arguments are incorrect. Use url and headers."}
			}
		}
		if url.Value == "" {
			return &object.Error{Message: "URL is required"}
		}

		var responseBody *bytes.Buffer
		if params != nil {
			booty := convertObjectToWhatever(params)

			jsonBody, err := json.Marshal(booty)

			if err != nil {
				return &object.Error{Message: "Your query is not formatted properly."}
			}

			responseBody = bytes.NewBuffer(jsonBody)
		}

		var req *http.Request
		var err error
		if responseBody != nil {
			req, err = http.NewRequest("GET", url.Value, responseBody)
		} else {
			req, err = http.NewRequest("GET", url.Value, nil)
		}
		if err != nil {
			return &object.Error{Message: "Failed to make the request"}
		}

		if headers != nil {
			for _, val := range headers.Pairs {
				req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
			}
		}
		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Failed to send the request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Failed to read the response."}
		}

		return &object.String{Value: string(respBody)}
	}

	if len(args) == 1 {
		url, ok := args[0].(*object.String)
		if !ok {
			return &object.Error{Message: "URL must be a string"}
		}
		req, err := http.NewRequest("GET", url.Value, nil)
		if err != nil {
			return &object.Error{Message: "Failed to make the request"}
		}

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Failed to send the request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Failed to read the response."}
		}

		return &object.String{Value: string(respBody)}
	}
	return &object.Error{Message: "Arguments are incorrect. Use url and headers."}
}

func postRequest(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		var url *object.String
		var headers, params *object.Dict
		for k, v := range defs {
			switch k {
			case "url":
				strUrl, ok := v.(*object.String)
				if !ok {
					return &object.Error{Message: "URL must be a string"}
				}
				url = strUrl
			case "headers":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Headers must be a dictionary"}
				}
				headers = dictHead
			case "body":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Body must be a dictionary"}
				}
				params = dictHead
			default:
				return &object.Error{Message: "Arguments are incorrect. Use url and headers."}
			}
		}
		if url.Value == "" {
			return &object.Error{Message: "URL is required"}
		}
		var responseBody *bytes.Buffer
		if params != nil {
			booty := convertObjectToWhatever(params)

			jsonBody, err := json.Marshal(booty)

			if err != nil {
				return &object.Error{Message: "Your query is not formatted properly."}
			}

			responseBody = bytes.NewBuffer(jsonBody)
		}
		var req *http.Request
		var err error
		if responseBody != nil {
			req, err = http.NewRequest("POST", url.Value, responseBody)
		} else {
			req, err = http.NewRequest("POST", url.Value, nil)
		}
		if err != nil {
			return &object.Error{Message: "Failed to make the request"}
		}
		if headers != nil {
			for _, val := range headers.Pairs {
				req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
			}
		}
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Failed to send the request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Failed to read the response."}
		}
		return &object.String{Value: string(respBody)}
	}
	return &object.Error{Message: "Arguments are incorrect. Use url and headers."}
}
