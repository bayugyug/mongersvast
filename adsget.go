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

//GetAdsInLine get the list of all inline
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

//GetAdsWrapper get the list of all wrapper
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

//GetAdsAdSystem get the list of all AdSystem
func (v *VAST) GetAdsAdSystem() []*AdSystem {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all []*AdSystem
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdSystem != nil {
			all = append(all, vv.InLine.AdSystem)
		} else if vv.Wrapper != nil && vv.Wrapper.AdSystem != nil {
			all = append(all, vv.Wrapper.AdSystem)
		}
	}
	//good ;-)
	return all
}
