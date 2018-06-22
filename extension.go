package mongersvast

//SetExtension add into the list of Extension
func (v *VAST) SetExtension(sType, sValue string, total *TotalAvailable, adverifications *AdVerifications) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Extension{
		Type:            sType,
		Value:           sValue,
		TotalAvailable:  total,
		AdVerifications: adverifications,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Extensions == nil {
			v.Ad[0].Wrapper.InLineWrapperData.Extensions = &Extensions{}
		}
		v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension = append(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension, data)
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Extensions == nil {
			v.Ad[0].InLine.InLineWrapperData.Extensions = &Extensions{}
		}
		v.Ad[0].InLine.InLineWrapperData.Extensions.Extension = append(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension, data)
	}
	//good ;-)
	return v
}

//SetExtensionTotalAvailable add into the list of Extension.TotalAvailable
func (v *VAST) SetExtensionTotalAvailable(sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &TotalAvailable{
		Value: sValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Extensions == nil {
			v.Ad[0].Wrapper.InLineWrapperData.Extensions = &Extensions{}
		}
		idx := len(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension)
		if idx > 0 {
			v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].TotalAvailable = data
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Extensions == nil {
			v.Ad[0].InLine.InLineWrapperData.Extensions = &Extensions{}
		}
		idx := len(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension)
		if idx > 0 {
			v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].TotalAvailable = data
		}
	}
	//good ;-)
	return v
}

//SetExtensionAdVerification add into the list of Extension.AdVerifications.Verification
func (v *VAST) SetExtensionAdVerification(js *JavaScriptResource, vp *VerificationParameters, tk *TrackingEvents) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Verification{
		JavaScriptResource:     js,
		VerificationParameters: vp,
		TrackingEvents:         tk,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Extensions == nil {
			v.Ad[0].Wrapper.InLineWrapperData.Extensions = &Extensions{}
		}
		idx := len(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension)
		if idx > 0 {
			if nil == v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications {
				v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications = &AdVerifications{}
			}
			v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification = append(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification, data)
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Extensions == nil {
			v.Ad[0].InLine.InLineWrapperData.Extensions = &Extensions{}
		}
		idx := len(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension)
		if idx > 0 {
			if nil == v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications {
				v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications = &AdVerifications{}
			}
			v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification = append(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification, data)
		}
	}
	//good ;-)
	return v
}

//SetExtensionJavaScriptResource add into the list of Extension.AdVerifications.Verification.JavaScriptResource
func (v *VAST) SetExtensionJavaScriptResource(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &JavaScriptResource{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Extensions == nil {
			v.Ad[0].Wrapper.InLineWrapperData.Extensions = &Extensions{}
		}
		idx := len(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension)
		if idx > 0 {
			if nil == v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications {
				v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications = &AdVerifications{}
			}
			idy := len(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification)
			if idy > 0 {
				v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].JavaScriptResource = data
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Extensions == nil {
			v.Ad[0].InLine.InLineWrapperData.Extensions = &Extensions{}
		}
		idx := len(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension)
		if idx > 0 {
			if nil == v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications {
				v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications = &AdVerifications{}
			}
			idy := len(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification)
			if idy > 0 {
				v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].JavaScriptResource = data
			}
		}
	}
	//good ;-)
	return v
}

//SetExtensionVerificationParameters add into the list of Extension.AdVerifications.Verification.VerificationParameters
func (v *VAST) SetExtensionVerificationParameters(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &VerificationParameters{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Extensions == nil {
			v.Ad[0].Wrapper.InLineWrapperData.Extensions = &Extensions{}
		}
		idx := len(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension)
		if idx > 0 {
			if nil == v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications {
				v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications = &AdVerifications{}
			}
			idy := len(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification)
			if idy > 0 {
				v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].VerificationParameters = data
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Extensions == nil {
			v.Ad[0].InLine.InLineWrapperData.Extensions = &Extensions{}
		}
		idx := len(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension)
		if idx > 0 {
			if nil == v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications {
				v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications = &AdVerifications{}
			}
			idy := len(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification)
			if idy > 0 {
				v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].VerificationParameters = data
			}
		}
	}
	//good ;-)
	return v
}

//SetExtensionTracking add into the list of Extension.AdVerifications.Verification.TrackingEvents.Tracking
func (v *VAST) SetExtensionTracking(sEvent, sOffset, sValue string) *VAST {
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
		if v.Ad[0].Wrapper.InLineWrapperData.Extensions == nil {
			v.Ad[0].Wrapper.InLineWrapperData.Extensions = &Extensions{}
		}
		idx := len(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension)
		if idx > 0 {
			if nil == v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications {
				v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications = &AdVerifications{}
			}
			idy := len(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification)
			if idy > 0 {
				if nil == v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].TrackingEvents {
					v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].TrackingEvents = &TrackingEvents{}
				}
				v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].TrackingEvents.Tracking = append(v.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].TrackingEvents.Tracking, data)
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Extensions == nil {
			v.Ad[0].InLine.InLineWrapperData.Extensions = &Extensions{}
		}
		idx := len(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension)
		if idx > 0 {
			if nil == v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications {
				v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications = &AdVerifications{}
			}
			idy := len(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification)
			if idy > 0 {
				if nil == v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].TrackingEvents {
					v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].TrackingEvents = &TrackingEvents{}
				}
				v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].TrackingEvents.Tracking = append(v.Ad[0].InLine.InLineWrapperData.Extensions.Extension[idx-1].AdVerifications.Verification[idy-1].TrackingEvents.Tracking, data)
			}
		}
	}
	//good ;-)
	return v
}
