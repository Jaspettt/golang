package delivery

import (
	"VinylShop/internal/vinyls/domain"
	ctx "VinylShop/pkg/types/context"
	"VinylShop/pkg/types/responses"
	"VinylShop/pkg/web/encdec"
	errs "VinylShop/pkg/web/errors"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (v *VinylsDeliveryImpl) AddVinylController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vinyl, err := encdec.ReadJSON[domain.Vinyl](r)
	if err != nil {
		encdec.SendErr(w, errs.ErrUnpocessableEntity)

	}
	id, error := v.usecase.InsertVinyl(ctx, &vinyl)
	if error != nil {
		encdec.SendErr(w, error)
		return
	}
	encdec.WriteJSON(w, 201, responses.Message{Message: fmt.Sprintf("inserted vinyl with id=%d", id)})
}
func (v *VinylsDeliveryImpl) GetVinylController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		encdec.SendErr(w, errs.ErrBadRequest)
		return
	}
	ctx := r.Context()
	vinyl, error := v.usecase.GetVinyl(ctx, id)
	if error != nil {
		encdec.SendErr(w, error)
		return
	}
	encdec.WriteJSON(w, 200, vinyl)
}
func (v *VinylsDeliveryImpl) GetVinylsController(w http.ResponseWriter, r *http.Request) {
	ctx := ctx.Context{
		Context: r.Context(),
		Data:    make(map[string]interface{}, 1),
	}
	filters := domain.ExtractFields(r)
	ctx.Data["url"] = fmt.Sprintf("http://%s%s", r.Host, r.URL.Path)
	response, error := v.usecase.GetVinyls(ctx, filters)
	if error != nil {
		encdec.SendErr(w, error)
		return
	}
	encdec.WriteJSON(w, 200, response)
}
func (v *VinylsDeliveryImpl) UpdateVinylController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		encdec.SendErr(w, errs.ErrBadRequest)
		return
	}
	vinyl, err := encdec.ReadJSON[domain.Vinyl](r)
	if err != nil {
		encdec.SendErr(w, errs.ErrUnpocessableEntity)
		return
	}
	newVinyl, error := v.usecase.UpdateVinyl(ctx, id, &vinyl)
	if error != nil {
		encdec.SendErr(w, error)
		return
	}
	encdec.WriteJSON(w, 200, newVinyl)
}
func (v *VinylsDeliveryImpl) DeleteVinylController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		encdec.SendErr(w, errs.ErrBadRequest)
		return
	}
	if error := v.usecase.DeleteVinyl(ctx, id); error != nil {
		encdec.SendErr(w, error)
		return
	}
	encdec.WriteJSON(w, 204, nil)
}
