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

func ConstantClass() ConstantClassLike {
	return constantClass()
}

// Constructor Methods

func (c *constantClass_) Constant(
	delimiter string,
	symbol string,
) ConstantLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &constant_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		symbol_:    symbol,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *constant_) GetClass() ConstantClassLike {
	return constantClass()
}

// Attribute Methods

func (v *constant_) GetDelimiter() string {
	return v.delimiter_
}

func (v *constant_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type constant_ struct {
	// Declare the instance attributes.
	delimiter_ string
	symbol_    string
}

// Class Structure

type constantClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constantClass() *constantClass_ {
	return constantClassReference_
}

var constantClassReference_ = &constantClass_{
	// Initialize the class constants.
}
