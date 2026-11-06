# V1EntertainmentSudokuGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**difficulty** | **string** | The difficulty level of the returned puzzle (easy, medium, or hard) | [optional] [default to undefined]
**puzzle** | **string** | 9×9 grid representing the puzzle — 0 means an empty cell to be filled in | [optional] [default to undefined]
**solution** | **string** | 9×9 grid containing the complete, valid solution | [optional] [default to undefined]

## Example

```typescript
import { V1EntertainmentSudokuGet200ResponseData } from 'api-client';

const instance: V1EntertainmentSudokuGet200ResponseData = {
    difficulty,
    puzzle,
    solution,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
