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

func CallClass() CallClassLike {
	return callClass()
}

// Constructor Methods

func (c *callClass_) Call(
	delimiter string,
	symbol string,
	optionalCardinality CardinalityLike,
) CallLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &call_{
		// Initialize the instance attributes.
		delimiter_:           delimiter,
		symbol_:              symbol,
		optionalCardinality_: optionalCardinality,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *call_) GetClass() CallClassLike {
	return callClass()
}

// Attribute Methods

func (v *call_) GetDelimiter() string {
	return v.delimiter_
}

func (v *call_) GetSymbol() string {
	return v.symbol_
}

func (v *call_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// PROTECTED INTERFACE

// Instance Structure

type call_ struct {
	// Declare the instance attributes.
	delimiter_           string
	symbol_              string
	optionalCardinality_ CardinalityLike
}

// Class Structure

type callClass_ struct {
	// Declare the class constants.
}

// Class Reference

func callClass() *callClass_ {
	return callClassReference_
}

var callClassReference_ = &callClass_{
	// Initialize the class constants.
}
