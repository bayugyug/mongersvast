package mongersvast

//SetCompanionAd add into the CompanionAds obj
func (v *VAST) SetCompanionAd(row *CompanionAds) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := row
	if data == nil {
		data = &CompanionAds{}
	}
	//check which type
	if v.IsAdWrapper() {
		if v.IsAdHasCreatives(AdTypeIsWrapper) {
			idx := v.LenCreative(AdTypeIsWrapper)
			if idx > 0 {
				v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds = data
			}
		}
	} else if v.IsAdInLine() {
		if v.IsAdHasCreatives(AdTypeIsInline) {
			idx := v.LenCreative(AdTypeIsInline)
			if idx > 0 {
				v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds = data
			}
		}
	}
	//good ;-)
	return v
}

//SetCompanion add into the CompanionAds.Companion obj
func (v *VAST) SetCompanion(sID, sWidth, sHeight, sAltText, sAssetWidth, sAssetHeight, sExpandedWidth, sExpandedHeight, sAPIFramework, sAdSlotID, sPxRatio string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Companion{
		ID:             sID,
		Width:          sWidth,
		Height:         sHeight,
		AssetWidth:     sAssetWidth,
		AssetHeight:    sAssetHeight,
		ExpandedWidth:  sExpandedWidth,
		ExpandedHeight: sExpandedHeight,
		APIFramework:   sAPIFramework,
		AdSlotID:       sAdSlotID,
		PxRatio:        sPxRatio,
	}
	data.AltText = &AltText{Value: sAltText}
	//check which type
	if v.IsAdWrapper() {
		if v.IsAdHasCreatives(AdTypeIsWrapper) {
			idx := v.LenCreative(AdTypeIsWrapper)
			if idx > 0 {
				//add to the list
				v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion, data)
			}
		}
	} else if v.IsAdInLine() {
		if v.IsAdHasCreatives(AdTypeIsInline) {
			idx := v.LenCreative(AdTypeIsInline)
			if idx > 0 {
				//add to the list
				v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion, data)
			}
		}
	}
	//good ;-)
	return v
}

//SetCompanionHTMLResource add into the CompanionAds.Companion.HTMLResource obj
func (v *VAST) SetCompanionHTMLResource(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &HTMLResource{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.IsAdWrapper() {
		if v.IsAdHasCreatives(AdTypeIsWrapper) {
			idx := v.LenCreative(AdTypeIsWrapper)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion)
				if idy > 0 {
					v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].HTMLResource = data
				}
			}
		}
	} else if v.IsAdInLine() {
		if v.IsAdHasCreatives(AdTypeIsInline) {
			idx := v.LenCreative(AdTypeIsInline)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion)
				if idy > 0 {
					v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].HTMLResource = data
				}
			}
		}
	}
	//good ;-)
	return v
}

//SetCompanionIFrameResource add into the CompanionAds.Companion.IFrameResource obj
func (v *VAST) SetCompanionIFrameResource(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &IFrameResource{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.IsAdWrapper() {
		if v.IsAdHasCreatives(AdTypeIsWrapper) {
			idx := v.LenCreative(AdTypeIsWrapper)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion)
				if idy > 0 {
					v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].IFrameResource = data
				}
			}
		}
	} else if v.IsAdInLine() {
		if v.IsAdHasCreatives(AdTypeIsInline) {
			idx := v.LenCreative(AdTypeIsInline)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion)
				if idy > 0 {
					v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].IFrameResource = data
				}
			}
		}
	}
	//good ;-)
	return v
}

//SetCompanionStaticResource add into the CompanionAds.Companion.StaticResource obj
func (v *VAST) SetCompanionStaticResource(sType, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &StaticResource{
		CreativeType: sType,
		Value:        sValue,
	}
	//check which type
	if v.IsAdWrapper() {
		if v.IsAdHasCreatives(AdTypeIsWrapper) {
			idx := v.LenCreative(AdTypeIsWrapper)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion)
				if idy > 0 {
					v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].StaticResource = data
				}
			}
		}
	} else if v.IsAdInLine() {
		if v.IsAdHasCreatives(AdTypeIsInline) {
			idx := v.LenCreative(AdTypeIsInline)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion)
				if idy > 0 {
					v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].StaticResource = data
				}
			}
		}
	}
	//good ;-)
	return v
}

//SetCompanionClickThrough add into the CompanionAds.Companion.CompanionClickThrough obj
func (v *VAST) SetCompanionClickThrough(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &CompanionClickThrough{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.IsAdWrapper() {
		if v.IsAdHasCreatives(AdTypeIsWrapper) {
			idx := v.LenCreative(AdTypeIsWrapper)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion)
				if idy > 0 {
					v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].CompanionClickThrough = data
				}
			}
		}
	} else if v.IsAdInLine() {
		if v.IsAdHasCreatives(AdTypeIsInline) {
			idx := v.LenCreative(AdTypeIsInline)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion)
				if idy > 0 {
					v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].CompanionClickThrough = data
				}
			}
		}
	}
	//good ;-)
	return v
}

//SetCompanionTracking add into the CompanionAds.Companion.TrackingEvents obj
func (v *VAST) SetCompanionTracking(sEvent, sOffset, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Tracking{Event: sEvent, Value: sValue, Offset: sOffset}
	//check which type
	if v.IsAdWrapper() {
		if v.IsAdHasCreatives(AdTypeIsWrapper) {
			idx := v.LenCreative(AdTypeIsWrapper)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion)
				if idy > 0 {
					if nil == v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].TrackingEvents {
						v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].TrackingEvents = &TrackingEvents{}
					}
					v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].TrackingEvents.Tracking = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].TrackingEvents.Tracking, data)
				}
			}
		}
	} else if v.IsAdInLine() {
		if v.IsAdHasCreatives(AdTypeIsInline) {
			idx := v.LenCreative(AdTypeIsInline)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion)
				if idy > 0 {
					if nil == v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].TrackingEvents {
						v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].TrackingEvents = &TrackingEvents{}
					}
					v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].TrackingEvents.Tracking = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].CompanionAds.Companion[idy-1].TrackingEvents.Tracking, data)
				}
			}
		}
	}
	//good ;-)
	return v
}
