/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we've switched to the design first approach! You can now help us improve the API whether it's by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
 *
 * API version: 1.0.20-SNAPSHOT
 * Contact: apiteam@swagger.io
 */

package v3

import (
	"net/http"
	"strings"

	"api/interfaces"
	"api/internal/domain"
	"api/internal/util"
	httpUtil "api/internal/util/http"
)

// TagAPIController binds http requests to an api tagService and writes the tagService results to the http response
type TagAPIController struct {
	tagService interfaces.TagService
}

// TagAPIOption for how the controller is set up.
type TagAPIOption func(*TagAPIController)

// NewTagAPIController creates a default api controller
func NewTagAPIController(tagService interfaces.TagService, opts ...TagAPIOption) *TagAPIController {
	controller := &TagAPIController{
		tagService: tagService,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// GetRoutes returns all the api routes for the TagAPIController
func (c *TagAPIController) GetRoutes() map[string]util.Route {
	return map[string]util.Route{
		"ListTags": {
			Method:      strings.ToUpper("Post"),
			Pattern:     "/v3/tag/list",
			HandlerFunc: c.ListTags,
			Protected:   true,
		},
	}
}

// ListTags - Finds Tags
func (c *TagAPIController) ListTags(w http.ResponseWriter, r *http.Request) error {
	result, err := c.tagService.ListTags(r.Context(), &domain.ListTagsRequest{})
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}
