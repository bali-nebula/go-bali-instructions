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

func PullClass() PullClassLike {
	return pullClass()
}

// Constructor Methods

func (c *pullClass_) Pull(
	value Modifier,
) PullLike {
	if uti.IsUndefined(value) {
		panic("The \"value\" attribute is required by this class.")
	}
	var instance = &pull_{
		// Initialize the instance attributes.
		value_: value,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *pull_) GetClass() PullClassLike {
	return pullClass()
}

func (v *pull_) AsSource() string {
	var source = "PULL "
	switch v.value_ {
	case Component:
		source += "COMPONENT"
	case Result:
		source += "RESULT"
	case Exception:
		source += "EXCEPTION"
	case Handler:
		source += "HANDLER"
	default:
		var message = fmt.Sprintf(
			"An invalid value was found: %v",
			v.value_,
		)
		panic(message)
	}
	return source
}

// Attribute Methods

func (v *pull_) GetValue() Modifier {
	return v.value_
}

// PROTECTED INTERFACE

// Instance Structure

type pull_ struct {
	// Declare the instance attributes.
	value_ Modifier
}

// Class Structure

type pullClass_ struct {
	// Declare the class constants.
}

// Class Reference

func pullClass() *pullClass_ {
	return pullClassReference_
}

var pullClassReference_ = &pullClass_{
	// Initialize the class constants.
}
