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
	"encoding/json"
	"net/http"

	"api/interfaces"
	"api/internal/api"
	"api/internal/domain"
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

// GetRoutes returns all the api routes for the PetAPIController
func (c *PetAPIController) GetRoutes() map[string]util.Route {
	return map[string]util.Route{
		"AddPet": {
			Method:      "POST",
			Pattern:     "/v3/pet/create",
			HandlerFunc: c.AddPet,
			Protected:   true,
		},
		"DeletePet": {
			Method:      "POST",
			Pattern:     "/v3/pet/delete",
			HandlerFunc: c.DeletePet,
			Protected:   true,
		},
		"UpdatePet": {
			Method:      "POST",
			Pattern:     "/v3/pet/update",
			HandlerFunc: c.UpdatePet,
			Protected:   true,
		},
		"GetPet": {
			Method:      "POST",
			Pattern:     "/v3/pet/get",
			HandlerFunc: c.GetPet,
			Protected:   true,
		},
		"ListPets": {
			Method:      "POST",
			Pattern:     "/v3/pet/list",
			HandlerFunc: c.ListPets,
			Protected:   true,
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
	result, err := c.petService.DeletePet(r.Context(), "")
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}

// ListPets - Finds Pets by status
func (c *PetAPIController) ListPets(w http.ResponseWriter, r *http.Request) error {
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
	result, err := c.petService.ListPets(r.Context(), &domain.ListPetsRequest{
		Status: domain.PetStatus(statusParam),
	})
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}

func (c *PetAPIController) GetPet(w http.ResponseWriter, r *http.Request) error {
	body := domain.GetPetRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&body); err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not update a pet")
	}

	result, err := c.petService.GetPet(r.Context(), &body)
	// If an error occurred, encode the error with the status code
	if err != nil {
		return err
	}
	// If no error, encode the body and the result code
	return httpUtil.EncodeJSONResponse(result, nil, w)
}

// UpdatePet - Update an existing pet
func (c *PetAPIController) UpdatePet(w http.ResponseWriter, r *http.Request) error {
	body := struct {
		Pet api.Pet `json:"pet"`
	}{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&body); err != nil {
		return syserr.Wrap(err, syserr.BadInputCode, "could not update a pet")
	}

	domainPet, err := body.Pet.ToDomain()
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
