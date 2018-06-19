package main

import (
	"fmt"

	mvast "github.com/bayugyug/mongersvast"
)

func main() {

	var xml string

	vst := &mvast.VAST{}

	vst.
		SetAd("1", "id01", "1", "false").
		SetInLineAd("in01").
		SetAdSystem("VAST Inline Simple With Non-Linear").
		SetAdTitle("Ad title here").
		SetAdServingID("", "ADID_INnonLINEARTEST_ABC123").
		SetDescription("Ad desc here ;-)").
		SetErrorURL("http://mongers.vast.utils/error").
		SetImpressionURL("imp1", "http://mongers.vast.utils/impression1").
		SetCreative("1", "", "", "").
		SetLinear(nil).
		SetLinearDuration("", "00:00:45").
		SetPricing("1", "CPM", "USD", "1.58").
		SetAdvertiser("Mongers-Adverts").
		SetCategory("1", "http://www.iabtechlab.com/categoryauthority", "Mongers-Categ 1").
		SetExtension("iab-Count", "", &mvast.TotalAvailable{Value: "2"}, nil)
	xml, _ = vst.ToString()
	fmt.Println(xml)

}
