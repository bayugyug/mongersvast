package mongersvast

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

//InLineAd inline ad template
func InLineAd(attrs AdAttributes, adSystem *AdSystem, title *AdTitle, desc *Description, verr *VASTError, imps []*Impression, creatives *Creatives) (req *VAST) {
	//minimal config
	req = &VAST{
		Version: VastXMLVer2,
		Ad: []*Ad{
			{InLine: &InLine{
				ID: fmtAdUUID("1"),
				InLineWrapperData: InLineWrapperData{
					AdSystem:    adSystem,
					AdTitle:     title,
					Description: desc,
					Error:       verr,
					Impression:  imps,
					Creatives:   creatives,
				},
			},
			},
		},
	}
	//options
	req.FormatAdAttrs(attrs)
	return
}

//WrapperAd wrapper ad template
func WrapperAd(attrs AdAttributes, adSystem *AdSystem, title *AdTitle, desc *Description, verr *VASTError, imps []*Impression, creatives *Creatives, adURI *VASTAdTagURI) (req *VAST) {
	//minimal config
	req = &VAST{
		Version: VastXMLVer2,
		Ad: []*Ad{
			{Wrapper: &Wrapper{
				ID: fmtAdUUID("2"),
				InLineWrapperData: InLineWrapperData{
					AdSystem:     adSystem,
					AdTitle:      title,
					Description:  desc,
					Error:        verr,
					Impression:   imps,
					Creatives:    creatives,
					VASTAdTagURI: adURI,
				},
			},
			},
		},
	}
	//options
	req.FormatAdAttrs(attrs)
	return
}

//fmtAdUUID make temp str
func fmtAdUUID(pfx string) string {
	if len(pfx) <= 0 {
		pfx = "070704"
	}
	return fmt.Sprintf("%s%05x%10x", pfx, rand.Intn(99999), time.Now().UTC().UnixNano())
}

//SetXMLHeaders set the xml headers simply
func SetXMLHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")
	w.Header().Set("Access-Control-Allow-Origin", "*") //Google HTML5 SDK CORS Header
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Max-Age", "10080")
}

//PushXML push content with proper xml hdrs
func PushXML(w http.ResponseWriter, xml string) {
	//just in case ;-)
	if v == nil {
		return
	}
	SetXMLHeaders(w)
	io.WriteString(w, xml)
}
