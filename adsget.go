package mongersvast

//GetAds get the list of ads
func (v *VAST) GetAds() []*Ad {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	//good ;-)
	return v.Ad
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

func (v *VAST) GetAdsInLine() []*InLine {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all []*InLine
	//just in case
	for _, vv := range v.Ad {
		all = append(all, vv.InLine)
	}
	//good ;-)
	return all
}

func (v *VAST) GetAdsWrapper() []*Wrapper {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all []*Wrapper
	//just in case
	for _, vv := range v.Ad {
		all = append(all, vv.Wrapper)
	}
	//good ;-)
	return all
}
