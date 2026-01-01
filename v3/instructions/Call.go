/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	argumentCount uint8,
) CallLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &call_{
		// Initialize the instance attributes.
		symbol_:        symbol,
		argumentCount_: argumentCount,
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
	switch v.argumentCount_ {
	case 0:
	case 1:
		source += " WITH 1 ARGUMENT"
	case 2:
		source += " WITH 2 ARGUMENTS"
	case 3:
		source += " WITH 3 ARGUMENTS"
	default:
		var message = fmt.Sprintf(
			"An invalid argument count was found: %v",
			v.argumentCount_,
		)
		panic(message)
	}
	return source
}

// Attribute Methods

func (v *call_) GetSymbol() string {
	return v.symbol_
}

func (v *call_) GetArgumentCount() uint8 {
	return v.argumentCount_
}

// PROTECTED INTERFACE

// Instance Structure

type call_ struct {
	// Declare the instance attributes.
	symbol_        string
	argumentCount_ uint8
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
