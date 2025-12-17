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
	var instructions = com.List[ins.InstructionLike]()
	var iterator = assembly.GetInstructions().GetIterator()
	var index = 1
	for iterator.HasNext() {
		fmt.Println(index)
		fmt.Println(v.stack_.GetLast())
		index++
		var instruction = v.stack_.RemoveLast().(ins.InstructionLike)
		fmt.Println(instruction.AsSource())
		instructions.AppendValue(instruction)
		iterator.GetNext()
	}
	instructions.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(ins.AssemblyClass().Assembly(instructions))
}

func (v *inflator_) PostprocessCall(
	call lan.CallLike,
	index_ uint,
	count_ uint,
) {
	var cardinality = ins.With0ArgumentsModifier
	if uti.IsDefined(call.GetOptionalCardinality()) {
		cardinality = v.stack_.RemoveLast().(ins.Modifier)
	}
	var symbol = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.CallClass().Call(symbol, cardinality))
}

func (v *inflator_) PostprocessCardinality(
	cardinality lan.CardinalityLike,
	index_ uint,
	count_ uint,
) {
	var modifier = cardinality.GetAny().(string)
	switch modifier {
	case "WITH 3 ARGUMENTS":
		v.stack_.AddValue(ins.With3ArgumentsModifier)
	case "WITH 2 ARGUMENTS":
		v.stack_.AddValue(ins.With2ArgumentsModifier)
	case "WITH 1 ARGUMENT":
		v.stack_.AddValue(ins.With1ArgumentModifier)
	default:
		v.stack_.AddValue(ins.With0ArgumentsModifier)
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
		v.stack_.AddValue(ins.DraftModifier)
	case "DOCUMENT":
		v.stack_.AddValue(ins.DocumentModifier)
	case "MESSAGE":
		v.stack_.AddValue(ins.MessageModifier)
	case "VARIABLE":
		v.stack_.AddValue(ins.VariableModifier)
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
		v.stack_.AddValue(ins.OnEmptyModifier)
	case "ON FALSE":
		v.stack_.AddValue(ins.OnFalseModifier)
	case "ON NONE":
		v.stack_.AddValue(ins.OnNoneModifier)
	default:
		v.stack_.AddValue(ins.OnAnyModifier)
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
		v.stack_.AddValue(ins.ComponentModifier)
	case "COMPONENT WITH ARGUMENTS":
		v.stack_.AddValue(ins.ComponentWithArgumentsModifier)
	case "DOCUMENT":
		v.stack_.AddValue(ins.DocumentModifier)
	case "DOCUMENT WITH ARGUMENTS":
		v.stack_.AddValue(ins.DocumentWithArgumentsModifier)
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
	var symbol = v.stack_.RemoveLast().(string)
	var component = v.stack_.RemoveLast().(ins.Modifier)
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
	var prefix string
	var action = v.stack_.RemoveLast()
	if uti.IsDefined(instruction.GetOptionalPrefix()) {
		prefix = v.stack_.RemoveLast().(string)
	}
	v.stack_.AddValue(ins.InstructionClass().Instruction(prefix, action))
}

func (v *inflator_) PostprocessJump(
	jump lan.JumpLike,
	index_ uint,
	count_ uint,
) {
	var condition = ins.OnAnyModifier
	if uti.IsDefined(jump.GetOptionalConditionally()) {
		condition = v.stack_.RemoveLast().(ins.Modifier)
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
	var component = v.stack_.RemoveLast().(ins.Modifier)
	v.stack_.AddValue(ins.LoadClass().Load(component, symbol))
}

func (v *inflator_) PostprocessNote(
	note lan.NoteLike,
	index_ uint,
	count_ uint,
) {
	var description = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(ins.NoteClass().Note(description))
}

func (v *inflator_) PostprocessPull(
	pull lan.PullLike,
	index_ uint,
	count_ uint,
) {
	var value = v.stack_.RemoveLast().(ins.Modifier)
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
	var component = v.stack_.RemoveLast().(ins.Modifier)
	v.stack_.AddValue(ins.SaveClass().Save(component, symbol))
}

func (v *inflator_) PostprocessSend(
	send lan.SendLike,
	index_ uint,
	count_ uint,
) {
	var destination = v.stack_.RemoveLast().(ins.Modifier)
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
	var modifier = value.GetAny().(string)
	switch modifier {
	case "COMPONENT":
		v.stack_.AddValue(ins.ComponentModifier)
	case "RESULT":
		v.stack_.AddValue(ins.ResultModifier)
	case "EXCEPTION":
		v.stack_.AddValue(ins.ExceptionModifier)
	case "HANDLER":
		v.stack_.AddValue(ins.HandlerModifier)
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
