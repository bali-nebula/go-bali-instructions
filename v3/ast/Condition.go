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

func ConditionClass() ConditionClassLike {
	return conditionClass()
}

// Constructor Methods

func (c *conditionClass_) Condition(
	any_ any,
) ConditionLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &condition_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *condition_) GetClass() ConditionClassLike {
	return conditionClass()
}

// Attribute Methods

func (v *condition_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type condition_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type conditionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func conditionClass() *conditionClass_ {
	return conditionClassReference_
}

var conditionClassReference_ = &conditionClass_{
	// Initialize the class constants.
}
