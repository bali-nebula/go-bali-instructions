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

package grammar

import (
	fmt "fmt"
	ast "github.com/bali-nebula/go-bali-instructions/v3/ast"
	com "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	mat "math"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ParserClass() ParserClassLike {
	return parserClass()
}

// Constructor Methods

func (c *parserClass_) Parser() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserClass()
}

func (v *parser_) ParseSource(
	source string,
) ast.AssemblyLike {
	v.source_ = sts.ReplaceAll(source, "\t", "    ")
	v.tokens_ = com.Queue[TokenLike]()
	v.next_ = com.Stack[TokenLike]()

	// The scanner runs in a separate Go routine.
	ScannerClass().Scanner(v.source_, v.tokens_)

	// Attempt to parse the assembly.
	var assembly, token, ok = v.parseAssembly()
	if !ok || !v.tokens_.IsEmpty() {
		var message = v.formatError("$Assembly", token)
		panic(message)
	}
	return assembly
}

// PROTECTED INTERFACE

// Private Methods

func (v *parser_) parseAction() (
	action ast.ActionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Note Action.
	var note ast.NoteLike
	note, token, ok = v.parseNote()
	if ok {
		// Found a single Note Action.
		action = ast.ActionClass().Action(note)
		return
	}

	// Attempt to parse a single Noop Action.
	var noop ast.NoopLike
	noop, token, ok = v.parseNoop()
	if ok {
		// Found a single Noop Action.
		action = ast.ActionClass().Action(noop)
		return
	}

	// Attempt to parse a single Jump Action.
	var jump ast.JumpLike
	jump, token, ok = v.parseJump()
	if ok {
		// Found a single Jump Action.
		action = ast.ActionClass().Action(jump)
		return
	}

	// Attempt to parse a single Push Action.
	var push ast.PushLike
	push, token, ok = v.parsePush()
	if ok {
		// Found a single Push Action.
		action = ast.ActionClass().Action(push)
		return
	}

	// Attempt to parse a single Pull Action.
	var pull ast.PullLike
	pull, token, ok = v.parsePull()
	if ok {
		// Found a single Pull Action.
		action = ast.ActionClass().Action(pull)
		return
	}

	// Attempt to parse a single Load Action.
	var load ast.LoadLike
	load, token, ok = v.parseLoad()
	if ok {
		// Found a single Load Action.
		action = ast.ActionClass().Action(load)
		return
	}

	// Attempt to parse a single Save Action.
	var save ast.SaveLike
	save, token, ok = v.parseSave()
	if ok {
		// Found a single Save Action.
		action = ast.ActionClass().Action(save)
		return
	}

	// Attempt to parse a single Drop Action.
	var drop ast.DropLike
	drop, token, ok = v.parseDrop()
	if ok {
		// Found a single Drop Action.
		action = ast.ActionClass().Action(drop)
		return
	}

	// Attempt to parse a single Call Action.
	var call ast.CallLike
	call, token, ok = v.parseCall()
	if ok {
		// Found a single Call Action.
		action = ast.ActionClass().Action(call)
		return
	}

	// Attempt to parse a single Send Action.
	var send ast.SendLike
	send, token, ok = v.parseSend()
	if ok {
		// Found a single Send Action.
		action = ast.ActionClass().Action(send)
		return
	}

	// This is not a single Action rule.
	return
}

func (v *parser_) parseArgument() (
	argument ast.ArgumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "ARGUMENT" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("ARGUMENT")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Argument rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Argument", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Argument", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Argument rule.
	ok = true
	v.remove(tokens)
	argument = ast.ArgumentClass().Argument(
		delimiter,
		symbol,
	)
	return
}

func (v *parser_) parseAssembly() (
	assembly ast.AssemblyLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse multiple Instruction rules.
	var instructions = com.List[ast.InstructionLike]()
instructionsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var instruction ast.InstructionLike
		instruction, token, ok = v.parseInstruction()
		if !ok {
			switch {
			case count_ >= 0:
				break instructionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Instruction rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Assembly", token)
				message += "0 or more Instruction rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		instructions.AppendValue(instruction)
	}

	// Found a single Assembly rule.
	ok = true
	v.remove(tokens)
	assembly = ast.AssemblyClass().Assembly(instructions)
	return
}

func (v *parser_) parseCall() (
	call ast.CallLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "CALL" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("CALL")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Call rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Call", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Call", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Cardinality rule.
	var optionalCardinality ast.CardinalityLike
	optionalCardinality, _, ok = v.parseCardinality()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Call rule.
	ok = true
	v.remove(tokens)
	call = ast.CallClass().Call(
		delimiter,
		symbol,
		optionalCardinality,
	)
	return
}

func (v *parser_) parseCardinality() (
	cardinality ast.CardinalityLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "WITH" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("WITH")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Cardinality rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Cardinality", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single count token.
	var count string
	count, token, ok = v.parseToken(CountToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single count token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Cardinality", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "ARGUMENTS" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("ARGUMENTS")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Cardinality rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Cardinality", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Cardinality rule.
	ok = true
	v.remove(tokens)
	cardinality = ast.CardinalityClass().Cardinality(
		delimiter1,
		count,
		delimiter2,
	)
	return
}

func (v *parser_) parseComponent() (
	component ast.ComponentLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "VARIABLE" delimiter.
	delimiter, token, ok = v.parseDelimiter("VARIABLE")
	if ok {
		// Found a single "VARIABLE" delimiter.
		component = ast.ComponentClass().Component(delimiter)
		return
	}

	// Attempt to parse a single "DOCUMENT" delimiter.
	delimiter, token, ok = v.parseDelimiter("DOCUMENT")
	if ok {
		// Found a single "DOCUMENT" delimiter.
		component = ast.ComponentClass().Component(delimiter)
		return
	}

	// Attempt to parse a single "CONTRACT" delimiter.
	delimiter, token, ok = v.parseDelimiter("CONTRACT")
	if ok {
		// Found a single "CONTRACT" delimiter.
		component = ast.ComponentClass().Component(delimiter)
		return
	}

	// Attempt to parse a single "MESSAGE" delimiter.
	delimiter, token, ok = v.parseDelimiter("MESSAGE")
	if ok {
		// Found a single "MESSAGE" delimiter.
		component = ast.ComponentClass().Component(delimiter)
		return
	}

	// This is not a single Component rule.
	return
}

func (v *parser_) parseCondition() (
	condition ast.ConditionLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "EMPTY" delimiter.
	delimiter, token, ok = v.parseDelimiter("EMPTY")
	if ok {
		// Found a single "EMPTY" delimiter.
		condition = ast.ConditionClass().Condition(delimiter)
		return
	}

	// Attempt to parse a single "NONE" delimiter.
	delimiter, token, ok = v.parseDelimiter("NONE")
	if ok {
		// Found a single "NONE" delimiter.
		condition = ast.ConditionClass().Condition(delimiter)
		return
	}

	// Attempt to parse a single "FALSE" delimiter.
	delimiter, token, ok = v.parseDelimiter("FALSE")
	if ok {
		// Found a single "FALSE" delimiter.
		condition = ast.ConditionClass().Condition(delimiter)
		return
	}

	// This is not a single Condition rule.
	return
}

func (v *parser_) parseConditionally() (
	conditionally ast.ConditionallyLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "ON" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("ON")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Conditionally rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Conditionally", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Condition rule.
	var condition ast.ConditionLike
	condition, token, ok = v.parseCondition()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Condition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Conditionally", token)
		panic(message)
	}

	// Found a single Conditionally rule.
	ok = true
	v.remove(tokens)
	conditionally = ast.ConditionallyClass().Conditionally(
		delimiter,
		condition,
	)
	return
}

func (v *parser_) parseConstant() (
	constant ast.ConstantLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "CONSTANT" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("CONSTANT")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Constant rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Constant", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Constant", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Constant rule.
	ok = true
	v.remove(tokens)
	constant = ast.ConstantClass().Constant(
		delimiter,
		symbol,
	)
	return
}

func (v *parser_) parseDestination() (
	destination ast.DestinationLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "COMPONENT" delimiter.
	delimiter, token, ok = v.parseDelimiter("COMPONENT")
	if ok {
		// Found a single "COMPONENT" delimiter.
		destination = ast.DestinationClass().Destination(delimiter)
		return
	}

	// Attempt to parse a single "DOCUMENT" delimiter.
	delimiter, token, ok = v.parseDelimiter("DOCUMENT")
	if ok {
		// Found a single "DOCUMENT" delimiter.
		destination = ast.DestinationClass().Destination(delimiter)
		return
	}

	// This is not a single Destination rule.
	return
}

func (v *parser_) parseDrop() (
	drop ast.DropLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "DROP" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("DROP")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Drop rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Drop", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Drop", token)
		panic(message)
	}

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Drop", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Drop rule.
	ok = true
	v.remove(tokens)
	drop = ast.DropClass().Drop(
		delimiter,
		component,
		symbol,
	)
	return
}

func (v *parser_) parseHandler() (
	handler ast.HandlerLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "HANDLER" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("HANDLER")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Handler rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Handler", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single label token.
	var label string
	label, token, ok = v.parseToken(LabelToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single label token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Handler", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Handler rule.
	ok = true
	v.remove(tokens)
	handler = ast.HandlerClass().Handler(
		delimiter,
		label,
	)
	return
}

func (v *parser_) parseInstruction() (
	instruction ast.InstructionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse an optional Prefix rule.
	var optionalPrefix ast.PrefixLike
	optionalPrefix, _, ok = v.parsePrefix()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single Action rule.
	var action ast.ActionLike
	action, token, ok = v.parseAction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Action rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Instruction", token)
		panic(message)
	}

	// Found a single Instruction rule.
	ok = true
	v.remove(tokens)
	instruction = ast.InstructionClass().Instruction(
		optionalPrefix,
		action,
	)
	return
}

func (v *parser_) parseJump() (
	jump ast.JumpLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "JUMP" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("JUMP")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Jump rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Jump", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "TO" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("TO")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Jump rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Jump", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single label token.
	var label string
	label, token, ok = v.parseToken(LabelToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single label token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Jump", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Conditionally rule.
	var optionalConditionally ast.ConditionallyLike
	optionalConditionally, _, ok = v.parseConditionally()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Jump rule.
	ok = true
	v.remove(tokens)
	jump = ast.JumpClass().Jump(
		delimiter1,
		delimiter2,
		label,
		optionalConditionally,
	)
	return
}

func (v *parser_) parseLiteral() (
	literal ast.LiteralLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "LITERAL" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("LITERAL")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Literal rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Literal", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single quoted token.
	var quoted string
	quoted, token, ok = v.parseToken(QuotedToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single quoted token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Literal", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Literal rule.
	ok = true
	v.remove(tokens)
	literal = ast.LiteralClass().Literal(
		delimiter,
		quoted,
	)
	return
}

func (v *parser_) parseLoad() (
	load ast.LoadLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "LOAD" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("LOAD")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Load rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Load", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Load", token)
		panic(message)
	}

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Load", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Load rule.
	ok = true
	v.remove(tokens)
	load = ast.LoadClass().Load(
		delimiter,
		component,
		symbol,
	)
	return
}

func (v *parser_) parseNoop() (
	noop ast.NoopLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "NOOP" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("NOOP")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Noop rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Noop", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Noop rule.
	ok = true
	v.remove(tokens)
	noop = ast.NoopClass().Noop(delimiter)
	return
}

func (v *parser_) parseNote() (
	note ast.NoteLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "NOTE" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("NOTE")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Note rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Note", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single explanation token.
	var explanation string
	explanation, token, ok = v.parseToken(ExplanationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single explanation token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Note", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Note rule.
	ok = true
	v.remove(tokens)
	note = ast.NoteClass().Note(
		delimiter,
		explanation,
	)
	return
}

func (v *parser_) parseParameterized() (
	parameterized ast.ParameterizedLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "WITH" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("WITH")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Parameterized rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameterized", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "ARGUMENTS" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("ARGUMENTS")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Parameterized rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameterized", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Parameterized rule.
	ok = true
	v.remove(tokens)
	parameterized = ast.ParameterizedClass().Parameterized(
		delimiter1,
		delimiter2,
	)
	return
}

func (v *parser_) parsePrefix() (
	prefix ast.PrefixLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single label token.
	var label string
	label, token, ok = v.parseToken(LabelToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single label token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Prefix", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ":" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Prefix rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Prefix", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Prefix rule.
	ok = true
	v.remove(tokens)
	prefix = ast.PrefixClass().Prefix(
		label,
		delimiter,
	)
	return
}

func (v *parser_) parsePull() (
	pull ast.PullLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "PULL" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("PULL")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Pull rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Pull", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Value rule.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Value rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Pull", token)
		panic(message)
	}

	// Found a single Pull rule.
	ok = true
	v.remove(tokens)
	pull = ast.PullClass().Pull(
		delimiter,
		value,
	)
	return
}

func (v *parser_) parsePush() (
	push ast.PushLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "PUSH" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("PUSH")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Push rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Push", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Source rule.
	var source ast.SourceLike
	source, token, ok = v.parseSource()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Source rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Push", token)
		panic(message)
	}

	// Found a single Push rule.
	ok = true
	v.remove(tokens)
	push = ast.PushClass().Push(
		delimiter,
		source,
	)
	return
}

func (v *parser_) parseSave() (
	save ast.SaveLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "SAVE" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("SAVE")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Save rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Save", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Save", token)
		panic(message)
	}

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Save", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Save rule.
	ok = true
	v.remove(tokens)
	save = ast.SaveClass().Save(
		delimiter,
		component,
		symbol,
	)
	return
}

func (v *parser_) parseSend() (
	send ast.SendLike,
	token TokenLike,
	ok bool,
) {
	var tokens = com.List[TokenLike]()

	// Attempt to parse a single "SEND" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("SEND")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Send rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Send", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Send", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "TO" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("TO")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Send rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Send", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Destination rule.
	var destination ast.DestinationLike
	destination, token, ok = v.parseDestination()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Destination rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Send", token)
		panic(message)
	}

	// Attempt to parse an optional Parameterized rule.
	var optionalParameterized ast.ParameterizedLike
	optionalParameterized, _, ok = v.parseParameterized()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Send rule.
	ok = true
	v.remove(tokens)
	send = ast.SendClass().Send(
		delimiter1,
		symbol,
		delimiter2,
		destination,
		optionalParameterized,
	)
	return
}

func (v *parser_) parseSource() (
	source ast.SourceLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Handler Source.
	var handler ast.HandlerLike
	handler, token, ok = v.parseHandler()
	if ok {
		// Found a single Handler Source.
		source = ast.SourceClass().Source(handler)
		return
	}

	// Attempt to parse a single Literal Source.
	var literal ast.LiteralLike
	literal, token, ok = v.parseLiteral()
	if ok {
		// Found a single Literal Source.
		source = ast.SourceClass().Source(literal)
		return
	}

	// Attempt to parse a single Constant Source.
	var constant ast.ConstantLike
	constant, token, ok = v.parseConstant()
	if ok {
		// Found a single Constant Source.
		source = ast.SourceClass().Source(constant)
		return
	}

	// Attempt to parse a single Argument Source.
	var argument ast.ArgumentLike
	argument, token, ok = v.parseArgument()
	if ok {
		// Found a single Argument Source.
		source = ast.SourceClass().Source(argument)
		return
	}

	// This is not a single Source rule.
	return
}

func (v *parser_) parseValue() (
	value ast.ValueLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "HANDLER" delimiter.
	delimiter, token, ok = v.parseDelimiter("HANDLER")
	if ok {
		// Found a single "HANDLER" delimiter.
		value = ast.ValueClass().Value(delimiter)
		return
	}

	// Attempt to parse a single "COMPONENT" delimiter.
	delimiter, token, ok = v.parseDelimiter("COMPONENT")
	if ok {
		// Found a single "COMPONENT" delimiter.
		value = ast.ValueClass().Value(delimiter)
		return
	}

	// Attempt to parse a single "RESULT" delimiter.
	delimiter, token, ok = v.parseDelimiter("RESULT")
	if ok {
		// Found a single "RESULT" delimiter.
		value = ast.ValueClass().Value(delimiter)
		return
	}

	// Attempt to parse a single "EXCEPTION" delimiter.
	delimiter, token, ok = v.parseDelimiter("EXCEPTION")
	if ok {
		// Found a single "EXCEPTION" delimiter.
		value = ast.ValueClass().Value(delimiter)
		return
	}

	// This is not a single Value rule.
	return
}

func (v *parser_) parseDelimiter(
	literal string,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == literal {
			// Found the desired delimiter.
			return
		}
		v.next_.AddValue(token)
		ok = false
	}

	// This is not the desired delimiter.
	return
}

func (v *parser_) parseToken(
	tokenType TokenType,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = com.List[TokenLike]()
	token = v.getNextToken()
	for token != nil {
		tokens.AppendValue(token)
		switch token.GetType() {
		case tokenType:
			// Found the desired token type.
			value = token.GetValue()
			ok = true
			return
		case SpaceToken, NewlineToken:
			// Ignore any unrequested whitespace.
			token = v.getNextToken()
		default:
			// This is not the desired token type.
			v.putBack(tokens)
			return
		}
	}

	// We are at the end-of-file marker.
	return
}

func (v *parser_) formatError(
	ruleName string,
	token TokenLike,
) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		ScannerClass().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source lines with the error in it.
	message += "\033[36m"
	for index := line - 3; index < line; index++ {
		if index > 1 {
			message += fmt.Sprintf("%04d: ", index) + string(lines[index-1]) + "\n"
		}
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if uti.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(
	ruleName string,
) string {
	var syntax = parserClass().syntax_
	var definition = syntax.GetValue(ruleName)
	return definition
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveLast()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveFirst() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError("", token)
		panic(message)
	}

	return token
}

func (v *parser_) putBack(
	tokens com.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

// NOTE:
// This method does nothing but must exist to satisfy the lint check on the
// generated parser code.  The generated code must call this method is some
// cases to make it look that the tokens variable is being used somewhere.
func (v *parser_) remove(
	tokens com.Sequential[TokenLike],
) {
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ com.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   com.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ com.CatalogLike[string, string]
}

// Class Reference

func parserClass() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: com.CatalogFromMap[string, string](
		map[string]string{
			"$Assembly":    `Instruction*`,
			"$Instruction": `Prefix? Action`,
			"$Prefix":      `label ":"`,
			"$Action": `
    Note
    Noop
    Jump
    Push
    Pull
    Load
    Save
    Drop
    Call
    Send`,
			"$Note":          `"NOTE" explanation`,
			"$Noop":          `"NOOP"`,
			"$Jump":          `"JUMP" "TO" label Conditionally?`,
			"$Conditionally": `"ON" Condition`,
			"$Condition": `
    "EMPTY"
    "NONE"
    "FALSE"`,
			"$Push": `"PUSH" Source`,
			"$Source": `
    Handler
    Literal
    Constant
    Argument`,
			"$Handler":  `"HANDLER" label`,
			"$Literal":  `"LITERAL" quoted`,
			"$Constant": `"CONSTANT" symbol`,
			"$Argument": `"ARGUMENT" symbol`,
			"$Pull":     `"PULL" Value`,
			"$Value": `
    "HANDLER"
    "COMPONENT"
    "RESULT"
    "EXCEPTION"`,
			"$Load": `"LOAD" Component symbol`,
			"$Save": `"SAVE" Component symbol`,
			"$Drop": `"DROP" Component symbol`,
			"$Component": `
    "VARIABLE"
    "DOCUMENT"
    "CONTRACT"
    "MESSAGE"`,
			"$Call":        `"CALL" symbol Cardinality?`,
			"$Cardinality": `"WITH" count "ARGUMENTS"  ! The argument count is in the range [1..3].`,
			"$Send":        `"SEND" symbol "TO" Destination Parameterized?`,
			"$Destination": `
    "COMPONENT"
    "DOCUMENT"`,
			"$Parameterized": `"WITH" "ARGUMENTS"`,
		},
	),
}
