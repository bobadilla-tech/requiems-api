# OpenapiClient::V1EntertainmentTriviaBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | List of trivia questions returned in the same order as the input filters. | [optional] |
| **results_category** | **String** | The topic filter used to retrieve the trivia question. | [optional] |
| **results_difficulty** | **String** | The complexity level filter used to retrieve the trivia question. | [optional] |
| **results_error** | **String** | Describes why the question could not be retrieved for this filter. Omitted if the query was successful. | [optional] |
| **results_data** | **Object** | The trivia question returned for this filter. Empty if the query failed. | [optional] |
| **results_data_question** | **String** | The trivia question text. | [optional] |
| **results_data_options** | **Array&lt;Object&gt;** | List of possible answers for the question. | [optional] |
| **results_data_answer** | **String** | The correct answer to the trivia question. | [optional] |
| **results_data_category** | **String** | The category the trivia question belongs to. | [optional] |
| **results_data_difficulty** | **String** | The difficulty level of the trivia question. | [optional] |
| **total** | **Integer** | Total number of trivia results returned. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1EntertainmentTriviaBatchPost200ResponseData.new(
  results: null,
  results_category: null,
  results_difficulty: null,
  results_error: null,
  results_data: null,
  results_data_question: null,
  results_data_options: null,
  results_data_answer: null,
  results_data_category: null,
  results_data_difficulty: null,
  total: null
)
```

