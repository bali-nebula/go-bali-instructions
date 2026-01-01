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

func DropClass() DropClassLike {
	return dropClass()
}

// Constructor Methods

func (c *dropClass_) Drop(
	component Component,
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
	case DraftComponent:
		source += "DRAFT"
	case DocumentComponent:
		source += "DOCUMENT"
	case MessageComponent:
		source += "MESSAGE"
	case VariableComponent:
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

func (v *drop_) GetComponent() Component {
	return v.component_
}

func (v *drop_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type drop_ struct {
	// Declare the instance attributes.
	component_ Component
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
