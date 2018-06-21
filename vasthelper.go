package mongersvast

import (
	"fmt"
	"math/rand"
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
