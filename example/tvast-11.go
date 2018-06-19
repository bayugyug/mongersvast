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
		SetLinearTracking(mvast.TrkEventStart, "", "http://mongers.vast.utils/start").
		SetLinearTracking(mvast.TrkEventFirstQuartile, "", "http://mongers.vast.utils/firstq").
		SetLinearTracking(mvast.TrkEventMidpoint, "", "http://mongers.vast.utils/midpoint").
		SetLinearTracking(mvast.TrkEventThirdQuartile, "", "http://mongers.vast.utils/thirdq").
		SetLinearTracking(mvast.TrkEventComplete, "", "http://mongers.vast.utils/complete").
		SetLinearClickThrough("1", "http://mongers.vast.utils/clickthrough").
		SetLinearClickTracking("1", "http://mongers.vast.utils/clicktracking").
		SetLinearMediaFile("media-01",
			"https://d1fudb3kxhcy38.cloudfront.net/5xjr/829/4eee446d1897557db60a9d7b3632d294_0001_640x360_700k.mp4",
			"progressive",
			"video/mp4",
			"640",
			"360",
			"784",
			"",
			"",
			"",
			"",
			"",
			"").
		SetLinearMediaFile("media-02",
			"https://iabtechlab.com/wp-content/uploads/2017/12/VAST-4.0-Short-Intro-low-resolution.mp4",
			"progressive",
			"video/mp4",
			"640",
			"360",
			"600",
			"500",
			"700",
			"1",
			"1",
			"0",
			"").
		SetPricing("1", "CPM", "USD", "1.58").
		SetAdvertiser("Mongers-Adverts").
		SetCategory("1", "http://www.iabtechlab.com/categoryauthority", "Mongers-Categ 1").
		SetExtension("iab-Count", "", &mvast.TotalAvailable{Value: "2"}, nil)
	xml, _ = vst.ToString()
	fmt.Println(xml)

}
