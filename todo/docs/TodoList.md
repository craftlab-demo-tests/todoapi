# TodoList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Todos** | Pointer to [**[]Todo**](Todo.md) |  | [optional] 

## Methods

### NewTodoList

`func NewTodoList() *TodoList`

NewTodoList instantiates a new TodoList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTodoListWithDefaults

`func NewTodoListWithDefaults() *TodoList`

NewTodoListWithDefaults instantiates a new TodoList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TodoList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TodoList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TodoList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TodoList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTodos

`func (o *TodoList) GetTodos() []Todo`

GetTodos returns the Todos field if non-nil, zero value otherwise.

### GetTodosOk

`func (o *TodoList) GetTodosOk() (*[]Todo, bool)`

GetTodosOk returns a tuple with the Todos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodos

`func (o *TodoList) SetTodos(v []Todo)`

SetTodos sets Todos field to given value.

### HasTodos

`func (o *TodoList) HasTodos() bool`

HasTodos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


