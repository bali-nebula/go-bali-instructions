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

package agents

import (
	fmt "fmt"
	ins "github.com/bali-nebula/go-bali-instructions/v3/instructions"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClass()
}

// Constructor Methods

func (c *visitorClass_) Visitor(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		processor_: processor,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClass()
}

func (v *visitor_) VisitMethod(
	method ins.MethodLike,
) {
	v.processor_.PreprocessMethod(
		method,
		0,
		0,
	)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(
		method,
		0,
		0,
	)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitArgument(
	argument ins.ArgumentLike,
) {
	var symbol = argument.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitCall(
	call ins.CallLike,
) {
	var symbol = call.GetSymbol()
	v.processor_.ProcessSymbol(symbol)

	// Visit slot 1 between terms.
	v.processor_.ProcessCallSlot(
		call,
		1,
	)

	var argumentCount = call.GetArgumentCount()
	v.processor_.ProcessArgumentCount(argumentCount)
}

func (v *visitor_) visitConstant(
	constant ins.ConstantLike,
) {
	var symbol = constant.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitDrop(
	drop ins.DropLike,
) {
	var component = drop.GetComponent()
	v.processor_.ProcessComponent(component)

	// Visit slot 1 between terms.
	v.processor_.ProcessDropSlot(
		drop,
		1,
	)

	var symbol = drop.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitHandler(
	handler ins.HandlerLike,
) {
	var label = handler.GetLabel()
	v.processor_.ProcessLabel(label)
}

func (v *visitor_) visitInstruction(
	instruction ins.InstructionLike,
) {
	var optionalPrefix = instruction.GetOptionalPrefix()
	if uti.IsDefined(optionalPrefix) {
		v.processor_.PreprocessPrefix(
			optionalPrefix,
			0,
			0,
		)
		v.visitPrefix(optionalPrefix)
		v.processor_.PostprocessPrefix(
			optionalPrefix,
			0,
			0,
		)
	}

	// Visit slot 1 between terms.
	v.processor_.ProcessInstructionSlot(
		instruction,
		1,
	)

	var action = instruction.GetAction()
	switch actual := action.(type) {
	case ins.NoteLike:
		v.processor_.PreprocessNote(
			actual,
			0,
			0,
		)
		v.visitNote(actual)
		v.processor_.PostprocessNote(
			actual,
			0,
			0,
		)
	case ins.SkipLike:
		v.processor_.PreprocessSkip(
			actual,
			0,
			0,
		)
		v.visitSkip(actual)
		v.processor_.PostprocessSkip(
			actual,
			0,
			0,
		)
	case ins.JumpLike:
		v.processor_.PreprocessJump(
			actual,
			0,
			0,
		)
		v.visitJump(actual)
		v.processor_.PostprocessJump(
			actual,
			0,
			0,
		)
	case ins.PushLike:
		v.processor_.PreprocessPush(
			actual,
			0,
			0,
		)
		v.visitPush(actual)
		v.processor_.PostprocessPush(
			actual,
			0,
			0,
		)
	case ins.PullLike:
		v.processor_.PreprocessPull(
			actual,
			0,
			0,
		)
		v.visitPull(actual)
		v.processor_.PostprocessPull(
			actual,
			0,
			0,
		)
	case ins.LoadLike:
		v.processor_.PreprocessLoad(
			actual,
			0,
			0,
		)
		v.visitLoad(actual)
		v.processor_.PostprocessLoad(
			actual,
			0,
			0,
		)
	case ins.SaveLike:
		v.processor_.PreprocessSave(
			actual,
			0,
			0,
		)
		v.visitSave(actual)
		v.processor_.PostprocessSave(
			actual,
			0,
			0,
		)
	case ins.DropLike:
		v.processor_.PreprocessDrop(
			actual,
			0,
			0,
		)
		v.visitDrop(actual)
		v.processor_.PostprocessDrop(
			actual,
			0,
			0,
		)
	case ins.CallLike:
		v.processor_.PreprocessCall(
			actual,
			0,
			0,
		)
		v.visitCall(actual)
		v.processor_.PostprocessCall(
			actual,
			0,
			0,
		)
	case ins.SendLike:
		v.processor_.PreprocessSend(
			actual,
			0,
			0,
		)
		v.visitSend(actual)
		v.processor_.PostprocessSend(
			actual,
			0,
			0,
		)
	default:
		var message = fmt.Sprintf(
			"An invalid action was found: %v",
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitJump(
	jump ins.JumpLike,
) {
	var label = jump.GetLabel()
	v.processor_.ProcessLabel(label)

	// Visit slot 1 between terms.
	v.processor_.ProcessJumpSlot(
		jump,
		1,
	)

	var condition = jump.GetCondition()
	v.processor_.ProcessCondition(condition)
}

func (v *visitor_) visitLiteral(
	literal ins.LiteralLike,
) {
	var quoted = literal.GetQuoted()
	v.processor_.ProcessQuoted(quoted)
}

func (v *visitor_) visitLoad(
	load ins.LoadLike,
) {
	var component = load.GetComponent()
	v.processor_.ProcessComponent(component)

	// Visit slot 1 between terms.
	v.processor_.ProcessLoadSlot(
		load,
		1,
	)

	var symbol = load.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitMethod(
	method ins.MethodLike,
) {
	var instructionsIndex uint
	var instructions = method.GetInstructions().GetIterator()
	var instructionsCount = uint(instructions.GetSize())
	for instructions.HasNext() {
		instructionsIndex++
		var rule = instructions.GetNext()
		v.processor_.PreprocessInstruction(
			rule,
			instructionsIndex,
			instructionsCount,
		)
		v.visitInstruction(rule)
		v.processor_.PostprocessInstruction(
			rule,
			instructionsIndex,
			instructionsCount,
		)
	}
}

func (v *visitor_) visitNote(
	note ins.NoteLike,
) {
	var description = note.GetDescription()
	v.processor_.ProcessDescription(description)
}

func (v *visitor_) visitPrefix(
	prefix ins.PrefixLike,
) {
	var label = prefix.GetLabel()
	v.processor_.ProcessLabel(label)
}

func (v *visitor_) visitPull(
	pull ins.PullLike,
) {
	var value = pull.GetValue()
	v.processor_.ProcessValue(value)
}

func (v *visitor_) visitPush(
	push ins.PushLike,
) {
	var source = push.GetSource()
	switch actual := source.(type) {
	case ins.LiteralLike:
		v.processor_.PreprocessLiteral(
			actual,
			0,
			0,
		)
		v.visitLiteral(actual)
		v.processor_.PostprocessLiteral(
			actual,
			0,
			0,
		)
	case ins.ConstantLike:
		v.processor_.PreprocessConstant(
			actual,
			0,
			0,
		)
		v.visitConstant(actual)
		v.processor_.PostprocessConstant(
			actual,
			0,
			0,
		)
	case ins.ArgumentLike:
		v.processor_.PreprocessArgument(
			actual,
			0,
			0,
		)
		v.visitArgument(actual)
		v.processor_.PostprocessArgument(
			actual,
			0,
			0,
		)
	case ins.HandlerLike:
		v.processor_.PreprocessHandler(
			actual,
			0,
			0,
		)
		v.visitHandler(actual)
		v.processor_.PostprocessHandler(
			actual,
			0,
			0,
		)
	default:
		var message = fmt.Sprintf(
			"An invalid source was found: %v",
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitSave(
	save ins.SaveLike,
) {
	var component = save.GetComponent()
	v.processor_.ProcessComponent(component)

	// Visit slot 1 between terms.
	v.processor_.ProcessSaveSlot(
		save,
		1,
	)

	var symbol = save.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitSend(
	send ins.SendLike,
) {
	var symbol = send.GetSymbol()
	v.processor_.ProcessSymbol(symbol)

	// Visit slot 1 between terms.
	v.processor_.ProcessSendSlot(
		send,
		1,
	)

	var destination = send.GetDestination()
	v.processor_.ProcessDestination(destination)
}

func (v *visitor_) visitSkip(
	skip ins.SkipLike,
) {
}

// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}

// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func visitorClass() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
