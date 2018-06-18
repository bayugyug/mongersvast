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
func (v *VAST) SetAd(adVersion, adID, adSequence, adConditional string) (req *VAST) {
	//minimal config
	req = &VAST{
		Version: adVersion,
		Ad:      []*Ad{},
	}
	//optional maybe ;-)
	if adID != "" {
		req.Ad[0].ID = adID
	}
	//optional maybe ;-)
	if adSequence != "" {
		req.Ad[0].Sequence = adSequence
	}
	//optional maybe ;-)
	if adConditional != "" {
		req.Ad[0].ConditionalAd = adConditional
	}
	//good ;-)
	return
}

//SetInLineAd set the minimum InLineAd
func (v *VAST) SetInLineAd(inlineID string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//add the inline
	req.Ad[0].InLine = &InLine{
		ID:                inlineID,
		InLineWrapperData: InLineWrapperData{},
	}
	//good ;-)
	return
}

//SetWrapperAd set the minimum WrapperAd
func (v *VAST) SetWrapperAd(wrapperID, followAdditionalWrappers, allowMultipleAds, fallbackOnNoAd string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//add the wrapper
	req.Ad[0].Wrapper = &Wrapper{
		ID: wrapperID,
		FollowAdditionalWrappers: followAdditionalWrappers,
		AllowMultipleAds:         allowMultipleAds,
		FallbackOnNoAd:           fallbackOnNoAd,
		InLineWrapperData:        InLineWrapperData{},
	}
	//good ;-)
	return
}

//SetAdSystem set the AdSystem
func (v *VAST) SetAdSystem(s string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := &AdSystem{
		Value: s,
	}
	//check which type
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.AdSystem = data
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.AdSystem = data
	}
	//good ;-)
	return
}

//SetAdTitle set the AdTitle
func (v *VAST) SetAdTitle(s string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := &AdTitle{
		Value: s,
	}
	//check which type
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.AdTitle = data
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.AdTitle = data
	}
	//good ;-)
	return
}

//SetDescription set the Description
func (v *VAST) SetDescription(s string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := &Description{
		Value: s,
	}
	//check which type
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.Description = data
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.Description = data
	}
	//good ;-)
	return
}

//SetErrorURL set the Error URL
func (v *VAST) SetErrorURL(s string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := &VASTError{
		Value: s,
	}
	//check which type
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.Error = data
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.Error = data
	}
	//good ;-)
	return
}

//SetImpressionURL set the Impression URL
func (v *VAST) SetImpressionURL(impID, impURL string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
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
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.Impression = append(req.Ad[0].Wrapper.InLineWrapperData.Impression, data)
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.Impression = append(req.Ad[0].InLine.InLineWrapperData.Impression, data)
	}
	//good ;-)
	return
}

//SetAdServingID set the AdServingID
func (v *VAST) SetAdServingID(adID, adValue string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
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
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.AdServingID = data
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.AdServingID = data
	}
	//good ;-)
	return
}

//SetSurvey set the Survey
func (v *VAST) SetSurvey(s string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := &Survey{
		Value: s,
	}
	//check which type
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.Survey = data
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.Survey = data
	}
	//good ;-)
	return
}

//SetViewableImpression add into the list of ViewableImpression
func (v *VAST) SetViewableImpression(sID string, viewable *Viewable, notviewable *NotViewable, undetermined *ViewUndetermined) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
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
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.ViewableImpression = append(req.Ad[0].Wrapper.InLineWrapperData.ViewableImpression, data)
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.ViewableImpression = append(req.Ad[0].InLine.InLineWrapperData.ViewableImpression, data)
	}
	//good ;-)
	return
}

//SetVASTAdTagURI set the VASTAdTagURI
func (v *VAST) SetVASTAdTagURI(adID, adValue string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
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
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.VASTAdTagURI = data
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.VASTAdTagURI = data
	}
	//good ;-)
	return
}

//SetPricing set the Pricing
func (v *VAST) SetPricing(adID, adModel, adCurr, adValue string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := &Pricing{
		ID:       adID,
		Model:    adModel,
		Currency: adCurr,
		Value:    adValue,
	}
	//check which type
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.Pricing = data
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.Pricing = data
	}
	//good ;-)
	return
}

//SetAdvertiser set the Advertiser
func (v *VAST) SetAdvertiser(s string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//set 1
	data := &Advertiser{
		Value: s,
	}
	//check which type
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.Advertiser = data
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.Advertiser = data
	}
	//good ;-)
	return
}

//SetCategory add into the list of Category
func (v *VAST) SetCategory(sID, sAuth, sValue string) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
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
	if req.Ad[0].Wrapper != nil {
		req.Ad[0].Wrapper.InLineWrapperData.Category = append(req.Ad[0].Wrapper.InLineWrapperData.Category, data)
	} else if req.Ad[0].InLine != nil {
		req.Ad[0].InLine.InLineWrapperData.Category = append(req.Ad[0].InLine.InLineWrapperData.Category, data)
	}
	//good ;-)
	return
}

//SetVerification add into the list of Verification
func (v *VAST) SetVerification(jscript *JavaScriptResource, verificationp *VerificationParameters, trkevents *TrackingEvents) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
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
	if req.Ad[0].Wrapper != nil {
		if req.Ad[0].Wrapper.InLineWrapperData.AdVerifications == nil {
			req.Ad[0].Wrapper.InLineWrapperData.AdVerifications = &AdVerifications{}
		}
		req.Ad[0].Wrapper.InLineWrapperData.AdVerifications.Verification = append(req.Ad[0].Wrapper.InLineWrapperData.AdVerifications.Verification, data)
	} else if req.Ad[0].InLine != nil {
		if req.Ad[0].InLine.InLineWrapperData.AdVerifications == nil {
			req.Ad[0].InLine.InLineWrapperData.AdVerifications = &AdVerifications{}
		}
		req.Ad[0].InLine.InLineWrapperData.AdVerifications.Verification = append(req.Ad[0].InLine.InLineWrapperData.AdVerifications.Verification, data)
	}
	//good ;-)
	return
}

//SetExtension add into the list of Extension
func (v *VAST) SetExtension(sType, sValue string, total *TotalAvailable, adverifications *AdVerifications) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
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
	if req.Ad[0].Wrapper != nil {
		if req.Ad[0].Wrapper.InLineWrapperData.Extensions == nil {
			req.Ad[0].Wrapper.InLineWrapperData.Extensions = &Extensions{}
		}
		req.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension = append(req.Ad[0].Wrapper.InLineWrapperData.Extensions.Extension, data)
	} else if req.Ad[0].InLine != nil {
		if req.Ad[0].InLine.InLineWrapperData.Extensions == nil {
			req.Ad[0].InLine.InLineWrapperData.Extensions = &Extensions{}
		}
		req.Ad[0].InLine.InLineWrapperData.Extensions.Extension = append(req.Ad[0].InLine.InLineWrapperData.Extensions.Extension, data)
	}
	//good ;-)
	return
}

//SetCreative add into the list of Creative
func (v *VAST) SetCreative(sID, sAdID, sSequence, sFramework string, linear *Linear, nonLinear *NonLinearAds, companion *CompanionAds, universal *UniversalAdID) (req *VAST) {
	//minimal config
	if v == nil {
		req = v.SetAd(VastXMLVer2, "", "", "")
	}
	//check flag
	if req == nil || len(req.Ad) <= 0 {
		req = v.SetAd(VastXMLVer2, "", "", "")
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
	if req.Ad[0].Wrapper != nil {
		if req.Ad[0].Wrapper.InLineWrapperData.Creatives == nil {
			req.Ad[0].Wrapper.InLineWrapperData.Creatives = &Creatives{}
		}
		req.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative = append(req.Ad[0].Wrapper.InLineWrapperData.Creatives.Creative, data)
	} else if req.Ad[0].InLine != nil {
		if req.Ad[0].InLine.InLineWrapperData.Creatives == nil {
			req.Ad[0].InLine.InLineWrapperData.Creatives = &Creatives{}
		}
		req.Ad[0].InLine.InLineWrapperData.Creatives.Creative = append(req.Ad[0].InLine.InLineWrapperData.Creatives.Creative, data)
	}
	//good ;-)
	return
}
