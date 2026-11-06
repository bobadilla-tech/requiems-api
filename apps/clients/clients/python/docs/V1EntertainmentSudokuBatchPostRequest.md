# V1EntertainmentSudokuBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**puzzles** | **List[object]** | Array of difficulty levels to generate (min: 1, max: 20). Each must be one of: easy, medium, hard. | 

## Example

```python
from openapi_client.models.v1_entertainment_sudoku_batch_post_request import V1EntertainmentSudokuBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentSudokuBatchPostRequest from a JSON string
v1_entertainment_sudoku_batch_post_request_instance = V1EntertainmentSudokuBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentSudokuBatchPostRequest.to_json())

# convert the object into a dict
v1_entertainment_sudoku_batch_post_request_dict = v1_entertainment_sudoku_batch_post_request_instance.to_dict()
# create an instance of V1EntertainmentSudokuBatchPostRequest from a dict
v1_entertainment_sudoku_batch_post_request_from_dict = V1EntertainmentSudokuBatchPostRequest.from_dict(v1_entertainment_sudoku_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


