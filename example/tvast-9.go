package main

import (
	"fmt"

	mvast "github.com/bayugyug/mongersvast"
)

func main() {

	var xml string
	//SIMPLE
	inAd := mvast.InLineAd(
		mvast.VastOptions{
			"ID":            "2007-07-04",
			"Version":       "4.0",
			"Sequence":      "1",
			"ConditionalAd": "false",
		},
		&mvast.AdSystem{
			Value: "VAST Inline Simple With Non-Linear",
		},
		&mvast.AdTitle{
			Value: "Ad title here",
		},
		&mvast.Description{
			Value: "VAST 4.0 sample tag for Non Linear ad (i.e Overlay ad). Change the StaticResources to have a tag with your own content. Change NonLinear tag's parameters accordingly to view desired results.",
		},
		&mvast.VASTError{
			Value: "http://mongers.vast.utils/error",
		},
		[]*mvast.Impression{
			{
				ID:    "imp-01",
				Value: "http://mongers.vast.utils/impression1",
			},
		},
		&mvast.Creatives{
			Creative: []*mvast.Creative{
				{
					AdID:     "2447226",
					ID:       "5480",
					Sequence: "1",
					NonLinearAds: &mvast.NonLinearAds{
						NonLinear: []*mvast.NonLinear{
							{
								ID: "1",
								StaticResource: &mvast.StaticResource{
									CreativeType: "image/png",
									Value:        "http://mms.businesswire.com/media/20150623005446/en/473787/21/iab_tech_lab.jpg",
								},
								NonLinearClickThrough: &mvast.NonLinearClickThrough{
									Value: "http://iabtechlab.com",
								},
								NonLinearClickTracking: &mvast.NonLinearClickTracking{
									Value: "http://example.com/trackingurl/clickTracking",
								},
							},
						},
					},
				},
			},
		})
	//INJECT additional vast element (Pricing)
	inAd.Ad[0].InLine.Pricing = &mvast.Pricing{
		Model:    "cpm",
		Currency: "USD",
		Value:    "5.88",
	}
	//INJECT additional vast element (Advertiser)
	inAd.Ad[0].InLine.Advertiser = &mvast.Advertiser{
		Value: "Mongers-Adverts",
	}
	//INJECT additional vast element (AdServingId)
	inAd.Ad[0].InLine.AdServingID = &mvast.AdServingID{
		Value: "ADID_INnonLINEARTEST_ABC123",
	}
	//INJECT additional vast element (Category)
	inAd.Ad[0].InLine.Category = []*mvast.Category{
		{
			Value:     "Mongers-Categ 1",
			Authority: "http://www.iabtechlab.com/categoryauthority",
		},
	}
	//INJECT additional vast element (UniversalAdId)
	inAd.Ad[0].InLine.Creatives.Creative[0].UniversalAdID = &mvast.UniversalAdID{
		IDRegistry: "Ad-ID",
		IDValue:    "8465",
		Value:      "8465",
	}

	//INJECT additional vast element (Extensions)
	inAd.Ad[0].InLine.Extensions = &mvast.Extensions{
		Extension: []*mvast.Extension{
			{Type: "iab-Count",
				TotalAvailable: &mvast.TotalAvailable{Value: "2"}},
		},
	}
	//stringify VAST obj
	xml, _ = inAd.ToString()
	fmt.Println(xml)

}
