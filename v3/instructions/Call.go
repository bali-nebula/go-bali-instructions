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

package instructions

import (
	fmt "fmt"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func CallClass() CallClassLike {
	return callClass()
}

// Constructor Methods

func (c *callClass_) Call(
	symbol string,
	cardinality Modifier,
) CallLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &call_{
		// Initialize the instance attributes.
		symbol_:      symbol,
		cardinality_: cardinality,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *call_) GetClass() CallClassLike {
	return callClass()
}

func (v *call_) AsSource() string {
	var source = "CALL " + v.symbol_
	switch v.cardinality_ {
	case With0ArgumentsModifier:
	case With1ArgumentModifier:
		source += " WITH 1 ARGUMENT"
	case With2ArgumentsModifier:
		source += " WITH 2 ARGUMENTS"
	case With3ArgumentsModifier:
		source += " WITH 3 ARGUMENTS"
	default:
		var message = fmt.Sprintf(
			"An invalid cardinality was found: %v",
			v.cardinality_,
		)
		panic(message)
	}
	return source
}

// Attribute Methods

func (v *call_) GetSymbol() string {
	return v.symbol_
}

func (v *call_) GetCardinality() Modifier {
	return v.cardinality_
}

// PROTECTED INTERFACE

// Instance Structure

type call_ struct {
	// Declare the instance attributes.
	symbol_      string
	cardinality_ Modifier
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
