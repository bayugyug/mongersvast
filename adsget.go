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
func (v *VAST) GetAdsAdSystem() map[string][]*AdSystem {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*AdSystem
	all = make(map[string][]*AdSystem)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdSystem != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.AdSystem)
		} else if vv.Wrapper != nil && vv.Wrapper.AdSystem != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.AdSystem)
		}
	}
	//good ;-)
	return all
}

//GetAdsAdTitle get the list of all AdTitle
func (v *VAST) GetAdsAdTitle() map[string][]*AdTitle {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*AdTitle
	all = make(map[string][]*AdTitle)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdTitle != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.AdTitle)
		} else if vv.Wrapper != nil && vv.Wrapper.AdTitle != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.AdTitle)
		}
	}
	//good ;-)
	return all
}

//GetAdsAdServingID get the list of all AdServingID
func (v *VAST) GetAdsAdServingID() map[string][]*AdServingID {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*AdServingID
	all = make(map[string][]*AdServingID)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdServingID != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.AdServingID)
		} else if vv.Wrapper != nil && vv.Wrapper.AdServingID != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.AdServingID)
		}
	}
	//good ;-)
	return all
}

//GetAdsDescription get the list of all Description
func (v *VAST) GetAdsDescription() map[string][]*Description {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Description
	all = make(map[string][]*Description)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Description != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Description)
		} else if vv.Wrapper != nil && vv.Wrapper.Description != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Description)
		}
	}
	//good ;-)
	return all
}
