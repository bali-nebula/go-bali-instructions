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

func LoadClass() LoadClassLike {
	return loadClass()
}

// Constructor Methods

func (c *loadClass_) Load(
	delimiter string,
	component ComponentLike,
	symbol string,
) LoadLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &load_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		component_: component,
		symbol_:    symbol,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *load_) GetClass() LoadClassLike {
	return loadClass()
}

// Attribute Methods

func (v *load_) GetDelimiter() string {
	return v.delimiter_
}

func (v *load_) GetComponent() ComponentLike {
	return v.component_
}

func (v *load_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type load_ struct {
	// Declare the instance attributes.
	delimiter_ string
	component_ ComponentLike
	symbol_    string
}

// Class Structure

type loadClass_ struct {
	// Declare the class constants.
}

// Class Reference

func loadClass() *loadClass_ {
	return loadClassReference_
}

var loadClassReference_ = &loadClass_{
	// Initialize the class constants.
}
