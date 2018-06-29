package mongersvast

//VASTErrorCodes list of standard vast errror codes
var VASTErrorCodes = map[string]string{
	"100": "XML parsing error",
	"101": "VAST schema validation error",
	"102": "VAST version of response not supported",
	"200": "Trafficking error Video player received an Ad type that it was not expecting and/or cannot display",
	"201": "Video player expecting different linearity",
	"202": "Video player expecting different duration",
	"203": "Video player expecting different size",
	"300": "General Wrapper error",
	"301": "Timeout of VAST URI provided in Wrapper element, or of VAST URI provided in a subsequent Wrapper element (URI was either unavailable or reached a timeout as defined by the video player)",
	"302": "Wrapper limit reached, as defined by the video player Too many Wrapper responses have  been received with no InLine response",
	"303": "No Ads VAST response after one or more Wrappers",
	"400": "General Linear error Video player is unable to display the Linear Ad",
	"401": "File not found Unable to find Linear/MediaFile from URI",
	"402": "Timeout of MediaFile URI",
	"403": "Couldn’t find MediaFile that is supported by this video player, based on the attributes of the MediaFile element",
	"405": "Problem displaying MediaFile Video player found a MediaFile with supported type but couldn’t display it MediaFile may include: unsupported codecs, different MIME type than MediaFile@type, unsupported delivery method, etc",
	"500": "General NonLinearAds error",
	"501": "Unable to display NonLinear Ad because creative dimensions do not align with creative display area (ie creative dimension too large)",
	"502": "Unable to fetch NonLinearAds/NonLinear resource",
	"503": "Couldn’t find NonLinear resource with supported type",
	"600": "General CompanionAds error",
	"601": "Unable to display Companion because creative dimensions do not fit within Companiondisplay area (ie, no available space)",
	"602": "Unable to display Required Companion",
	"603": "Unable to fetch CompanionAds/Companion resource",
	"604": "Couldn’t find Companion resource with supported type",
	"900": "Undefined Error",
	"901": "General VPAID error",
}

//constant Generic Macros
const (
	MacroUnknown     = -1
	MacroUnavailable = -2
	AdsPreRoll       = 1
	AdsMidRoll       = 2
	AdsPostRoll      = 3
)

//VASTMacros generic macros for VAST
var VASTMacros = map[string]string{
	"Timestamp":           `[TIMESTAMP]`,           //string  ,  Unencoded: 2016-01-17T8:15:07.127-05
	"CacheBusting":        `[CACHEBUSTING]`,        //integer ,  12345678
	"ContentplayHead":     `[CONTENTPLAYHEAD]`,     //timecode,  Unencoded: 00:05:21.123
	"MediaplayHead":       `[MEDIAPLAYHEAD]`,       //timecode,  Unencoded: 00:05:21.123
	"BreakPosition":       `[BREAKPOSITION]`,       //integer ,  1 for pre-roll, 2 for mid-roll, 3 for post-roll
	"BlockedAdCategories": `[BLOCKEDADCATEGORIES]`, //list<string>, IAB1-6,IAB1-7
	"AdCategories":        `[ADCATEGORIES]`,        //list<string>, IAB1-6,IAB1-7
	"AdCount":             `[ADCOUNT]`,             //integer ,  1
	"TransactionID":       `[TRANSACTIONID]`,       //string  ,  123e4567-e89b-12d3-a456-426655440000
	"PlacementType":       `[PLACEMENTTYPE]`,       //integer ,  1
	"AdType":              `[ADTYPE]`,              //string  ,  video,display,banner
	"UniversalAdID":       `[UNIVERSALADID]`,       //string  ,  Unencoded: ad-id.org CNPA0484000H
	"IFA":                 `[IFA]`,                 //string  ,  123e4567-e89b-12d3-a456-426655440000
	"IFAtype":             `[IFATYPE]`,             //string  ,  rida
	"ClientUA":            `[CLIENTUA]`,            //string  ,  Unencoded: MyPlayer/7.1 MyPlayerVastPlugin/1.1.2
	"ServerUA":            `[SERVERUA]`,            //string  ,  Unencoded: AdsBot-Google (+http:, //www.google.com/adsbot.html)
	"DeviceUA":            `[DEVICEUA]`,            //string  ,  Unencoded: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36
	"DeviceIP":            `[DEVICEIP]`,            //string  ,  Unencoded: IPv4: 8.8.8.8
	"LatLong":             `[LATLONG]`,             //number  ,  51.004703,3.754806
	"Domain":              `[DOMAIN]`,              //string  ,  Unencoded: mydomain.com
	"PageURL":             `[PAGEURL]`,             //string  ,  Unencoded: https:, //www.mydomain.com/article/page
	"AppName":             `[APPNAME]`,             //string  ,  Unencoded: com.example.myapp
	"VastVersions":        `[VASTVERSIONS]`,        //list<integer>, 2,3,5,6,7,8,310  #310 for VAST 4.1 #311 for VAST 4.1 Wrapper
	"ApiFrameworks":       `[APIFRAMEWORKS]`,       //list<integer>, 2,3
	"Extensions":          `[EXTENSIONS]`,          //list<string> , AdVerifications,extensionA,extensionB
	"VerificationVendors": `[VERIFICATIONVENDORS]`, //list<string> , moat.com-omid,ias.com-omid,doubleverify.com-omid
	"MediaMime":           `[MEDIAMIME]`,           //list<string> , Unencoded: video/mp4,application/x-mpegURL
	"PlayerCapabilities":  `[PLAYERCAPABILITIES]`,  //list<string> , mautoplay,fullscreen,icon
	"ClickType":           `[CLICKTYPE]`,           //integer      , 1 #0: not clickable, 1: clickable on full area of video, 2: clickable only on associated button/link, 3: clickable with confirmation dialog
	"PlayerState":         `[PLAYERSTATE]`,         //list<string> , autoplayed,fullscreen
	"PlayerSize":          `[PLAYERSIZE]`,          //integer      , 640,360
	"AdPlayHead":          `[ADPLAYHEAD]`,          //timecode     , Unencoded: 00:05:21.123
	"AssetURI":            `[ASSETURI]`,            //string       , Unencoded:  https:, //myadserver.com/video.mp4
	"PodSequence":         `[PODSEQUENCE]`,         //integer      , 1
	"AdServingID":         `[ADSERVINGID]`,         //string       , ServerName-47ed3bac-1768-4b9a-9d0e-0b92422ab066
	"ClickPos":            `[CLICKPOS]`,            //integer      , 315,204
	"ErrorCode":           `[ERRORCODE]`,           //integer      , 900
	"Reason":              `[REASON]`,              //integer      , 1
	"LimitAdTracking":     `[LIMITADTRACKING]`,     //integer      , 1
	"Regulations":         `[REGULATIONS]`,         //list<string> , gdpr,coppa
	"GdprConsent":         `[GDPRCONSENT]`,         //string       , BOLqFHuOLqFHuAABAENAAAAAAAAoAAA
}
