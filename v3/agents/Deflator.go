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

func DeflatorClass() DeflatorClassLike {
	return deflatorClass()
}

// Constructor Methods

func (c *deflatorClass_) Deflator() DeflatorLike {
	var instance = &deflator_{
		// Initialize the instance attributes.
		stack_: com.StackWithCapacity[any](256),

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *deflator_) GetClass() DeflatorClassLike {
	return deflatorClass()
}

func (v *deflator_) DeflateAssembly(
	assembly ins.AssemblyLike,
) lan.AssemblyLike {
	VisitorClass().Visitor(v).VisitAssembly(assembly)
	return v.stack_.RemoveLast().(lan.AssemblyLike)
}

// Attribute Methods

// Methodical Methods

func (v *deflator_) ProcessArgumentCount(
	argumentCount uint8,
) {
	v.stack_.AddValue(argumentCount)
}

func (v *deflator_) ProcessDescription(
	description string,
) {
	v.stack_.AddValue(description)
}

func (v *deflator_) ProcessLabel(
	label string,
) {
	v.stack_.AddValue(label)
}

func (v *deflator_) ProcessModifier(
	modifier ins.Modifier,
) {
	var delimiter string
	switch modifier {
	case ins.OnEmptyModifier:
		delimiter = "ON EMPTY"
	case ins.OnNoneModifier:
		delimiter = "ON NONE"
	case ins.OnFalseModifier:
		delimiter = "ON FALSE"
	case ins.HandlerModifier:
		delimiter = "HANDLER"
	case ins.LiteralModifier:
		delimiter = "LITERAL"
	case ins.ConstantModifier:
		delimiter = "CONSTANT"
	case ins.VariableModifier:
		delimiter = "VARIABLE"
	case ins.ArgumentModifier:
		delimiter = "ARGUMENT"
	case ins.MessageModifier:
		delimiter = "MESSAGE"
	case ins.ResultModifier:
		delimiter = "RESULT"
	case ins.ExceptionModifier:
		delimiter = "EXCEPTION"
	case ins.DraftModifier:
		delimiter = "DRAFT"
	case ins.ComponentModifier:
		delimiter = "COMPONENT"
	case ins.ComponentWithArgumentsModifier:
		delimiter = "COMPONENT WITH ARGUMENTS"
	case ins.DocumentModifier:
		delimiter = "DOCUMENT"
	case ins.DocumentWithArgumentsModifier:
		delimiter = "DOCUMENT WITH ARGUMENTS"
	}
	v.stack_.AddValue(delimiter)
}

func (v *deflator_) ProcessQuoted(
	quoted string,
) {
	v.stack_.AddValue(quoted)
}

func (v *deflator_) ProcessSymbol(
	symbol string,
) {
	v.stack_.AddValue(symbol)
}

func (v *deflator_) PostprocessArgument(
	argument ins.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Argument("ARGUMENT", symbol))
}

func (v *deflator_) PostprocessAssembly(
	assembly ins.AssemblyLike,
	index_ uint,
	count_ uint,
) {
	var instructions = v.stack_.RemoveLast().(com.Sequential[lan.InstructionLike])
	v.stack_.AddValue(lan.Assembly(instructions))
	if v.stack_.GetSize() != 1 {
		var message = fmt.Sprintf(
			"Internal Error: the deflator stack is corrupted: %v",
			v.stack_,
		)
		panic(message)
	}
}

func (v *deflator_) PostprocessCall(
	call ins.CallLike,
	index_ uint,
	count_ uint,
) {
	var context lan.ContextLike
	if call.GetArgumentCount() > 0 {
		context = v.stack_.RemoveLast().(lan.ContextLike)
	}
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Call("CALL", symbol, context))
}

func (v *deflator_) PostprocessConstant(
	constant ins.ConstantLike,
	index_ uint,
	count_ uint,
) {
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Constant("CONSTANT", symbol))
}

func (v *deflator_) PostprocessDrop(
	drop ins.DropLike,
	index_ uint,
	count_ uint,
) {
	var component = v.stack_.RemoveLast().(lan.ComponentLike)
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Drop("DROP", component, symbol))
}

func (v *deflator_) PostprocessHandler(
	handler ins.HandlerLike,
	index_ uint,
	count_ uint,
) {
	var label = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Handler("HANDLER", label))
}

func (v *deflator_) PostprocessInstruction(
	instruction ins.InstructionLike,
	index_ uint,
	count_ uint,
) {
	var action = lan.Action(v.stack_.RemoveLast())
	var prefix lan.PrefixLike
	if uti.IsDefined(instruction.GetOptionalPrefix()) {
		prefix = v.stack_.RemoveLast().(lan.PrefixLike)
	}
	v.stack_.AddValue(lan.Instruction(prefix, action))
}

func (v *deflator_) PostprocessJump(
	jump ins.JumpLike,
	index_ uint,
	count_ uint,
) {
	var condition lan.ConditionLike
	if jump.GetCondition() != ins.OnAnyModifier {
		condition = v.stack_.RemoveLast().(lan.ConditionLike)
	}
	var label = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Jump("JUMP TO", label, condition))
}

func (v *deflator_) PostprocessLiteral(
	literal ins.LiteralLike,
	index_ uint,
	count_ uint,
) {
	var quoted = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Literal("LITERAL", quoted))
}

func (v *deflator_) PostprocessLoad(
	load ins.LoadLike,
	index_ uint,
	count_ uint,
) {
	var component = v.stack_.RemoveLast().(lan.ComponentLike)
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Load("LOAD", component, symbol))
}

func (v *deflator_) PostprocessNote(
	note ins.NoteLike,
	index_ uint,
	count_ uint,
) {
	var description = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Note("NOTE", description))
}

func (v *deflator_) PostprocessPrefix(
	prefix ins.PrefixLike,
	index_ uint,
	count_ uint,
) {
	var label = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Prefix(label, ":"))
}

func (v *deflator_) PostprocessPull(
	pull ins.PullLike,
	index_ uint,
	count_ uint,
) {
	var value = v.stack_.RemoveLast().(lan.ValueLike)
	v.stack_.AddValue(lan.Pull("PULL", value))
}

func (v *deflator_) PostprocessPush(
	push ins.PushLike,
	index_ uint,
	count_ uint,
) {
	var source = v.stack_.RemoveLast().(lan.SourceLike)
	v.stack_.AddValue(lan.Push("PUSH", source))
}

func (v *deflator_) PostprocessSave(
	save ins.SaveLike,
	index_ uint,
	count_ uint,
) {
	var component = v.stack_.RemoveLast().(lan.ComponentLike)
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Save("SAVE", component, symbol))
}

func (v *deflator_) PostprocessSend(
	send ins.SendLike,
	index_ uint,
	count_ uint,
) {
	var destination = v.stack_.RemoveLast().(lan.DestinationLike)
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(lan.Send("SEND", symbol, "TO", destination))
}

func (v *deflator_) PostprocessSkip(
	skip ins.SkipLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(lan.Skip("SKIP"))
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type deflator_ struct {
	// Declare the instance attributes.
	stack_ com.StackLike[any]

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type deflatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func deflatorClass() *deflatorClass_ {
	return deflatorClassReference_
}

var deflatorClassReference_ = &deflatorClass_{
	// Initialize the class constants.
}
