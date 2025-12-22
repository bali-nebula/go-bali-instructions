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
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func MethodClass() MethodClassLike {
	return methodClass()
}

// Constructor Methods

func (c *methodClass_) Method(
	instructions com.Sequential[InstructionLike],
) MethodLike {
	if uti.IsUndefined(instructions) {
		panic("The \"instructions\" attribute is required by this class.")
	}
	var instance = &method_{
		// Initialize the instance attributes.
		instructions_: instructions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *method_) GetClass() MethodClassLike {
	return methodClass()
}

func (v *method_) GetInstructions() com.Sequential[InstructionLike] {
	return v.instructions_
}

// PROTECTED INTERFACE

// Instance Structure

type method_ struct {
	// Declare the instance attributes.
	instructions_ com.Sequential[InstructionLike]
}

// Class Structure

type methodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func methodClass() *methodClass_ {
	return methodClassReference_
}

var methodClassReference_ = &methodClass_{
	// Initialize the class constants.
}
