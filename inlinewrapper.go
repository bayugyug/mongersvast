package mongersvast

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
	v.Ad[v.GetAdPos()].InLine = &InLine{
		ID:                inlineID,
		InLineWrapperData: InLineWrapperData{},
	}
	v.Ad[v.GetAdPos()].Wrapper = nil
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdSystem = data
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdSystem = data
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdTitle = data
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdTitle = data
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Description = data
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Description = data
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Error = data
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Error = data
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Impression = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Impression, data)
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Impression = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Impression, data)
	}
	//good ;-)
	return v
}

//SetImpression set the Impression URL
func (v *VAST) SetImpression(impID, impURL string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	//good ;-)
	return v.SetImpressionURL(impID, impURL)
}

//SetAdServing set the AdServingID
func (v *VAST) SetAdServing(adID, adValue string) *VAST {
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.AdServingID = data
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.AdServingID = data
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Survey = data
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Survey = data
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression, data)
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression, data)
	}
	//good ;-)
	return v
}

//SetViewableImpressionViewable add into the list of ViewableImpression.Viewable
func (v *VAST) SetViewableImpressionViewable(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &Viewable{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.IsAdWrapper() {
		idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression)
		if idx <= 0 {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression, &ViewableImpression{})
		}
		idx = len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression)
		if idx > 0 {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression[idx-1].Viewable = data
		}
	} else if v.IsAdInLine() {
		idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression)
		if idx <= 0 {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression, &ViewableImpression{})
		}
		idx = len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression)
		if idx > 0 {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression[idx-1].Viewable = data
		}
	}
	//good ;-)
	return v
}

//SetViewableImpressionNotViewable add into the list of ViewableImpression.NotViewable
func (v *VAST) SetViewableImpressionNotViewable(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &NotViewable{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.IsAdWrapper() {
		idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression)
		if idx <= 0 {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression, &ViewableImpression{})
		}
		idx = len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression)
		if idx > 0 {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression[idx-1].NotViewable = data
		}
	} else if v.IsAdInLine() {
		idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression)
		if idx <= 0 {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression, &ViewableImpression{})
		}
		idx = len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression)
		if idx > 0 {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression[idx-1].NotViewable = data
		}
	}
	//good ;-)
	return v
}

//SetViewableImpressionViewUndetermined add into the list of ViewableImpression.ViewUndetermined
func (v *VAST) SetViewableImpressionViewUndetermined(sID, sValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &ViewUndetermined{
		ID:    sID,
		Value: sValue,
	}
	//check which type
	if v.IsAdWrapper() {
		idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression)
		if idx <= 0 {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression, &ViewableImpression{})
		}
		idx = len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression)
		if idx > 0 {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.ViewableImpression[idx-1].ViewUndetermined = data
		}
	} else if v.IsAdInLine() {
		idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression)
		if idx <= 0 {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression, &ViewableImpression{})
		}
		idx = len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression)
		if idx > 0 {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.ViewableImpression[idx-1].ViewUndetermined = data
		}
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Pricing = data
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Pricing = data
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Advertiser = data
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Advertiser = data
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Category = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Category, data)
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Category = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Category, data)
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
	if v.IsAdWrapper() {
		if !v.IsAdHasCreatives(AdTypeIsWrapper) {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives = &Creatives{}
		}
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative, data)
	} else if v.IsAdInLine() {
		if !v.IsAdHasCreatives(AdTypeIsInline) {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives = &Creatives{}
		}
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative, data)
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
	if v.IsAdWrapper() {
		if !v.IsAdHasCreatives(AdTypeIsWrapper) {
			v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives = &Creatives{}
		}
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative = append(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative, data)
	} else if v.IsAdInLine() {
		if !v.IsAdHasCreatives(AdTypeIsInline) {
			v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives = &Creatives{}
		}
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative = append(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative, data)
	}
	//good ;-)
	return v
}

//SetUniversalAd add into the UniversalAdID obj
func (v *VAST) SetUniversalAd(sID, sIDRegistry, sIDValue, sValue string) *VAST {
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
	if v.IsAdWrapper() {
		if v.IsAdHasCreatives(AdTypeIsWrapper) {
			idx := len(v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.Creatives.Creative[idx-1].UniversalAdID = data
			}
		}
	} else if v.IsAdInLine() {
		if v.IsAdHasCreatives(AdTypeIsInline) {
			idx := len(v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative)
			if idx > 0 {
				v.Ad[v.GetAdPos()].InLine.InLineWrapperData.Creatives.Creative[idx-1].UniversalAdID = data
			}
		}
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
	v.Ad[v.GetAdPos()].Wrapper = &Wrapper{
		ID: wrapperID,
		FollowAdditionalWrappers: followAdditionalWrappers,
		AllowMultipleAds:         allowMultipleAds,
		FallbackOnNoAd:           fallbackOnNoAd,
		InLineWrapperData:        InLineWrapperData{},
	}
	v.Ad[v.GetAdPos()].InLine = nil
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
	if v.IsAdWrapper() {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.VASTAdTagURI = data
	} else if v.IsAdInLine() {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.VASTAdTagURI = data
	}
	//good ;-)
	return v
}
