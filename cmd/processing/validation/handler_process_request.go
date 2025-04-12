package validation

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)

type BovineDuplicateRequest struct {
	UserID               int               `json:"user_id"`
	EventCode            string            `json:"event_code"`
	InstitutionOrName    string            `json:"institution_or_name"`
	FiscalOrPersonalCode string            `json:"fiscal_or_personal_code"`
	FarmName             string            `json:"farm_name"`
	Address              string            `json:"address"`
	Locality             string            `json:"locality"`
	RepresentativeName   string            `json:"representative_name"`
	Duplicates           []BovineDuplicate `json:"duplicates"`
	Phone                string            `json:"phone"`
	Email                string            `json:"email"`
	FullName             string            `json:"full_name"`
}

type BovineDuplicate struct {
	TagNumber string `json:"tag_number"`
}

func HandlerProcessRequestBovine(ctx *fasthttp.RequestCtx) {
	if !ctx.IsPost() {
		ctx.Error("method not allowed", fasthttp.StatusMethodNotAllowed)
		return
	}

	requestBody := ctx.PostBody()

	var request *BovineDuplicateRequest

	err := json.Unmarshal(requestBody, &request)
	if err != nil {
		log.Println(err)
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	err = ProcessRequest(ctx, request)
	if err != nil {
		log.Print(err.Error())
		return
	}

	ctx.SetStatusCode(fasthttp.StatusNoContent)
	return
}
