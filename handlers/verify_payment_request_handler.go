package handlers

import (
	"fmt"
	"net/http"

	"github.com/aliforever/go-irankish"
	"github.com/aliforever/go-irankish-proxy-api/responses"
)

func VerifyPaymentRequestHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	resp := responses.Response{}
	if err != nil {
		resp.Ok = false
		resp.Message = "parsing_form_error"
		fmt.Println("error parsing form, " + err.Error())
		_, _ = fmt.Fprint(w, resp.JSON())
		return
	}
	merchantId := r.Form.Get("merchant_id")
	if merchantId == "" {
		resp.Ok = false
		resp.Message = "empty_merchant_id_error"
		_, _ = fmt.Fprint(w, resp.JSON())
		return
	}
	shaKey1 := r.Form.Get("sha1_key")
	if shaKey1 == "" {
		resp.Ok = false
		resp.Message = "empty_sha1_key_error"
		_, _ = fmt.Fprint(w, resp.JSON())
		return
	}
	token := r.Form.Get("token")
	if token == "" {
		resp.Ok = false
		resp.Message = "empty_token_error"
		_, _ = fmt.Fprint(w, resp.JSON())
		return
	}
	referenceNumber := r.Form.Get("reference_number")
	if referenceNumber == "" {
		resp.Ok = false
		resp.Message = "empty_reference_number_error"
		_, _ = fmt.Fprint(w, resp.JSON())
		return
	}
	ik := irankish.IranKish{MerchantId: merchantId, Sha1Key: shaKey1}
	payment := irankish.VerifyPayment{}
	payment.Token = token
	payment.ReferenceNumber = referenceNumber
	ik.Verify = &payment
	verify, err := ik.VerifyPayment()
	if err != nil {
		resp.Ok = false
		resp.Message = "verify_payment_err"
		resp.Error = err.Error()
		_, _ = fmt.Fprint(w, resp.JSON())
		return
	}
	resp.Ok = true
	resp.Message = "Success"
	resp.Result = verify
	_, _ = fmt.Fprint(w, resp.JSON())
}
