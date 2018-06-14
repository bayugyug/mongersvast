package mongersvast

import (
	"errors"
)

//misc vars here
var (
	ErrFailedToString         = errors.New("String Format failed")
	ErrFailedToStringNilValue = errors.New("String Format failed (nil value found)")
	ErrFailedFileOpen         = errors.New("XML file open failed")
	ErrFailedFileSave         = errors.New("XML file save failed")
	VastXMLVer1               = "1.0"
	VastXMLVer2               = "2.0"
	VastXMLVer3               = "3.0"
	VastXMLVer4               = "4.0"
	VastXMLHeader             = `<?xml version="1.0" encoding="UTF-8"?>`
)

//const for sample vast xmls
const (
	VastInlineLinear = 1 + iota
	VastInlineNonlinear
	VastWrapperLinear1
	VastWrapperLinear2
	VastWrapperNonlinear1
	VastWrapperNonlinear2
)

//const more
const (
	TrkEventCreativeView     = "creativeView"
	TrkEventExpand           = "expand"
	TrkEventCollapse         = "collapse"
	TrkEventAcceptInvitation = "acceptInvitation"
	TrkEventClose            = "close"
	TrkEventStart            = "start"
	TrkEventFirstQuartile    = "firstQuartile"
	TrkEventMidpoint         = "midpoint"
	TrkEventThirdQuartile    = "thirdQuartile"
	TrkEventComplete         = "complete"
	TrkEventPause            = "pause"
	TrkEventResume           = "resume"
	TrkEventMute             = "mute"
	TrkEventUnMute           = "unmute"
	TrkEventFullscreen       = "fullscreen"
	VastXMLNs                = "http://www.iab.com/VAST"
	VastXMLNsXs              = "http://www.w3.org/2001/XMLSchema"
)

//TrackingEventTypes list of known event types of Tracking
var TrackingEventTypes = map[string]string{
	"CreativeView":     TrkEventCreativeView,
	"Expand":           TrkEventExpand,
	"Collapse":         TrkEventCollapse,
	"AcceptInvitation": TrkEventAcceptInvitation,
	"Close":            TrkEventClose,
	"Start":            TrkEventStart,
	"FirstQuartile":    TrkEventFirstQuartile,
	"Midpoint":         TrkEventMidpoint,
	"ThirdQuartile":    TrkEventThirdQuartile,
	"Complete":         TrkEventComplete,
	"Pause":            TrkEventPause,
	"Resume":           TrkEventResume,
	"Mute":             TrkEventMute,
	"UnMute":           TrkEventUnMute,
	"Fullscreen":       TrkEventFullscreen,
}

//AdAttributes attrs for Ad object
type AdAttributes map[string]string

//VAST the root element of the XML
type VAST struct {
	Version string `xml:"version,attr,omitempty"`
	XMLNs   string `xml:"xmlns,attr,omitempty"`
	XMLNsXs string `xml:"xmlns:xs,attr,omitempty"`
	Ad      []*Ad  `xml:",omitempty"`
}

//VideoAdServingTemplate compat for ver1.0
type VideoAdServingTemplate struct {
	Version string `xml:"version,attr,omitempty"`
	Ad      []*Ad  `xml:",omitempty"`
}

//Ad is an element of the VAST structure
type Ad struct {
	ID            string   `xml:"id,attr,omitempty"`
	Sequence      string   `xml:"sequence,attr,omitempty"`
	ConditionalAd string   `xml:"conditionalAd,attr,omitempty"`
	InLine        *InLine  `xml:",omitempty"`
	Wrapper       *Wrapper `xml:",omitempty"`
}

//AdSystem is an element of the VAST structure
type AdSystem struct {
	Version string `xml:"version,attr,omitempty"`
	Value   string `xml:",cdata"`
}

//InLine is an element of the VAST structure (LINEAR)
type InLine struct {
	InLineWrapperData
	ID string `xml:"id,attr,omitempty"`
}

//Wrapper is an element of the VAST structure (WRAPPER)
type Wrapper struct {
	InLineWrapperData
	ID                       string `xml:"id,attr,omitempty"`
	FollowAdditionalWrappers string `xml:"followAdditionalWrappers,attr,omitempty"`
	AllowMultipleAds         string `xml:"allowMultipleAds,attr,omitempty"`
	FallbackOnNoAd           string `xml:"fallbackOnNoAd,attr,omitempty"`
}

//InLineWrapperData common data for InLine/Wrapper
type InLineWrapperData struct {
	AdSystem           *AdSystem             `xml:",omitempty"`
	AdTitle            *AdTitle              `xml:",omitempty"`
	Description        *Description          `xml:",omitempty"`
	Survey             *Survey               `xml:",omitempty"`
	VASTAdTagURI       *VASTAdTagURI         `xml:",omitempty"`
	Error              *VASTError            `xml:",omitempty"`
	Impression         []*Impression         `xml:",omitempty"`
	Creatives          *Creatives            `xml:",omitempty"`
	Extensions         *Extensions           `xml:",omitempty"`
	Pricing            *Pricing              `xml:",omitempty"`
	AdVerifications    *AdVerifications      `xml:",omitempty"`
	Advertiser         *Advertiser           `xml:",omitempty"`
	Category           []*Category           `xml:",omitempty"`
	ViewableImpression []*ViewableImpression `xml:",omitempty"`
}

//Extensions is an element list
type Extensions struct {
	Extension []*Extension `xml:",omitempty"`
}

//Extension is an element of the VAST structure
type Extension struct {
	Type            string           `xml:"type,attr,omitempty"`
	TotalAvailable  *TotalAvailable  `xml:"total_available,omitempty"`
	Value           string           `xml:",cdata"`
	AdVerifications *AdVerifications `xml:",omitempty"`
}

//StaticResource is an element of the VAST structure
type StaticResource struct {
	CreativeType string `xml:"creativeType,attr,omitempty"`
	Value        string `xml:",cdata"`
}

//CompanionAds is an element list
type CompanionAds struct {
	Companion []*Companion `xml:",omitempty"`
}

//Companion is an element of the VAST structure
type Companion struct {
	ID                    string                 `xml:"id,attr,omitempty"`
	Width                 string                 `xml:"width,attr,omitempty"`
	Height                string                 `xml:"height,attr,omitempty"`
	AltText               string                 `xml:"altText,attr,omitempty"`
	AssetWidth            string                 `xml:"assetWidth,attr,omitempty"`
	AssetHeight           string                 `xml:"assetHeight,attr,omitempty"`
	ExpandedWidth         string                 `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight        string                 `xml:"expandedHeight,attr,omitempty"`
	APIFramework          string                 `xml:"apiFramework,attr,omitempty"`
	AdSlotID              string                 `xml:"adSlotID,attr,omitempty"`
	PxRatio               string                 `xml:"pxratio,attr,omitempty"`
	HTMLResource          *HTMLResource          `xml:",omitempty"`
	IFrameResource        *IFrameResource        `xml:",omitempty"`
	CompanionClickThrough *CompanionClickThrough `xml:",omitempty"`
	StaticResource        *StaticResource        `xml:",omitempty"`
	TrackingEvents        *TrackingEvents        `xml:",omitempty"`
}

//MediaFiles is an element list
type MediaFiles struct {
	MediaFile               []*MediaFile             `xml:",omitempty"`
	Mezzanine               *Mezzanine               `xml:",omitempty"`
	InteractiveCreativeFile *InteractiveCreativeFile `xml:",omitempty"`
}

//MediaFile is an element of the VAST structure
type MediaFile struct {
	ID                  string `xml:"id,attr,omitempty"`
	Delivery            string `xml:"delivery,attr,omitempty"`
	Type                string `xml:"type,attr,omitempty"`
	Width               string `xml:"width,attr,omitempty"`
	Height              string `xml:"height,attr,omitempty"`
	Bitrate             string `xml:"bitrate,attr,omitempty"`
	MinBitrate          string `xml:"minBitrate,attr,omitempty"`
	MaxBitrate          string `xml:"maxBitrate,attr,omitempty"`
	Scalable            string `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio string `xml:"maintainAspectRatio,attr,omitempty"`
	Codec               string `xml:"codec,attr,omitempty"`
	APIFramework        string `xml:"apiFramework,attr,omitempty"`
	Value               string `xml:",cdata"`
}

//VideoClicks is an element of the VAST structure
type VideoClicks struct {
	ClickThrough  *ClickThrough  `xml:",omitempty"`
	ClickTracking *ClickTracking `xml:",omitempty"`
	CustomClick   *CustomClick   `xml:",omitempty"`
}

//Tracking is an element of the VAST structure
type Tracking struct {
	Event  string `xml:"event,attr,omitempty"`
	Offset string `xml:"offset,attr,omitempty"`
	Value  string `xml:",cdata"`
	URL    []*URL `xml:",omitempty"`
}

//TrackingEvents is an element of the VAST structure
type TrackingEvents struct {
	Tracking []*Tracking `xml:",omitempty"`
}

//Linear is an element of the VAST structure
type Linear struct {
	Duration       *Duration       `xml:",omitempty"`
	TrackingEvents *TrackingEvents `xml:",omitempty"`
	VideoClicks    *VideoClicks    `xml:",omitempty"`
	MediaFiles     *MediaFiles     `xml:",omitempty"`
}

//Creative is an element of the VAST structure
type Creative struct {
	ID            string         `xml:"id,attr,omitempty"`
	AdID          string         `xml:"adID,attr,omitempty"`
	Sequence      string         `xml:"sequence,attr,omitempty"`
	APIFramework  string         `xml:"apiFramework,attr,omitempty"`
	Linear        *Linear        `xml:",omitempty"`
	NonLinearAds  *NonLinearAds  `xml:",omitempty"`
	CompanionAds  *CompanionAds  `xml:",omitempty"`
	UniversalAdID *UniversalAdID `xml:",omitempty"`
}

//Creatives is an element of the VAST structure
type Creatives struct {
	Creative []*Creative `xml:"Creative,omitempty"`
}

//NonLinear is an element of the VAST structure
type NonLinear struct {
	ID                     string                  `xml:"id,attr,omitempty"`
	APIFramework           string                  `xml:"apiFramework,attr,omitempty"`
	Height                 string                  `xml:"height,attr,omitempty"`
	Width                  string                  `xml:"width,attr,omitempty"`
	MinSuggestedDuration   string                  `xml:"minSuggestedDuration,attr,omitempty"`
	StaticResource         *StaticResource         `xml:",omitempty"`
	NonLinearClickThrough  *NonLinearClickThrough  `xml:",omitempty"`
	NonLinearClickTracking *NonLinearClickTracking `xml:",omitempty"`
	URL                    []*URL                  `xml:",omitempty"`
}

//NonLinearAds is an element of the VAST structure
type NonLinearAds struct {
	TrackingEvents *TrackingEvents `xml:",omitempty"`
	NonLinear      []*NonLinear    `xml:",omitempty"`
}

//Duration string representation of the creative in seconds
type Duration struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

//Pricing price tag of the vast creative
type Pricing struct {
	ID       string `xml:"id,attr,omitempty"`
	Model    string `xml:"model,attr,omitempty"`
	Currency string `xml:"currency,attr,omitempty"`
	Value    string `xml:",cdata"`
}

//VASTAdTagURI URL of the wrapper
type VASTAdTagURI struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//VASTError  error url
type VASTError struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//Impression vast url
type Impression struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
	URL   []*URL `xml:",omitempty"`
}

//URL vast
type URL struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//Survey vast
type Survey struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//AdTitle vast
type AdTitle struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//Description vast
type Description struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//CompanionClickThrough vast url
type CompanionClickThrough struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
	URL   []*URL `xml:",omitempty"`
}

//ClickThrough vast url
type ClickThrough struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
	URL   []*URL `xml:",omitempty"`
}

//ClickTracking vast url
type ClickTracking struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
	URL   []*URL `xml:",omitempty"`
}

//CustomClick vast url
type CustomClick struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
	URL   []*URL `xml:",omitempty"`
}

//HTMLResource source url
type HTMLResource struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//IFrameResource source url
type IFrameResource struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//TotalAvailable total avail count
type TotalAvailable struct {
	Value string `xml:",cdata"`
}

//NonLinearClickThrough vast url
type NonLinearClickThrough struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
	URL   []*URL `xml:",omitempty"`
}

//NonLinearClickTracking vast url
type NonLinearClickTracking struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
	URL   []*URL `xml:",omitempty"`
}

//UniversalAdID ad id
type UniversalAdID struct {
	ID         string `xml:"id,attr,omitempty"`
	IDRegistry string `xml:"idRegistry,attr,omitempty"`
	IDValue    string `xml:"idValue,attr,omitempty"`
	Value      string `xml:",cdata"`
}

//Category name of the category in creative
type Category struct {
	ID        string `xml:"id,attr,omitempty"`
	Authority string `xml:"authority,attr,omitempty"`
	Value     string `xml:",cdata"`
}

//Advertiser name of the advertiser
type Advertiser struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//JavaScriptResource vast url for javascript res
type JavaScriptResource struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//Verification is a JavaScriptResource
type Verification struct {
	JavaScriptResource     *JavaScriptResource     `xml:",omitempty"`
	VerificationParameters *VerificationParameters `xml:",omitempty"`
	TrackingEvents         *TrackingEvents         `xml:",omitempty"`
}

//AdVerifications list of Verification
type AdVerifications struct {
	Verification []*Verification `xml:",omitempty"`
}

//VerificationParameters vast element
type VerificationParameters struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//Mezzanine vast url
type Mezzanine struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//InteractiveCreativeFile creative is interactive
type InteractiveCreativeFile struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//Viewable creative can be viewed
type Viewable struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//NotViewable creative cant be viewed
type NotViewable struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//ViewUndetermined view not determined
type ViewUndetermined struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

//ViewableImpression viewable imp
type ViewableImpression struct {
	ID               string            `xml:"id,attr,omitempty"`
	Viewable         *Viewable         `xml:",omitempty"`
	NotViewable      *NotViewable      `xml:",omitempty"`
	ViewUndetermined *ViewUndetermined `xml:",omitempty"`
}

//Code vast element
type Code struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",cdata"`
}

// xml members that are for CDATA
