#TODO


2.1.5.1 crossdomain.xml for Flash


http://adserver.com/crossdomain.xml


<cross-domain-policy>
<allow-access-from domain=”*”>
<cross-domain-policy>

## VAST HEIRARCHY

VAST
    - Ad
        - Inline
        - Wrapper

            ** (Required)
            - AdSystem
            - AdTitle
            - Impression
            - Creatives
                - Creative
                    - CompanionAds
                        - Companion
                            - StaticResource
                            - CompanionClickThrough
                    - Linear
                        - Duration
                        - TrackingEvents
                            - Tracking
                        - VideoClicks
                            - ClickThrough
                            - ClickTracking
                            - CustomClick
                        - MediaFiles
                            - MediaFile


            ** (Optional)
            - Description
            - Advertiser
            - Survey
            - Error
            - Pricing
            - Extensions


            
