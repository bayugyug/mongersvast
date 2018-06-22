package mongersvast

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
		v.Ad[0].ID = kk
	}
	//Ad attrs
	if kk, _ := attrs["Sequence"]; kk != "" {
		v.Ad[0].Sequence = kk
	}
	//Ad attrs
	if kk, _ := attrs["ConditionalAd"]; kk != "" {
		v.Ad[0].ConditionalAd = kk
	}
	//Wrapper attrs
	if kk, _ := attrs["FollowAdditionalWrappers"]; v.Ad[0].Wrapper != nil && kk != "" {
		v.Ad[0].Wrapper.FollowAdditionalWrappers = kk
	}
	//Wrapper attrs
	if kk, _ := attrs["AllowMultipleAds"]; v.Ad[0].Wrapper != nil && kk != "" {
		v.Ad[0].Wrapper.AllowMultipleAds = kk
	}
	//Wrapper attrs
	if kk, _ := attrs["FallbackOnNoAd"]; v.Ad[0].Wrapper != nil && kk != "" {
		v.Ad[0].Wrapper.FallbackOnNoAd = kk
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

//SetInLineAd set the minimum InLineAd
func (v *VAST) SetInLineAd(inlineID string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//add the inline
	v.Ad[0].InLine = &InLine{
		ID:                inlineID,
		InLineWrapperData: InLineWrapperData{},
	}
	//good ;-)
	return v
}

//SetWrapperAd set the minimum WrapperAd
func (v *VAST) SetWrapperAd(wrapperID, followAdditionalWrappers, allowMultipleAds, fallbackOnNoAd string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//add the wrapper
	v.Ad[0].Wrapper = &Wrapper{
		ID: wrapperID,
		FollowAdditionalWrappers: followAdditionalWrappers,
		AllowMultipleAds:         allowMultipleAds,
		FallbackOnNoAd:           fallbackOnNoAd,
		InLineWrapperData:        InLineWrapperData{},
	}
	//good ;-)
	return v
}

//SetAdSystem set the AdSystem
func (v *VAST) SetAdSystem(s string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &AdSystem{
		Value: s,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.AdSystem = data
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.AdSystem = data
	}
	//good ;-)
	return v
}

//SetAdTitle set the AdTitle
func (v *VAST) SetAdTitle(s string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &AdTitle{
		Value: s,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.AdTitle = data
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.AdTitle = data
	}
	//good ;-)
	return v
}

//SetDescription set the Description
func (v *VAST) SetDescription(s string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Description{
		Value: s,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.Description = data
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.Description = data
	}
	//good ;-)
	return v
}

//SetErrorURL set the Error URL
func (v *VAST) SetErrorURL(s string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &VASTError{
		Value: s,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.Error = data
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.Error = data
	}
	//good ;-)
	return v
}

//SetImpressionURL set the Impression URL
func (v *VAST) SetImpressionURL(impID, impURL string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Impression{
		ID:    impID,
		Value: impURL,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.Impression = append(v.Ad[0].Wrapper.InLineWrapperData.Impression, data)
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.Impression = append(v.Ad[0].InLine.InLineWrapperData.Impression, data)
	}
	//good ;-)
	return v
}

//SetAdServingID set the AdServingID
func (v *VAST) SetAdServingID(adID, adValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &AdServingID{
		ID:    adID,
		Value: adValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.AdServingID = data
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.AdServingID = data
	}
	//good ;-)
	return v
}

//SetSurvey set the Survey
func (v *VAST) SetSurvey(s string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Survey{
		Value: s,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.Survey = data
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.Survey = data
	}
	//good ;-)
	return v
}

//SetViewableImpression add into the list of ViewableImpression
func (v *VAST) SetViewableImpression(sID string, viewable *Viewable, notviewable *NotViewable, undetermined *ViewUndetermined) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &ViewableImpression{
		ID:               sID,
		Viewable:         viewable,
		NotViewable:      notviewable,
		ViewUndetermined: undetermined,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.ViewableImpression = append(v.Ad[0].Wrapper.InLineWrapperData.ViewableImpression, data)
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.ViewableImpression = append(v.Ad[0].InLine.InLineWrapperData.ViewableImpression, data)
	}
	//good ;-)
	return v
}

//SetVASTAdTagURI set the VASTAdTagURI
func (v *VAST) SetVASTAdTagURI(adID, adValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &VASTAdTagURI{
		ID:    adID,
		Value: adValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.VASTAdTagURI = data
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.VASTAdTagURI = data
	}
	//good ;-)
	return v
}

//SetPricing set the Pricing
func (v *VAST) SetPricing(adID, adModel, adCurr, adValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Pricing{
		ID:       adID,
		Model:    adModel,
		Currency: adCurr,
		Value:    adValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.Pricing = data
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.Pricing = data
	}
	//good ;-)
	return v
}

//SetAdvertiser set the Advertiser
func (v *VAST) SetAdvertiser(s string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Advertiser{
		Value: s,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.Advertiser = data
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.Advertiser = data
	}
	//good ;-)
	return v
}

//SetCategory add into the list of Category
func (v *VAST) SetCategory(sID, sAuth, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Category{
		ID:        sID,
		Authority: sAuth,
		Value:     sValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		v.Ad[0].Wrapper.InLineWrapperData.Category = append(v.Ad[0].Wrapper.InLineWrapperData.Category, data)
	} else if v.Ad[0].InLine != nil {
		v.Ad[0].InLine.InLineWrapperData.Category = append(v.Ad[0].InLine.InLineWrapperData.Category, data)
	}
	//good ;-)
	return v
}

//SetCreativeRow add into the list of Creative
func (v *VAST) SetCreativeRow(sID, sAdID, sSequence, sFramework string, linear *Linear, nonLinear *NonLinearAds, companion *CompanionAds, universal *UniversalAdID) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Creative{
		ID:            sID,
		AdID:          sAdID,
		Sequence:      sSequence,
		APIFramework:  sFramework,
		Linear:        linear,
		NonLinearAds:  nonLinear,
		CompanionAds:  companion,
		UniversalAdID: universal,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives == nil {
			v.Ad[0].Wrapper.InLineWrapperData.Creatives = &Creatives{}
		}
		v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative = append(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative, data)
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives == nil {
			v.Ad[0].InLine.InLineWrapperData.Creatives = &Creatives{}
		}
		v.Ad[0].InLine.InLineWrapperData.Creatives.Creative = append(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative, data)
	}
	//good ;-)
	return v
}

//SetCreative add into the list of Creative
func (v *VAST) SetCreative(sID, sAdID, sSequence, sFramework string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Creative{
		ID:           sID,
		AdID:         sAdID,
		Sequence:     sSequence,
		APIFramework: sFramework,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives == nil {
			v.Ad[0].Wrapper.InLineWrapperData.Creatives = &Creatives{}
		}
		v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative = append(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative, data)
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives == nil {
			v.Ad[0].InLine.InLineWrapperData.Creatives = &Creatives{}
		}
		v.Ad[0].InLine.InLineWrapperData.Creatives.Creative = append(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative, data)
	}
	//good ;-)
	return v
}

//SetUniversalAdID add into the UniversalAdID obj
func (v *VAST) SetUniversalAdID(sID, sIDRegistry, sIDValue, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &UniversalAdID{
		ID:         sID,
		IDRegistry: sIDRegistry,
		IDValue:    sIDValue,
		Value:      sValue,
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				v.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].UniversalAdID = data
			}
		}
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.Creatives != nil {
			idx := len(v.Ad[0].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				v.Ad[0].InLine.InLineWrapperData.Creatives.Creative[idx-1].UniversalAdID = data
			}
		}
	}
	//good ;-)
	return v
}
