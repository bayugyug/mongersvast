package mongersvast

//SetNonLinear add into the  NonLinearAds obj
func (v *VAST) SetNonLinear(row *NonLinearAds) *VAST {
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
		data = &NonLinearAds{}
	}
	//check which type
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds = data
			}
		}
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds = data
			}
		}
	}
	//good ;-)
	return v
}

//SetNonLinearTracking add into the  NonLinearAds.TrackingEvents obj
func (v *VAST) SetNonLinearTracking(sEvent, sOffset, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Tracking{
		Event:  sEvent,
		Value:  sValue,
		Offset: sOffset,
	}
	//check which type
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.TrackingEvents == nil {
					v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.TrackingEvents = &TrackingEvents{}
				}
				//add to the list
				v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.TrackingEvents.Tracking = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.TrackingEvents.Tracking, data)
			}
		}
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.TrackingEvents == nil {
					v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.TrackingEvents = &TrackingEvents{}
				}
				//add to the list
				v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.TrackingEvents.Tracking = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.TrackingEvents.Tracking, data)
			}
		}
	}
	//good ;-)
	return v
}

//SetNonLinearAd add into the  NonLinearAds.NonLinear obj
func (v *VAST) SetNonLinearAd(sID, sAPIFramework, sWidth, sHeight, sMinSuggestedDuration, sScalable, sMaintainAspectRatio string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &NonLinear{
		ID:                   sID,
		APIFramework:         sAPIFramework,
		Width:                sWidth,
		Height:               sHeight,
		Scalable:             sScalable,
		MaintainAspectRatio:  sMaintainAspectRatio,
		MinSuggestedDuration: sMinSuggestedDuration,
	}
	//check which type
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				//add to the list
				v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear, data)
			}
		}
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				//add to the list
				v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear, data)
			}
		}
	}
	//good ;-)
	return v
}

//SetNonLinearStaticResource  add into the  NonLinearAds.NonLinear.StaticResource obj
func (v *VAST) SetNonLinearStaticResource(sType, sValue string) *VAST {
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
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear)
				if idy > 0 {
					v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear[idy-1].StaticResource = data
				}
			}
		}
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear)
				if idy > 0 {
					v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear[idy-1].StaticResource = data
				}
			}
		}
	}
	//good ;-)
	return v
}

//SetNonLinearClickThrough add into the  NonLinearAds.NonLinear.NonLinearClickThrough obj
func (v *VAST) SetNonLinearClickThrough(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &NonLinearClickThrough{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear)
				if idy > 0 {
					v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear[idy-1].NonLinearClickThrough = data
				}
			}
		}
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear)
				if idy > 0 {
					v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear[idy-1].NonLinearClickThrough = data
				}
			}
		}
	}
	//good ;-)
	return v
}

//SetNonLinearClickTracking add into the  NonLinearAds.NonLinear.NonLinearClickTracking obj
func (v *VAST) SetNonLinearClickTracking(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &NonLinearClickTracking{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear)
				if idy > 0 {
					v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear[idy-1].NonLinearClickTracking = data
				}
			}
		}
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				//add to the list
				idy := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear)
				if idy > 0 {
					v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].NonLinearAds.NonLinear[idy-1].NonLinearClickTracking = data
				}
			}
		}
	}
	//good ;-)
	return v
}
