/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we've switched to the design first approach! You can now help us improve the API whether it's by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
 *
 * API version: 1.0.20-SNAPSHOT
 * Contact: apiteam@swagger.io
 */

package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"

	"api/interfaces"
	"api/internal/api"
	"api/internal/util"
	httpUtil "api/internal/util/http"
	"api/pkg/syserr"
)

// PetAPIController binds http requests to an api petService and writes the petService results to the http response
type PetAPIController struct {
	petService interfaces.PetService
}

// PetAPIOption for how the controller is set up.
type PetAPIOption func(*PetAPIController)

// NewPetAPIController creates a default api controller
func NewPetAPIController(petService interfaces.PetService, opts ...PetAPIOption) *PetAPIController {
	controller := &PetAPIController{
		petService: petService,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PetAPIController
func (c *PetAPIController) GetRoutes() map[string]util.Route {
	return map[string]util.Route{
		"AddPet": util.Route{
			strings.ToUpper("Post"),
			"/v3/pet",
			c.AddPet,
		},
		"DeletePet": util.Route{
			strings.ToUpper("Delete"),
			"/v3/pet/{petId}",
			c.DeletePet,
		},
		"FindPetsByStatus": util.Route{
			strings.ToUpper("Get"),
			"/v3/pet/findByStatus",
			c.FindPetsByStatus,
		},
		"FindPetsByTags": util.Route{
			strings.ToUpper("Get"),
			"/v3/pet/findByTags",
			c.FindPetsByTags,
		},
		"GetPetById": util.Route{
			strings.ToUpper("Get"),
			"/v3/pet/{petId}",
			c.GetPetById,
		},
		"UpdatePet": util.Route{
			strings.ToUpper("Put"),
			"/v3/pet",
			c.UpdatePet,
		},
		"UpdatePetWithForm": util.Route{
			strings.ToUpper("Post"),
			"/v3/pet/{petId}",
			c.UpdatePetWithForm,
		},
		"UploadFile": util.Route{
			strings.ToUpper("Post"),
			"/v3/pet/{petId}/uploadImage",
			c.UploadFile,
		},
	}
}

// AddPet - Add a new pet to the store
func (c *PetAPIController) AddPet(w http.ResponseWriter, r *http.Request) error {
	petParam := api.Pet{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&petParam); err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not add a pet")
	}
	if err := api.AssertPetRequired(petParam); err != nil {
		return err
	}
	if err := api.AssertPetConstraints(petParam); err != nil {
		return err
	}
	domainPet, err := petParam.ToDomain()
	if err != nil {
		return err
	}
	result, err := c.petService.AddPet(r.Context(), domainPet)
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}

// DeletePet - Deletes a pet
func (c *PetAPIController) DeletePet(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	petIdParam, err := httpUtil.ParseNumericParameter[int64](
		params["petId"],
		httpUtil.WithRequire[int64](httpUtil.ParseInt64),
	)
	if err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not delete a pet", syserr.F("param", "petID"))
	}
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.petService.DeletePet(r.Context(), petIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}

// FindPetsByStatus - Finds Pets by status
func (c *PetAPIController) FindPetsByStatus(w http.ResponseWriter, r *http.Request) error {
	query, err := httpUtil.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not find pets")
	}
	var statusParam string
	if query.Has("status") {
		param := query.Get("status")

		statusParam = param
	} else {
		param := "available"
		statusParam = param
	}
	result, err := c.petService.FindPetsByStatus(r.Context(), statusParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}

// FindPetsByTags - Finds Pets by tags
func (c *PetAPIController) FindPetsByTags(w http.ResponseWriter, r *http.Request) error {
	query, err := httpUtil.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not find pets by tags")
	}
	var tagsParam []string
	if query.Has("tags") {
		tagsParam = strings.Split(query.Get("tags"), ",")
	}
	result, err := c.petService.FindPetsByTags(r.Context(), tagsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}

// GetPetById - Find pet by ID
func (c *PetAPIController) GetPetById(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	petIdParam, err := httpUtil.ParseNumericParameter[int64](
		params["petId"],
		httpUtil.WithRequire[int64](httpUtil.ParseInt64),
	)
	if err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not get a pet by id", syserr.F("param", "petID"))
	}
	result, err := c.petService.GetPetById(r.Context(), petIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}

// UpdatePet - Update an existing pet
func (c *PetAPIController) UpdatePet(w http.ResponseWriter, r *http.Request) error {
	petParam := api.Pet{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&petParam); err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not update a pet")
	}
	if err := api.AssertPetRequired(petParam); err != nil {
		return err
	}
	if err := api.AssertPetConstraints(petParam); err != nil {
		return err
	}
	domainPet, err := petParam.ToDomain()
	if err != nil {
		return err
	}
	result, err := c.petService.UpdatePet(r.Context(), domainPet)
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}

// UpdatePetWithForm - Updates a pet in the store with form data
func (c *PetAPIController) UpdatePetWithForm(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	query, err := httpUtil.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not update a pet with form")
	}
	petIdParam, err := httpUtil.ParseNumericParameter[int64](
		params["petId"],
		httpUtil.WithRequire[int64](httpUtil.ParseInt64),
	)
	if err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not update a pet with form", syserr.F("param", "petId"))
	}
	var nameParam string
	if query.Has("name") {
		param := query.Get("name")

		nameParam = param
	} else {
	}
	var statusParam string
	if query.Has("status") {
		param := query.Get("status")

		statusParam = param
	} else {
	}
	result, err := c.petService.UpdatePetWithForm(r.Context(), petIdParam, nameParam, statusParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}

// UploadFile - uploads an image
func (c *PetAPIController) UploadFile(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	query, err := httpUtil.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not update file")
	}
	petIdParam, err := httpUtil.ParseNumericParameter[int64](
		params["petId"],
		httpUtil.WithRequire[int64](httpUtil.ParseInt64),
	)
	if err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not upload file form", syserr.F("param", "petId"))
	}
	var additionalMetadataParam string
	if query.Has("additionalMetadata") {
		param := query.Get("additionalMetadata")

		additionalMetadataParam = param
	} else {
	}
	bodyParam := &os.File{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		return syserr.Wrap(err, syserr.BadInputCode, "could not update file")
	}
	result, err := c.petService.UploadFile(r.Context(), petIdParam, additionalMetadataParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}
