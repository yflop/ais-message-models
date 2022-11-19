/*
Ais-Stream WebsocketObjects

A sample API to illustrate OpenAPI concepts

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aisStream

import (
	"encoding/json"
)

// AddessedSafetyMessage struct for AddessedSafetyMessage
type AddessedSafetyMessage struct {
	MessageID int32 `json:"MessageID"`
	RepeatIndicator int32 `json:"RepeatIndicator"`
	UserID int32 `json:"UserID"`
	Valid bool `json:"Valid"`
	Sequenceinteger int32 `json:"Sequenceinteger"`
	DestinationID int32 `json:"DestinationID"`
	Retransmission bool `json:"Retransmission"`
	Spare bool `json:"Spare"`
	Text string `json:"Text"`
}

// NewAddessedSafetyMessage instantiates a new AddessedSafetyMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddessedSafetyMessage(messageID int32, repeatIndicator int32, userID int32, valid bool, sequenceinteger int32, destinationID int32, retransmission bool, spare bool, text string) *AddessedSafetyMessage {
	this := AddessedSafetyMessage{}
	this.MessageID = messageID
	this.RepeatIndicator = repeatIndicator
	this.UserID = userID
	this.Valid = valid
	this.Sequenceinteger = sequenceinteger
	this.DestinationID = destinationID
	this.Retransmission = retransmission
	this.Spare = spare
	this.Text = text
	return &this
}

// NewAddessedSafetyMessageWithDefaults instantiates a new AddessedSafetyMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddessedSafetyMessageWithDefaults() *AddessedSafetyMessage {
	this := AddessedSafetyMessage{}
	return &this
}

// GetMessageID returns the MessageID field value
func (o *AddessedSafetyMessage) GetMessageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value
// and a boolean to check if the value has been set.
func (o *AddessedSafetyMessage) GetMessageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageID, true
}

// SetMessageID sets field value
func (o *AddessedSafetyMessage) SetMessageID(v int32) {
	o.MessageID = v
}

// GetRepeatIndicator returns the RepeatIndicator field value
func (o *AddessedSafetyMessage) GetRepeatIndicator() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RepeatIndicator
}

// GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field value
// and a boolean to check if the value has been set.
func (o *AddessedSafetyMessage) GetRepeatIndicatorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepeatIndicator, true
}

// SetRepeatIndicator sets field value
func (o *AddessedSafetyMessage) SetRepeatIndicator(v int32) {
	o.RepeatIndicator = v
}

// GetUserID returns the UserID field value
func (o *AddessedSafetyMessage) GetUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value
// and a boolean to check if the value has been set.
func (o *AddessedSafetyMessage) GetUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserID, true
}

// SetUserID sets field value
func (o *AddessedSafetyMessage) SetUserID(v int32) {
	o.UserID = v
}

// GetValid returns the Valid field value
func (o *AddessedSafetyMessage) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *AddessedSafetyMessage) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *AddessedSafetyMessage) SetValid(v bool) {
	o.Valid = v
}

// GetSequenceinteger returns the Sequenceinteger field value
func (o *AddessedSafetyMessage) GetSequenceinteger() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sequenceinteger
}

// GetSequenceintegerOk returns a tuple with the Sequenceinteger field value
// and a boolean to check if the value has been set.
func (o *AddessedSafetyMessage) GetSequenceintegerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequenceinteger, true
}

// SetSequenceinteger sets field value
func (o *AddessedSafetyMessage) SetSequenceinteger(v int32) {
	o.Sequenceinteger = v
}

// GetDestinationID returns the DestinationID field value
func (o *AddessedSafetyMessage) GetDestinationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DestinationID
}

// GetDestinationIDOk returns a tuple with the DestinationID field value
// and a boolean to check if the value has been set.
func (o *AddessedSafetyMessage) GetDestinationIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationID, true
}

// SetDestinationID sets field value
func (o *AddessedSafetyMessage) SetDestinationID(v int32) {
	o.DestinationID = v
}

// GetRetransmission returns the Retransmission field value
func (o *AddessedSafetyMessage) GetRetransmission() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Retransmission
}

// GetRetransmissionOk returns a tuple with the Retransmission field value
// and a boolean to check if the value has been set.
func (o *AddessedSafetyMessage) GetRetransmissionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retransmission, true
}

// SetRetransmission sets field value
func (o *AddessedSafetyMessage) SetRetransmission(v bool) {
	o.Retransmission = v
}

// GetSpare returns the Spare field value
func (o *AddessedSafetyMessage) GetSpare() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Spare
}

// GetSpareOk returns a tuple with the Spare field value
// and a boolean to check if the value has been set.
func (o *AddessedSafetyMessage) GetSpareOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spare, true
}

// SetSpare sets field value
func (o *AddessedSafetyMessage) SetSpare(v bool) {
	o.Spare = v
}

// GetText returns the Text field value
func (o *AddessedSafetyMessage) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *AddessedSafetyMessage) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *AddessedSafetyMessage) SetText(v string) {
	o.Text = v
}

func (o AddessedSafetyMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MessageID"] = o.MessageID
	}
	if true {
		toSerialize["RepeatIndicator"] = o.RepeatIndicator
	}
	if true {
		toSerialize["UserID"] = o.UserID
	}
	if true {
		toSerialize["Valid"] = o.Valid
	}
	if true {
		toSerialize["Sequenceinteger"] = o.Sequenceinteger
	}
	if true {
		toSerialize["DestinationID"] = o.DestinationID
	}
	if true {
		toSerialize["Retransmission"] = o.Retransmission
	}
	if true {
		toSerialize["Spare"] = o.Spare
	}
	if true {
		toSerialize["Text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableAddessedSafetyMessage struct {
	value *AddessedSafetyMessage
	isSet bool
}

func (v NullableAddessedSafetyMessage) Get() *AddessedSafetyMessage {
	return v.value
}

func (v *NullableAddessedSafetyMessage) Set(val *AddessedSafetyMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAddessedSafetyMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAddessedSafetyMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddessedSafetyMessage(val *AddessedSafetyMessage) *NullableAddessedSafetyMessage {
	return &NullableAddessedSafetyMessage{value: val, isSet: true}
}

func (v NullableAddessedSafetyMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddessedSafetyMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

