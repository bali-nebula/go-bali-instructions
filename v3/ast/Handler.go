/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func HandlerClass() HandlerClassLike {
	return handlerClass()
}

// Constructor Methods

func (c *handlerClass_) Handler(
	delimiter string,
	label string,
) HandlerLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(label) {
		panic("The \"label\" attribute is required by this class.")
	}
	var instance = &handler_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		label_:     label,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *handler_) GetClass() HandlerClassLike {
	return handlerClass()
}

// Attribute Methods

func (v *handler_) GetDelimiter() string {
	return v.delimiter_
}

func (v *handler_) GetLabel() string {
	return v.label_
}

// PROTECTED INTERFACE

// Instance Structure

type handler_ struct {
	// Declare the instance attributes.
	delimiter_ string
	label_     string
}

// Class Structure

type handlerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func handlerClass() *handlerClass_ {
	return handlerClassReference_
}

var handlerClassReference_ = &handlerClass_{
	// Initialize the class constants.
}
