package main

import (
	"fmt"

	mvast "github.com/bayugyug/mongersvast"
)

func main() {

	var xml string
	//INLINE LINEAR AD
	inAd := mvast.InLineAd(
		mvast.AdAttributes{"ID": "2007-07-04"},
		&mvast.AdSystem{Value: "Ads for VAST"},
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
								{Event: mvast.TrkEventStart, Value: "http://mongers.vast.utils/start"},
								{Event: mvast.TrkEventFirstQuartile, Value: "http://mongers.vast.utils/firstq"},
								{Event: mvast.TrkEventMidpoint, Value: "http://mongers.vast.utils/midpoint"},
								{Event: mvast.TrkEventThirdQuartile, Value: "http://mongers.vast.utils/thirdq"},
								{Event: mvast.TrkEventComplete, Value: "http://mongers.vast.utils/complete"},
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
							},
						},
					}},
			},
		})
	xml, _ = inAd.ToString()
	fmt.Println(xml)

}
