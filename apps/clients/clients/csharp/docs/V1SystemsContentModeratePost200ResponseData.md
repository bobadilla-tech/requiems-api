# Org.OpenAPITools.Model.V1SystemsContentModeratePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSafe** | **bool** | False when any moderation category is triggered. True otherwise. | [optional] 
**ToxicityScore** | **string** | Reserved for future use. Currently always null. | [optional] 
**Sentiment** | **string** | Detected sentiment: positive, negative, or neutral | [optional] 
**Language** | **string** | Detected or provided ISO 639-1 language code. Null when detection confidence is 0. | [optional] 
**LanguageConfidence** | **string** | Confidence in the detected language from 0.0 to 1.0. 1.0 when language was explicitly provided. | [optional] 
**Flags** | **List&lt;Object&gt;** | Triggered moderation flags. Currently possible: text_profanity | [optional] 
**Categories** | **Object** | Per-category boolean flags: profanity, hate_speech, spam, violence | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

