## mongersvast

- [x] This is a simple golang utility library on formatting/creating of VAST XML
- [x] Supports VAST 1.0, 2.0, 3.0


### Install

```sh

go get -u -v github.com/bayugyug/mongersvast

```

### Mini-How-To


### Load XML from string

```go

package main

import (
        "fmt"
        mvast "github.com/bayugyug/mongersvast"
       )


func main() {

    var xml,body string

    //LOAD from a sample xml string
    body = mvast.XMLInlineLinear
    vt := mvast.VAST{}
    vt.FromString(body)
    xml, _ = vt.ToString()
    fmt.Println(xml)

}

```

### Output

```xml

<?xml version="1.0" encoding="UTF-8"?>
  <VAST version="2.0">
      <Ad id="2007-07-04">
          <InLine>
              <AdSystem><![CDATA[Acudeo Compatible]]></AdSystem>
              <AdTitle><![CDATA[VAST 2.0 Instream Test 1]]></AdTitle>
              <Description><![CDATA[VAST 2.0 Instream Test 1]]></Description>
              <Error><![CDATA[http://mongers.vast.utils/error]]></Error>
              <Impression><![CDATA[http://mongers.vast.utils/impression]]></Impression>
              <Creatives>
                  <Creative adID="601364">
                      <Linear>
                          <Duration>00:00:30</Duration>
                          <TrackingEvents>
                              <Tracking event="creativeView"><![CDATA[http://mongers.vast.utils/creativeView]]></Tracking>
                              <Tracking event="start"><![CDATA[http://mongers.vast.utils/start]]></Tracking>
                              <Tracking event="midpoint"><![CDATA[http://mongers.vast.utils/midpoint]]></Tracking>
                              <Tracking event="firstQuartile"><![CDATA[http://mongers.vast.utils/firstQuartile]]></Tracking>
                              <Tracking event="thirdQuartile"><![CDATA[http://mongers.vast.utils/thirdQuartile]]></Tracking>
                              <Tracking event="complete"><![CDATA[http://mongers.vast.utils/complete]]></Tracking>
                          </TrackingEvents>
                          <VideoClicks>
                              <ClickThrough><![CDATA[http://mongers.vast.utils]]></ClickThrough>
                              <ClickTracking><![CDATA[http://mongers.vast.utils/click]]></ClickTracking>
                          </VideoClicks>
                          <MediaFiles>
                              <MediaFile delivery="progressive" type="video/x-flv" width="400" height="300" bitrate="500" scalable="true" maintainAspectRatio="true"><![CDATA[http://cdnp.tremormedia.com/video/acudeo/Carrot_400x300_500kb.flv]]></MediaFile>
                          </MediaFiles>
                      </Linear>
                  </Creative>
                  <Creative adID="601364-Companion">
                      <CompanionAds>
                          <Companion width="300" height="250">
                              <CompanionClickThrough><![CDATA[http://mongers.vast.utils]]></CompanionClickThrough>
                              <StaticResource creativeType="image/jpeg"><![CDATA[http://demo-mongers.vast.utils/proddev/vast/Blistex1.jpg]]></StaticResource>
                              <TrackingEvents>
                                  <Tracking event="creativeView"><![CDATA[http://mongers.vast.utils/firstCompanionCreativeView]]></Tracking>
                              </TrackingEvents>
                          </Companion>
                          <Companion width="728" height="90">
                              <CompanionClickThrough><![CDATA[http://mongers.vast.utils]]></CompanionClickThrough>
                              <StaticResource creativeType="image/jpeg"><![CDATA[http://demo-mongers.vast.utils/proddev/vast/728x90_banner1.jpg]]></StaticResource>
                          </Companion>
                      </CompanionAds>
                  </Creative>
              </Creatives>
          </InLine>
      </Ad>
  </VAST>

```

### Create an InLine Linear Ad

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
        xml, _ = inAd.ToString()
        fmt.Println(xml)

}
```


### Output

```xml

<?xml version="1.0" encoding="UTF-8"?>
  <VAST version="2.0">
      <Ad id="2007-07-04">
          <InLine>
              <AdSystem><![CDATA[Ads for VAST]]></AdSystem>
              <AdTitle><![CDATA[Ad title here]]></AdTitle>
              <Description><![CDATA[Ad remarks here]]></Description>
              <Error><![CDATA[http://mongers.vast.utils/error]]></Error>
              <Impression id="imp-01"><![CDATA[http://mongers.vast.utils/impression1]]></Impression>
              <Creatives>
                  <Creative adID="ad01">
                      <Linear>
                          <Duration>00:00:30</Duration>
                          <TrackingEvents>
                              <Tracking event="start"><![CDATA[http://mongers.vast.utils/start]]></Tracking>
                              <Tracking event="firstQuartile"><![CDATA[http://mongers.vast.utils/firstq]]></Tracking>
                              <Tracking event="midpoint"><![CDATA[http://mongers.vast.utils/midpoint]]></Tracking>
                              <Tracking event="thirdQuartile"><![CDATA[http://mongers.vast.utils/thirdq]]></Tracking>
                              <Tracking event="complete"><![CDATA[http://mongers.vast.utils/complete]]></Tracking>
                          </TrackingEvents>
                          <VideoClicks>
                              <ClickThrough><![CDATA[http://mongers.vast.utils/clickthrough]]></ClickThrough>
                              <ClickTracking><![CDATA[http://mongers.vast.utils/clicktracking]]></ClickTracking>
                          </VideoClicks>
                          <MediaFiles>
                              <MediaFile id="media-01" delivery="progressive" type="video/mp4" width="640" height="360" bitrate="784"><![CDATA[https://d1fudb3kxhcy38.cloudfront.net/5xjr/829/4eee446d1897557db60a9d7b3632d294_0001_640x360_700k.mp4]]></MediaFile>
                          </MediaFiles>
                      </Linear>
                  </Creative>
              </Creatives>
          </InLine>
      </Ad>
  </VAST>
```



### Create a Wrapper Linear Ad

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
    xml, _ = wrpAd.ToString()
    fmt.Println(xml)

}
```


### Output

```xml

<?xml version="1.0" encoding="UTF-8"?>
  <VAST version="2.0">
      <Ad id="2007-07-04">
          <Wrapper>
              <AdSystem><![CDATA[Ads for VAST]]></AdSystem>
              <AdTitle><![CDATA[Ad title here]]></AdTitle>
              <Description><![CDATA[Ad remarks here]]></Description>
              <VASTAdTagURI id="adwrapper-01"><![CDATA[http://mongers.vast.utils/2nd-vast-here.xml]]></VASTAdTagURI>
              <Error><![CDATA[http://mongers.vast.utils/error]]></Error>
              <Impression id="imp-01"><![CDATA[http://mongers.vast.utils/impression1]]></Impression>
              <Creatives>
                  <Creative adID="ad01">
                      <Linear>
                          <Duration>00:00:30</Duration>
                          <TrackingEvents>
                              <Tracking event="start"><![CDATA[http://mongers.vast.utils/start]]></Tracking>
                              <Tracking event="firstQuartile"><![CDATA[http://mongers.vast.utils/firstq]]></Tracking>
                              <Tracking event="midpoint"><![CDATA[http://mongers.vast.utils/midpoint]]></Tracking>
                              <Tracking event="thirdQuartile"><![CDATA[http://mongers.vast.utils/thirdq]]></Tracking>
                              <Tracking event="complete"><![CDATA[http://mongers.vast.utils/complete]]></Tracking>
                          </TrackingEvents>
                          <VideoClicks>
                              <ClickThrough><![CDATA[http://mongers.vast.utils/clickthrough]]></ClickThrough>
                              <ClickTracking><![CDATA[http://mongers.vast.utils/clicktracking]]></ClickTracking>
                          </VideoClicks>
                      </Linear>
                  </Creative>
              </Creatives>
          </Wrapper>
      </Ad>
  </VAST>

  
```

### Load from an XML file

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
   //show
   xml, _ = vastf.ToString()
   fmt.Println(xml)

}
```

[Download Sample XML](https://github.com/bayugyug/mongersvast/blob/master/example/tsample.vast.xml)


### Output

```xml

<?xml version="1.0" encoding="UTF-8"?>
  <VAST version="2.0">
      <Ad id="pre-roll-0">
          <InLine>
              <AdSystem><![CDATA[2.0]]></AdSystem>
              <AdTitle><![CDATA[Sample]]></AdTitle>
              <Impression></Impression>
              <Creatives>
                  <Creative id="2" sequence="1">
                      <Linear>
                          <Duration>00:02:00</Duration>
                          <MediaFiles>
                              <MediaFile delivery="progressive" type="video/mp4" bitrate="400">
                                <![CDATA[http://mongers.vast.utils/demo-sample.mp4]]>
                              </MediaFile>
                          </MediaFiles>
                      </Linear>
                  </Creative>
              </Creatives>
              <Extensions>
                  <Extension type="iab-Count">
                      <total_available><![CDATA[2]]></total_available>
                  </Extension>
              </Extensions>
              <Pricing model="CPM" currency="USD">0.99</Pricing>
          </InLine>
      </Ad>
  </VAST>
  
```



### Simple InLine Linear Ad

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

```

### Output

```xml

<?xml version="1.0" encoding="UTF-8"?>
  <VAST version="2.0">
      <Ad id="2007-07-04">
          <InLine id="1">
              <AdSystem><![CDATA[Ads for VAST Inline_Simple]]></AdSystem>
              <AdTitle><![CDATA[Ad title here]]></AdTitle>
              <Description><![CDATA[Ad remarks here]]></Description>
              <Error><![CDATA[http://mongers.vast.utils/error]]></Error>
              <Impression id="imp-01"><![CDATA[http://mongers.vast.utils/impression1]]></Impression>
              <Creatives>
                  <Creative adID="ad01">
                      <Linear>
                          <Duration>00:00:30</Duration>
                          <TrackingEvents>
                              <Tracking event="start"><![CDATA[http://mongers.vast.utils/start]]></Tracking>
                              <Tracking event="firstQuartile"><![CDATA[http://mongers.vast.utils/firstq]]></Tracking>
                              <Tracking event="midpoint"><![CDATA[http://mongers.vast.utils/midpoint]]></Tracking>
                              <Tracking event="thirdQuartile"><![CDATA[http://mongers.vast.utils/thirdq]]></Tracking>
                              <Tracking event="complete"><![CDATA[http://mongers.vast.utils/complete]]></Tracking>
                          </TrackingEvents>
                          <VideoClicks>
                              <ClickThrough><![CDATA[http://mongers.vast.utils/clickthrough]]></ClickThrough>
                              <ClickTracking><![CDATA[http://mongers.vast.utils/clicktracking]]></ClickTracking>
                          </VideoClicks>
                          <MediaFiles>
                              <MediaFile id="media-01" delivery="progressive" type="video/mp4" width="640" height="360" bitrate="784"><![CDATA[https://d1fudb3kxhcy38.cloudfront.net/5xjr/829/4eee446d1897557db60a9d7b3632d294_0001_640x360_700k.mp4]]></MediaFile>
                              <MediaFile id="media-02" delivery="progressive" type="video/mp4" width="1280" height="720" bitrate="2000" minBitrate="1500" maxBitrate="2500" scalable="1" maintainAspectRatio="1" codec="0"><![CDATA[https://iabtechlab.com/wp-content/uploads/2016/07/VAST-4.0-Short-Intro.mp4]]></MediaFile>
                              <MediaFile id="media-03" delivery="progressive" type="video/mp4" width="854" height="480" bitrate="1000" minBitrate="700" maxBitrate="1500" scalable="1" maintainAspectRatio="1" codec="0"><![CDATA[https://iabtechlab.com/wp-content/uploads/2017/12/VAST-4.0-Short-Intro-mid-resolution.mp4]]></MediaFile>
                              <MediaFile id="media-03" delivery="progressive" type="video/mp4" width="640" height="360" bitrate="600" minBitrate="500" maxBitrate="700" scalable="1" maintainAspectRatio="1" codec="0"><![CDATA[https://iabtechlab.com/wp-content/uploads/2017/12/VAST-4.0-Short-Intro-low-resolution.mp4]]></MediaFile>
                          </MediaFiles>
                      </Linear>
                      <UniversalAdID idRegistry="Ad-ID" idValue="8465"><![CDATA[8465]]></UniversalAdID>
                  </Creative>
              </Creatives>
              <Pricing model="CPM" currency="USD"><![CDATA[2.99]]></Pricing>
              <Advertiser><![CDATA[Mongers-Adverts]]></Advertiser>
              <Category><![CDATA[Mongers-Categ 1]]></Category>
              <Category><![CDATA[Mongers-Categ 2]]></Category>
          </InLine>
      </Ad>
  </VAST>

  
```












### Reference

[VAST INSIGHTS](https://www.iab.com/insights/vast-2-0-xml-samples-for-testing/)



### License

[MIT](https://bayugyug.mit-license.org/)

