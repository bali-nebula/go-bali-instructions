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

func SendClass() SendClassLike {
	return sendClass()
}

// Constructor Methods

func (c *sendClass_) Send(
	delimiter1 string,
	symbol string,
	delimiter2 string,
	destination DestinationLike,
	optionalParameterized ParameterizedLike,
) SendLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(destination) {
		panic("The \"destination\" attribute is required by this class.")
	}
	var instance = &send_{
		// Initialize the instance attributes.
		delimiter1_:            delimiter1,
		symbol_:                symbol,
		delimiter2_:            delimiter2,
		destination_:           destination,
		optionalParameterized_: optionalParameterized,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *send_) GetClass() SendClassLike {
	return sendClass()
}

// Attribute Methods

func (v *send_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *send_) GetSymbol() string {
	return v.symbol_
}

func (v *send_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *send_) GetDestination() DestinationLike {
	return v.destination_
}

func (v *send_) GetOptionalParameterized() ParameterizedLike {
	return v.optionalParameterized_
}

// PROTECTED INTERFACE

// Instance Structure

type send_ struct {
	// Declare the instance attributes.
	delimiter1_            string
	symbol_                string
	delimiter2_            string
	destination_           DestinationLike
	optionalParameterized_ ParameterizedLike
}

// Class Structure

type sendClass_ struct {
	// Declare the class constants.
}

// Class Reference

func sendClass() *sendClass_ {
	return sendClassReference_
}

var sendClassReference_ = &sendClass_{
	// Initialize the class constants.
}
