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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func InstructionClass() InstructionClassLike {
	return instructionClass()
}

// Constructor Methods

func (c *instructionClass_) Instruction(
	optionalPrefix PrefixLike,
	action ActionLike,
) InstructionLike {
	if uti.IsUndefined(action) {
		panic("The \"action\" attribute is required by this class.")
	}
	var instance = &instruction_{
		// Initialize the instance attributes.
		optionalPrefix_: optionalPrefix,
		action_:         action,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *instruction_) GetClass() InstructionClassLike {
	return instructionClass()
}

// Attribute Methods

func (v *instruction_) GetOptionalPrefix() PrefixLike {
	return v.optionalPrefix_
}

func (v *instruction_) GetAction() ActionLike {
	return v.action_
}

// PROTECTED INTERFACE

// Instance Structure

type instruction_ struct {
	// Declare the instance attributes.
	optionalPrefix_ PrefixLike
	action_         ActionLike
}

// Class Structure

type instructionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func instructionClass() *instructionClass_ {
	return instructionClassReference_
}

var instructionClassReference_ = &instructionClass_{
	// Initialize the class constants.
}
