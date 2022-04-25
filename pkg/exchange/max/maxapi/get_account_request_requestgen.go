// Code generated by "requestgen -method GET -url v2/members/accounts/:currency -type GetAccountRequest -responseType .Account"; DO NOT EDIT.

package max

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

func (g *GetAccountRequest) Currency(currency string) *GetAccountRequest {
	g.currency = currency
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetAccountRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for k, v := range params {
		query.Add(k, fmt.Sprintf("%v", v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetAccountRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetAccountRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for k, v := range params {
		if g.isVarSlice(v) {
			g.iterateSlice(v, func(it interface{}) {
				query.Add(k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(k, fmt.Sprintf("%v", v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetAccountRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetAccountRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check currency field -> json key currency
	currency := g.currency

	// assign parameter of currency
	params["currency"] = currency

	return params, nil
}

func (g *GetAccountRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for k, v := range slugs {
		needleRE := regexp.MustCompile(":" + k + "\\b")
		url = needleRE.ReplaceAllString(url, v)
	}

	return url
}

func (g *GetAccountRequest) iterateSlice(slice interface{}, f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		it := sliceValue.Index(i).Interface()
		f(it)
	}
}

func (g *GetAccountRequest) isVarSlice(v interface{}) bool {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetAccountRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for k, v := range params {
		slugs[k] = fmt.Sprintf("%v", v)
	}

	return slugs, nil
}

func (g *GetAccountRequest) Do(ctx context.Context) (*Account, error) {

	// no body params
	var params interface{}
	query := url.Values{}

	apiURL := "v2/members/accounts/:currency"
	slugs, err := g.GetSlugsMap()
	if err != nil {
		return nil, err
	}

	apiURL = g.applySlugsToUrl(apiURL, slugs)

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse Account
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}