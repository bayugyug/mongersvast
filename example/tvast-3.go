package main

import (
	"fmt"

	mvast "github.com/bayugyug/mongersvast"
)

func main() {

	var xml string
	//INLINE SIMPLE
	inAd := mvast.InLineAd(
		mvast.AdAttributes{"ID": "2007-07-04", "Version": "4.0", "Sequence": "1", "ConditionalAd": "false"},
		&mvast.AdSystem{Value: "Ads for VAST Inline_Simple"},
		&mvast.AdTitle{Value: "Ad title here"},
		&mvast.Description{Value: "Ad remarks here"},
		&mvast.VASTError{Value: "http://mongers.vast.utils/error"},
		[]*mvast.Impression{
			{ID: "imp-01", Value: "http://mongers.vast.utils/impression1"},
		},
		&mvast.Creatives{
			Creative: []*mvast.Creative{
				{AdID: "ad01",
					Linear: &mvast.Linear{
						Duration: &mvast.Duration{Value: "00:00:30"},
						TrackingEvents: &mvast.TrackingEvents{
							Tracking: []*mvast.Tracking{
								{Event: mvast.TrackingEventTypes["Start"], Value: "http://mongers.vast.utils/start"},
								{Event: mvast.TrackingEventTypes["FirstQuartile"], Value: "http://mongers.vast.utils/firstq"},
								{Event: mvast.TrackingEventTypes["Midpoint"], Value: "http://mongers.vast.utils/midpoint"},
								{Event: mvast.TrackingEventTypes["ThirdQuartile"], Value: "http://mongers.vast.utils/thirdq"},
								{Event: mvast.TrackingEventTypes["Complete"], Value: "http://mongers.vast.utils/complete"},
							},
						},
						VideoClicks: &mvast.VideoClicks{
							ClickThrough: &mvast.ClickThrough{
								Value: "http://mongers.vast.utils/clickthrough"},
							ClickTracking: &mvast.ClickTracking{
								Value: "http://mongers.vast.utils/clicktracking"},
						},
						MediaFiles: &mvast.MediaFiles{
							MediaFile: []*mvast.MediaFile{
								{
									ID:       "media-01",
									Delivery: "progressive",
									Type:     "video/mp4",
									Width:    "640",
									Height:   "360",
									Bitrate:  "784",
									Value:    "https://d1fudb3kxhcy38.cloudfront.net/5xjr/829/4eee446d1897557db60a9d7b3632d294_0001_640x360_700k.mp4",
								},
								{
									ID:                  "media-02",
									Delivery:            "progressive",
									Type:                "video/mp4",
									Width:               "1280",
									Height:              "720",
									Bitrate:             "2000",
									MinBitrate:          "1500",
									MaxBitrate:          "2500",
									Scalable:            "1",
									MaintainAspectRatio: "1",
									Codec:               "0",
									Value:               "https://iabtechlab.com/wp-content/uploads/2016/07/VAST-4.0-Short-Intro.mp4",
								},
								{
									ID:                  "media-03",
									Delivery:            "progressive",
									Type:                "video/mp4",
									Width:               "854",
									Height:              "480",
									Bitrate:             "1000",
									MinBitrate:          "700",
									MaxBitrate:          "1500",
									Scalable:            "1",
									MaintainAspectRatio: "1",
									Codec:               "0",
									Value:               "https://iabtechlab.com/wp-content/uploads/2017/12/VAST-4.0-Short-Intro-mid-resolution.mp4",
								},
								{
									ID:                  "media-03",
									Delivery:            "progressive",
									Type:                "video/mp4",
									Width:               "640",
									Height:              "360",
									Bitrate:             "600",
									MinBitrate:          "500",
									MaxBitrate:          "700",
									Scalable:            "1",
									MaintainAspectRatio: "1",
									Codec:               "0",
									Value:               "https://iabtechlab.com/wp-content/uploads/2017/12/VAST-4.0-Short-Intro-low-resolution.mp4",
								},
							},
						},
					}},
			},
		})
	//INJECT additional vast element (Pricing)
	inAd.Ad[0].InLine.Pricing = &mvast.Pricing{
		Model:    "CPM",
		Currency: "USD",
		Value:    "2.99",
	}
	//INJECT additional vast element (Advertiser)
	inAd.Ad[0].InLine.Advertiser = &mvast.Advertiser{
		Value: "Mongers-Adverts",
	}
	//INJECT additional vast element (Category)
	inAd.Ad[0].InLine.Category = []*mvast.Category{
		{Value: "Mongers-Categ 1"},
		{Value: "Mongers-Categ 2"},
	}
	//INJECT additional vast element (UniversalAdId )
	inAd.Ad[0].InLine.Creatives.Creative[0].UniversalAdID = &mvast.UniversalAdID{
		IDRegistry: "Ad-ID",
		IDValue:    "8465",
		Value:      "8465",
	}
	//convert & show
	xml, _ = inAd.ToString()
	fmt.Println(xml)

}
