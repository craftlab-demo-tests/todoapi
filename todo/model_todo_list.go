/*
 * todo-service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package todoclient

import (
	"encoding/json"
)

// TodoList struct for TodoList
type TodoList struct {
	Count *int32 `json:"count,omitempty"`
	Todos *[]Todo `json:"todos,omitempty"`
}

// NewTodoList instantiates a new TodoList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTodoList() *TodoList {
	this := TodoList{}
	return &this
}

// NewTodoListWithDefaults instantiates a new TodoList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTodoListWithDefaults() *TodoList {
	this := TodoList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TodoList) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TodoList) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TodoList) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *TodoList) SetCount(v int32) {
	o.Count = &v
}

// GetTodos returns the Todos field value if set, zero value otherwise.
func (o *TodoList) GetTodos() []Todo {
	if o == nil || o.Todos == nil {
		var ret []Todo
		return ret
	}
	return *o.Todos
}

// GetTodosOk returns a tuple with the Todos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TodoList) GetTodosOk() (*[]Todo, bool) {
	if o == nil || o.Todos == nil {
		return nil, false
	}
	return o.Todos, true
}

// HasTodos returns a boolean if a field has been set.
func (o *TodoList) HasTodos() bool {
	if o != nil && o.Todos != nil {
		return true
	}

	return false
}

// SetTodos gets a reference to the given []Todo and assigns it to the Todos field.
func (o *TodoList) SetTodos(v []Todo) {
	o.Todos = &v
}

func (o TodoList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Todos != nil {
		toSerialize["todos"] = o.Todos
	}
	return json.Marshal(toSerialize)
}

type NullableTodoList struct {
	value *TodoList
	isSet bool
}

func (v NullableTodoList) Get() *TodoList {
	return v.value
}

func (v *NullableTodoList) Set(val *TodoList) {
	v.value = val
	v.isSet = true
}

func (v NullableTodoList) IsSet() bool {
	return v.isSet
}

func (v *NullableTodoList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTodoList(val *TodoList) *NullableTodoList {
	return &NullableTodoList{value: val, isSet: true}
}

func (v NullableTodoList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTodoList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


