// Package adyen provides redirect URL generation to set up a payment session for Adyen HPP
//
// For specification see:
//
// https://support.adyen.com/index.php?/Knowledgebase/Article/View/1301/101/integration-manual-v178
package adyen

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
)

// SignStr generates the signing string with the required parameters in the specified order.
// It leaves unspecified parameters blank.
func SignStr(v url.Values) string {
	return v.Get("paymentAmount") + v.Get("currencyCode") + v.Get("shipBeforeDate") + v.Get("merchantReference") + v.Get("skinCode") +
		v.Get("merchantAccount") + v.Get("sessionValidity") + v.Get("shopperEmail") + v.Get("shopperReference") + v.Get("recurringContract") +
		v.Get("allowedMethods") + v.Get("blockedMethods") + v.Get("shopperStatement") + v.Get("merchantReturnData") +
		v.Get("billingAddressType") + v.Get("deliveryAddressType") + v.Get("offset")
}

// Signature generates the HMAC signature from the signing string using SHA-1.
// It returns the base64 encoded signature to be set as "merchantSig".
func Signature(key string, signStr string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(signStr))
	sum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum)
}
