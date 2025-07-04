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

func ArgumentClass() ArgumentClassLike {
	return argumentClass()
}

// Constructor Methods

func (c *argumentClass_) Argument(
	delimiter string,
	symbol string,
) ArgumentLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &argument_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		symbol_:    symbol,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *argument_) GetClass() ArgumentClassLike {
	return argumentClass()
}

// Attribute Methods

func (v *argument_) GetDelimiter() string {
	return v.delimiter_
}

func (v *argument_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type argument_ struct {
	// Declare the instance attributes.
	delimiter_ string
	symbol_    string
}

// Class Structure

type argumentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func argumentClass() *argumentClass_ {
	return argumentClassReference_
}

var argumentClassReference_ = &argumentClass_{
	// Initialize the class constants.
}
