# Org.OpenAPITools.Model.V1EntertainmentSudokuBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | Generated puzzles in the same order as the input array. Each item has the same fields as the single-puzzle endpoint. | [optional] 
**ResultsDifficulty** | **string** | The difficulty level of the puzzle (easy, medium, or hard) | [optional] 
**ResultsPuzzle** | **string** | 9×9 grid representing the puzzle — 0 means an empty cell to be filled in | [optional] 
**ResultsSolution** | **string** | 9×9 grid containing the complete, valid solution | [optional] 
**Total** | **int** | Number of puzzles returned. Matches the length of the input array. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

