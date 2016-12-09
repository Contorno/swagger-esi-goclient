/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.1
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"net/url"
	"strings"
	"time"

	"encoding/json"
	"fmt"
)

type SearchApiService service

/**
 * Search on a string
 * Search for entities that match a given sub-string.  ---  Alternate route: &#x60;/v2/characters/{character_id}/search/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param characterId An EVE character ID
 * @param search The string to search on
 * @param categories Type of entities to search for
 * @param language(string) Search locale
 * @param strict(bool) Whether the search should be a strict match
 * @param datasource(string) The server name you would like data from
 * @return *GetCharactersCharacterIdSearchOk
 */
func (a SearchApiService) GetCharactersCharacterIdSearch(ts TokenSource, characterId int32, search string, categories []string, language interface{}, strict interface{}, datasource interface{}) (*GetCharactersCharacterIdSearchOk, time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/{character_id}/search/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(language, "string", "language"); err != nil {
		return nil, time.Now(), err
	}
	if err := a.client.typeCheckParameter(strict, "bool", "strict"); err != nil {
		return nil, time.Now(), err
	}
	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	localVarQueryParams.Add("search", a.client.parameterToString(search, ""))
	localVarQueryParams.Add("categories", a.client.parameterToString(categories, "csv"))
	if language != nil {
		localVarQueryParams.Add("language", a.client.parameterToString(language, ""))
	}
	if strict != nil {
		localVarQueryParams.Add("strict", a.client.parameterToString(strict, ""))
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GetCharactersCharacterIdSearchOk)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return successPayload, time.Now(), err
	}

	if ts != nil {
		if t, err := ts.Token(); err != nil {
			return successPayload, time.Now(), err
		} else if t != nil {
			t.SetAuthHeader(r)
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, time.Now(), err
	}
	defer localVarHttpResponse.Body.Close()

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return successPayload, expires, err
}

/**
 * Search on a string
 * Search for entities that match a given sub-string.  ---  Alternate route: &#x60;/v1/search/&#x60;  Alternate route: &#x60;/legacy/search/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param search The string to search on
 * @param categories Type of entities to search for
 * @param language(string) Search locale
 * @param strict(bool) Whether the search should be a strict match
 * @param datasource(string) The server name you would like data from
 * @return *GetSearchOk
 */
func (a SearchApiService) GetSearch(search string, categories []string, language interface{}, strict interface{}, datasource interface{}) (*GetSearchOk, time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/search/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(language, "string", "language"); err != nil {
		return nil, time.Now(), err
	}
	if err := a.client.typeCheckParameter(strict, "bool", "strict"); err != nil {
		return nil, time.Now(), err
	}
	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	localVarQueryParams.Add("search", a.client.parameterToString(search, ""))
	localVarQueryParams.Add("categories", a.client.parameterToString(categories, "csv"))
	if language != nil {
		localVarQueryParams.Add("language", a.client.parameterToString(language, ""))
	}
	if strict != nil {
		localVarQueryParams.Add("strict", a.client.parameterToString(strict, ""))
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GetSearchOk)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return successPayload, time.Now(), err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, time.Now(), err
	}
	defer localVarHttpResponse.Body.Close()

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return successPayload, expires, err
}
