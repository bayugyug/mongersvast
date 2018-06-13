package mongersvast

//SAMPLE VAST XMLS
var (
	XMLInlineNonLinear = `
<?xml version="1.0" encoding="UTF-8"?>
<VAST version="2.0">
  <Ad id="602678">
  <InLine>
    <AdSystem>Acudeo Compatible</AdSystem>
    <AdTitle>NonLinear Test Campaign 1</AdTitle>
    <Description>NonLinear Test Campaign 1</Description>
    <Survey>http://mySurveyURL/survey</Survey>
    <Error>http://mongers.vast.utils/error</Error>
    <Impression>http://mongers.vast.utils/impression</Impression>
	<Creatives>
		<Creative AdID="602678-NonLinear">
			<NonLinearAds>
				<TrackingEvents>
					<Tracking event="creativeView">http://mongers.vast.utils/nonlinear/creativeView</Tracking>
					<Tracking event="expand">http://mongers.vast.utils/nonlinear/expand</Tracking>
					<Tracking event="collapse">http://mongers.vast.utils/nonlinear/collapse</Tracking>
					<Tracking event="acceptInvitation">http://mongers.vast.utils/nonlinear/acceptInvitation</Tracking>
					<Tracking event="close">http://mongers.vast.utils/nonlinear/close</Tracking>
				</TrackingEvents>
				<NonLinear height="50" width="300" minSuggestedDuration="00:00:15">
					<StaticResource creativeType="image/jpeg">
					http://demo-mongers.vast.utils/proddev/vast/50x300_static.jpg
					</StaticResource>
					<NonLinearClickThrough>http://mongers.vast.utils</NonLinearClickThrough>
				</NonLinear>
				<NonLinear height="50" width="450" minSuggestedDuration="00:00:20">
					<StaticResource creativeType="image/jpeg">
					http://demo-mongers.vast.utils/proddev/vast/50x450_static.jpg
					</StaticResource>
					<NonLinearClickThrough>http://mongers.vast.utils</NonLinearClickThrough>
				</NonLinear>
			</NonLinearAds>
		</Creative>
		<Creative AdID="602678-Companion">
			<CompanionAds>
				<Companion width="300" height="250">
					<StaticResource creativeType="application/x-shockwave-flash">http://demo-mongers.vast.utils/proddev/vast/300x250_companion_1.swf</StaticResource>
 
					<CompanionClickThrough>http://mongers.vast.utils</CompanionClickThrough>
				</Companion>
				<Companion width="728" height="90">
					<StaticResource creativeType="image/jpeg">http://demo-mongers.vast.utils/proddev/vast/728x90_banner1.jpg</StaticResource>
					<TrackingEvents>
						<Tracking event="creativeView">http://mongers.vast.utils/secondCompanion</Tracking>
					</TrackingEvents>
					<CompanionClickThrough>http://mongers.vast.utils</CompanionClickThrough>
				</Companion>
			</CompanionAds>
		</Creative>
	</Creatives>
  </InLine>
  </Ad>
</VAST>
`
	XMLInlineLinear = `
<?xml version="1.0" encoding="UTF-8"?>
<VAST version="2.0">
  <Ad id="601364">
  <InLine>
    <AdSystem>Acudeo Compatible</AdSystem>
    <AdTitle>VAST 2.0 Instream Test 1</AdTitle>
    <Description>VAST 2.0 Instream Test 1</Description>
    <Error>http://mongers.vast.utils/error</Error>
    <Impression>http://mongers.vast.utils/impression</Impression>
	<Creatives>
		<Creative AdID="601364">
			<Linear>
				<Duration>00:00:30</Duration>
				<TrackingEvents>
					<Tracking event="creativeView">http://mongers.vast.utils/creativeView</Tracking>
					<Tracking event="start">http://mongers.vast.utils/start</Tracking>
					<Tracking event="midpoint">http://mongers.vast.utils/midpoint</Tracking>
					<Tracking event="firstQuartile">http://mongers.vast.utils/firstQuartile</Tracking>
					<Tracking event="thirdQuartile">http://mongers.vast.utils/thirdQuartile</Tracking>
					<Tracking event="complete">http://mongers.vast.utils/complete</Tracking>
				</TrackingEvents>
				<VideoClicks>
					<ClickThrough>http://mongers.vast.utils</ClickThrough>
					<ClickTracking>http://mongers.vast.utils/click</ClickTracking>
				</VideoClicks>
				<MediaFiles>
					<MediaFile delivery="progressive" type="video/x-flv" bitrate="500" width="400" height="300" scalable="true" maintainAspectRatio="true">http://cdnp.tremormedia.com/video/acudeo/Carrot_400x300_500kb.flv</MediaFile>
				</MediaFiles>
			</Linear>
		</Creative>
		<Creative AdID="601364-Companion">
			<CompanionAds>
				<Companion width="300" height="250">
					<StaticResource creativeType="image/jpeg">http://demo-mongers.vast.utils/proddev/vast/Blistex1.jpg</StaticResource>
					<TrackingEvents>
						<Tracking event="creativeView">http://mongers.vast.utils/firstCompanionCreativeView</Tracking>
					</TrackingEvents>

					<CompanionClickThrough>http://mongers.vast.utils</CompanionClickThrough>
				</Companion>
				<Companion width="728" height="90">
					<StaticResource creativeType="image/jpeg">http://demo-mongers.vast.utils/proddev/vast/728x90_banner1.jpg</StaticResource>
					<CompanionClickThrough>http://mongers.vast.utils</CompanionClickThrough>
				</Companion>
			</CompanionAds>
		</Creative>
	</Creatives>
  </InLine>
  </Ad>
</VAST>
`

	XMLWrapperLinear1 = `
<?xml version="1.0" encoding="UTF-8"?>
<VAST version="2.0">
  <Ad id="602833">
  <Wrapper>
    <AdSystem>Acudeo Compatible</AdSystem>
    <VASTAdTagURI>http://demo-mongers.vast.utils/proddev/vast/vast_inline_linear.xml</VASTAdTagURI>
    <Error>http://mongers.vast.utils/wrapper/error</Error>
    <Impression>http://mongers.vast.utils/wrapper/impression</Impression>
	<Creatives>
		<Creative AdID="602833">
			<Linear>
				<TrackingEvents>
					<Tracking event="creativeView">http://mongers.vast.utils/wrapper/creativeView</Tracking>
					<Tracking event="start">http://mongers.vast.utils/wrapper/start</Tracking>
					<Tracking event="midpoint">http://mongers.vast.utils/wrapper/midpoint</Tracking>
					<Tracking event="firstQuartile">http://mongers.vast.utils/wrapper/firstQuartile</Tracking>
					<Tracking event="thirdQuartile">http://mongers.vast.utils/wrapper/thirdQuartile</Tracking>
					<Tracking event="complete">http://mongers.vast.utils/wrapper/complete</Tracking>
					<Tracking event="mute">http://mongers.vast.utils/wrapper/mute</Tracking>
					<Tracking event="unmute">http://mongers.vast.utils/wrapper/unmute</Tracking>
					<Tracking event="pause">http://mongers.vast.utils/wrapper/pause</Tracking>
					<Tracking event="resume">http://mongers.vast.utils/wrapper/resume</Tracking>
					<Tracking event="fullscreen">http://mongers.vast.utils/wrapper/fullscreen</Tracking>
				</TrackingEvents>
			</Linear>
		</Creative>
		<Creative>
			<Linear>
				<VideoClicks>
					<ClickTracking>http://mongers.vast.utils/wrapper/click</ClickTracking>
				</VideoClicks>
			</Linear>
		</Creative>
		<Creative AdID="602833-NonLinearTracking">
			<NonLinearAds>
				<TrackingEvents>
				</TrackingEvents>
			</NonLinearAds>
		</Creative>
	</Creatives>
  </Wrapper>
  </Ad>
</VAST>
`

	XMLWrapperLinear2 = `
	<?xml version="1.0" encoding="UTF-8"?>
<VAST version="2.0">
  <Ad id="602833">
  <Wrapper>
    <AdSystem>Acudeo Compatible</AdSystem>
    <VASTAdTagURI>http://demo-mongers.vast.utils/proddev/vast/vast_inline_linear.xml</VASTAdTagURI>
    <Impression>http://mongers.vast.utils/wrapper/impression</Impression>
	<Creatives>
		<Creative AdID="602833">
			<Linear>
				<TrackingEvents>
				</TrackingEvents>
			</Linear>
		</Creative>
		<Creative AdID="602833-Companion">
			<CompanionAds>
				<Companion width="300" height="250">
					<StaticResource creativeType="image/jpeg">http://demo-mongers.vast.utils/proddev/vast/300x250_banner1.jpg</StaticResource>
					<TrackingEvents>
						<Tracking event="creativeView">http://mongers.vast.utils/wrapper/firstCompanionCreativeView</Tracking>
					</TrackingEvents>
					<CompanionClickThrough>http://mongers.vast.utils</CompanionClickThrough>
				</Companion>
				<Companion width="728" height="90">
					<StaticResource creativeType="image/jpeg">http://demo-mongers.vast.utils/proddev/vast/728x90_banner1.jpg</StaticResource>
					<CompanionClickThrough>http://mongers.vast.utils</CompanionClickThrough>
				</Companion>
			</CompanionAds>
		</Creative>
	</Creatives>
  </Wrapper>
  </Ad>
</VAST>
`

	XMLWrapperNonLinear1 = `
	<?xml version="1.0" encoding="UTF-8"?>
<VAST version="2.0">
  <Ad id="602867">
  <Wrapper>
    <AdSystem>Acudeo Compatible</AdSystem>
    <VASTAdTagURI>http://demo-mongers.vast.utils/proddev/vast/vast_inline_nonlinear2.xml</VASTAdTagURI>
    <Error>http://mongers.vast.utils/wrapper/error</Error>
    <Impression>http://mongers.vast.utils/wrapper/impression</Impression>
	<Creatives>
		<Creative AdID="602867">
			<Linear>
				<TrackingEvents>
				</TrackingEvents>
			</Linear>
		</Creative>
		<Creative AdID="602867-NonLinearTracking">
			<NonLinearAds>
				<TrackingEvents>
					<Tracking event="creativeView">http://mongers.vast.utils/wrapper/nonlinear/creativeView/creativeView</Tracking>
					<Tracking event="expand">http://mongers.vast.utils/wrapper/nonlinear/creativeView/expand</Tracking>
					<Tracking event="collapse">http://mongers.vast.utils/wrapper/nonlinear/creativeView/collapse</Tracking>
					<Tracking event="acceptInvitation">http://mongers.vast.utils/wrapper/nonlinear/creativeView/acceptInvitation</Tracking>
					<Tracking event="close">http://mongers.vast.utils/wrapper/nonlinear/creativeView/close</Tracking>
				</TrackingEvents>
			</NonLinearAds>
		</Creative>
	</Creatives>
  </Wrapper>
  </Ad>
</VAST>
`

	XMLWrapperNonLinear2 = `
	<?xml version="1.0" encoding="UTF-8"?>
<VAST version="2.0">
  <Ad id="602867">
  <Wrapper>
    <AdSystem>Acudeo Compatible</AdSystem>
    <VASTAdTagURI>http://demo-mongers.vast.utils/proddev/vast/vast_inline_nonlinear3.xml</VASTAdTagURI>
    <Impression>http://mongers.vast.utils/wrapper/impression</Impression>
	<Creatives>
		<Creative AdID="602867">
			<Linear>
				<TrackingEvents>
				</TrackingEvents>
			</Linear>
		</Creative>
		<Creative AdID="602867-NonLinearTracking">
			<NonLinearAds>
				<TrackingEvents>
				</TrackingEvents>
			</NonLinearAds>
		</Creative>
		<Creative AdID="602867-Companion">
			<CompanionAds>
				<Companion width="300" height="250">
					<StaticResource creativeType="image/jpeg">http://demo-mongers.vast.utils/proddev/vast/300x250_banner1.jpg</StaticResource>
					<TrackingEvents>
						<Tracking event="creativeView">http://mongers.vast.utils/wrapper/firstCompanionCreativeView</Tracking>
					</TrackingEvents>
					<CompanionClickThrough>http://mongers.vast.utils</CompanionClickThrough>
				</Companion>
				<Companion width="728" height="90">
					<StaticResource creativeType="image/jpeg">http://demo-mongers.vast.utils/proddev/vast/728x90_banner1.jpg</StaticResource>
					<CompanionClickThrough>http://mongers.vast.utils</CompanionClickThrough>
				</Companion>
			</CompanionAds>
		</Creative>
	</Creatives>
  </Wrapper>
  </Ad>
</VAST>
`
)
