package main

import (
	"fmt"
)

func main() {

	body := XMLInlineLinear
	/**
		body = XML_INLINE_NON_LINEAR
		body = XML_INLINE_LINEAR
		body = XML_WRAPPER_LINEAR_1
		body = XML_WRAPPER_LINEAR_2
		body = XML_WRAPPER_NON_LINEAR_1
		body = XML_WRAPPER_NON_LINEAR_2
	**/
	body = XMLWrapperNonLinear2
	vt := VAST{}
	vt.FromString(body)
	xml, _ := vt.ToString()
	fmt.Println(xml)
	fmt.Println("==================================================")
	v3 := vt.InLineAd(
		AdAttributes{"ID": "324234"},
		&AdSystem{Value: "Ads for VAST"},
		&AdTitle{Value: "Ad title here"},
		&Description{Value: "Ad remarks here"},
		&VASTError{Value: "http://mongers.vast.utils/error"},
		[]*Impression{
			{ID: "imp-01", Value: "http://mongers.vast.utils/impression1"},
		},
		&Creatives{
			Creative: []*Creative{
				{AdID: "ad01",
					Linear: &Linear{
						Duration: &Duration{Value: "00:00:30"},
						TrackingEvents: &TrackingEvents{
							Tracking: []*Tracking{
								{Event: TrackingEventTypes["Start"], Value: "http://mongers.vast.utils/start"},
								{Event: TrackingEventTypes["FirstQuartile"], Value: "http://mongers.vast.utils/firstq"},
								{Event: TrackingEventTypes["Midpoint"], Value: "http://mongers.vast.utils/midpoint"},
								{Event: TrackingEventTypes["ThirdQuartile"], Value: "http://mongers.vast.utils/thirdq"},
								{Event: TrackingEventTypes["Complete"], Value: "http://mongers.vast.utils/complete"},
							},
						},
						VideoClicks: &VideoClicks{
							ClickThrough: &ClickThrough{
								Value: "http://mongers.vast.utils/clickthrough"},
							ClickTracking: &ClickTracking{
								Value: "http://mongers.vast.utils/clicktracking"},
						},
						MediaFiles: &MediaFiles{
							MediaFile: []*MediaFile{
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
	xml3, _ := v3.ToString()
	fmt.Println(xml3)
}
