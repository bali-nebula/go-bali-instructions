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

func DropClass() DropClassLike {
	return dropClass()
}

// Constructor Methods

func (c *dropClass_) Drop(
	component Modifier,
	symbol string,
) DropLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &drop_{
		// Initialize the instance attributes.
		component_: component,
		symbol_:    symbol,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *drop_) GetClass() DropClassLike {
	return dropClass()
}

func (v *drop_) AsSource() string {
	var source = "DROP "
	switch v.component_ {
	case Draft:
		source += "DRAFT"
	case Document:
		source += "DOCUMENT"
	case Message:
		source += "MESSAGE"
	case Variable:
		source += "VARIABLE"
	default:
		var message = fmt.Sprintf(
			"An invalid component was found: %v",
			v.component_,
		)
		panic(message)
	}
	return source + " " + v.symbol_
}

// Attribute Methods

func (v *drop_) GetComponent() Modifier {
	return v.component_
}

func (v *drop_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type drop_ struct {
	// Declare the instance attributes.
	component_ Modifier
	symbol_    string
}

// Class Structure

type dropClass_ struct {
	// Declare the class constants.
}

// Class Reference

func dropClass() *dropClass_ {
	return dropClassReference_
}

var dropClassReference_ = &dropClass_{
	// Initialize the class constants.
}
