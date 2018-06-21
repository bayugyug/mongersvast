package main

import (
	"fmt"

	mvast "github.com/bayugyug/mongersvast"
)

func main() {

	var xml string
	//SIMPLE
	wrAd := mvast.WrapperAd(
		mvast.AdAttributes{
			"ID":                       "2007-07-04", //Ad.id
			"Version":                  "4.0",        //Ad.version
			"Sequence":                 "1",          //Ad.sequence
			"ConditionalAd":            "false",      //Ad.conditionalAd
			"FollowAdditionalWrappers": "0",          //Wrapper.followAdditionalWrappers
			"AllowMultipleAds":         "1",          //Wrapper.allowMultipleAds
			"FallbackOnNoAd":           "0",          //Wrapper.fallbackOnNoAd
		},
		&mvast.AdSystem{Version: "4.0", Value: "VAST Wrapper Tag with Viewable Impression"},
		&mvast.AdTitle{Value: "Ad title here"},
		&mvast.Description{Value: "Ad remarks here"},
		&mvast.VASTError{Value: "http://mongers.vast.utils/error"},
		[]*mvast.Impression{
			{ID: "imp-01", Value: "http://mongers.vast.utils/impression1"},
		},
		&mvast.Creatives{
			Creative: []*mvast.Creative{
				{
					AdID:     "2447226",
					ID:       "5480",
					Sequence: "1",
					Linear: &mvast.Linear{
						TrackingEvents: &mvast.TrackingEvents{
							Tracking: []*mvast.Tracking{
								{
									Event:  mvast.TrkEventStart,
									Offset: "09:00:10",
									Value:  "http://mongers.vast.utils/start",
								},
							},
						},
					},
				},
			},
		},
		&mvast.VASTAdTagURI{Value: "https://raw.githubusercontent.com/InteractiveAdvertisingBureau/VAST_Samples/master/VAST%204.0%20Samples/Inline_Companion_Tag-test.xml"},
	)
	//INJECT additional vast element (ViewableImpression)
	wrAd.Ad[0].Wrapper.ViewableImpression = []*mvast.ViewableImpression{
		{
			ID:               "1",
			Viewable:         &mvast.Viewable{Value: "http://search.iabtechlab.com/error?errcode=102&imprid=s5-ea2f7f298e28c0c98374491aec3dfeb1&ts=1243"},
			NotViewable:      &mvast.NotViewable{Value: "http://search.iabtechlab.com/error?errcode=103&imprid=s5-ea2f7f298e28c0c98374491aec3dfeb1&ts=1243"},
			ViewUndetermined: &mvast.ViewUndetermined{Value: "http://search.iabtechlab.com/error?errcode=104&imprid=s5-ea2f7f298e28c0c98374491aec3dfeb1&ts=1243"},
		},
	}
	//INJECT additional vast element (Pricing)
	wrAd.Ad[0].Wrapper.Pricing = &mvast.Pricing{
		Model:    "cpm",
		Currency: "USD",
		Value:    "5.99",
	}
	//INJECT additional vast element (Advertiser)
	wrAd.Ad[0].Wrapper.Advertiser = &mvast.Advertiser{
		Value: "Mongers-Adverts",
	}
	//INJECT additional vast element (Category)
	wrAd.Ad[0].Wrapper.Category = []*mvast.Category{
		{Value: "Mongers-Categ 1"},
	}
	//INJECT additional vast element (UniversalAdId)
	wrAd.Ad[0].Wrapper.Creatives.Creative[0].UniversalAdID = &mvast.UniversalAdID{
		IDRegistry: "Ad-ID",
		IDValue:    "8465",
		Value:      "8465",
	}

	//stringify VAST obj
	xml, _ = wrAd.ToString()
	fmt.Println(xml)

}
