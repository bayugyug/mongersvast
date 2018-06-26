package mongersvast

//SetVerification add into the list of Verification
func (v *VAST) SetVerification(jscript *JavaScriptResource, verificationp *VerificationParameters, trkevents *TrackingEvents) *VAST {
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
		JavaScriptResource:     jscript,
		VerificationParameters: verificationp,
		TrackingEvents:         trkevents,
	}
	//check which type
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications == nil {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications = &AdVerifications{}
		}
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification, data)
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications == nil {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications = &AdVerifications{}
		}
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification, data)
	}
	//good ;-)
	return v
}

//SetVerificationJavaScriptResource add into the list of Verification.JavaScriptResource
func (v *VAST) SetVerificationJavaScriptResource(sID, sValue string) *VAST {
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
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications == nil {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications = &AdVerifications{}
		}

		idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification)
		if idx > 0 {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification[idx-1].JavaScriptResource = data
		}
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications == nil {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications = &AdVerifications{}
		}
		idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification)
		if idx > 0 {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification[idx-1].JavaScriptResource = data
		}
	}
	//good ;-)
	return v
}

//SetVerificationVerificationParameters add into the list of Verification.VerificationParameters
func (v *VAST) SetVerificationVerificationParameters(sID, sValue string) *VAST {
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
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications == nil {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications = &AdVerifications{}
		}

		idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification)
		if idx > 0 {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification[idx-1].VerificationParameters = data
		}
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications == nil {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications = &AdVerifications{}
		}
		idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification)
		if idx > 0 {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification[idx-1].VerificationParameters = data
		}
	}
	//good ;-)
	return v
}

//SetVerificationTracking add into the list of Verification.TrackingEvents.Tracking
func (v *VAST) SetVerificationTracking(sEvent, sOffset, sValue string) *VAST {
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
		if v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications == nil {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications = &AdVerifications{}
		}

		idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification)
		if idx > 0 {
			if nil == v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification[idx-1].TrackingEvents {
				v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification[idx-1].TrackingEvents = &TrackingEvents{}
			}
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification[idx-1].TrackingEvents.Tracking = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdVerifications.Verification[idx-1].TrackingEvents.Tracking, data)
		}
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		if v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications == nil {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications = &AdVerifications{}
		}
		idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification)
		if idx > 0 {
			if nil == v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification[idx-1].TrackingEvents {
				v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification[idx-1].TrackingEvents = &TrackingEvents{}
			}
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification[idx-1].TrackingEvents.Tracking = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdVerifications.Verification[idx-1].TrackingEvents.Tracking, data)
		}
	}
	//good ;-)
	return v
}
