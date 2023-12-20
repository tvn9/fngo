package skus

import "testing"

func TestSkuFormatCheck(t *testing.T) {
	skuNum := "tn0001"
	if !SkuFormatCheck(skuNum) {
		t.Error(`SkuFormatCheck("tn0001") = false`)
	}
}
