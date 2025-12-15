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

func SendClass() SendClassLike {
	return sendClass()
}

// Constructor Methods

func (c *sendClass_) Send(
	symbol string,
	destination Modifier,
) SendLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	if uti.IsUndefined(destination) {
		panic("The \"destination\" attribute is required by this class.")
	}
	var instance = &send_{
		// Initialize the instance attributes.
		symbol_:      symbol,
		destination_: destination,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *send_) GetClass() SendClassLike {
	return sendClass()
}

func (v *send_) AsSource() string {
	var source = "SEND " + v.symbol_ + " TO "
	switch v.destination_ {
	case ComponentModifier:
		source += "COMPONENT"
	case ComponentWithArgumentsModifier:
		source += "COMPONENT WITH ARGUMENTS"
	case DocumentModifier:
		source += "DOCUMENT"
	case DocumentWithArgumentsModifier:
		source += "DOCUMENT WITH ARGUMENTS"
	default:
		var message = fmt.Sprintf(
			"An invalid destination was found: %v",
			v.destination_,
		)
		panic(message)
	}
	return source
}

// Attribute Methods

func (v *send_) GetSymbol() string {
	return v.symbol_
}

func (v *send_) GetDestination() Modifier {
	return v.destination_
}

// PROTECTED INTERFACE

// Instance Structure

type send_ struct {
	// Declare the instance attributes.
	symbol_      string
	destination_ Modifier
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
