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
	var context lan.ContextLike
	switch argumentCount {
	case 1:
		context = lan.Context("WITH 1 ARGUMENT")
	case 2:
		context = lan.Context("WITH 2 ARGUMENTS")
	case 3:
		context = lan.Context("WITH 3 ARGUMENTS")
	}
	if uti.IsDefined(context) {
		v.stack_.AddValue(context)
	}
}

func (v *deflator_) ProcessComponent(
	component ins.Component,
) {
	var delimiter string
	switch component {
	case ins.VariableComponent:
		delimiter = "VARIABLE"
	case ins.DraftComponent:
		delimiter = "DRAFT"
	case ins.DocumentComponent:
		delimiter = "DOCUMENT"
	case ins.MessageComponent:
		delimiter = "MESSAGE"
	default:
		var message = fmt.Sprintf(
			"An invalid component type was found: %v",
			component,
		)
		panic(message)
	}
	v.stack_.AddValue(delimiter)
}

func (v *deflator_) ProcessCondition(
	condition ins.Condition,
) {
	var delimiter string
	switch condition {
	case ins.OnAnyCondition:
	case ins.OnEmptyCondition:
		delimiter = "ON EMPTY"
	case ins.OnNoneCondition:
		delimiter = "ON NONE"
	case ins.OnFalseCondition:
		delimiter = "ON FALSE"
	default:
		var message = fmt.Sprintf(
			"An invalid condition was found: %v",
			condition,
		)
		panic(message)
	}
	if uti.IsDefined(delimiter) {
		v.stack_.AddValue(delimiter)
	}
}

func (v *deflator_) ProcessDescription(
	description string,
) {
	v.stack_.AddValue(description)
}

func (v *deflator_) ProcessDestination(
	destination ins.Destination,
) {
	var delimiter string
	switch destination {
	case ins.ComponentDestination:
		delimiter = "COMPONENT"
	case ins.ComponentWithArgumentsDestination:
		delimiter = "COMPONENT WITH ARGUMENTS"
	case ins.DocumentDestination:
		delimiter = "DOCUMENT"
	case ins.DocumentWithArgumentsDestination:
		delimiter = "DOCUMENT WITH ARGUMENTS"
	default:
		var message = fmt.Sprintf(
			"An invalid destination was found: %v",
			destination,
		)
		panic(message)
	}
	v.stack_.AddValue(delimiter)
}

func (v *deflator_) ProcessLabel(
	label string,
) {
	v.stack_.AddValue(label)
}

func (v *deflator_) ProcessQuoted(
	quoted string,
) {
	v.stack_.AddValue(quoted)
}

func (v *deflator_) ProcessSource(
	source ins.Source,
) {
	var delimiter string
	switch source {
	case ins.LiteralSource:
		delimiter = "LITERAL"
	case ins.ConstantSource:
		delimiter = "CONSTANT"
	case ins.ArgumentSource:
		delimiter = "ARGUMENT"
	case ins.HandlerSource:
		delimiter = "HANDLER"
	default:
		var message = fmt.Sprintf(
			"An invalid source was found: %v",
			source,
		)
		panic(message)
	}
	v.stack_.AddValue(delimiter)
}

func (v *deflator_) ProcessSymbol(
	symbol string,
) {
	v.stack_.AddValue(symbol)
}

func (v *deflator_) ProcessValue(
	value ins.Value,
) {
	var delimiter string
	switch value {
	case ins.ComponentValue:
		delimiter = "COMPONENT"
	case ins.ResultValue:
		delimiter = "RESULT"
	case ins.ExceptionValue:
		delimiter = "EXCEPTION"
	case ins.HandlerValue:
		delimiter = "HANDLER"
	default:
		var message = fmt.Sprintf(
			"An invalid value was found: %v",
			value,
		)
		panic(message)
	}
	v.stack_.AddValue(delimiter)
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
	var instructions = com.List[lan.InstructionLike]()
	var iterator = assembly.GetInstructions().GetIterator()
	for iterator.HasNext() {
		var instruction = v.stack_.RemoveLast().(lan.InstructionLike)
		instructions.AppendValue(instruction)
		iterator.GetNext()
	}
	instructions.ReverseValues()
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
	var symbol = v.stack_.RemoveLast().(string)
	var component = lan.Component(v.stack_.RemoveLast().(string))
	v.stack_.AddValue(lan.Load("DROP", component, symbol))
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
	if jump.GetCondition() != ins.OnAnyCondition {
		condition = lan.Condition(v.stack_.RemoveLast().(string))
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
	var symbol = v.stack_.RemoveLast().(string)
	var component = lan.Component(v.stack_.RemoveLast().(string))
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
	var value = lan.Value(v.stack_.RemoveLast().(string))
	v.stack_.AddValue(lan.Pull("PULL", value))
}

func (v *deflator_) PostprocessPush(
	push ins.PushLike,
	index_ uint,
	count_ uint,
) {
	var source = lan.Source(v.stack_.RemoveLast())
	v.stack_.AddValue(lan.Push("PUSH", source))
}

func (v *deflator_) PostprocessSave(
	save ins.SaveLike,
	index_ uint,
	count_ uint,
) {
	var symbol = v.stack_.RemoveLast().(string)
	var component = lan.Component(v.stack_.RemoveLast().(string))
	v.stack_.AddValue(lan.Load("SAVE", component, symbol))
}

func (v *deflator_) PostprocessSend(
	send ins.SendLike,
	index_ uint,
	count_ uint,
) {
	var destination = lan.Destination(v.stack_.RemoveLast().(string))
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
