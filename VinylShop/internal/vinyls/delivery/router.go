package delivery

import (
	"github.com/go-chi/chi"
)

func (v *VinylsDeliveryImpl) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", v.GetVinylsController)
	r.Get("/{id}", v.GetVinylController)
	r.Post("/", v.AddVinylController)
	r.Put("/{id}", v.UpdateVinylController)
	r.Delete("/{id}", v.DeleteVinylController)
	return r
}
