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

func (v *formatter_) PreprocessAction(
	action ast.ActionLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessActionSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessAction(
	action ast.ActionLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessArgument(
	argument ast.ArgumentLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessArgumentSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessArgument(
	argument ast.ArgumentLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessAssembly(
	assembly ast.AssemblyLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessAssemblySlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessAssembly(
	assembly ast.AssemblyLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessCall(
	call ast.CallLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessCallSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessCall(
	call ast.CallLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessComponent(
	component ast.ComponentLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessComponentSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessComponent(
	component ast.ComponentLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessCondition(
	condition ast.ConditionLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessConditionSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessCondition(
	condition ast.ConditionLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessConditionally(
	conditionally ast.ConditionallyLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessConditionallySlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessConditionally(
	conditionally ast.ConditionallyLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessConstant(
	constant ast.ConstantLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessConstantSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessConstant(
	constant ast.ConstantLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessContext(
	context ast.ContextLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessContextSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessContext(
	context ast.ContextLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessDestination(
	destination ast.DestinationLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessDestinationSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessDestination(
	destination ast.DestinationLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessDrop(
	drop ast.DropLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessDropSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessDrop(
	drop ast.DropLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessHandler(
	handler ast.HandlerLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessHandlerSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessHandler(
	handler ast.HandlerLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessInstruction(
	instruction ast.InstructionLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessInstructionSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessInstruction(
	instruction ast.InstructionLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessItem(
	item ast.ItemLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessItemSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessItem(
	item ast.ItemLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessJump(
	jump ast.JumpLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessJumpSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessJump(
	jump ast.JumpLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessLiteral(
	literal ast.LiteralLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessLiteralSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessLiteral(
	literal ast.LiteralLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessLoad(
	load ast.LoadLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessLoadSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessLoad(
	load ast.LoadLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessNoop(
	noop ast.NoopLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessNoopSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessNoop(
	noop ast.NoopLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessNote(
	note ast.NoteLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessNoteSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessNote(
	note ast.NoteLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessPrefix(
	prefix ast.PrefixLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessPrefixSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessPrefix(
	prefix ast.PrefixLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessPull(
	pull ast.PullLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessPullSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessPull(
	pull ast.PullLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessPush(
	push ast.PushLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessPushSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessPush(
	push ast.PushLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessSave(
	save ast.SaveLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessSaveSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessSave(
	save ast.SaveLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessSend(
	send ast.SendLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessSendSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessSend(
	send ast.SendLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessSource(
	source ast.SourceLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessSourceSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessSource(
	source ast.SourceLike,
	index uint,
	count uint,
) {
	// TBD - Add formatting of the rule.
}

const _indentation = "\t"

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
