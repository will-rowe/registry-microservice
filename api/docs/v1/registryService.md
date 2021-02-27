# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/proto/v1/registryService.proto](#api/proto/v1/registryService.proto)
    - [CreateRequest](#v1.CreateRequest)
    - [CreateResponse](#v1.CreateResponse)
    - [DeleteRequest](#v1.DeleteRequest)
    - [DeleteResponse](#v1.DeleteResponse)
    - [Participant](#v1.Participant)
    - [RetrieveRequest](#v1.RetrieveRequest)
    - [RetrieveResponse](#v1.RetrieveResponse)
    - [UpdateRequest](#v1.UpdateRequest)
    - [UpdateResponse](#v1.UpdateResponse)
  
    - [RegistryService](#v1.RegistryService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api/proto/v1/registryService.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/proto/v1/registryService.proto



<a name="v1.CreateRequest"></a>

### CreateRequest
CreateRequest will request a participant is created
in the registry.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_version | [string](#string) |  | api version |
| participant | [Participant](#v1.Participant) |  | participant to create |






<a name="v1.CreateResponse"></a>

### CreateResponse
CreateResponse contains the status of
the create operation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_version | [string](#string) |  | api version |
| created | [bool](#bool) |  | created is true if participant was created |






<a name="v1.DeleteRequest"></a>

### DeleteRequest
DeleteRequest will request a participant to
be deleted in the registry.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_version | [string](#string) |  | api version |
| id | [string](#string) |  | unique string reference number for the requested participant |






<a name="v1.DeleteResponse"></a>

### DeleteResponse
DeleteResponse contains the status of
the delete operation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_version | [string](#string) |  | api version |
| deleted | [bool](#bool) |  | deleted is true if participant was deleted |






<a name="v1.Participant"></a>

### Participant
Participant describes a study participant
that needs to be recorded in the registry.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique string reference number for the participant |
| dob | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | date of birth |
| phone | [string](#string) |  | phone number |
| address | [string](#string) |  | address |






<a name="v1.RetrieveRequest"></a>

### RetrieveRequest
RetrieveRequest will request a participant
from the registry using the provided id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_version | [string](#string) |  | api version |
| id | [string](#string) |  | unique string reference number for the requested participant |






<a name="v1.RetrieveResponse"></a>

### RetrieveResponse
RetrieveResponse contains the participant data
held in the registry.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_version | [string](#string) |  | api version |
| participant | [Participant](#v1.Participant) |  | participant to return |






<a name="v1.UpdateRequest"></a>

### UpdateRequest
UpdateRequest will request a participant to
be updated in the registry.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_version | [string](#string) |  | api version |
| participant | [Participant](#v1.Participant) |  | participant to update |






<a name="v1.UpdateResponse"></a>

### UpdateResponse
UpdateResponse contains the status of 
the update operation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_version | [string](#string) |  | api version |
| updated | [bool](#bool) |  | updated is true if participant was updated |





 

 

 


<a name="v1.RegistryService"></a>

### RegistryService
RegistryService manages CRUD operations for study participants.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#v1.CreateRequest) | [CreateResponse](#v1.CreateResponse) | Create a new participant |
| Retrieve | [RetrieveRequest](#v1.RetrieveRequest) | [RetrieveResponse](#v1.RetrieveResponse) | Retrieve participant from registry |
| Update | [UpdateRequest](#v1.UpdateRequest) | [UpdateResponse](#v1.UpdateResponse) | Update participant details |
| Delete | [DeleteRequest](#v1.DeleteRequest) | [DeleteResponse](#v1.DeleteResponse) | Delete participant from registry |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

