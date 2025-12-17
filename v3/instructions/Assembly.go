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

func AssemblyClass() AssemblyClassLike {
	return assemblyClass()
}

// Constructor Methods

func (c *assemblyClass_) Assembly(
	instructions com.Sequential[InstructionLike],
) AssemblyLike {
	if uti.IsUndefined(instructions) {
		panic("The \"instructions\" attribute is required by this class.")
	}
	var instance = &assembly_{
		// Initialize the instance attributes.
		instructions_: instructions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *assembly_) GetClass() AssemblyClassLike {
	return assemblyClass()
}

func (v *assembly_) GetInstructions() com.Sequential[InstructionLike] {
	return v.instructions_
}

// PROTECTED INTERFACE

// Instance Structure

type assembly_ struct {
	// Declare the instance attributes.
	instructions_ com.Sequential[InstructionLike]
}

// Class Structure

type assemblyClass_ struct {
	// Declare the class constants.
}

// Class Reference

func assemblyClass() *assemblyClass_ {
	return assemblyClassReference_
}

var assemblyClassReference_ = &assemblyClass_{
	// Initialize the class constants.
}
