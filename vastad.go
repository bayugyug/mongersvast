package mongersvast

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

//FromString load and unmarshal from string
func (v *VAST) FromString(body string) {
	body = strings.Replace(body, "\n", "", -1)
	body = strings.Replace(body, "\t", "", -1)
	body = strings.Replace(body, "\r", "", -1)
	xml.Unmarshal([]byte(body), v)
}

//ToString convert to string
func (v *VAST) ToString() (string, error) {
	//sanity check
	if v == nil {
		return "", ErrFailedToStringNilValue
	}
	w := new(bytes.Buffer)
	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		return "", fmt.Errorf("%s , %s", ErrFailedToString.Error(), err.Error())
	}
	return strings.TrimSpace(VastXMLHeader + "\n" + w.String()), nil
}

//FromXML an alias to fromString
func (v *VAST) FromXML(body string) {
	v.FromString(body)
}

//ToXML an alias to toString
func (v *VAST) ToXML() (string, error) {
	return v.ToString()
}

//Stringify an alias to toString
func (v *VAST) Stringify() (string, error) {
	return v.ToString()
}

//FromFile load from file
func (v *VAST) FromFile(filename string) {
	content, _ := ioutil.ReadFile(filename) //make sure the xml is readable and exists
	v.FromString(strings.TrimSpace(string(content)))
}

//ToFile save the xml into a file
func (v *VAST) ToFile(filename, body string) (bool, error) {
	var f *os.File
	var err error
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return false, fmt.Errorf("%s , %s", ErrFailedFileSave.Error(), err.Error())
	}
	defer f.Close()
	_, err = f.Write([]byte(body))
	if err != nil {
		return false, fmt.Errorf("%s , %s", ErrFailedFileSave.Error(), err.Error())
	}
	return true, nil
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

//SetXMLHeaders set the xml headers simply
func (v *VAST) SetXMLHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")
	w.Header().Set("Access-Control-Allow-Origin", "*") //Google HTML5 SDK CORS Header
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Max-Age", "10080")
}

//PushXML push content with proper xml hdrs
func (v *VAST) PushXML(w http.ResponseWriter) {
	//just in case ;-)
	if v == nil {
		return
	}
	xml, _ := v.Stringify()
	v.SetXMLHeaders(w)
	io.WriteString(w, xml)
}

//VideoDuration convert duration seconds
func (v *VAST) VideoDuration(secs int) *Duration {
	//just in case ;-)
	if v == nil {
		return nil
	}
	ts := time.Duration(secs) * time.Second
	tm := strings.TrimSpace(fmt.Sprintf("%02d:%02d:%02d", int(math.Mod(ts.Hours(), 12)), int(math.Mod(ts.Minutes(), 60)), int(math.Mod(ts.Seconds(), 60))))
	return &Duration{Value: tm}
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
	v.Ad = append(v.Ad, &Ad{})
	//optional maybe ;-)
	if adID != "" {
		v.Ad[0].ID = adID
	}
	//optional maybe ;-)
	if adSequence != "" {
		v.Ad[0].Sequence = adSequence
	}
	//optional maybe ;-)
	if adConditional != "" {
		v.Ad[0].ConditionalAd = adConditional
	}
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
	//set 1
	data := &Impression{
		Value: impURL,
	}
	//optional maybe ;-)
	if impID != "" {
		data.ID = impID
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
	//set 1
	data := &AdServingID{
		Value: adValue,
	}
	//optional maybe ;-)
	if adID != "" {
		data.ID = adID
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
	//set 1
	data := &ViewableImpression{
		ID: sID,
	}
	//optional maybe ;-)
	if viewable != nil {
		data.Viewable = viewable
	}
	//optional maybe ;-)
	if notviewable != nil {
		data.NotViewable = notviewable
	}
	//optional maybe ;-)
	if undetermined != nil {
		data.ViewUndetermined = undetermined
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
	//set 1
	data := &VASTAdTagURI{
		Value: adValue,
	}
	//optional maybe ;-)
	if adID != "" {
		data.ID = adID
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
	//set 1
	data := &Category{
		Value: sValue,
	}
	//optional maybe ;-)
	if sID != "" {
		data.ID = sID
	}
	//optional maybe ;-)
	if sAuth != "" {
		data.Authority = sAuth
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

//SetVerification add into the list of Verification
func (v *VAST) SetVerification(jscript *JavaScriptResource, verificationp *VerificationParameters, trkevents *TrackingEvents) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := &Verification{}
	//optional maybe ;-)
	if jscript != nil {
		data.JavaScriptResource = jscript
	}
	//optional maybe ;-)
	if verificationp != nil {
		data.VerificationParameters = verificationp
	}
	//optional maybe ;-)
	if trkevents != nil {
		data.TrackingEvents = trkevents
	}
	//check which type
	if v.Ad[0].Wrapper != nil {
		if v.Ad[0].Wrapper.InLineWrapperData.AdVerifications == nil {
			v.Ad[0].Wrapper.InLineWrapperData.AdVerifications = &AdVerifications{}
		}
		v.Ad[0].Wrapper.InLineWrapperData.AdVerifications.Verification = append(v.Ad[0].Wrapper.InLineWrapperData.AdVerifications.Verification, data)
	} else if v.Ad[0].InLine != nil {
		if v.Ad[0].InLine.InLineWrapperData.AdVerifications == nil {
			v.Ad[0].InLine.InLineWrapperData.AdVerifications = &AdVerifications{}
		}
		v.Ad[0].InLine.InLineWrapperData.AdVerifications.Verification = append(v.Ad[0].InLine.InLineWrapperData.AdVerifications.Verification, data)
	}
	//good ;-)
	return v
}

//SetExtension add into the list of Extension
func (v *VAST) SetExtension(sType, sValue string, total *TotalAvailable, adverifications *AdVerifications) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := &Extension{}
	//optional maybe ;-)
	if sType != "" {
		data.Type = sType
	}
	//optional maybe ;-)
	if sValue != "" {
		data.Value = sValue
	}
	//optional maybe ;-)
	if total != nil {
		data.TotalAvailable = total
	}
	//optional maybe ;-)
	if adverifications != nil {
		data.AdVerifications = adverifications
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

//SetCreativeRow add into the list of Creative
func (v *VAST) SetCreativeRow(sID, sAdID, sSequence, sFramework string, linear *Linear, nonLinear *NonLinearAds, companion *CompanionAds, universal *UniversalAdID) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := &Creative{}
	//optional maybe ;-)
	if sID != "" {
		data.ID = sID
	}
	//optional maybe ;-)
	if sAdID != "" {
		data.AdID = sAdID
	}
	//optional maybe ;-)
	if sSequence != "" {
		data.Sequence = sSequence
	}
	//optional maybe ;-)
	if sFramework != "" {
		data.APIFramework = sFramework
	}
	//optional maybe ;-)
	if linear != nil {
		data.Linear = linear
	}
	//optional maybe ;-)
	if nonLinear != nil {
		data.NonLinearAds = nonLinear
	}
	//optional maybe ;-)
	if companion != nil {
		data.CompanionAds = companion
	}
	//optional maybe ;-)
	if universal != nil {
		data.UniversalAdID = universal
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
func (v *VAST) SetCreative(creative *Creative) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := creative
	if data == nil {
		data = &Creative{}
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

//SetLinear add into the Linear obj
func (v *VAST) SetLinear(row *Linear) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
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
	//set 1
	data := &Tracking{Event: sEvent, Value: sValue}
	//optional maybe;-)
	if sOffset != "" {
		data.Offset = sOffset
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
	//set 1
	data := &ClickThrough{Value: sValue}
	//optional maybe;-)
	if sID != "" {
		data.ID = sID
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
	//set 1
	data := &ClickTracking{Value: sValue}
	//optional maybe;-)
	if sID != "" {
		data.ID = sID
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
	//set 1
	data := &CustomClick{Value: sValue}
	//optional maybe;-)
	if sID != "" {
		data.ID = sID
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
	//set 1
	data := &MediaFile{
		Value:    sValue,
		Delivery: sDelivery,
		Type:     sType,
		Width:    sWidth,
		Height:   sHeight,
		Bitrate:  sBitrate,
	}
	//optional maybe;-)
	if sID != "" {
		data.ID = sID
	}
	//optional maybe;-)
	if sMinBitrate != "" {
		data.MinBitrate = sMinBitrate
	}
	//optional maybe;-)
	if sMaxBitrate != "" {
		data.MaxBitrate = sMaxBitrate
	}
	//optional maybe;-)
	if sScalable != "" {
		data.Scalable = sScalable
	}
	//optional maybe;-)
	if sMaintainAspectRatio != "" {
		data.MaintainAspectRatio = sMaintainAspectRatio
	}
	//optional maybe;-)
	if sCodec != "" {
		data.Codec = sCodec
	}
	//optional maybe;-)
	if sAPIFramework != "" {
		data.APIFramework = sAPIFramework
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
	//set 1
	data := &Mezzanine{
		Value: sValue,
	}
	//optional maybe;-)
	if sID != "" {
		data.ID = sID
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
	//set 1
	data := &InteractiveCreativeFile{
		Value: sValue,
	}
	//optional maybe;-)
	if sID != "" {
		data.ID = sID
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
