package mongersvast

import "strings"

//NewVAST get instance on VAST object
func (v *VAST) NewVAST(version string) *VAST {
	//minimal config
	v = &VAST{}
	v.SetVersion(version)
	//good ;-)
	return v
}

//SetVersion set the VAST version
func (v *VAST) SetVersion(version string) *VAST {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: version,
		}
	}
	//options
	switch version {
	case VastXMLVer3:
		v.Version = VastXMLVer3
		v.XMLNsXs = VastXMLNsXs
	case VastXMLVer4:
		v.Version = VastXMLVer4
		v.XMLNsXs = VastXMLNsXs
		v.XMLNs = VastXMLNs
	default:
		v.Version = VastXMLVer2
	}
	//good ;-)
	return v
}

//FormatAd set the minimum Ad
func (v *VAST) FormatAd() *VAST {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	//just in case
	if len(v.Ad) <= 0 {
		v.SetAd(VastXMLVer2, "", "", "")
	}
	//good ;-)
	return v
}

//GetAdPos get the last index
func (v *VAST) GetAdPos() int {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	//just in case
	idx := len(v.Ad)
	if idx > 0 {
		return (idx - 1)
	}
	//good ;-)
	return 0
}

//FormatAdAttrs sync all possible options/attrs
func (v *VAST) FormatAdAttrs(attrs AdAttributes) {
	//just in case ;-)
	if v == nil {
		return
	}
	if len(v.Ad) <= 0 {
		return
	}
	//Ad attrs
	if kk, _ := attrs["ID"]; kk != "" {
		v.Ad[v.GetAdPos()].ID = kk
	}
	//Ad attrs
	if kk, _ := attrs["Sequence"]; kk != "" {
		v.Ad[v.GetAdPos()].Sequence = kk
	}
	//Ad attrs
	if kk, _ := attrs["ConditionalAd"]; kk != "" {
		v.Ad[v.GetAdPos()].ConditionalAd = kk
	}
	//Wrapper attrs
	if kk, _ := attrs["FollowAdditionalWrappers"]; v.Ad[v.GetAdPos()].Wrapper != nil && kk != "" {
		v.Ad[v.GetAdPos()].Wrapper.FollowAdditionalWrappers = kk
	}
	//Wrapper attrs
	if kk, _ := attrs["AllowMultipleAds"]; v.Ad[v.GetAdPos()].Wrapper != nil && kk != "" {
		v.Ad[v.GetAdPos()].Wrapper.AllowMultipleAds = kk
	}
	//Wrapper attrs
	if kk, _ := attrs["FallbackOnNoAd"]; v.Ad[v.GetAdPos()].Wrapper != nil && kk != "" {
		v.Ad[v.GetAdPos()].Wrapper.FallbackOnNoAd = kk
	}
	//VAST version
	if kk, _ := attrs["Version"]; kk != "" {
		switch kk {
		case VastXMLVer3:
			v.Version = VastXMLVer3
			v.XMLNsXs = VastXMLNsXs
		case VastXMLVer4:
			v.Version = VastXMLVer4
			v.XMLNsXs = VastXMLNsXs
			v.XMLNs = VastXMLNs
		default:
			v.Version = VastXMLVer2
		}
	}
}

//SetAd set the minimum Ad
func (v *VAST) SetAd(adVersion, adID, adSequence, adConditional string) *VAST {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: adVersion,
		}
	}
	//add 1
	v.Ad = append(v.Ad, &Ad{
		ID:            adID,
		Sequence:      adSequence,
		ConditionalAd: adConditional,
	})
	//good ;-)
	return v
}

//IsAdWrapper check ad type
func (v *VAST) IsAdWrapper() bool {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	//just in case
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		return true
	}
	//good ;-)
	return false
}

//IsAdInLine check ad type
func (v *VAST) IsAdInLine() bool {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	//just in case
	if v.Ad[v.GetAdPos()].InLine != nil {
		return true
	}
	//good ;-)
	return false
}

//IsAdHasCreatives check if creative element is valid
func (v *VAST) IsAdHasCreatives(s string) bool {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	//just in case
	if strings.EqualFold(s, AdTypeIsWrapper) && v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives != nil {
		return true
	} else if strings.EqualFold(s, AdTypeIsInline) && v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives != nil {
		return true
	}
	//good ;-)
	return false
}

//LenCreative get total length of creatives.creative
func (v *VAST) LenCreative(s string) int {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	//just in case
	if v.IsAdHasCreatives(s) {
		return len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative)
	} else if v.IsAdHasCreatives(s) {
		return len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative)
	}
	//good ;-)
	return 0
}
