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
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func NoteClass() NoteClassLike {
	return noteClass()
}

// Constructor Methods

func (c *noteClass_) Note(
	delimiter string,
	explanation string,
) NoteLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(explanation) {
		panic("The \"explanation\" attribute is required by this class.")
	}
	var instance = &note_{
		// Initialize the instance attributes.
		delimiter_:   delimiter,
		explanation_: explanation,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *note_) GetClass() NoteClassLike {
	return noteClass()
}

// Attribute Methods

func (v *note_) GetDelimiter() string {
	return v.delimiter_
}

func (v *note_) GetExplanation() string {
	return v.explanation_
}

// PROTECTED INTERFACE

// Instance Structure

type note_ struct {
	// Declare the instance attributes.
	delimiter_   string
	explanation_ string
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
