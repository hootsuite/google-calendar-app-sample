//go:build go1.22

// Package planned_content provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package planned_content

import (
	"fmt"
	"net/http"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for PlannedContentItemMediaType.
const (
	Folder         PlannedContentItemMediaType = "folder"
	PlannedContent PlannedContentItemMediaType = "plannedContent"
)

// Attachment defines model for Attachment.
type Attachment struct {
	DurationInSeconds *int               `json:"durationInSeconds,omitempty"`
	Id                string             `json:"id"`
	MediaType         *string            `json:"mediaType,omitempty"`
	MimeType          *string            `json:"mimeType,omitempty"`
	Name              *string            `json:"name,omitempty"`
	Original          *ImageMetadata     `json:"original,omitempty"`
	SubtitleFileName  *string            `json:"subtitleFileName,omitempty"`
	Subtitles         *string            `json:"subtitles,omitempty"`
	Thumbnail         *ThumbnailMetadata `json:"thumbnail,omitempty"`
}

// AuthorMetadata defines model for AuthorMetadata.
type AuthorMetadata struct {
	AvatarUrl *string `json:"avatarUrl,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	Id        *string `json:"id,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// Dates defines model for Dates.
type Dates struct {
	Created *string `json:"created,omitempty"`
	Due     *string `json:"due,omitempty"`
	Updated *string `json:"updated,omitempty"`
}

// ImageMetadata defines model for ImageMetadata.
type ImageMetadata struct {
	Height      *int    `json:"height,omitempty"`
	SizeInBytes *int    `json:"sizeInBytes,omitempty"`
	Url         *string `json:"url,omitempty"`
	Width       *int    `json:"width,omitempty"`
}

// Meta defines model for Meta.
type Meta struct {
	Cursor *string `json:"cursor,omitempty"`
}

// PlannedContentItem defines model for PlannedContentItem.
type PlannedContentItem struct {
	Attachments *[]Attachment                `json:"attachments,omitempty"`
	Author      *AuthorMetadata              `json:"author,omitempty"`
	Body        *string                      `json:"body,omitempty"`
	Dates       *Dates                       `json:"dates,omitempty"`
	Id          string                       `json:"id"`
	Links       *[]string                    `json:"links,omitempty"`
	MediaType   *PlannedContentItemMediaType `json:"mediaType,omitempty"`
	Name        *string                      `json:"name,omitempty"`
	Permalink   *string                      `json:"permalink,omitempty"`
	Tags        *[]string                    `json:"tags,omitempty"`
}

// PlannedContentItemMediaType defines model for PlannedContentItem.MediaType.
type PlannedContentItemMediaType string

// PlannedContentResponse defines model for PlannedContentResponse.
type PlannedContentResponse struct {
	Data []PlannedContentItem `json:"data"`
	Meta *Meta                `json:"meta,omitempty"`
}

// ThumbnailMetadata defines model for ThumbnailMetadata.
type ThumbnailMetadata struct {
	Height *int    `json:"height,omitempty"`
	Url    *string `json:"url,omitempty"`
	Width  *int    `json:"width,omitempty"`
}

// GetPlannedContentParams defines parameters for GetPlannedContent.
type GetPlannedContentParams struct {
	StartDate     *openapi_types.Date `form:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate       *openapi_types.Date `form:"endDate,omitempty" json:"endDate,omitempty"`
	ParentId      *string             `form:"parentId,omitempty" json:"parentId,omitempty"`
	Cursor        *string             `form:"cursor,omitempty" json:"cursor,omitempty"`
	Authorization string              `json:"Authorization"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// fetch planned content
	// (GET /planned-content)
	GetPlannedContent(w http.ResponseWriter, r *http.Request, params GetPlannedContentParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetPlannedContent operation middleware
func (siw *ServerInterfaceWrapper) GetPlannedContent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPlannedContentParams

	// ------------- Optional query parameter "startDate" -------------

	err = runtime.BindQueryParameter("form", true, false, "startDate", r.URL.Query(), &params.StartDate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "startDate", Err: err})
		return
	}

	// ------------- Optional query parameter "endDate" -------------

	err = runtime.BindQueryParameter("form", true, false, "endDate", r.URL.Query(), &params.EndDate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "endDate", Err: err})
		return
	}

	// ------------- Optional query parameter "parentId" -------------

	err = runtime.BindQueryParameter("form", true, false, "parentId", r.URL.Query(), &params.ParentId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "parentId", Err: err})
		return
	}

	// ------------- Optional query parameter "cursor" -------------

	err = runtime.BindQueryParameter("form", true, false, "cursor", r.URL.Query(), &params.Cursor)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "cursor", Err: err})
		return
	}

	headers := r.Header

	// ------------- Required header parameter "Authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Authorization")]; found {
		var Authorization string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "Authorization", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "Authorization", valueList[0], &Authorization, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: true})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "Authorization", Err: err})
			return
		}

		params.Authorization = Authorization

	} else {
		err := fmt.Errorf("Header parameter Authorization is required, but not found")
		siw.ErrorHandlerFunc(w, r, &RequiredHeaderError{ParamName: "Authorization", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPlannedContent(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       *http.ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m *http.ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m *http.ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/planned-content", wrapper.GetPlannedContent)

	return m
}
