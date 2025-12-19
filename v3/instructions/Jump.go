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

func JumpClass() JumpClassLike {
	return jumpClass()
}

// Constructor Methods

func (c *jumpClass_) Jump(
	label string,
	condition Condition,
) JumpLike {
	if uti.IsUndefined(label) {
		panic("The \"label\" attribute is required by this class.")
	}
	var instance = &jump_{
		// Initialize the instance attributes.
		label_:     label,
		condition_: condition,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *jump_) GetClass() JumpClassLike {
	return jumpClass()
}

func (v *jump_) AsSource() string {
	var source = "JUMP TO " + v.label_
	switch v.condition_ {
	case OnAnyCondition:
	case OnNoneCondition:
		source += " ON NONE"
	case OnFalseCondition:
		source += " ON FALSE"
	case OnEmptyCondition:
		source += " ON EMPTY"
	default:
		var message = fmt.Sprintf(
			"An invalid condition was found: %v",
			v.condition_,
		)
		panic(message)
	}
	return source
}

// Attribute Methods

func (v *jump_) GetLabel() string {
	return v.label_
}

func (v *jump_) GetCondition() Condition {
	return v.condition_
}

// PROTECTED INTERFACE

// Instance Structure

type jump_ struct {
	// Declare the instance attributes.
	label_     string
	condition_ Condition
}

// Class Structure

type jumpClass_ struct {
	// Declare the class constants.
}

// Class Reference

func jumpClass() *jumpClass_ {
	return jumpClassReference_
}

var jumpClassReference_ = &jumpClass_{
	// Initialize the class constants.
}
