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
│            This "package_api.go" file was automatically generated.           │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

For detailed documentation on this package refer to the wiki:
  - https://github.com/bali-nebula/go-bali-instructions/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package grammar

import (
	ast "github.com/bali-nebula/go-bali-instructions/v3/ast"
	col "github.com/craterdog/go-collection-framework/v7"
)

// TYPE DECLARATIONS

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	CountToken
	DelimiterToken
	ExplanationToken
	LabelToken
	NewlineToken
	QuotedToken
	SpaceToken
	SymbolToken
)

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
FormatterClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor Methods
	Formatter() FormatterLike
}

/*
ParserClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor Methods
	Parser() ParserLike
}

/*
ProcessorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor Methods
	Processor() ProcessorLike
}

/*
ScannerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructor Methods
	Scanner(
		source string,
		tokens col.QueueLike[TokenLike],
	) ScannerLike

	// Function Methods
	FormatToken(
		token TokenLike,
	) string
	FormatType(
		tokenType TokenType,
	) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type TokenClassLike interface {
	// Constructor Methods
	Token(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor Methods
	Validator() ValidatorLike
}

/*
VisitorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor Methods
	Visitor(
		processor Methodical,
	) VisitorLike
}

// INSTANCE DECLARATIONS

/*
FormatterLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Principal Methods
	GetClass() FormatterClassLike
	FormatAssembly(
		assembly ast.AssemblyLike,
	) string

	// Aspect Interfaces
	Methodical
}

/*
ParserLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Principal Methods
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.AssemblyLike
}

/*
ProcessorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Principal Methods
	GetClass() ProcessorClassLike

	// Aspect Interfaces
	Methodical
}

/*
ScannerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Principal Methods
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Principal Methods
	GetClass() TokenClassLike

	// Attribute Methods
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Principal Methods
	GetClass() ValidatorClassLike
	ValidateAssembly(
		assembly ast.AssemblyLike,
	)

	// Aspect Interfaces
	Methodical
}

/*
VisitorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Principal Methods
	GetClass() VisitorClassLike
	VisitAssembly(
		assembly ast.AssemblyLike,
	)
}

// ASPECT DECLARATIONS

/*
Methodical declares the set of method signatures that must be supported by
all methodical processors.
*/
type Methodical interface {
	ProcessCount(
		count string,
	)
	ProcessDelimiter(
		delimiter string,
	)
	ProcessExplanation(
		explanation string,
	)
	ProcessLabel(
		label string,
	)
	ProcessNewline(
		newline string,
	)
	ProcessQuoted(
		quoted string,
	)
	ProcessSpace(
		space string,
	)
	ProcessSymbol(
		symbol string,
	)
	PreprocessAction(
		action ast.ActionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAction(
		action ast.ActionLike,
		index_ uint,
		count_ uint,
	)
	ProcessActionSlot(
		action ast.ActionLike,
		slot_ uint,
	)
	PreprocessArgument(
		argument ast.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessArgument(
		argument ast.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessArgumentSlot(
		argument ast.ArgumentLike,
		slot_ uint,
	)
	PreprocessAssembly(
		assembly ast.AssemblyLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAssembly(
		assembly ast.AssemblyLike,
		index_ uint,
		count_ uint,
	)
	ProcessAssemblySlot(
		assembly ast.AssemblyLike,
		slot_ uint,
	)
	PreprocessCall(
		call ast.CallLike,
		index_ uint,
		count_ uint,
	)
	PostprocessCall(
		call ast.CallLike,
		index_ uint,
		count_ uint,
	)
	ProcessCallSlot(
		call ast.CallLike,
		slot_ uint,
	)
	PreprocessComponent(
		component ast.ComponentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessComponent(
		component ast.ComponentLike,
		index_ uint,
		count_ uint,
	)
	ProcessComponentSlot(
		component ast.ComponentLike,
		slot_ uint,
	)
	PreprocessCondition(
		condition ast.ConditionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessCondition(
		condition ast.ConditionLike,
		index_ uint,
		count_ uint,
	)
	ProcessConditionSlot(
		condition ast.ConditionLike,
		slot_ uint,
	)
	PreprocessConditionally(
		conditionally ast.ConditionallyLike,
		index_ uint,
		count_ uint,
	)
	PostprocessConditionally(
		conditionally ast.ConditionallyLike,
		index_ uint,
		count_ uint,
	)
	ProcessConditionallySlot(
		conditionally ast.ConditionallyLike,
		slot_ uint,
	)
	PreprocessConstant(
		constant ast.ConstantLike,
		index_ uint,
		count_ uint,
	)
	PostprocessConstant(
		constant ast.ConstantLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstantSlot(
		constant ast.ConstantLike,
		slot_ uint,
	)
	PreprocessContext(
		context ast.ContextLike,
		index_ uint,
		count_ uint,
	)
	PostprocessContext(
		context ast.ContextLike,
		index_ uint,
		count_ uint,
	)
	ProcessContextSlot(
		context ast.ContextLike,
		slot_ uint,
	)
	PreprocessDestination(
		destination ast.DestinationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessDestination(
		destination ast.DestinationLike,
		index_ uint,
		count_ uint,
	)
	ProcessDestinationSlot(
		destination ast.DestinationLike,
		slot_ uint,
	)
	PreprocessDrop(
		drop ast.DropLike,
		index_ uint,
		count_ uint,
	)
	PostprocessDrop(
		drop ast.DropLike,
		index_ uint,
		count_ uint,
	)
	ProcessDropSlot(
		drop ast.DropLike,
		slot_ uint,
	)
	PreprocessHandler(
		handler ast.HandlerLike,
		index_ uint,
		count_ uint,
	)
	PostprocessHandler(
		handler ast.HandlerLike,
		index_ uint,
		count_ uint,
	)
	ProcessHandlerSlot(
		handler ast.HandlerLike,
		slot_ uint,
	)
	PreprocessInstruction(
		instruction ast.InstructionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessInstruction(
		instruction ast.InstructionLike,
		index_ uint,
		count_ uint,
	)
	ProcessInstructionSlot(
		instruction ast.InstructionLike,
		slot_ uint,
	)
	PreprocessItem(
		item ast.ItemLike,
		index_ uint,
		count_ uint,
	)
	PostprocessItem(
		item ast.ItemLike,
		index_ uint,
		count_ uint,
	)
	ProcessItemSlot(
		item ast.ItemLike,
		slot_ uint,
	)
	PreprocessJump(
		jump ast.JumpLike,
		index_ uint,
		count_ uint,
	)
	PostprocessJump(
		jump ast.JumpLike,
		index_ uint,
		count_ uint,
	)
	ProcessJumpSlot(
		jump ast.JumpLike,
		slot_ uint,
	)
	PreprocessLiteral(
		literal ast.LiteralLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLiteral(
		literal ast.LiteralLike,
		index_ uint,
		count_ uint,
	)
	ProcessLiteralSlot(
		literal ast.LiteralLike,
		slot_ uint,
	)
	PreprocessLoad(
		load ast.LoadLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLoad(
		load ast.LoadLike,
		index_ uint,
		count_ uint,
	)
	ProcessLoadSlot(
		load ast.LoadLike,
		slot_ uint,
	)
	PreprocessNoop(
		noop ast.NoopLike,
		index_ uint,
		count_ uint,
	)
	PostprocessNoop(
		noop ast.NoopLike,
		index_ uint,
		count_ uint,
	)
	ProcessNoopSlot(
		noop ast.NoopLike,
		slot_ uint,
	)
	PreprocessNote(
		note ast.NoteLike,
		index_ uint,
		count_ uint,
	)
	PostprocessNote(
		note ast.NoteLike,
		index_ uint,
		count_ uint,
	)
	ProcessNoteSlot(
		note ast.NoteLike,
		slot_ uint,
	)
	PreprocessPrefix(
		prefix ast.PrefixLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPrefix(
		prefix ast.PrefixLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrefixSlot(
		prefix ast.PrefixLike,
		slot_ uint,
	)
	PreprocessPull(
		pull ast.PullLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPull(
		pull ast.PullLike,
		index_ uint,
		count_ uint,
	)
	ProcessPullSlot(
		pull ast.PullLike,
		slot_ uint,
	)
	PreprocessPush(
		push ast.PushLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPush(
		push ast.PushLike,
		index_ uint,
		count_ uint,
	)
	ProcessPushSlot(
		push ast.PushLike,
		slot_ uint,
	)
	PreprocessSave(
		save ast.SaveLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSave(
		save ast.SaveLike,
		index_ uint,
		count_ uint,
	)
	ProcessSaveSlot(
		save ast.SaveLike,
		slot_ uint,
	)
	PreprocessSend(
		send ast.SendLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSend(
		send ast.SendLike,
		index_ uint,
		count_ uint,
	)
	ProcessSendSlot(
		send ast.SendLike,
		slot_ uint,
	)
	PreprocessSource(
		source ast.SourceLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSource(
		source ast.SourceLike,
		index_ uint,
		count_ uint,
	)
	ProcessSourceSlot(
		source ast.SourceLike,
		slot_ uint,
	)
}
