# OpenapiClient::V1EntertainmentTriviaGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **question** | **String** | The trivia question text | [optional] |
| **options** | **String** | Four multiple-choice answer options | [optional] |
| **answer** | **String** | The correct answer — always one of the values in options | [optional] |
| **category** | **String** | The category the question belongs to | [optional] |
| **difficulty** | **String** | The difficulty level of the question (easy, medium, or hard) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1EntertainmentTriviaGet200ResponseData.new(
  question: null,
  options: null,
  answer: null,
  category: null,
  difficulty: null
)
```

