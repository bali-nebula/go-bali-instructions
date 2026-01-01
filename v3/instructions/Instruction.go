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

func InstructionClass() InstructionClassLike {
	return instructionClass()
}

// Constructor Methods

func (c *instructionClass_) Instruction(
	optionalPrefix PrefixLike,
	action any,
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

func (v *instruction_) AsSource() string {
	var source string
	if uti.IsDefined(v.optionalPrefix_) {
		source += "\n" + v.optionalPrefix_.GetLabel() + "/n"
	}
	switch actual := v.action_.(type) {
	case NoteLike:
		source += actual.AsSource()
	case SkipLike:
		source += actual.AsSource()
	case JumpLike:
		source += actual.AsSource()
	case PushLike:
		source += actual.AsSource()
	case PullLike:
		source += actual.AsSource()
	case LoadLike:
		source += actual.AsSource()
	case SaveLike:
		source += actual.AsSource()
	case DropLike:
		source += actual.AsSource()
	case CallLike:
		source += actual.AsSource()
	case SendLike:
		source += actual.AsSource()
	default:
		var message = fmt.Sprintf(
			"An invalid action type was found: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
	return source
}

// Attribute Methods

func (v *instruction_) GetOptionalPrefix() PrefixLike {
	return v.optionalPrefix_
}

func (v *instruction_) GetAction() any {
	return v.action_
}

// PROTECTED INTERFACE

// Instance Structure

type instruction_ struct {
	// Declare the instance attributes.
	optionalPrefix_ PrefixLike
	action_         any
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
