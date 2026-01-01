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

func LoadClass() LoadClassLike {
	return loadClass()
}

// Constructor Methods

func (c *loadClass_) Load(
	component Component,
	symbol string,
) LoadLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &load_{
		// Initialize the instance attributes.
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

func (v *load_) AsSource() string {
	var source = "LOAD "
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

func (v *load_) GetComponent() Component {
	return v.component_
}

func (v *load_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type load_ struct {
	// Declare the instance attributes.
	component_ Component
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
