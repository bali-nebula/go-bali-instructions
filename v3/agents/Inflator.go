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

func (v *inflator_) InflateMethod(
	method lan.MethodLike,
) ins.MethodLike {
	lan.VisitorClass().Visitor(v).VisitMethod(method)
	return v.stack_.RemoveLast().(ins.MethodLike)
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

func (v *inflator_) PostprocessCall(
	call lan.CallLike,
	index_ uint,
	count_ uint,
) {
	var argumentCount uint8
	if uti.IsDefined(call.GetOptionalContext()) {
		argumentCount = v.stack_.RemoveLast().(uint8)
	}
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.CallClass().Call(symbol, argumentCount))
}

func (v *inflator_) PostprocessContext(
	context lan.ContextLike,
	index_ uint,
	count_ uint,
) {
	var argumentCount uint8
	var delimiter = context.GetAny().(string)
	switch delimiter {
	case "":
		argumentCount = 0
	case "WITH 1 ARGUMENT":
		argumentCount = 1
	case "WITH 2 ARGUMENTS":
		argumentCount = 2
	case "WITH 3 ARGUMENTS":
		argumentCount = 3
	default:
		var message = fmt.Sprintf(
			"Found an unexpected argument count in a switch statement: %s",
			delimiter,
		)
		panic(message)
	}
	v.stack_.AddValue(argumentCount)
}

func (v *inflator_) PostprocessComponent(
	component lan.ComponentLike,
	index_ uint,
	count_ uint,
) {
	var delimiter = component.GetAny().(string)
	switch delimiter {
	case "DRAFT":
		v.stack_.AddValue(ins.DraftComponent)
	case "DOCUMENT":
		v.stack_.AddValue(ins.DocumentComponent)
	case "MESSAGE":
		v.stack_.AddValue(ins.MessageComponent)
	case "VARIABLE":
		v.stack_.AddValue(ins.VariableComponent)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected component in a switch statement: %s",
			delimiter,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessCondition(
	condition lan.ConditionLike,
	index_ uint,
	count_ uint,
) {
	var delimiter = condition.GetAny().(string)
	switch delimiter {
	case "":
		v.stack_.AddValue(ins.OnAnyCondition)
	case "ON EMPTY":
		v.stack_.AddValue(ins.OnEmptyCondition)
	case "ON FALSE":
		v.stack_.AddValue(ins.OnFalseCondition)
	case "ON NONE":
		v.stack_.AddValue(ins.OnNoneCondition)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected condition in a switch statement: %s",
			delimiter,
		)
		panic(message)
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
	var delimiter = destination.GetAny().(string)
	switch delimiter {
	case "COMPONENT":
		v.stack_.AddValue(ins.ComponentDestination)
	case "COMPONENT WITH ARGUMENTS":
		v.stack_.AddValue(ins.ComponentWithArgumentsDestination)
	case "DOCUMENT":
		v.stack_.AddValue(ins.DocumentDestination)
	case "DOCUMENT WITH ARGUMENTS":
		v.stack_.AddValue(ins.DocumentWithArgumentsDestination)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected destination in a switch statement: %s",
			delimiter,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessDrop(
	drop lan.DropLike,
	index_ uint,
	count_ uint,
) {
	var symbol = v.stack_.RemoveLast().(string)
	var component = v.stack_.RemoveLast().(ins.Component)
	v.stack_.AddValue(ins.DropClass().Drop(component, symbol))
}

func (v *inflator_) PostprocessHandler(
	handler lan.HandlerLike,
	index_ uint,
	count_ uint,
) {
	var label = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.HandlerClass().Handler(label))
}

func (v *inflator_) PostprocessInstruction(
	instruction lan.InstructionLike,
	index_ uint,
	count_ uint,
) {
	var prefix ins.PrefixLike
	var action = v.stack_.RemoveLast()
	if uti.IsDefined(instruction.GetOptionalPrefix()) {
		prefix = v.stack_.RemoveLast().(ins.PrefixLike)
	}
	v.stack_.AddValue(ins.InstructionClass().Instruction(prefix, action))
}

func (v *inflator_) PostprocessJump(
	jump lan.JumpLike,
	index_ uint,
	count_ uint,
) {
	var condition ins.Condition
	if uti.IsDefined(jump.GetOptionalCondition()) {
		condition = v.stack_.RemoveLast().(ins.Condition)
	}
	var label = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.JumpClass().Jump(label, condition))
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
	var symbol = v.stack_.RemoveLast().(string)
	var component = v.stack_.RemoveLast().(ins.Component)
	v.stack_.AddValue(ins.LoadClass().Load(component, symbol))
}

func (v *inflator_) PostprocessMethod(
	method lan.MethodLike,
	index_ uint,
	count_ uint,
) {
	var instructions = com.List[ins.InstructionLike]()
	var iterator = method.GetInstructions().GetIterator()
	for iterator.HasNext() {
		var instruction = v.stack_.RemoveLast().(ins.InstructionLike)
		instructions.AppendValue(instruction)
		iterator.GetNext()
	}
	instructions.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(ins.MethodClass().Method(instructions))
	if v.stack_.GetSize() != 1 {
		var message = fmt.Sprintf(
			"Internal Error: the inflator stack is corrupted: %v",
			v.stack_,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessNote(
	note lan.NoteLike,
	index_ uint,
	count_ uint,
) {
	var description = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.NoteClass().Note(description))
}

func (v *inflator_) PostprocessPrefix(
	prefix lan.PrefixLike,
	index_ uint,
	count_ uint,
) {
	var label = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.PrefixClass().Prefix(label))
}

func (v *inflator_) PostprocessPull(
	pull lan.PullLike,
	index_ uint,
	count_ uint,
) {
	var value = v.stack_.RemoveLast().(ins.Value)
	v.stack_.AddValue(ins.PullClass().Pull(value))
}

func (v *inflator_) PostprocessPush(
	push lan.PushLike,
	index_ uint,
	count_ uint,
) {
	var source = v.stack_.RemoveLast()
	v.stack_.AddValue(ins.PushClass().Push(source))
}

func (v *inflator_) PostprocessSave(
	save lan.SaveLike,
	index_ uint,
	count_ uint,
) {
	var symbol = v.stack_.RemoveLast().(string)
	var component = v.stack_.RemoveLast().(ins.Component)
	v.stack_.AddValue(ins.SaveClass().Save(component, symbol))
}

func (v *inflator_) PostprocessSend(
	send lan.SendLike,
	index_ uint,
	count_ uint,
) {
	var destination = v.stack_.RemoveLast().(ins.Destination)
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.SendClass().Send(symbol, destination))
}

func (v *inflator_) PostprocessSkip(
	skip lan.SkipLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(ins.SkipClass().Skip())
}

func (v *inflator_) PostprocessValue(
	value lan.ValueLike,
	index_ uint,
	count_ uint,
) {
	var delimiter = value.GetAny().(string)
	switch delimiter {
	case "COMPONENT":
		v.stack_.AddValue(ins.ComponentValue)
	case "RESULT":
		v.stack_.AddValue(ins.ResultValue)
	case "EXCEPTION":
		v.stack_.AddValue(ins.ExceptionValue)
	case "HANDLER":
		v.stack_.AddValue(ins.HandlerValue)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %s",
			delimiter,
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
