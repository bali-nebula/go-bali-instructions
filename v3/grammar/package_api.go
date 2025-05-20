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
		index uint,
		count uint,
	)
	ProcessActionSlot(
		slot uint,
	)
	PostprocessAction(
		action ast.ActionLike,
		index uint,
		count uint,
	)
	PreprocessArgument(
		argument ast.ArgumentLike,
		index uint,
		count uint,
	)
	ProcessArgumentSlot(
		slot uint,
	)
	PostprocessArgument(
		argument ast.ArgumentLike,
		index uint,
		count uint,
	)
	PreprocessAssembly(
		assembly ast.AssemblyLike,
		index uint,
		count uint,
	)
	ProcessAssemblySlot(
		slot uint,
	)
	PostprocessAssembly(
		assembly ast.AssemblyLike,
		index uint,
		count uint,
	)
	PreprocessCall(
		call ast.CallLike,
		index uint,
		count uint,
	)
	ProcessCallSlot(
		slot uint,
	)
	PostprocessCall(
		call ast.CallLike,
		index uint,
		count uint,
	)
	PreprocessComponent(
		component ast.ComponentLike,
		index uint,
		count uint,
	)
	ProcessComponentSlot(
		slot uint,
	)
	PostprocessComponent(
		component ast.ComponentLike,
		index uint,
		count uint,
	)
	PreprocessCondition(
		condition ast.ConditionLike,
		index uint,
		count uint,
	)
	ProcessConditionSlot(
		slot uint,
	)
	PostprocessCondition(
		condition ast.ConditionLike,
		index uint,
		count uint,
	)
	PreprocessConditionally(
		conditionally ast.ConditionallyLike,
		index uint,
		count uint,
	)
	ProcessConditionallySlot(
		slot uint,
	)
	PostprocessConditionally(
		conditionally ast.ConditionallyLike,
		index uint,
		count uint,
	)
	PreprocessConstant(
		constant ast.ConstantLike,
		index uint,
		count uint,
	)
	ProcessConstantSlot(
		slot uint,
	)
	PostprocessConstant(
		constant ast.ConstantLike,
		index uint,
		count uint,
	)
	PreprocessContext(
		context ast.ContextLike,
		index uint,
		count uint,
	)
	ProcessContextSlot(
		slot uint,
	)
	PostprocessContext(
		context ast.ContextLike,
		index uint,
		count uint,
	)
	PreprocessDestination(
		destination ast.DestinationLike,
		index uint,
		count uint,
	)
	ProcessDestinationSlot(
		slot uint,
	)
	PostprocessDestination(
		destination ast.DestinationLike,
		index uint,
		count uint,
	)
	PreprocessDrop(
		drop ast.DropLike,
		index uint,
		count uint,
	)
	ProcessDropSlot(
		slot uint,
	)
	PostprocessDrop(
		drop ast.DropLike,
		index uint,
		count uint,
	)
	PreprocessHandler(
		handler ast.HandlerLike,
		index uint,
		count uint,
	)
	ProcessHandlerSlot(
		slot uint,
	)
	PostprocessHandler(
		handler ast.HandlerLike,
		index uint,
		count uint,
	)
	PreprocessInstruction(
		instruction ast.InstructionLike,
		index uint,
		count uint,
	)
	ProcessInstructionSlot(
		slot uint,
	)
	PostprocessInstruction(
		instruction ast.InstructionLike,
		index uint,
		count uint,
	)
	PreprocessItem(
		item ast.ItemLike,
		index uint,
		count uint,
	)
	ProcessItemSlot(
		slot uint,
	)
	PostprocessItem(
		item ast.ItemLike,
		index uint,
		count uint,
	)
	PreprocessJump(
		jump ast.JumpLike,
		index uint,
		count uint,
	)
	ProcessJumpSlot(
		slot uint,
	)
	PostprocessJump(
		jump ast.JumpLike,
		index uint,
		count uint,
	)
	PreprocessLiteral(
		literal ast.LiteralLike,
		index uint,
		count uint,
	)
	ProcessLiteralSlot(
		slot uint,
	)
	PostprocessLiteral(
		literal ast.LiteralLike,
		index uint,
		count uint,
	)
	PreprocessLoad(
		load ast.LoadLike,
		index uint,
		count uint,
	)
	ProcessLoadSlot(
		slot uint,
	)
	PostprocessLoad(
		load ast.LoadLike,
		index uint,
		count uint,
	)
	PreprocessNoop(
		noop ast.NoopLike,
		index uint,
		count uint,
	)
	ProcessNoopSlot(
		slot uint,
	)
	PostprocessNoop(
		noop ast.NoopLike,
		index uint,
		count uint,
	)
	PreprocessNote(
		note ast.NoteLike,
		index uint,
		count uint,
	)
	ProcessNoteSlot(
		slot uint,
	)
	PostprocessNote(
		note ast.NoteLike,
		index uint,
		count uint,
	)
	PreprocessPrefix(
		prefix ast.PrefixLike,
		index uint,
		count uint,
	)
	ProcessPrefixSlot(
		slot uint,
	)
	PostprocessPrefix(
		prefix ast.PrefixLike,
		index uint,
		count uint,
	)
	PreprocessPull(
		pull ast.PullLike,
		index uint,
		count uint,
	)
	ProcessPullSlot(
		slot uint,
	)
	PostprocessPull(
		pull ast.PullLike,
		index uint,
		count uint,
	)
	PreprocessPush(
		push ast.PushLike,
		index uint,
		count uint,
	)
	ProcessPushSlot(
		slot uint,
	)
	PostprocessPush(
		push ast.PushLike,
		index uint,
		count uint,
	)
	PreprocessSave(
		save ast.SaveLike,
		index uint,
		count uint,
	)
	ProcessSaveSlot(
		slot uint,
	)
	PostprocessSave(
		save ast.SaveLike,
		index uint,
		count uint,
	)
	PreprocessSend(
		send ast.SendLike,
		index uint,
		count uint,
	)
	ProcessSendSlot(
		slot uint,
	)
	PostprocessSend(
		send ast.SendLike,
		index uint,
		count uint,
	)
	PreprocessSource(
		source ast.SourceLike,
		index uint,
		count uint,
	)
	ProcessSourceSlot(
		slot uint,
	)
	PostprocessSource(
		source ast.SourceLike,
		index uint,
		count uint,
	)
}
