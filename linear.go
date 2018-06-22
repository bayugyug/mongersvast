package mongersvast

//SetLinear add into the Linear obj
func (v *VAST) SetLinear(row *Linear) *VAST {
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
		data = &Linear{}
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear = data
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear = data
			}
		}
	}
	//good ;-)
	return v
}

//SetLinearDuration add into the Linear.Duration obj
func (v *VAST) SetLinearDuration(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Duration{Value: sValue}
	if sID != "" {
		data.ID = sID
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.Duration = data
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.Duration = data
			}
		}
	}
	//good ;-)
	return v
}

//SetLinearTracking add into the Linear.TrackingEvents.Tracking obj
func (v *VAST) SetLinearTracking(sEvent, sOffset, sValue string) *VAST {
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
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.TrackingEvents == nil {
					v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.TrackingEvents = &TrackingEvents{}
				}
				//add to the list
				v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.TrackingEvents.Tracking = append(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.TrackingEvents.Tracking, data)
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.TrackingEvents == nil {
					v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.TrackingEvents = &TrackingEvents{}
				}
				//add to the list
				v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.TrackingEvents.Tracking = append(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.TrackingEvents.Tracking, data)
			}
		}
	}
	//good ;-)
	return v
}

//SetLinearClickThrough add into the Linear.VideoClicks.ClickThrough obj
func (v *VAST) SetLinearClickThrough(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &ClickThrough{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks == nil {
					v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks = &VideoClicks{}
				}
				//add to the list
				v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks.ClickThrough = data
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks == nil {
					v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks = &VideoClicks{}
				}
				//add to the list
				v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks.ClickThrough = data
			}
		}
	}
	//good ;-)
	return v
}

//SetLinearClickTracking add into the Linear.VideoClicks.ClickTracking obj
func (v *VAST) SetLinearClickTracking(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &ClickTracking{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks == nil {
					v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks = &VideoClicks{}
				}
				//add to the list
				v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks.ClickTracking = data
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks == nil {
					v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks = &VideoClicks{}
				}
				//add to the list
				v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks.ClickTracking = data
			}
		}
	}
	//good ;-)
	return v
}

//SetLinearCustomClick add into the Linear.VideoClicks.CustomClick obj
func (v *VAST) SetLinearCustomClick(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &CustomClick{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks == nil {
					v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks = &VideoClicks{}
				}
				//add to the list
				v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks.CustomClick = data
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks == nil {
					v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks = &VideoClicks{}
				}
				//add to the list
				v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.VideoClicks.CustomClick = data
			}
		}
	}
	//good ;-)
	return v
}

//SetLinearMediaFile add into the Linear.MediaFiles.MediaFile obj
func (v *VAST) SetLinearMediaFile(sID, sValue, sDelivery, sType, sWidth, sHeight, sBitrate, sMinBitrate, sMaxBitrate, sScalable, sMaintainAspectRatio, sCodec, sAPIFramework string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &MediaFile{
		Value:               sValue,
		Delivery:            sDelivery,
		Type:                sType,
		Width:               sWidth,
		Height:              sHeight,
		Bitrate:             sBitrate,
		ID:                  sID,
		MinBitrate:          sMinBitrate,
		MaxBitrate:          sMaxBitrate,
		Scalable:            sScalable,
		MaintainAspectRatio: sMaintainAspectRatio,
		Codec:               sCodec,
		APIFramework:        sAPIFramework,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles == nil {
					v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles = &MediaFiles{}
				}
				//add to the list
				v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles.MediaFile = append(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles.MediaFile, data)
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles == nil {
					v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles = &MediaFiles{}
				}
				//add to the list
				v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles.MediaFile = append(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles.MediaFile, data)
			}
		}
	}
	//good ;-)
	return v
}

//SetLinearMezzanine add into the Linear.MediaFiles.Mezzanine obj
func (v *VAST) SetLinearMezzanine(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Mezzanine{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles == nil {
					v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles = &MediaFiles{}
				}
				//add to the list
				v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles.Mezzanine = data
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles == nil {
					v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles = &MediaFiles{}
				}
				//add to the list
				v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles.Mezzanine = data
			}
		}
	}
	//good ;-)
	return v
}

//SetLinearInteractiveCreativeFile add into the Linear.MediaFiles.InteractiveCreativeFile obj
func (v *VAST) SetLinearInteractiveCreativeFile(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &InteractiveCreativeFile{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles == nil {
					v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles = &MediaFiles{}
				}
				//add to the list
				v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles.InteractiveCreativeFile = data
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				if v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles == nil {
					v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles = &MediaFiles{}
				}
				//add to the list
				v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].Linear.MediaFiles.InteractiveCreativeFile = data
			}
		}
	}
	//good ;-)
	return v
}
