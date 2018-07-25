// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "RestaurantType": Application Controllers
//
// Command:
// $ goagen
// --design=Restaurant/design
// --out=$(GOPATH)\src\Restaurant
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// RestauranttypesController is the controller interface for the Restauranttypes actions.
type RestauranttypesController interface {
	goa.Muxer
	Create(*CreateRestauranttypesContext) error
	Delete(*DeleteRestauranttypesContext) error
	Get(*GetRestauranttypesContext) error
	List(*ListRestauranttypesContext) error
	Update(*UpdateRestauranttypesContext) error
}

// MountRestauranttypesController "mounts" a Restauranttypes resource controller on the given service.
func MountRestauranttypesController(service *goa.Service, ctrl RestauranttypesController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/V1/restauranttypes/", ctrl.MuxHandler("preflight", handleRestauranttypesOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/V1/restauranttypes/:id", ctrl.MuxHandler("preflight", handleRestauranttypesOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateRestauranttypesContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateRestauranttypesPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleRestauranttypesOrigin(h)
	service.Mux.Handle("POST", "/V1/restauranttypes/", ctrl.MuxHandler("create", h, unmarshalCreateRestauranttypesPayload))
	service.LogInfo("mount", "ctrl", "Restauranttypes", "action", "Create", "route", "POST /V1/restauranttypes/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteRestauranttypesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleRestauranttypesOrigin(h)
	service.Mux.Handle("DELETE", "/V1/restauranttypes/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Restauranttypes", "action", "Delete", "route", "DELETE /V1/restauranttypes/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetRestauranttypesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleRestauranttypesOrigin(h)
	service.Mux.Handle("GET", "/V1/restauranttypes/:id", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Restauranttypes", "action", "Get", "route", "GET /V1/restauranttypes/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListRestauranttypesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleRestauranttypesOrigin(h)
	service.Mux.Handle("GET", "/V1/restauranttypes/", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Restauranttypes", "action", "List", "route", "GET /V1/restauranttypes/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateRestauranttypesContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateRestauranttypesPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleRestauranttypesOrigin(h)
	service.Mux.Handle("PUT", "/V1/restauranttypes/", ctrl.MuxHandler("update", h, unmarshalUpdateRestauranttypesPayload))
	service.LogInfo("mount", "ctrl", "Restauranttypes", "action", "Update", "route", "PUT /V1/restauranttypes/")
}

// handleRestauranttypesOrigin applies the CORS response headers corresponding to the origin.
func handleRestauranttypesOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "*")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateRestauranttypesPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateRestauranttypesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createRestauranttypesPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateRestauranttypesPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateRestauranttypesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateRestauranttypesPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
