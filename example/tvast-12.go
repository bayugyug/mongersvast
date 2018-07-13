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
		SetAdTitle("NonLinear Image").
		SetAdServingID("", "ADID_INnonLINEARTEST_ABC123").
		SetDescription("VAST 3.0 sample tag for Non Linear ad (i.e Overlay ad). Change the StaticResources to have a tag with your own content. Change NonLinear tag's parameters accordingly to view desired results.").
		SetErrorURL("http://mongers.vast.utils/error").
		SetImpressionURL("imp1", "http://mongers.vast.utils/impression1").
		SetCreative("1", "", "", "").
		SetNonLinear(nil).
		SetNonLinearAd("1", "", "480", "150", "00:00:05", "", "").
		SetNonLinearStaticResource("image/png", "http://mms.businesswire.com/media/20150623005446/en/473787/21/iab_tech_lab.jpg").
		SetNonLinearClickTracking("", "http://example.com/trackingurl/clickTracking").
		SetNonLinearClickThrough("", "http://iabtechlab.com").
		SetPricing("1", "CPM", "USD", "1.58").
		SetUniversalAd("1", "", "", "univer id is here").
		SetAdvertiser("Mongers-Adverts").
		SetCategory("1", "http://www.iabtechlab.com/categoryauthority", "Mongers-Categ 1").
		SetExtension("iab-Count", "", &mvast.TotalAvailable{Value: "2"}, nil)
	xml, _ = vst.ToString()
	fmt.Println(xml)

}
