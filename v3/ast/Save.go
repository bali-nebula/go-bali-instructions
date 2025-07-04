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

func SaveClass() SaveClassLike {
	return saveClass()
}

// Constructor Methods

func (c *saveClass_) Save(
	delimiter string,
	component ComponentLike,
	symbol string,
) SaveLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &save_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		component_: component,
		symbol_:    symbol,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *save_) GetClass() SaveClassLike {
	return saveClass()
}

// Attribute Methods

func (v *save_) GetDelimiter() string {
	return v.delimiter_
}

func (v *save_) GetComponent() ComponentLike {
	return v.component_
}

func (v *save_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type save_ struct {
	// Declare the instance attributes.
	delimiter_ string
	component_ ComponentLike
	symbol_    string
}

// Class Structure

type saveClass_ struct {
	// Declare the class constants.
}

// Class Reference

func saveClass() *saveClass_ {
	return saveClassReference_
}

var saveClassReference_ = &saveClass_{
	// Initialize the class constants.
}
