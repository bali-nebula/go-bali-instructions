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

package agents

import (
	fmt "fmt"
	lan "github.com/bali-nebula/go-assembly-language/v3"
	ins "github.com/bali-nebula/go-bali-instructions/v3/instructions"
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func InflatorClass() InflatorClassLike {
	return inflatorClass()
}

// Constructor Methods

func (c *inflatorClass_) Inflator() InflatorLike {
	var instance = &inflator_{
		// Initialize the instance attributes.
		stack_: com.StackWithCapacity[any](256),

		// Initialize the inherited aspects.
		Methodical: lan.ProcessorClass().Processor(),
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *inflator_) GetClass() InflatorClassLike {
	return inflatorClass()
}

func (v *inflator_) InflateAssembly(
	assembly lan.AssemblyLike,
) ins.AssemblyLike {
	lan.VisitorClass().Visitor(v).VisitAssembly(assembly)
	if v.stack_.GetSize() != 1 {
		var message = fmt.Sprintf(
			"Internal Error: the inflator stack is corrupted: %v",
			v.stack_,
		)
		panic(message)
	}
	return v.stack_.RemoveLast().(ins.AssemblyLike)
}

// Attribute Methods

// lan.Methodical Methods

func (v *inflator_) ProcessDescription(
	description string,
) {
	v.stack_.AddValue(description)
}

func (v *inflator_) ProcessLabel(
	label string,
) {
	v.stack_.AddValue(label)
}

func (v *inflator_) ProcessQuoted(
	quoted string,
) {
	v.stack_.AddValue(quoted)
}

func (v *inflator_) ProcessSymbol(
	symbol string,
) {
	v.stack_.AddValue(symbol)
}

func (v *inflator_) PostprocessArgument(
	argument lan.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.ArgumentClass().Argument(symbol))
}

func (v *inflator_) PostprocessAssembly(
	assembly lan.AssemblyLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessAssemblySlot(
	assembly lan.AssemblyLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessCall(
	call lan.CallLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessCallSlot(
	call lan.CallLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessCardinality(
	cardinality lan.CardinalityLike,
	index_ uint,
	count_ uint,
) {
	var modifier = cardinality.GetAny().(string)
	switch modifier {
	case "WITH 3 ARGUMENTS":
		v.stack_.AddValue(ins.With3Arguments)
	case "WITH 2 ARGUMENTS":
		v.stack_.AddValue(ins.With2Arguments)
	case "WITH 1 ARGUMENT":
		v.stack_.AddValue(ins.With1Argument)
	default:
		v.stack_.AddValue(ins.With0Arguments)
	}
}

func (v *inflator_) PostprocessComponent(
	component lan.ComponentLike,
	index_ uint,
	count_ uint,
) {
	var modifier = component.GetAny().(string)
	switch modifier {
	case "DRAFT":
		v.stack_.AddValue(ins.Draft)
	case "DOCUMENT":
		v.stack_.AddValue(ins.Document)
	case "MESSAGE":
		v.stack_.AddValue(ins.Message)
	case "VARIABLE":
		v.stack_.AddValue(ins.Variable)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected component in a switch statement: %s",
			modifier,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessConditionally(
	conditionally lan.ConditionallyLike,
	index_ uint,
	count_ uint,
) {
	var modifier = conditionally.GetAny().(string)
	switch modifier {
	case "ON EMPTY":
		v.stack_.AddValue(ins.OnEmpty)
	case "ON FALSE":
		v.stack_.AddValue(ins.OnFalse)
	case "ON NONE":
		v.stack_.AddValue(ins.OnNone)
	default:
		v.stack_.AddValue(ins.OnAny)
	}
}

func (v *inflator_) PostprocessConstant(
	constant lan.ConstantLike,
	index_ uint,
	count_ uint,
) {
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.ConstantClass().Constant(symbol))
}

func (v *inflator_) PostprocessDestination(
	destination lan.DestinationLike,
	index_ uint,
	count_ uint,
) {
	var modifier = destination.GetAny().(string)
	switch modifier {
	case "COMPONENT":
		v.stack_.AddValue(ins.Component)
	case "COMPONENT WITH ARGUMENTS":
		v.stack_.AddValue(ins.ComponentWithArguments)
	case "DOCUMENT":
		v.stack_.AddValue(ins.Document)
	case "DOCUMENT WITH ARGUMENTS":
		v.stack_.AddValue(ins.DocumentWithArguments)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected destination in a switch statement: %s",
			modifier,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessDrop(
	drop lan.DropLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessDropSlot(
	drop lan.DropLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessHandler(
	handler lan.HandlerLike,
	index_ uint,
	count_ uint,
) {
	var label = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.HandlerClass().Handler(label))
}

func (v *inflator_) PreprocessInstruction(
	instruction lan.InstructionLike,
	index_ uint,
	count_ uint,
) {
	var label string
	var prefix = instruction.GetOptionalPrefix()
	if uti.IsUndefined(prefix) {
		// We only add it if it is not defined, otherwise ProcessLabel adds it.
		v.stack_.AddValue(label)
	}
}

func (v *inflator_) PostprocessInstruction(
	instruction lan.InstructionLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessInstructionSlot(
	instruction lan.InstructionLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessJump(
	jump lan.JumpLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessJumpSlot(
	jump lan.JumpLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessLiteral(
	literal lan.LiteralLike,
	index_ uint,
	count_ uint,
) {
	var quoted = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.LiteralClass().Literal(quoted))
}

func (v *inflator_) PostprocessLoad(
	load lan.LoadLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessLoadSlot(
	load lan.LoadLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessNote(
	note lan.NoteLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessNoteSlot(
	note lan.NoteLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessPrefix(
	prefix lan.PrefixLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessPrefixSlot(
	prefix lan.PrefixLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessPull(
	pull lan.PullLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessPullSlot(
	pull lan.PullLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessPush(
	push lan.PushLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessPushSlot(
	push lan.PushLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessSave(
	save lan.SaveLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessSaveSlot(
	save lan.SaveLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessSend(
	send lan.SendLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessSendSlot(
	send lan.SendLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessSkip(
	skip lan.SkipLike,
	index_ uint,
	count_ uint,
) {
}

func (v *inflator_) ProcessSkipSlot(
	skip lan.SkipLike,
	slot_ uint,
) {
}

func (v *inflator_) PostprocessValue(
	value lan.ValueLike,
	index_ uint,
	count_ uint,
) {
	var modifier = value.GetAny().(string)
	switch modifier {
	case "COMPONENT":
		v.stack_.AddValue(ins.Component)
	case "RESULT":
		v.stack_.AddValue(ins.Result)
	case "EXCEPTION":
		v.stack_.AddValue(ins.Exception)
	case "HANDLER":
		v.stack_.AddValue(ins.Handler)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %s",
			modifier,
		)
		panic(message)
	}
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type inflator_ struct {
	// Declare the instance attributes.
	stack_ com.StackLike[any]

	// Declare the inherited aspects.
	lan.Methodical
}

// Class Structure

type inflatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inflatorClass() *inflatorClass_ {
	return inflatorClassReference_
}

var inflatorClassReference_ = &inflatorClass_{
	// Initialize the class constants.
}
