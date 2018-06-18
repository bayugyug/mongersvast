## mongersvast

- [x] Simple golang utility library on formatting/manipulating of VAST XML
- [x] Supports VAST 1.0, 2.0, 3.0


### Install

```sh

go get -u -v github.com/bayugyug/mongersvast

```

### Mini-How-To


### > Load XML from string

```go

package main

import (
        "fmt"
        mvast "github.com/bayugyug/mongersvast"
       )


func main() {

    var xml string

    //LOAD from a sample xml string
    vt := mvast.VAST{}
    vt.FromString(mvast.XMLInlineLinear)
    //stringify VAST obj
    xml, _ = vt.ToString()
    fmt.Println(xml)

}

```

### -- Output

[Download XML](https://github.com/bayugyug/mongersvast/blob/master/example/output/sample01-load-xml-from-string.xml)




### > InLine Linear Ad

```go

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
                                                    },
                                            },
                                    }},
                    },
            })
        //stringify VAST obj
        xml, _ = inAd.ToString()
        fmt.Println(xml)

}
```


### -- Output


[Download XML](https://github.com/bayugyug/mongersvast/blob/master/example/output/sample02-inline-linear-ad.xml)


### > Wrapper Linear Ad

```go

package main

import (
        "fmt"
        mvast "github.com/bayugyug/mongersvast"
       )


func main() {

   var xml string
   //WRAPPER LINEAR AD
   wrpAd := mvast.WrapperAd(
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
                                    }},
                    },
            },
            &mvast.VASTAdTagURI{
                    ID:    "adwrapper-01",
                    Value: "http://mongers.vast.utils/2nd-vast-here.xml"},
    )
    //stringify VAST obj
    xml, _ = wrpAd.ToString()
    fmt.Println(xml)

}
```


### -- Output

[Download XML](https://github.com/bayugyug/mongersvast/blob/master/example/output/sample03-wrapper-linear-ad.xml)


### > Load from an XML file

```go
package main

import (
        "fmt"
        mvast "github.com/bayugyug/mongersvast"
       )

func main() {

   var xml string
   //Load the XML file from directory
   vastf := mvast.VAST{}
   vastf.FromFile("./tsample.vast.xml")
   //INJECT additional vast element (Pricing)
   vastf.Ad[0].InLine.Pricing = &mvast.Pricing{
            Model:    "CPM",
            Currency: "USD",
            Value:    "0.99",
    }
   //INJECT additional vast element (Extensions)
   vastf.Ad[0].InLine.Extensions = &mvast.Extensions{
            Extension: []*mvast.Extension{
                    {Type: "iab-Count",
                            TotalAvailable: &mvast.TotalAvailable{Value: "2"}},
            },
    }
   //stringify VAST obj
   xml, _ = vastf.ToString()
   fmt.Println(xml)

}
```

[Input XML](https://github.com/bayugyug/mongersvast/blob/master/example/tsample.vast.xml)


### -- Output

[Download XML](https://github.com/bayugyug/mongersvast/blob/master/example/output/sample04-load-from-an-xml-file.xml)



### > Simple InLine Linear Ad

```go

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
                                                                        ID:                  "media-04",
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
        //stringify VAST obj
        xml, _ = inAd.ToString()
        fmt.Println(xml)

}

```

### -- Output

[Download XML](https://github.com/bayugyug/mongersvast/blob/master/example/output/sample05-simple-inline-linear-ad.xml)




### > Simple Wrapper Tag with Viewable Impression

```go

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

```

### -- Output

[Download XML](https://github.com/bayugyug/mongersvast/blob/master/example/output/sample06-simple-wrapper-tag-with-viewable-impression.xml)


### > Simple Inline Tag With Non-Linear

```go

package main

import (
	"fmt"
	mvast "github.com/bayugyug/mongersvast"
)

func main() {

        var xml string
        //SIMPLE
        inAd := mvast.InLineAd(
                mvast.AdAttributes{
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


```

### -- Output

[Download XML](https://github.com/bayugyug/mongersvast/blob/master/example/output/sample07-simple-inline-tag-with-nonlinear.xml)















### Reference

[VAST INSIGHTS](https://www.iab.com/insights/vast-2-0-xml-samples-for-testing/)



### License

[MIT](https://bayugyug.mit-license.org/)

