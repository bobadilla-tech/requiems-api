# V1EntertainmentSudokuBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Generated puzzles in the same order as the input array. Each item has the same fields as the single-puzzle endpoint. | [optional] [default to undefined]
**results_difficulty** | **string** | The difficulty level of the puzzle (easy, medium, or hard) | [optional] [default to undefined]
**results_puzzle** | **string** | 9×9 grid representing the puzzle — 0 means an empty cell to be filled in | [optional] [default to undefined]
**results_solution** | **string** | 9×9 grid containing the complete, valid solution | [optional] [default to undefined]
**total** | **number** | Number of puzzles returned. Matches the length of the input array. | [optional] [default to undefined]

## Example

```typescript
import { V1EntertainmentSudokuBatchPost200ResponseData } from 'api-client';

const instance: V1EntertainmentSudokuBatchPost200ResponseData = {
    results,
    results_difficulty,
    results_puzzle,
    results_solution,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
