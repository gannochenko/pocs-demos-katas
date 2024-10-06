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

// CategoryAPIController binds http requests to an api categoryService and writes the categoryService results to the http response
type CategoryAPIController struct {
	categoryService interfaces.CategoryService
}

// CategoryAPIOption for how the controller is set up.
type CategoryAPIOption func(*CategoryAPIController)

// NewCategoryAPIController creates a default api controller
func NewCategoryAPIController(categoryService interfaces.CategoryService, opts ...CategoryAPIOption) *CategoryAPIController {
	controller := &CategoryAPIController{
		categoryService: categoryService,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// GetRoutes returns all the api routes for the CategoryAPIController
func (c *CategoryAPIController) GetRoutes() map[string]util.Route {
	return map[string]util.Route{
		"ListTags": {
			Method:      strings.ToUpper("Post"),
			Pattern:     "/v3/categories/list",
			HandlerFunc: c.ListCategories,
			Protected:   true,
		},
	}
}

// ListCategories - Finds Categories
func (c *CategoryAPIController) ListCategories(w http.ResponseWriter, r *http.Request) error {
	result, err := c.categoryService.ListCategories(r.Context(), &domain.ListCategoriesRequest{})
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}
