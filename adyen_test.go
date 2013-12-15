package adyen

import (
	"fmt"
	"net/url"
	"testing"
)

func fixture() (v url.Values) {

	v = url.Values{}

	v.Set("paymentAmount", "10000")
	v.Set("currencyCode", "GBP")
	v.Set("shipBeforeDate", "2007-10-20")
	v.Set("merchantReference", "Internet Order 12345")
	v.Set("skinCode", "4aD37dJA")

	v.Set("merchantAccount", "TestMerchant")
	v.Set("sessionValidity", "2007-10-11T11:00:00Z")

	v.Set("orderData", "H4sIAAAAAAAAALMpsOPlCkssyswvLVZIz89PKVZIzEtRKE4tKstMTi3W4+Wy0S+wAwDOGUCXJgAAAA==")
	v.Set("shopperLocale", "en_GB")

	return
}

func TestSignStr(t *testing.T) {
	should := "10000GBP2007-10-20Internet Order 123454aD37dJATestMerchant2007-10-11T11:00:00Z"
	if str := string(SignStr(fixture())); str != should {
		t.Fatal("\n"+str, "\n"+should)
	}
}

func TestSignature(t *testing.T) {
	key := "Kah942*$7sdp0)"
	should := "x58ZcRVL1H6y+XSeBGrySJ9ACVo="
	if str := Signature(key, SignStr(fixture())); str != should {
		t.Fatal("\n"+str, "\n"+should)
	}
}

func Example() {
	endPoint := "https://test.adyen.com/hpp/select.shtml?"
	key := "Kah942*$7sdp0)"
	v := url.Values{}

	// set parameters
	v.Set("merchantAccount", "TestMerchant")
	v.Set("skinCode", "4aD37dJA")

	v.Set("merchantReference", "Internet Order 12345")
	v.Set("paymentAmount", "10000")
	v.Set("currencyCode", "GBP")

	v.Set("shipBeforeDate", "2007-10-20")
	v.Set("sessionValidity", "2007-10-11T11:00:00Z")

	v.Set("orderData", "H4sIAAAAAAAAALMpsOPlCkssyswvLVZIz89PKVZIzEtRKE4tKstMTi3W4+Wy0S+wAwDOGUCXJgAAAA==")
	v.Set("shopperLocale", "en_GB")

	// generate signing string
	signStr := SignStr(v)

	// generate signature string
	signature := Signature(key, signStr)

	// set signature parameter
	v.Set("merchantSig", signature)

	fmt.Println(endPoint + v.Encode())

	// Output: https://test.adyen.com/hpp/select.shtml?currencyCode=GBP&merchantAccount=TestMerchant&merchantReference=Internet+Order+12345&merchantSig=x58ZcRVL1H6y%2BXSeBGrySJ9ACVo%3D&orderData=H4sIAAAAAAAAALMpsOPlCkssyswvLVZIz89PKVZIzEtRKE4tKstMTi3W4%2BWy0S%2BwAwDOGUCXJgAAAA%3D%3D&paymentAmount=10000&sessionValidity=2007-10-11T11%3A00%3A00Z&shipBeforeDate=2007-10-20&shopperLocale=en_GB&skinCode=4aD37dJA
}
