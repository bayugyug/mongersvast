package main

import (
	"fmt"

	mvast "github.com/bayugyug/mongersvast"
)

func main() {

	var xml string
	//INLINE SIMPLE
	wrAd := mvast.WrapperAd(
		mvast.AdAttributes{"ID": "2007-07-04", "Version": "4.0", "FollowAdditionalWrappers": "0", "AllowMultipleAds": "1", "FallbackOnNoAd": "0"},
		&mvast.AdSystem{Value: "VAST Wrapper Tag"},
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
					CompanionAds: &mvast.CompanionAds{
						Companion: []*mvast.Companion{
							{
								ID:             "01",
								Width:          "100",
								Height:         "150",
								AssetWidth:     "250",
								AssetHeight:    "200",
								AdSlotID:       "3214",
								PxRatio:        "1400",
								ExpandedWidth:  "350",
								ExpandedHeight: "250",
								APIFramework:   "VPAID",
								StaticResource: &mvast.StaticResource{
									CreativeType: "image/png",
									Value:        "https://www.iab.com/wp-content/uploads/2014/09/iab-tech-lab-6-644x290.png",
								},
								CompanionClickThrough: &mvast.CompanionClickThrough{
									Value: "https://iabtechlab.com",
								},
							},
						},
					},
				},
			},
		},
		&mvast.VASTAdTagURI{Value: "https://raw.githubusercontent.com/InteractiveAdvertisingBureau/VAST_Samples/master/VAST%204.0%20Samples/Inline_Companion_Tag-test.xml"},
	)
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
	//INJECT additional vast element (UniversalAdId )
	wrAd.Ad[0].Wrapper.Creatives.Creative[0].UniversalAdID = &mvast.UniversalAdID{
		IDRegistry: "Ad-ID",
		IDValue:    "8465",
		Value:      "8465",
	}
	//convert & show
	xml, _ = wrAd.ToString()
	fmt.Println(xml)

}
