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

//GetAdsAdServing get the list of all AdServingID
func (v *VAST) GetAdsAdServing() map[string][]*AdServingID {
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

//GetAdsErrorURL get the list of all VASTError
func (v *VAST) GetAdsErrorURL() map[string][]*VASTError {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*VASTError
	all = make(map[string][]*VASTError)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Error != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Error)
		} else if vv.Wrapper != nil && vv.Wrapper.Error != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Error)
		}
	}
	//good ;-)
	return all
}

//GetAdsError get the list of all VASTError
func (v *VAST) GetAdsError() map[string][]*VASTError {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	//good ;-)
	return v.GetAdsErrorURL()
}

//GetAdsSurvey get the list of all Survey
func (v *VAST) GetAdsSurvey() map[string][]*Survey {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Survey
	all = make(map[string][]*Survey)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Survey != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Survey)
		} else if vv.Wrapper != nil && vv.Wrapper.Survey != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Survey)
		}
	}
	//good ;-)
	return all
}

//GetAdsExpires get the list of all Expires
func (v *VAST) GetAdsExpires() map[string][]*Expires {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Expires
	all = make(map[string][]*Expires)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Expires != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Expires)
		} else if vv.Wrapper != nil && vv.Wrapper.Expires != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Expires)
		}
	}
	//good ;-)
	return all
}

//GetAdsCategory get the list of all Category
func (v *VAST) GetAdsCategory() map[string][]*Category {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Category
	all = make(map[string][]*Category)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && len(vv.InLine.Category) > 0 {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Category...)
		} else if vv.Wrapper != nil && len(vv.Wrapper.Category) > 0 {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Category...)
		}
	}
	//good ;-)
	return all
}

//GetAdsAdvertiser get the list of all Advertiser
func (v *VAST) GetAdsAdvertiser() map[string][]*Advertiser {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Advertiser
	all = make(map[string][]*Advertiser)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Advertiser != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Advertiser)
		} else if vv.Wrapper != nil && vv.Wrapper.Advertiser != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Advertiser)
		}
	}
	//good ;-)
	return all
}

//GetAdsPricing get the list of all Pricing
func (v *VAST) GetAdsPricing() map[string][]*Pricing {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Pricing
	all = make(map[string][]*Pricing)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Pricing != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Pricing)
		} else if vv.Wrapper != nil && vv.Wrapper.Pricing != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Pricing)
		}
	}
	//good ;-)
	return all
}

//GetAdsImpression get the list of all Impression
func (v *VAST) GetAdsImpression() map[string][]*Impression {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Impression
	all = make(map[string][]*Impression)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && len(vv.InLine.Impression) > 0 {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Impression...)
		} else if vv.Wrapper != nil && len(vv.Wrapper.Impression) > 0 {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Impression...)
		}
	}
	//good ;-)
	return all
}

//GetAdsVASTAdTagURI get the list of all VASTAdTagURI
func (v *VAST) GetAdsVASTAdTagURI() map[string][]*VASTAdTagURI {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*VASTAdTagURI
	all = make(map[string][]*VASTAdTagURI)
	//just in case
	for _, vv := range v.Ad {
		if vv.Wrapper != nil && vv.Wrapper.VASTAdTagURI != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.VASTAdTagURI)
		}
	}
	//good ;-)
	return all
}

//GetAdsAdVerifications get the list of all AdVerifications
func (v *VAST) GetAdsAdVerifications() map[string][]*AdVerifications {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*AdVerifications
	all = make(map[string][]*AdVerifications)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.AdVerifications)
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.AdVerifications)
		}
	}

	//good ;-)
	return all
}

//GetAdsAdVerification get the list of all AdVerifications.Verification
func (v *VAST) GetAdsAdVerification() map[string][]*Verification {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Verification
	all = make(map[string][]*Verification)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil && len(vv.InLine.AdVerifications.Verification) > 0 {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.AdVerifications.Verification...)
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil && len(vv.Wrapper.AdVerifications.Verification) > 0 {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.AdVerifications.Verification...)
		}
	}

	//good ;-)
	return all
}

//GetAdsAdVerificationJavaScriptResource get the list of all AdVerifications.Verification.JavaScriptResource
func (v *VAST) GetAdsAdVerificationJavaScriptResource() map[string][]*JavaScriptResource {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*JavaScriptResource
	all = make(map[string][]*JavaScriptResource)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil && len(vv.InLine.AdVerifications.Verification) > 0 {
			for _, kv := range vv.InLine.AdVerifications.Verification {
				if kv.JavaScriptResource != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.JavaScriptResource)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil && len(vv.Wrapper.AdVerifications.Verification) > 0 {
			for _, kv := range vv.Wrapper.AdVerifications.Verification {
				if kv.JavaScriptResource != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.JavaScriptResource)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsAdVerificationFlashResource get the list of all AdVerifications.Verification.FlashResource
func (v *VAST) GetAdsAdVerificationFlashResource() map[string][]*FlashResource {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*FlashResource
	all = make(map[string][]*FlashResource)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil && len(vv.InLine.AdVerifications.Verification) > 0 {
			for _, kv := range vv.InLine.AdVerifications.Verification {
				if kv.FlashResource != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.FlashResource)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil && len(vv.Wrapper.AdVerifications.Verification) > 0 {
			for _, kv := range vv.Wrapper.AdVerifications.Verification {
				if kv.FlashResource != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.FlashResource)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsAdVerificationExecutableResource get the list of all AdVerifications.Verification.ExecutableResource
func (v *VAST) GetAdsAdVerificationExecutableResource() map[string][]*ExecutableResource {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*ExecutableResource
	all = make(map[string][]*ExecutableResource)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil && len(vv.InLine.AdVerifications.Verification) > 0 {
			for _, kv := range vv.InLine.AdVerifications.Verification {
				if kv.ExecutableResource != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.ExecutableResource)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil && len(vv.Wrapper.AdVerifications.Verification) > 0 {
			for _, kv := range vv.Wrapper.AdVerifications.Verification {
				if kv.ExecutableResource != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.ExecutableResource)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsAdVerificationVerificationParameters get the list of all AdVerifications.Verification.VerificationParameters
func (v *VAST) GetAdsAdVerificationVerificationParameters() map[string][]*VerificationParameters {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*VerificationParameters
	all = make(map[string][]*VerificationParameters)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil && len(vv.InLine.AdVerifications.Verification) > 0 {
			for _, kv := range vv.InLine.AdVerifications.Verification {
				if kv.VerificationParameters != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.VerificationParameters)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil && len(vv.Wrapper.AdVerifications.Verification) > 0 {
			for _, kv := range vv.Wrapper.AdVerifications.Verification {
				if kv.VerificationParameters != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.VerificationParameters)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsAdVerificationTracking get the list of all AdVerifications.Verification.TrackingEvents.Tracking
func (v *VAST) GetAdsAdVerificationTracking() map[string][]*Tracking {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Tracking
	all = make(map[string][]*Tracking)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil && len(vv.InLine.AdVerifications.Verification) > 0 {
			for _, kv := range vv.InLine.AdVerifications.Verification {
				if kv.TrackingEvents != nil && len(kv.TrackingEvents.Tracking) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.TrackingEvents.Tracking...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil && len(vv.Wrapper.AdVerifications.Verification) > 0 {
			for _, kv := range vv.Wrapper.AdVerifications.Verification {
				if kv.TrackingEvents != nil && len(kv.TrackingEvents.Tracking) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.TrackingEvents.Tracking...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsAdVerificationViewableImpression get the list of all AdVerifications.Verification.ViewableImpression
func (v *VAST) GetAdsAdVerificationViewableImpression() map[string][]*ViewableImpression {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*ViewableImpression
	all = make(map[string][]*ViewableImpression)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil && len(vv.InLine.AdVerifications.Verification) > 0 {
			for _, kv := range vv.InLine.AdVerifications.Verification {
				if len(kv.ViewableImpression) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.ViewableImpression...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil && len(vv.Wrapper.AdVerifications.Verification) > 0 {
			for _, kv := range vv.Wrapper.AdVerifications.Verification {
				if len(kv.ViewableImpression) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.ViewableImpression...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsAdVerificationViewable get the list of all AdVerifications.Verification.ViewableImpression.Viewable
func (v *VAST) GetAdsAdVerificationViewable() map[string][]*Viewable {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Viewable
	all = make(map[string][]*Viewable)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil && len(vv.InLine.AdVerifications.Verification) > 0 {
			for _, kv := range vv.InLine.AdVerifications.Verification {
				for _, sv := range kv.ViewableImpression {
					if sv.Viewable != nil {
						all[AdTypeIsInline] = append(all[AdTypeIsInline], sv.Viewable)
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil && len(vv.Wrapper.AdVerifications.Verification) > 0 {
			for _, kv := range vv.Wrapper.AdVerifications.Verification {
				for _, sv := range kv.ViewableImpression {
					if sv.Viewable != nil {
						all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], sv.Viewable)
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsAdVerificationNotViewable get the list of all AdVerifications.Verification.ViewableImpression.NotViewable
func (v *VAST) GetAdsAdVerificationNotViewable() map[string][]*NotViewable {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*NotViewable
	all = make(map[string][]*NotViewable)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil && len(vv.InLine.AdVerifications.Verification) > 0 {
			for _, kv := range vv.InLine.AdVerifications.Verification {
				for _, sv := range kv.ViewableImpression {
					if sv.NotViewable != nil {
						all[AdTypeIsInline] = append(all[AdTypeIsInline], sv.NotViewable)
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil && len(vv.Wrapper.AdVerifications.Verification) > 0 {
			for _, kv := range vv.Wrapper.AdVerifications.Verification {
				for _, sv := range kv.ViewableImpression {
					if sv.NotViewable != nil {
						all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], sv.NotViewable)
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsAdVerificationViewUndetermined get the list of all AdVerifications.Verification.ViewableImpression.ViewUndetermined
func (v *VAST) GetAdsAdVerificationViewUndetermined() map[string][]*ViewUndetermined {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*ViewUndetermined
	all = make(map[string][]*ViewUndetermined)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.AdVerifications != nil && len(vv.InLine.AdVerifications.Verification) > 0 {
			for _, kv := range vv.InLine.AdVerifications.Verification {
				for _, sv := range kv.ViewableImpression {
					if sv.ViewUndetermined != nil {
						all[AdTypeIsInline] = append(all[AdTypeIsInline], sv.ViewUndetermined)
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.AdVerifications != nil && len(vv.Wrapper.AdVerifications.Verification) > 0 {
			for _, kv := range vv.Wrapper.AdVerifications.Verification {
				for _, sv := range kv.ViewableImpression {
					if sv.ViewUndetermined != nil {
						all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], sv.ViewUndetermined)
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsExtensions get the list of all Extensions
func (v *VAST) GetAdsExtensions() map[string][]*Extensions {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Extensions
	all = make(map[string][]*Extensions)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Extensions != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Extensions)
		} else if vv.Wrapper != nil && vv.Wrapper.Extensions != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Extensions)
		}
	}

	//good ;-)
	return all
}

//GetAdsExtension get the list of all Extensions.Extension
func (v *VAST) GetAdsExtension() map[string][]*Extension {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Extension
	all = make(map[string][]*Extension)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Extensions != nil && len(vv.InLine.Extensions.Extension) > 0 {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Extensions.Extension...)
		} else if vv.Wrapper != nil && vv.Wrapper.Extensions != nil && len(vv.Wrapper.Extensions.Extension) > 0 {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Extensions.Extension...)
		}
	}

	//good ;-)
	return all
}

//GetAdsExtensionTotalAvailable get the list of all Extensions.Extension.TotalAvailable
func (v *VAST) GetAdsExtensionTotalAvailable() map[string][]*TotalAvailable {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*TotalAvailable
	all = make(map[string][]*TotalAvailable)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Extensions != nil && len(vv.InLine.Extensions.Extension) > 0 {
			for _, kv := range vv.InLine.Extensions.Extension {
				if kv.TotalAvailable != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.TotalAvailable)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Extensions != nil && len(vv.Wrapper.Extensions.Extension) > 0 {
			for _, kv := range vv.Wrapper.Extensions.Extension {
				if kv.TotalAvailable != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.TotalAvailable)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsExtensionCustomTracking get the list of all Extensions.Extension.CustomTracking.Tracking
func (v *VAST) GetAdsExtensionCustomTracking() map[string][]*Tracking {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Tracking
	all = make(map[string][]*Tracking)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Extensions != nil && len(vv.InLine.Extensions.Extension) > 0 {
			for _, kv := range vv.InLine.Extensions.Extension {
				if kv.CustomTracking != nil && len(kv.CustomTracking.Tracking) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.CustomTracking.Tracking...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Extensions != nil && len(vv.Wrapper.Extensions.Extension) > 0 {
			for _, kv := range vv.Wrapper.Extensions.Extension {
				if kv.CustomTracking != nil && len(kv.CustomTracking.Tracking) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.CustomTracking.Tracking...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsExtensionAdVerification get the list of all Extensions.Extension.AdVerifications.Verification
func (v *VAST) GetAdsExtensionAdVerification() map[string][]*Verification {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Verification
	all = make(map[string][]*Verification)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Extensions != nil && len(vv.InLine.Extensions.Extension) > 0 {
			for _, kv := range vv.InLine.Extensions.Extension {
				if kv.AdVerifications != nil && len(kv.AdVerifications.Verification) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.AdVerifications.Verification...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Extensions != nil && len(vv.Wrapper.Extensions.Extension) > 0 {
			for _, kv := range vv.Wrapper.Extensions.Extension {
				if kv.AdVerifications != nil && len(kv.AdVerifications.Verification) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.AdVerifications.Verification...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsViewable get the list of all Viewable
func (v *VAST) GetAdsViewable() map[string][]*Viewable {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Viewable
	all = make(map[string][]*Viewable)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && len(vv.InLine.ViewableImpression) > 0 {
			for _, kk := range vv.InLine.ViewableImpression {
				if kk.Viewable != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Viewable)
				}
			}
		} else if vv.Wrapper != nil && len(vv.Wrapper.ViewableImpression) > 0 {
			for _, kk := range vv.Wrapper.ViewableImpression {
				if kk.Viewable != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Viewable)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsNotViewable get the list of all Viewable
func (v *VAST) GetAdsNotViewable() map[string][]*NotViewable {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*NotViewable
	all = make(map[string][]*NotViewable)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && len(vv.InLine.ViewableImpression) > 0 {
			for _, kk := range vv.InLine.ViewableImpression {
				if kk.NotViewable != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.NotViewable)
				}
			}
		} else if vv.Wrapper != nil && len(vv.Wrapper.ViewableImpression) > 0 {
			for _, kk := range vv.Wrapper.ViewableImpression {
				if kk.NotViewable != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.NotViewable)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsViewUndetermined get the list of all Viewable
func (v *VAST) GetAdsViewUndetermined() map[string][]*ViewUndetermined {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*ViewUndetermined
	all = make(map[string][]*ViewUndetermined)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && len(vv.InLine.ViewableImpression) > 0 {
			for _, kk := range vv.InLine.ViewableImpression {
				if kk.ViewUndetermined != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.ViewUndetermined)
				}
			}
		} else if vv.Wrapper != nil && len(vv.Wrapper.ViewableImpression) > 0 {
			for _, kk := range vv.Wrapper.ViewableImpression {
				if kk.ViewUndetermined != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.ViewUndetermined)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsViewableImpression get the list of all ViewableImpression
func (v *VAST) GetAdsViewableImpression() map[string][]*ViewableImpression {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*ViewableImpression
	all = make(map[string][]*ViewableImpression)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && len(vv.InLine.ViewableImpression) > 0 {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.ViewableImpression...)
		} else if vv.Wrapper != nil && len(vv.Wrapper.ViewableImpression) > 0 {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.ViewableImpression...)
		}
	}

	//good ;-)
	return all
}

//GetAdsViewableImpressionViewable get the list of all ViewableImpression.Viewable
func (v *VAST) GetAdsViewableImpressionViewable() map[string][]*Viewable {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Viewable
	all = make(map[string][]*Viewable)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && len(vv.InLine.ViewableImpression) > 0 {
			for _, kv := range vv.InLine.ViewableImpression {
				if kv.Viewable != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.Viewable)
				}
			}
		} else if vv.Wrapper != nil && len(vv.Wrapper.ViewableImpression) > 0 {
			for _, kv := range vv.Wrapper.ViewableImpression {
				if kv.Viewable != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.Viewable)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsViewableImpressionNotViewable get the list of all ViewableImpression.NotViewable
func (v *VAST) GetAdsViewableImpressionNotViewable() map[string][]*NotViewable {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*NotViewable
	all = make(map[string][]*NotViewable)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && len(vv.InLine.ViewableImpression) > 0 {
			for _, kv := range vv.InLine.ViewableImpression {
				if kv.NotViewable != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.NotViewable)
				}
			}
		} else if vv.Wrapper != nil && len(vv.Wrapper.ViewableImpression) > 0 {
			for _, kv := range vv.Wrapper.ViewableImpression {
				if kv.NotViewable != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.NotViewable)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsViewableImpressionViewUndetermined get the list of all ViewableImpression.ViewUndetermined
func (v *VAST) GetAdsViewableImpressionViewUndetermined() map[string][]*ViewUndetermined {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*ViewUndetermined
	all = make(map[string][]*ViewUndetermined)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && len(vv.InLine.ViewableImpression) > 0 {
			for _, kv := range vv.InLine.ViewableImpression {
				if kv.ViewUndetermined != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kv.ViewUndetermined)
				}
			}
		} else if vv.Wrapper != nil && len(vv.Wrapper.ViewableImpression) > 0 {
			for _, kv := range vv.Wrapper.ViewableImpression {
				if kv.ViewUndetermined != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kv.ViewUndetermined)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreatives get the list of all Creatives
func (v *VAST) GetAdsCreatives() map[string][]*Creatives {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Creatives
	all = make(map[string][]*Creatives)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Creatives)
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Creatives)
		}
	}

	//good ;-)
	return all
}

//GetAdsCreative get the list of all Creative
func (v *VAST) GetAdsCreative() map[string][]*Creative {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Creative
	all = make(map[string][]*Creative)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			all[AdTypeIsInline] = append(all[AdTypeIsInline], vv.InLine.Creatives.Creative...)
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], vv.Wrapper.Creatives.Creative...)
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeExtensions get the list of all Creative.CreativeExtensions
func (v *VAST) GetAdsCreativeExtensions() map[string][]*CreativeExtensions {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*CreativeExtensions
	all = make(map[string][]*CreativeExtensions)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CreativeExtensions != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.CreativeExtensions)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CreativeExtensions != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.CreativeExtensions)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeExtension get the list of all Creative.CreativeExtensions.CreativeExtension
func (v *VAST) GetAdsCreativeExtension() map[string][]*CreativeExtension {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*CreativeExtension
	all = make(map[string][]*CreativeExtension)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CreativeExtensions != nil && len(kk.CreativeExtensions.CreativeExtension) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.CreativeExtensions.CreativeExtension...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CreativeExtensions != nil && len(kk.CreativeExtensions.CreativeExtension) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.CreativeExtensions.CreativeExtension...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeUniversalAd get the list of all Creative.UniversalAdID
func (v *VAST) GetAdsCreativeUniversalAd() map[string][]*UniversalAdID {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*UniversalAdID
	all = make(map[string][]*UniversalAdID)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.UniversalAdID != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.UniversalAdID)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.UniversalAdID != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.UniversalAdID)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinear get the list of all Creative.Linear
func (v *VAST) GetAdsCreativeLinear() map[string][]*Linear {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Linear
	all = make(map[string][]*Linear)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearDuration get the list of all Creative.Linear.Duration
func (v *VAST) GetAdsCreativeLinearDuration() map[string][]*Duration {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Duration
	all = make(map[string][]*Duration)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.Duration != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.Duration)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.Duration != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.Duration)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearTrackingEvents get the list of all Creative.Linear.TrackingEvents
func (v *VAST) GetAdsCreativeLinearTrackingEvents() map[string][]*TrackingEvents {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*TrackingEvents
	all = make(map[string][]*TrackingEvents)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.TrackingEvents != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.TrackingEvents)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.TrackingEvents != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.TrackingEvents)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearTracking get the list of all Creative.Linear.TrackingEvents.Tracking
func (v *VAST) GetAdsCreativeLinearTracking() map[string][]*Tracking {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Tracking
	all = make(map[string][]*Tracking)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.TrackingEvents != nil && len(kk.Linear.TrackingEvents.Tracking) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.TrackingEvents.Tracking...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.TrackingEvents != nil && len(kk.Linear.TrackingEvents.Tracking) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.TrackingEvents.Tracking...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearVideoClicks get the list of all Creative.Linear.VideoClicks
func (v *VAST) GetAdsCreativeLinearVideoClicks() map[string][]*VideoClicks {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*VideoClicks
	all = make(map[string][]*VideoClicks)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.VideoClicks != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.VideoClicks)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.VideoClicks != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.VideoClicks)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearVideoClickThrough get the list of all Creative.Linear.VideoClicks.ClickThrough
func (v *VAST) GetAdsCreativeLinearVideoClickThrough() map[string][]*ClickThrough {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*ClickThrough
	all = make(map[string][]*ClickThrough)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.VideoClicks != nil && kk.Linear.VideoClicks.ClickThrough != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.VideoClicks.ClickThrough)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.VideoClicks != nil && kk.Linear.VideoClicks.ClickThrough != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.VideoClicks.ClickThrough)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearVideoClickTracking get the list of all Creative.Linear.VideoClicks.ClickTracking
func (v *VAST) GetAdsCreativeLinearVideoClickTracking() map[string][]*ClickTracking {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*ClickTracking
	all = make(map[string][]*ClickTracking)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.VideoClicks != nil && kk.Linear.VideoClicks.ClickTracking != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.VideoClicks.ClickTracking)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.VideoClicks != nil && kk.Linear.VideoClicks.ClickTracking != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.VideoClicks.ClickTracking)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearVideoCustomClick get the list of all Creative.Linear.VideoClicks.CustomClick
func (v *VAST) GetAdsCreativeLinearVideoCustomClick() map[string][]*CustomClick {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*CustomClick
	all = make(map[string][]*CustomClick)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.VideoClicks != nil && kk.Linear.VideoClicks.CustomClick != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.VideoClicks.CustomClick)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.VideoClicks != nil && kk.Linear.VideoClicks.CustomClick != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.VideoClicks.CustomClick)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearMediaFiles get the list of all Creative.Linear.MediaFiles
func (v *VAST) GetAdsCreativeLinearMediaFiles() map[string][]*MediaFiles {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*MediaFiles
	all = make(map[string][]*MediaFiles)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.MediaFiles)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.MediaFiles)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearInteractiveCreativeFile get the list of all Creative.Linear.MediaFiles.InteractiveCreativeFile
func (v *VAST) GetAdsCreativeLinearInteractiveCreativeFile() map[string][]*InteractiveCreativeFile {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*InteractiveCreativeFile
	all = make(map[string][]*InteractiveCreativeFile)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil && kk.Linear.MediaFiles.InteractiveCreativeFile != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.MediaFiles.InteractiveCreativeFile)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil && kk.Linear.MediaFiles.InteractiveCreativeFile != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.MediaFiles.InteractiveCreativeFile)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearMezzanine get the list of all Creative.Linear.MediaFiles.Mezzanine
func (v *VAST) GetAdsCreativeLinearMezzanine() map[string][]*Mezzanine {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Mezzanine
	all = make(map[string][]*Mezzanine)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil && kk.Linear.MediaFiles.Mezzanine != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.MediaFiles.Mezzanine)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil && kk.Linear.MediaFiles.Mezzanine != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.MediaFiles.Mezzanine)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearClosedCaptionFiles get the list of all Creative.Linear.MediaFiles.ClosedCaptionFiles
func (v *VAST) GetAdsCreativeLinearClosedCaptionFiles() map[string][]*ClosedCaptionFiles {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*ClosedCaptionFiles
	all = make(map[string][]*ClosedCaptionFiles)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil && kk.Linear.MediaFiles.ClosedCaptionFiles != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.MediaFiles.ClosedCaptionFiles)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil && kk.Linear.MediaFiles.ClosedCaptionFiles != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.MediaFiles.ClosedCaptionFiles)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearClosedCaptionFile get the list of all Creative.Linear.MediaFiles.ClosedCaptionFiles.ClosedCaptionFile
func (v *VAST) GetAdsCreativeLinearClosedCaptionFile() map[string][]*ClosedCaptionFile {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*ClosedCaptionFile
	all = make(map[string][]*ClosedCaptionFile)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil && kk.Linear.MediaFiles.ClosedCaptionFiles != nil && len(kk.Linear.MediaFiles.ClosedCaptionFiles.ClosedCaptionFile) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.MediaFiles.ClosedCaptionFiles.ClosedCaptionFile...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil && kk.Linear.MediaFiles.ClosedCaptionFiles != nil && len(kk.Linear.MediaFiles.ClosedCaptionFiles.ClosedCaptionFile) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.MediaFiles.ClosedCaptionFiles.ClosedCaptionFile...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearMediaFile get the list of all Creative.Linear.MediaFiles.MediaFile
func (v *VAST) GetAdsCreativeLinearMediaFile() map[string][]*MediaFile {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*MediaFile
	all = make(map[string][]*MediaFile)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil && len(kk.Linear.MediaFiles.MediaFile) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.MediaFiles.MediaFile...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.MediaFiles != nil && len(kk.Linear.MediaFiles.MediaFile) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.MediaFiles.MediaFile...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearIcons get the list of all Creative.Linear.Icons
func (v *VAST) GetAdsCreativeLinearIcons() map[string][]*Icons {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Icons
	all = make(map[string][]*Icons)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.Icons != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.Icons)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.Icons != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.Icons)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeLinearIcon get the list of all Creative.Linear.Icons.Icon
func (v *VAST) GetAdsCreativeLinearIcon() map[string][]*Icon {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Icon
	all = make(map[string][]*Icon)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.Icons != nil && len(kk.Linear.Icons.Icon) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.Linear.Icons.Icon...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.Linear != nil && kk.Linear.Icons != nil && len(kk.Linear.Icons.Icon) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.Linear.Icons.Icon...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeNonLinearAds get the list of all Creative.NonLinearAds
func (v *VAST) GetAdsCreativeNonLinearAds() map[string][]*NonLinearAds {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*NonLinearAds
	all = make(map[string][]*NonLinearAds)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.NonLinearAds != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.NonLinearAds)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.NonLinearAds != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.NonLinearAds)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeNonLinear get the list of all Creative.NonLinearAds.NonLinear
func (v *VAST) GetAdsCreativeNonLinear() map[string][]*NonLinear {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*NonLinear
	all = make(map[string][]*NonLinear)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.NonLinearAds != nil && len(kk.NonLinearAds.NonLinear) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.NonLinearAds.NonLinear...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.NonLinearAds != nil && len(kk.NonLinearAds.NonLinear) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.NonLinearAds.NonLinear...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeNonLinearTracking get the list of all Creative.NonLinearAds.TrackingEvents.Tracking
func (v *VAST) GetAdsCreativeNonLinearTracking() map[string][]*Tracking {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Tracking
	all = make(map[string][]*Tracking)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.NonLinearAds != nil && kk.NonLinearAds.TrackingEvents != nil && len(kk.NonLinearAds.TrackingEvents.Tracking) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.NonLinearAds.TrackingEvents.Tracking...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.NonLinearAds != nil && kk.NonLinearAds.TrackingEvents != nil && len(kk.NonLinearAds.TrackingEvents.Tracking) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.NonLinearAds.TrackingEvents.Tracking...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeNonLinearStaticResource get the list of all Creative.NonLinearAds.NonLinear.StaticResource
func (v *VAST) GetAdsCreativeNonLinearStaticResource() map[string][]*StaticResource {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*StaticResource
	all = make(map[string][]*StaticResource)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.NonLinearAds != nil && len(kk.NonLinearAds.NonLinear) > 0 {
					for _, xk := range kk.NonLinearAds.NonLinear {
						all[AdTypeIsInline] = append(all[AdTypeIsInline], xk.StaticResource)
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.NonLinearAds != nil && len(kk.NonLinearAds.NonLinear) > 0 {
					for _, xk := range kk.NonLinearAds.NonLinear {
						all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], xk.StaticResource)
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeNonLinearClickThrough get the list of all Creative.NonLinearAds.NonLinear.NonLinearClickThrough
func (v *VAST) GetAdsCreativeNonLinearClickThrough() map[string][]*NonLinearClickThrough {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*NonLinearClickThrough
	all = make(map[string][]*NonLinearClickThrough)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.NonLinearAds != nil && len(kk.NonLinearAds.NonLinear) > 0 {
					for _, xk := range kk.NonLinearAds.NonLinear {
						all[AdTypeIsInline] = append(all[AdTypeIsInline], xk.NonLinearClickThrough)
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.NonLinearAds != nil && len(kk.NonLinearAds.NonLinear) > 0 {
					for _, xk := range kk.NonLinearAds.NonLinear {
						all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], xk.NonLinearClickThrough)
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeNonLinearClickTracking get the list of all Creative.NonLinearAds.NonLinear.NonLinearClickTracking
func (v *VAST) GetAdsCreativeNonLinearClickTracking() map[string][]*NonLinearClickTracking {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*NonLinearClickTracking
	all = make(map[string][]*NonLinearClickTracking)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.NonLinearAds != nil && len(kk.NonLinearAds.NonLinear) > 0 {
					for _, xk := range kk.NonLinearAds.NonLinear {
						all[AdTypeIsInline] = append(all[AdTypeIsInline], xk.NonLinearClickTracking)
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.NonLinearAds != nil && len(kk.NonLinearAds.NonLinear) > 0 {
					for _, xk := range kk.NonLinearAds.NonLinear {
						all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], xk.NonLinearClickTracking)
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeCompanionAds get the list of all Creative.CompanionAds
func (v *VAST) GetAdsCreativeCompanionAds() map[string][]*CompanionAds {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*CompanionAds
	all = make(map[string][]*CompanionAds)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CompanionAds != nil {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.CompanionAds)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CompanionAds != nil {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.CompanionAds)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeCompanion get the list of all Creative.CompanionAds.Companion
func (v *VAST) GetAdsCreativeCompanion() map[string][]*Companion {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Companion
	all = make(map[string][]*Companion)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					all[AdTypeIsInline] = append(all[AdTypeIsInline], kk.CompanionAds.Companion...)
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], kk.CompanionAds.Companion...)
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeCompanionAltText get the list of all Creative.CompanionAds.Companion.AltText
func (v *VAST) GetAdsCreativeCompanionAltText() map[string][]*AltText {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*AltText
	all = make(map[string][]*AltText)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.AltText != nil {
							all[AdTypeIsInline] = append(all[AdTypeIsInline], va.AltText)
						}
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.AltText != nil {
							all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], va.AltText)
						}
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeCompanionAdParameters get the list of all Creative.CompanionAds.Companion.AdParameters
func (v *VAST) GetAdsCreativeCompanionAdParameters() map[string][]*AdParameters {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*AdParameters
	all = make(map[string][]*AdParameters)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.AdParameters != nil {
							all[AdTypeIsInline] = append(all[AdTypeIsInline], va.AdParameters)
						}
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.AdParameters != nil {
							all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], va.AdParameters)
						}
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeCompanionClickThrough get the list of all Creative.CompanionAds.Companion.CompanionClickThrough
func (v *VAST) GetAdsCreativeCompanionClickThrough() map[string][]*CompanionClickThrough {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*CompanionClickThrough
	all = make(map[string][]*CompanionClickThrough)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.CompanionClickThrough != nil {
							all[AdTypeIsInline] = append(all[AdTypeIsInline], va.CompanionClickThrough)
						}
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.CompanionClickThrough != nil {
							all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], va.CompanionClickThrough)
						}
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeCompanionHTMLResource get the list of all Creative.CompanionAds.Companion.HTMLResource
func (v *VAST) GetAdsCreativeCompanionHTMLResource() map[string][]*HTMLResource {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*HTMLResource
	all = make(map[string][]*HTMLResource)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.HTMLResource != nil {
							all[AdTypeIsInline] = append(all[AdTypeIsInline], va.HTMLResource)
						}
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.HTMLResource != nil {
							all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], va.HTMLResource)
						}
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeCompanionIFrameResource get the list of all Creative.CompanionAds.Companion.IFrameResource
func (v *VAST) GetAdsCreativeCompanionIFrameResource() map[string][]*IFrameResource {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*IFrameResource
	all = make(map[string][]*IFrameResource)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.IFrameResource != nil {
							all[AdTypeIsInline] = append(all[AdTypeIsInline], va.IFrameResource)
						}
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.IFrameResource != nil {
							all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], va.IFrameResource)
						}
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeCompanionStaticResource get the list of all Creative.CompanionAds.Companion.StaticResource
func (v *VAST) GetAdsCreativeCompanionStaticResource() map[string][]*StaticResource {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*StaticResource
	all = make(map[string][]*StaticResource)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.StaticResource != nil {
							all[AdTypeIsInline] = append(all[AdTypeIsInline], va.StaticResource)
						}
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.StaticResource != nil {
							all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], va.StaticResource)
						}
					}
				}
			}
		}
	}

	//good ;-)
	return all
}

//GetAdsCreativeCompanionTracking get the list of all Creative.CompanionAds.Companion.TrackingEvents.Tracking
func (v *VAST) GetAdsCreativeCompanionTracking() map[string][]*Tracking {
	//minimal config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
	}
	var all map[string][]*Tracking
	all = make(map[string][]*Tracking)
	//just in case
	for _, vv := range v.Ad {
		if vv.InLine != nil && vv.InLine.Creatives != nil && len(vv.InLine.Creatives.Creative) > 0 {
			for _, kk := range vv.InLine.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.TrackingEvents != nil && len(va.TrackingEvents.Tracking) > 0 {
							all[AdTypeIsInline] = append(all[AdTypeIsInline], va.TrackingEvents.Tracking...)
						}
					}
				}
			}
		} else if vv.Wrapper != nil && vv.Wrapper.Creatives != nil && len(vv.Wrapper.Creatives.Creative) > 0 {
			for _, kk := range vv.Wrapper.Creatives.Creative {
				if kk.CompanionAds != nil && len(kk.CompanionAds.Companion) > 0 {
					for _, va := range kk.CompanionAds.Companion {
						if va.TrackingEvents != nil && len(va.TrackingEvents.Tracking) > 0 {
							all[AdTypeIsWrapper] = append(all[AdTypeIsWrapper], va.TrackingEvents.Tracking...)
						}
					}
				}
			}
		}
	}

	//good ;-)
	return all
}
