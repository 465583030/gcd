// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Runtime types.
// API Version: 1.1
package main

 
// Unique object identifier.
type ChromeRuntimeRemoteObjectId string 
 
 
// Mirror object referencing original JavaScript object.
type ChromeRuntimeRemoteObject struct {
	Type string `json:"type"` // Object type.
	Subtype string `json:"subtype,omitempty"` // Object subtype hint. Specified for <code>object</code> type values only.
	ClassName string `json:"className,omitempty"` // Object class (constructor) name. Specified for <code>object</code> type values only.
	Value string `json:"value,omitempty"` // Remote object value in case of primitive values or JSON values (if it was requested), or description string if the value can not be JSON-stringified (like NaN, Infinity, -Infinity, -0).
	Description string `json:"description,omitempty"` // String representation of the object.
	ObjectId *ChromeRuntimeRemoteObjectId `json:"objectId,omitempty"` // Unique object identifier (for non-primitive values).
	Preview *ChromeRuntimeObjectPreview `json:"preview,omitempty"` // Preview containing abbreviated property values.
}
 
 
// Object containing abbreviated remote object value.
type ChromeRuntimeObjectPreview struct {
	Lossless bool `json:"lossless"` // Determines whether preview is lossless (contains all information of the original object).
	Overflow bool `json:"overflow"` // True iff some of the properties of the original did not fit.
	Properties []*ChromeRuntimePropertyPreview `json:"properties"` // List of the properties.
}
 
 
// 
type ChromeRuntimePropertyPreview struct {
	Name string `json:"name"` // Property name.
	Type string `json:"type"` // Object type. Accessor means that the property itself is an accessor property.
	Value string `json:"value,omitempty"` // User-friendly property value string.
	ValuePreview *ChromeRuntimeObjectPreview `json:"valuePreview,omitempty"` // Nested value preview.
	Subtype string `json:"subtype,omitempty"` // Object subtype hint. Specified for <code>object</code> type values only.
}
 
 
// Object property descriptor.
type ChromeRuntimePropertyDescriptor struct {
	Name string `json:"name"` // Property name or symbol description.
	Value *ChromeRuntimeRemoteObject `json:"value,omitempty"` // The value associated with the property.
	Writable bool `json:"writable,omitempty"` // True if the value associated with the property may be changed (data descriptors only).
	Get *ChromeRuntimeRemoteObject `json:"get,omitempty"` // A function which serves as a getter for the property, or <code>undefined</code> if there is no getter (accessor descriptors only).
	Set *ChromeRuntimeRemoteObject `json:"set,omitempty"` // A function which serves as a setter for the property, or <code>undefined</code> if there is no setter (accessor descriptors only).
	Configurable bool `json:"configurable"` // True if the type of this property descriptor may be changed and if the property may be deleted from the corresponding object.
	Enumerable bool `json:"enumerable"` // True if this property shows up during enumeration of the properties on the corresponding object.
	WasThrown bool `json:"wasThrown,omitempty"` // True if the result was thrown during the evaluation.
	IsOwn bool `json:"isOwn,omitempty"` // True if the property is owned for the object.
	Symbol *ChromeRuntimeRemoteObject `json:"symbol,omitempty"` // Property symbol object, if the property is of the <code>symbol</code> type.
}
 
 
// Object internal property descriptor. This property isn't normally visible in JavaScript code.
type ChromeRuntimeInternalPropertyDescriptor struct {
	Name string `json:"name"` // Conventional property name.
	Value *ChromeRuntimeRemoteObject `json:"value,omitempty"` // The value associated with the property.
}
 
 
// Represents function call argument. Either remote object id <code>objectId</code> or primitive <code>value</code> or neither of (for undefined) them should be specified.
type ChromeRuntimeCallArgument struct {
	Value string `json:"value,omitempty"` // Primitive value, or description string if the value can not be JSON-stringified (like NaN, Infinity, -Infinity, -0).
	ObjectId *ChromeRuntimeRemoteObjectId `json:"objectId,omitempty"` // Remote object handle.
	Type string `json:"type,omitempty"` // Object type.
}
 
 
// Id of an execution context.
type ChromeRuntimeExecutionContextId int 
 
 
// Description of an isolated world.
type ChromeRuntimeExecutionContextDescription struct {
	Id *ChromeRuntimeExecutionContextId `json:"id"` // Unique id of the execution context. It can be used to specify in which execution context script evaluation should be performed.
	IsPageContext bool `json:"isPageContext"` // True if this is a context where inpspected web page scripts run. False if it is a content script isolated context.
	Name string `json:"name"` // Human readable name describing given context.
	FrameId *ChromePageFrameId `json:"frameId"` // Id of the owning frame.
}
 
