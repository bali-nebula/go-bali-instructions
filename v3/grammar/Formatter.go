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
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	ast "github.com/bali-nebula/go-bali-instructions/v3/ast"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func FormatterClass() FormatterClassLike {
	return formatterClass()
}

// Constructor Methods

func (c *formatterClass_) Formatter() FormatterLike {
	var instance = &formatter_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *formatter_) GetClass() FormatterClassLike {
	return formatterClass()
}

func (v *formatter_) FormatAssembly(assembly ast.AssemblyLike) string {
	v.visitor_.VisitAssembly(assembly)
	return v.getResult()
}

// Methodical Methods

func (v *formatter_) ProcessCount(
	count string,
) {
	v.appendString(count)
}

func (v *formatter_) ProcessDelimiter(
	delimiter string,
) {
	v.appendString(delimiter)
}

func (v *formatter_) ProcessExplanation(
	explanation string,
) {
	v.appendString(explanation)
}

func (v *formatter_) ProcessLabel(
	label string,
) {
	v.appendString(label)
}

func (v *formatter_) ProcessNewline(
	newline string,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessQuoted(
	quoted string,
) {
	v.appendString(quoted)
}

func (v *formatter_) ProcessSpace(
	space string,
) {
	v.appendString(space)
}

func (v *formatter_) ProcessSymbol(
	symbol string,
) {
	v.appendString(symbol)
}

func (v *formatter_) ProcessArgumentSlot(
	argument ast.ArgumentLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessCallSlot(
	call ast.CallLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessCardinality(
	cardinality ast.CardinalityLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessCardinalitySlot(
	cardinality ast.CardinalityLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessConditionally(
	conditionally ast.ConditionallyLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessConditionallySlot(
	conditionally ast.ConditionallyLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessConstantSlot(
	constant ast.ConstantLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessDropSlot(
	drop ast.DropLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessHandlerSlot(
	handler ast.HandlerLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PostprocessInstruction(
	instruction ast.InstructionLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessJumpSlot(
	jump ast.JumpLike,
	slot_ uint,
) {
	switch slot_ {
	case 1, 2:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessLiteralSlot(
	literal ast.LiteralLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessLoadSlot(
	load ast.LoadLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessNoteSlot(
	note ast.NoteLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessParameterized(
	parameterized ast.ParameterizedLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessParameterizedSlot(
	parameterized ast.ParameterizedLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessPrefix(
	prefix ast.PrefixLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessPrefix(
	prefix ast.PrefixLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessPullSlot(
	pull ast.PullLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessPushSlot(
	push ast.PushLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessSaveSlot(
	save ast.SaveLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessSendSlot(
	send ast.SendLike,
	slot_ uint,
) {
	switch slot_ {
	case 1, 2, 3:
		v.appendString(" ")
	}
}

const _indentation = "    "

// PROTECTED INTERFACE

// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var level uint
	for ; level < v.depth_; level++ {
		newline += _indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(
	string_ string,
) {
	v.result_.WriteString(string_)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}

// Instance Structure

type formatter_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type formatterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func formatterClass() *formatterClass_ {
	return formatterClassReference_
}

var formatterClassReference_ = &formatterClass_{
	// Initialize the class constants.
}
