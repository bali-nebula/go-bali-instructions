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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AssemblyClass() AssemblyClassLike {
	return assemblyClass()
}

// Constructor Methods

func (c *assemblyClass_) Assembly(
	instructions fra.ListLike[InstructionLike],
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

// Attribute Methods

func (v *assembly_) GetInstructions() fra.ListLike[InstructionLike] {
	return v.instructions_
}

// PROTECTED INTERFACE

// Instance Structure

type assembly_ struct {
	// Declare the instance attributes.
	instructions_ fra.ListLike[InstructionLike]
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
