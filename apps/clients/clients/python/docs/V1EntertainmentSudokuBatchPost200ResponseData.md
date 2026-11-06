# V1EntertainmentSudokuBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Generated puzzles in the same order as the input array. Each item has the same fields as the single-puzzle endpoint. | [optional] 
**results_difficulty** | **str** | The difficulty level of the puzzle (easy, medium, or hard) | [optional] 
**results_puzzle** | **str** | 9×9 grid representing the puzzle — 0 means an empty cell to be filled in | [optional] 
**results_solution** | **str** | 9×9 grid containing the complete, valid solution | [optional] 
**total** | **int** | Number of puzzles returned. Matches the length of the input array. | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_sudoku_batch_post200_response_data import V1EntertainmentSudokuBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentSudokuBatchPost200ResponseData from a JSON string
v1_entertainment_sudoku_batch_post200_response_data_instance = V1EntertainmentSudokuBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentSudokuBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_entertainment_sudoku_batch_post200_response_data_dict = v1_entertainment_sudoku_batch_post200_response_data_instance.to_dict()
# create an instance of V1EntertainmentSudokuBatchPost200ResponseData from a dict
v1_entertainment_sudoku_batch_post200_response_data_from_dict = V1EntertainmentSudokuBatchPost200ResponseData.from_dict(v1_entertainment_sudoku_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


