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

//GetAdsCreativeUniversalAdID get the list of all Creative.UniversalAdID
func (v *VAST) GetAdsCreativeUniversalAdID() map[string][]*UniversalAdID {
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
