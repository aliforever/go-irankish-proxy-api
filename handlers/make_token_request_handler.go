package handlers

import (
	"fmt"
	"net/http"

	"github.com/aliforever/go-irankish"
	"github.com/aliforever/go-irankish-proxy-api/responses"
)

func MakeTokenRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("heey")
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
	amount := r.Form.Get("amount")
	if amount == "" {
		resp.Ok = false
		resp.Message = "empty_amount_error"
		_, _ = fmt.Fprint(w, resp.JSON())
		return
	}
	invoiceId := r.Form.Get("invoice_id")
	if amount == "" {
		resp.Ok = false
		resp.Message = "empty_invoice_id_error"
		_, _ = fmt.Fprint(w, resp.JSON())
		return
	}
	callbackUrl := r.Form.Get("callback_url")
	if amount == "" {
		resp.Ok = false
		resp.Message = "empty_callback_url_error"
		_, _ = fmt.Fprint(w, resp.JSON())
		return
	}
	ik := irankish.IranKish{MerchantId: merchantId}
	payment := irankish.Payment{}
	payment.Amount = amount
	payment.InvoiceId = invoiceId
	payment.CallbackUrl = callbackUrl
	ik.Payment = &payment
	token, err := ik.MakeToken()
	if err != nil {
		resp.Ok = false
		resp.Message = "make_token_err"
		resp.Error = err.Error()
		_, _ = fmt.Fprint(w, resp.JSON())
		return
	}
	resp.Ok = true
	resp.Message = "Success"
	resp.Result = token
	_, _ = fmt.Fprint(w, resp.JSON())
}
