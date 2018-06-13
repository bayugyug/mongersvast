## mongersvast

- [x] This is a simple golang utility library on formatting/creating of VAST XML
- [x] Supports VAST 1.0, 2.0, 3.0


## Install

```sh

go get -u -v github.com/bayugyug/mongersvast

```

## How-To


### Load XML from string

```go
    package main

    import (
            "fmt"
            mvast "github.com/bayugyug/mongersvast"
    )

    body := mvast.XMLInlineLinear
    vt := mvast.VAST{}
    vt.FromString(body)
    xml, _ := vt.ToString()
    fmt.Println(xml)


```

### Output

```xml
<?xml version="1.0" encoding="UTF-8"?>
  <VAST version="2.0">
      <Ad id="602867">
          <Wrapper>
              <AdSystem><![CDATA[Acudeo Compatible]]></AdSystem>
              <VASTAdTagURI><![CDATA[http://demo-mongers.vast.utils/proddev/vast/vast_inline_nonlinear3.xml]]></VASTAdTagURI>
              <Impression><![CDATA[http://mongers.vast.utils/wrapper/impression]]></Impression>
              <Creatives>
                  <Creative AdID="602867">
                      <Linear>
                          <TrackingEvents></TrackingEvents>
                      </Linear>
                  </Creative>
                  <Creative AdID="602867-NonLinearTracking">
                      <NonLinearAds>
                          <TrackingEvents></TrackingEvents>
                      </NonLinearAds>
                  </Creative>
                  <Creative AdID="602867-Companion">
                      <CompanionAds>
                          <Companion width="300" height="250">
                              <CompanionClickThrough><![CDATA[http://mongers.vast.utils]]></CompanionClickThrough>
                              <StaticResource creativeType="image/jpeg"><![CDATA[http://demo-mongers.vast.utils/proddev/vast/300x250_banner1.jpg]]></StaticResource>
                              <TrackingEvents>
                                  <Tracking event="creativeView"><![CDATA[http://mongers.vast.utils/wrapper/firstCompanionCreativeView]]></Tracking>
                              </TrackingEvents>
                          </Companion>
                          <Companion width="728" height="90">
                              <CompanionClickThrough><![CDATA[http://mongers.vast.utils]]></CompanionClickThrough>
                              <StaticResource creativeType="image/jpeg"><![CDATA[http://demo-mongers.vast.utils/proddev/vast/728x90_banner1.jpg]]></StaticResource>
                          </Companion>
                      </CompanionAds>
                  </Creative>
              </Creatives>
          </Wrapper>
      </Ad>
  </VAST>
```













### Reference

[iab.com VAST INSIGHTS](https://www.iab.com/insights/vast-2-0-xml-samples-for-testing/)



### License

[MIT](https://bayugyug.mit-license.org/)

