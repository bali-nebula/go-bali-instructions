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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func NoteClass() NoteClassLike {
	return noteClass()
}

// Constructor Methods

func (c *noteClass_) Note(
	description string,
) NoteLike {
	if uti.IsUndefined(description) {
		panic("The \"description\" attribute is required by this class.")
	}
	var instance = &note_{
		// Initialize the instance attributes.
		description_: description,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *note_) GetClass() NoteClassLike {
	return noteClass()
}

func (v *note_) AsSource() string {
	return "NOTE " + v.description_
}

// Attribute Methods

func (v *note_) GetDescription() string {
	return v.description_
}

// PROTECTED INTERFACE

// Instance Structure

type note_ struct {
	// Declare the instance attributes.
	description_ string
}

// Class Structure

type noteClass_ struct {
	// Declare the class constants.
}

// Class Reference

func noteClass() *noteClass_ {
	return noteClassReference_
}

var noteClassReference_ = &noteClass_{
	// Initialize the class constants.
}
