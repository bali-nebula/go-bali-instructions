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
	doc "github.com/bali-nebula/go-bali-documents/v3/documents"
	not "github.com/bali-nebula/go-document-notation/v3"
	com "github.com/craterdog/go-essential-composites/v8"
	pri "github.com/craterdog/go-essential-primitives/v8"
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

func (v *deflator_) DeflateDocument(
	document doc.DocumentLike,
) not.DocumentLike {
	VisitorClass().Visitor(v).VisitDocument(document)
	if v.stack_.GetSize() != 1 {
		var message = fmt.Sprintf(
			"Internal Error: the deflator stack is corrupted: %v",
			v.stack_,
		)
		panic(message)
	}
	return v.stack_.RemoveLast().(not.DocumentLike)
}

// Attribute Methods

// Methodical Methods

func (v *deflator_) ProcessAngle(
	angle pri.AngleLike,
) {
	v.stack_.AddValue(not.Element(angle.AsSource()))
}

func (v *deflator_) ProcessAnnotation(
	annotation string,
) {
	v.stack_.AddValue(not.Annotation(annotation))
}

func (v *deflator_) ProcessAssignment(
	assignment doc.Assignment,
) {
	var operation string
	switch assignment {
	case doc.DefaultEquals:
		operation = "?="
	case doc.AssignEquals:
		operation = ":="
	case doc.PlusEquals:
		operation = "+="
	case doc.MinusEquals:
		operation = "-="
	case doc.TimesEquals:
		operation = "*="
	case doc.DivideEquals:
		operation = "/="
	case doc.ChainEquals:
		operation = "&="
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			assignment,
			assignment,
		)
		panic(message)
	}
	v.stack_.AddValue(not.Assignment(operation))
}

func (v *deflator_) ProcessBinary(
	binary pri.BinaryLike,
) {
	v.stack_.AddValue(not.Sequence(binary.AsSource()))
}

func (v *deflator_) ProcessBoolean(
	boolean pri.BooleanLike,
) {
	v.stack_.AddValue(not.Element(boolean.AsSource()))
}

func (v *deflator_) ProcessBytecode(
	bytecode pri.BytecodeLike,
) {
	v.stack_.AddValue(not.Sequence(bytecode.AsSource()))
}

func (v *deflator_) ProcessComment(
	comment string,
) {
	v.stack_.AddValue(comment)
}

func (v *deflator_) ProcessDuration(
	duration pri.DurationLike,
) {
	v.stack_.AddValue(not.Element(duration.AsSource()))
}

func (v *deflator_) ProcessGlyph(
	glyph pri.GlyphLike,
) {
	v.stack_.AddValue(not.Element(glyph.AsSource()))
}

func (v *deflator_) ProcessIdentifier(
	identifier string,
) {
	v.stack_.AddValue(identifier)
}

func (v *deflator_) ProcessInverse(
	inverse doc.Inverse,
) {
	var operation string
	switch inverse {
	case doc.Additive:
		operation = "-"
	case doc.Multiplicative:
		operation = "/"
	case doc.Conjugate:
		operation = "*"
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			inverse,
			inverse,
		)
		panic(message)
	}
	v.stack_.AddValue(not.Inverse(operation))
}

func (v *deflator_) ProcessInvoke(
	invoke doc.Invoke,
) {
	var operation string
	switch invoke {
	case doc.Synchronous:
		operation = "<-"
	case doc.Asynchronous:
		operation = "<~"
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			invoke,
			invoke,
		)
		panic(message)
	}
	v.stack_.AddValue(not.Invoke(operation))
}

func (v *deflator_) ProcessMoment(
	moment pri.MomentLike,
) {
	v.stack_.AddValue(not.Element(moment.AsSource()))
}

func (v *deflator_) ProcessName(
	name pri.NameLike,
) {
	v.stack_.AddValue(not.Sequence(name.AsSource()))
}

func (v *deflator_) ProcessNarrative(
	narrative pri.NarrativeLike,
) {
	v.stack_.AddValue(not.Sequence(narrative.AsSource()))
}

func (v *deflator_) ProcessNote(
	note string,
) {
	v.stack_.AddValue(note)
}

func (v *deflator_) ProcessNumber(
	number pri.NumberLike,
) {
	v.stack_.AddValue(not.Element(number.AsSource()))
}

func (v *deflator_) ProcessOperator(
	operator doc.Operator,
) {
	var operation any
	switch operator {
	case doc.Less:
		operation = not.Comparison("<")
	case doc.Equal:
		operation = not.Comparison("=")
	case doc.More:
		operation = not.Comparison(">")
	case doc.Is:
		operation = not.Comparison("is")
	case doc.Matches:
		operation = not.Comparison("matches")
	case doc.And:
		operation = not.Logical("and")
	case doc.San:
		operation = not.Logical("san")
	case doc.Ior:
		operation = not.Logical("ior")
	case doc.Xor:
		operation = not.Logical("xor")
	case doc.Plus:
		operation = not.Arithmetic("+")
	case doc.Minus:
		operation = not.Arithmetic("-")
	case doc.Times:
		operation = not.Arithmetic("*")
	case doc.Divide:
		operation = not.Arithmetic("/")
	case doc.Remainder:
		operation = not.Arithmetic("%")
	case doc.Power:
		operation = not.Arithmetic("^")
	case doc.Chain:
		operation = not.Lexical("&")
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			operator,
			operator,
		)
		panic(message)
	}
	v.stack_.AddValue(not.Operator(operation))
}

func (v *deflator_) ProcessPattern(
	pattern pri.PatternLike,
) {
	v.stack_.AddValue(not.Sequence(pattern.AsSource()))
}

func (v *deflator_) ProcessPercentage(
	percentage pri.PercentageLike,
) {
	v.stack_.AddValue(not.Element(percentage.AsSource()))
}

func (v *deflator_) ProcessProbability(
	probability pri.ProbabilityLike,
) {
	v.stack_.AddValue(not.Element(probability.AsSource()))
}

func (v *deflator_) ProcessQuote(
	quote pri.QuoteLike,
) {
	v.stack_.AddValue(not.Sequence(quote.AsSource()))
}

func (v *deflator_) ProcessResource(
	resource pri.ResourceLike,
) {
	v.stack_.AddValue(not.Element(resource.AsSource()))
}

func (v *deflator_) ProcessSymbol(
	symbol pri.SymbolLike,
) {
	v.stack_.AddValue(not.Sequence(symbol.AsSource()))
}

func (v *deflator_) ProcessTag(
	tag pri.TagLike,
) {
	v.stack_.AddValue(not.Sequence(tag.AsSource()))
}

func (v *deflator_) ProcessVersion(
	version pri.VersionLike,
) {
	v.stack_.AddValue(not.Sequence(version.AsSource()))
}

func (v *deflator_) PostprocessAcceptClause(
	acceptClause doc.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var bag = not.Bag(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.AcceptClause("accept", message, "from", bag),
		),
	)
}

func (v *deflator_) PostprocessAttributes(
	attributes doc.AttributesLike,
	index_ uint,
	count_ uint,
) {
	var associations = com.List[not.AssociationLike]()
	var iterator = attributes.GetAssociations().GetIterator()
	for iterator.HasNext() {
		var content = v.stack_.RemoveLast().(not.ContentLike)
		var primitive = not.Primitive(v.stack_.RemoveLast())
		var association = not.Association(primitive, ":", content)
		associations.AppendValue(association)
		iterator.GetNext()
	}
	associations.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Collection(not.Attributes("[", associations, "]")))
}

func (v *deflator_) PostprocessBreakClause(
	breakClause doc.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(
		not.FlowControl(
			not.BreakClause("break", "loop"),
		),
	)
}

func (v *deflator_) ProcessCheckoutClauseSlot(
	checkoutClause doc.CheckoutClauseLike,
	slot_ uint,
) {
	switch slot_ {
	case 2:
		if uti.IsUndefined(checkoutClause.GetOptionalAtLevel()) {
			v.stack_.AddValue(nil)
		}
	}
}

func (v *deflator_) PostprocessCheckoutClause(
	checkoutClause doc.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var location = not.Location(expression)
	var atLevel not.AtLevelLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		var expression = optional.(not.ExpressionLike)
		atLevel = not.AtLevel("at", "level", expression)
	}
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.SequenceLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.CheckoutClause("checkout", recipient, atLevel, "from", location),
		),
	)
}

func (v *deflator_) PostprocessComplement(
	complement doc.ComplementLike,
	index_ uint,
	count_ uint,
) {
	var reversible not.ReversibleLike
	switch actual := v.stack_.RemoveLast().(type) {
	case string:
		reversible = not.Reversible(not.Value(actual))
	default:
		reversible = not.Reversible(actual)
	}
	v.stack_.AddValue(
		not.Complement("not", reversible),
	)
}

func (v *deflator_) ProcessCompositeSlot(
	composite doc.Composite,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(composite.GetOptionalGenerics()) {
			var generics not.GenericsLike
			v.stack_.AddValue(generics)
		}
	}
}

func (v *deflator_) PostprocessComposite(
	composite doc.Composite,
	index_ uint,
	count_ uint,
) {
	var generics not.GenericsLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		generics = optional.(not.GenericsLike)
	}
	var entity = not.Entity(v.stack_.RemoveLast())
	v.stack_.AddValue(
		not.Component(
			entity,
			generics,
		),
	)
}

func (v *deflator_) ProcessContentSlot(
	content doc.ContentLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(content.GetOptionalNote()) {
			var note string
			v.stack_.AddValue(note)
		}
	}
}

func (v *deflator_) PostprocessContent(
	content doc.ContentLike,
	index_ uint,
	count_ uint,
) {
	var note = v.stack_.RemoveLast().(string)
	var component = v.stack_.RemoveLast().(not.ComponentLike)
	v.stack_.AddValue(
		not.Content(
			component,
			note,
		),
	)
}

func (v *deflator_) ProcessConstraintSlot(
	constraint doc.ConstraintLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(constraint.GetOptionalGenerics()) {
			var generics not.GenericsLike
			v.stack_.AddValue(generics)
		}
	}
}

func (v *deflator_) PostprocessConstraint(
	constraint doc.ConstraintLike,
	index_ uint,
	count_ uint,
) {
	var generics not.GenericsLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		generics = optional.(not.GenericsLike)
	}
	var metadata = not.Metadata(v.stack_.RemoveLast())
	v.stack_.AddValue(
		not.Constraint(
			metadata,
			generics,
		),
	)
}

func (v *deflator_) PostprocessContinueClause(
	continueClause doc.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(
		not.FlowControl(
			not.ContinueClause("continue", "loop"),
		),
	)
}

func (v *deflator_) PostprocessDiscardClause(
	discardClause doc.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var citation = not.Citation(expression)
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.DiscardClause("discard", citation),
		),
	)
}

func (v *deflator_) PostprocessDoClause(
	doClause doc.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	var method = v.stack_.RemoveLast().(not.MethodLike)
	v.stack_.AddValue(
		not.ActionInduction(
			not.DoClause("do", method),
		),
	)
}

func (v *deflator_) PreprocessDocument(
	document doc.DocumentLike,
	index_ uint,
	count_ uint,
) {
	var comment = document.GetOptionalComment()
	if uti.IsUndefined(comment) {
		// We only add it if it is not defined, otherwise ProcessComment adds it.
		v.stack_.AddValue(comment)
	}
}

func (v *deflator_) PostprocessDocument(
	document doc.DocumentLike,
	index_ uint,
	count_ uint,
) {
	var component = v.stack_.RemoveLast().(not.ComponentLike)
	var heading not.HeadingLike
	var comment = v.stack_.RemoveLast().(string)
	if uti.IsDefined(comment) {
		heading = not.Heading(comment)
	}
	v.stack_.AddValue(
		not.Document(
			heading,
			component,
		),
	)
}

func (v *deflator_) PostprocessExpression(
	expression doc.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	var predicates = com.List[not.PredicateLike]()
	var iterator = expression.GetPredicates().GetIterator()
	for iterator.HasNext() {
		var predicate = v.stack_.RemoveLast().(not.PredicateLike)
		predicates.AppendValue(predicate)
		iterator.GetNext()
	}
	predicates.ReverseValues() // They were pulled off the stack in reverse order.
	var subject not.SubjectLike
	switch actual := v.stack_.RemoveLast().(type) {
	case string:
		subject = not.Subject(not.Value(actual))
	default:
		subject = not.Subject(actual)
	}
	v.stack_.AddValue(not.Expression(subject, predicates))
}

func (v *deflator_) PostprocessFunction(
	function doc.FunctionLike,
	index_ uint,
	count_ uint,
) {
	var arguments = com.List[not.ArgumentLike]()
	var iterator = function.GetArguments().GetIterator()
	for iterator.HasNext() {
		var argument not.ArgumentLike
		switch actual := v.stack_.RemoveLast().(type) {
		case string:
			argument = not.Argument(not.Value(actual))
		default:
			argument = not.Argument(not.Primitive(actual))
		}
		arguments.AppendValue(argument)
		iterator.GetNext()
	}
	arguments.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(not.Function(identifier, "(", arguments, ")"))
}

func (v *deflator_) PostprocessIfClause(
	ifClause doc.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.FlowControl(
			not.IfClause("if", expression, "do", procedure),
		),
	)
}

func (v *deflator_) PostprocessInversion(
	inversion doc.InversionLike,
	index_ uint,
	count_ uint,
) {
	var numerical not.NumericalLike
	switch actual := v.stack_.RemoveLast().(type) {
	case string:
		numerical = not.Numerical(not.Value(actual))
	default:
		numerical = not.Numerical(actual)
	}
	var inverse = v.stack_.RemoveLast().(not.InverseLike)
	v.stack_.AddValue(
		not.Inversion(
			inverse,
			numerical,
		),
	)
}

func (v *deflator_) PostprocessItems(
	items doc.ItemsLike,
	index_ uint,
	count_ uint,
) {
	var contents = com.List[not.ContentLike]()
	var iterator = items.GetContents().GetIterator()
	for iterator.HasNext() {
		var content = v.stack_.RemoveLast().(not.ContentLike)
		contents.AppendValue(content)
		iterator.GetNext()
	}
	contents.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Collection(not.Items("[", contents, "]")))
}

func (v *deflator_) PostprocessLetClause(
	letClause doc.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var assignment = v.stack_.RemoveLast().(not.AssignmentLike)
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.SequenceLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.ActionInduction(
			not.LetClause("let", recipient, assignment, expression),
		),
	)
}

func (v *deflator_) PostprocessMagnitude(
	magnitude doc.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.Magnitude("|", expression, "|"),
	)
}

func (v *deflator_) PostprocessMatchingClause(
	matchingClause doc.MatchingClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.MatchingClause("matching", expression, "do", procedure),
	)
}

func (v *deflator_) PostprocessMethod(
	method doc.MethodLike,
	index_ uint,
	count_ uint,
) {
	var arguments = com.List[not.ArgumentLike]()
	var iterator = method.GetArguments().GetIterator()
	for iterator.HasNext() {
		var argument not.ArgumentLike
		switch actual := v.stack_.RemoveLast().(type) {
		case string:
			argument = not.Argument(not.Value(actual))
		default:
			argument = not.Argument(not.Primitive(actual))
		}
		arguments.AppendValue(argument)
		iterator.GetNext()
	}
	arguments.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	var invoke = v.stack_.RemoveLast().(not.InvokeLike)
	var target = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(not.Method(target, invoke, identifier, "(", arguments, ")"))
}

func (v *deflator_) PostprocessNotarizeClause(
	notarizeClause doc.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var location = not.Location(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var draft = not.Draft(expression)
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.NotarizeClause("notarize", draft, "as", location),
		),
	)
}

func (v *deflator_) PostprocessOnClause(
	onClause doc.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	var matchingClauses = com.List[not.MatchingClauseLike]()
	var iterator = onClause.GetMatchingClauses().GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(not.MatchingClauseLike)
		matchingClauses.AppendValue(matchingClause)
		iterator.GetNext()
	}
	matchingClauses.ReverseValues() // They were pulled off the stack in reverse order.
	var sequence = v.stack_.RemoveLast().(not.SequenceLike)
	var symbol = sequence.GetAny().(string)
	v.stack_.AddValue(
		not.OnClause("on", symbol, matchingClauses),
	)
}

func (v *deflator_) PostprocessGenerics(
	generics doc.GenericsLike,
	index_ uint,
	count_ uint,
) {
	var parameters = com.List[not.ParameterLike]()
	var iterator = generics.GetParameters().GetIterator()
	for iterator.HasNext() {
		var constraint = v.stack_.RemoveLast().(not.ConstraintLike)
		var sequence = v.stack_.RemoveLast().(not.SequenceLike)
		var symbol = sequence.GetAny().(string)
		var parameter = not.Parameter(symbol, ":", constraint)
		parameters.AppendValue(parameter)
		iterator.GetNext()
	}
	parameters.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Generics("(", parameters, ")"))
}

func (v *deflator_) PostprocessPrecedence(
	precedence doc.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.Precedence("(", expression, ")"),
	)
}

func (v *deflator_) PostprocessPredicate(
	predicate doc.PredicateLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var operator = v.stack_.RemoveLast().(not.OperatorLike)
	v.stack_.AddValue(
		not.Predicate(operator, expression),
	)
}

func (v *deflator_) PostprocessProcedure(
	procedure doc.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	var lines = com.List[not.LineLike]()
	var iterator = procedure.GetLines().GetIterator()
	for iterator.HasNext() {
		var line = not.Line(v.stack_.RemoveLast())
		lines.AppendValue(line)
		iterator.GetNext()
	}
	lines.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Procedure("{", lines, "}"))
}

func (v *deflator_) PostprocessPublishClause(
	publishClause doc.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.PublishClause("publish", message),
		),
	)
}

func (v *deflator_) PostprocessRange(
	range_ doc.RangeLike,
	index_ uint,
	count_ uint,
) {
	var right not.RightLike
	var extent = range_.GetRight()
	switch extent {
	case com.Inclusive:
		right = not.Right("]")
	case com.Exclusive:
		right = not.Right(")")
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			extent,
			extent,
		)
		panic(message)
	}
	var primitive2 = not.Primitive(v.stack_.RemoveLast())
	var primitive1 = not.Primitive(v.stack_.RemoveLast())
	var left not.LeftLike
	extent = range_.GetLeft()
	switch extent {
	case com.Inclusive:
		left = not.Left("[")
	case com.Exclusive:
		left = not.Left("(")
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			extent,
			extent,
		)
		panic(message)
	}
	v.stack_.AddValue(
		not.Range(left, primitive1, "..", primitive2, right),
	)
}

func (v *deflator_) PostprocessReceiveClause(
	receiveClause doc.ReceiveClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var bag = not.Bag(expression)
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.SequenceLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.MessageHandling(
			not.ReceiveClause("receive", recipient, "from", bag),
		),
	)
}

func (v *deflator_) PostprocessReferent(
	referent doc.ReferentLike,
	index_ uint,
	count_ uint,
) {
	var reference not.ReferenceLike
	switch actual := v.stack_.RemoveLast().(type) {
	case string:
		reference = not.Reference(not.Value(actual))
	default:
		reference = not.Reference(actual)
	}
	v.stack_.AddValue(not.Referent("@", reference))
}

func (v *deflator_) PostprocessRejectClause(
	rejectClause doc.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var bag = not.Bag(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.RejectClause("reject", message, "from", bag),
		),
	)
}

func (v *deflator_) PostprocessReturnClause(
	returnClause doc.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.FlowControl(
			not.ReturnClause("return", expression),
		),
	)
}

func (v *deflator_) PostprocessSaveClause(
	saveClause doc.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.SequenceLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var draft = not.Draft(expression)
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.SaveClause("save", draft, "to", recipient),
		),
	)
}

func (v *deflator_) PostprocessSelectClause(
	selectClause doc.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	var matchingClauses = com.List[not.MatchingClauseLike]()
	var iterator = selectClause.GetMatchingClauses().GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(not.MatchingClauseLike)
		matchingClauses.AppendValue(matchingClause)
		iterator.GetNext()
	}
	matchingClauses.ReverseValues() // They were pulled off the stack in reverse order.
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.FlowControl(
			not.SelectClause("select", expression, matchingClauses),
		),
	)
}

func (v *deflator_) PostprocessSendClause(
	sendClause doc.SendClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var bag = not.Bag(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.SendClause("send", message, "to", bag),
		),
	)
}

func (v *deflator_) ProcessStatementSlot(
	statement doc.StatementLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(statement.GetOptionalOnClause()) {
			v.stack_.AddValue(nil)
		}
	}
}

func (v *deflator_) PostprocessStatement(
	statement doc.StatementLike,
	index_ uint,
	count_ uint,
) {
	var onClause not.OnClauseLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		onClause = optional.(not.OnClauseLike)
	}
	var mainClause = not.MainClause(v.stack_.RemoveLast())
	v.stack_.AddValue(not.Statement(mainClause, onClause))
}

func (v *deflator_) PostprocessSubcomponent(
	subcomponent doc.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	var indexes = com.List[not.IndexLike]()
	var iterator = subcomponent.GetIndexes().GetIterator()
	for iterator.HasNext() {
		var index not.IndexLike
		switch actual := v.stack_.RemoveLast().(type) {
		case string:
			index = not.Index(not.Value(actual))
		default:
			index = not.Index(not.Primitive(actual))
		}
		indexes.AppendValue(index)
		iterator.GetNext()
	}
	indexes.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(not.Subcomponent(identifier, "[", indexes, "]"))
}

func (v *deflator_) PostprocessThrowClause(
	throwClause doc.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.FlowControl(
			not.ThrowClause("throw", expression),
		),
	)
}

func (v *deflator_) PostprocessWhileClause(
	whileClause doc.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.FlowControl(
			not.WhileClause("while", expression, "do", procedure),
		),
	)
}

func (v *deflator_) PostprocessWithClause(
	withClause doc.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var sequence = v.stack_.RemoveLast().(not.SequenceLike)
	var symbol = sequence.GetAny().(string)
	v.stack_.AddValue(
		not.FlowControl(
			not.WithClause(
				"with",
				"each",
				symbol,
				"in",
				expression,
				"do",
				procedure,
			),
		),
	)
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
