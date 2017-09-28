// Generated from LDE.g4 by ANTLR 4.7.

package parser // LDE

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LDEListener is a complete listener for a parse tree produced by LDEParser.
type LDEListener interface {
	antlr.ParseTreeListener

	// EnterRules is called when entering the rules production.
	EnterRules(c *RulesContext)

	// EnterAtomicRule is called when entering the atomicRule production.
	EnterAtomicRule(c *AtomicRuleContext)

	// EnterBaseAction is called when entering the baseAction production.
	EnterBaseAction(c *BaseActionContext)

	// EnterAtomicAction is called when entering the atomicAction production.
	EnterAtomicAction(c *AtomicActionContext)

	// EnterPassStringPrefix is called when entering the passStringPrefix production.
	EnterPassStringPrefix(c *PassStringPrefixContext)

	// EnterPassCharPrefix is called when entering the passCharPrefix production.
	EnterPassCharPrefix(c *PassCharPrefixContext)

	// EnterMayPassStringPrefix is called when entering the mayPassStringPrefix production.
	EnterMayPassStringPrefix(c *MayPassStringPrefixContext)

	// EnterMayPassCharPrefix is called when entering the mayPassCharPrefix production.
	EnterMayPassCharPrefix(c *MayPassCharPrefixContext)

	// EnterPassChars is called when entering the passChars production.
	EnterPassChars(c *PassCharsContext)

	// EnterPassUntil is called when entering the passUntil production.
	EnterPassUntil(c *PassUntilContext)

	// EnterMayPassUntil is called when entering the mayPassUntil production.
	EnterMayPassUntil(c *MayPassUntilContext)

	// EnterTakeUntil is called when entering the takeUntil production.
	EnterTakeUntil(c *TakeUntilContext)

	// EnterTakeUntilOrRest is called when entering the takeUntilOrRest production.
	EnterTakeUntilOrRest(c *TakeUntilOrRestContext)

	// EnterTakeUntilRest is called when entering the takeUntilRest production.
	EnterTakeUntilRest(c *TakeUntilRestContext)

	// EnterOptionalNamedArea is called when entering the optionalNamedArea production.
	EnterOptionalNamedArea(c *OptionalNamedAreaContext)

	// EnterOptionalArea is called when entering the optionalArea production.
	EnterOptionalArea(c *OptionalAreaContext)

	// EnterAtEnd is called when entering the atEnd production.
	EnterAtEnd(c *AtEndContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterTargetLit is called when entering the targetLit production.
	EnterTargetLit(c *TargetLitContext)

	// EnterBound is called when entering the bound production.
	EnterBound(c *BoundContext)

	// EnterLimit is called when entering the limit production.
	EnterLimit(c *LimitContext)

	// EnterExact is called when entering the exact production.
	EnterExact(c *ExactContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// ExitRules is called when exiting the rules production.
	ExitRules(c *RulesContext)

	// ExitAtomicRule is called when exiting the atomicRule production.
	ExitAtomicRule(c *AtomicRuleContext)

	// ExitBaseAction is called when exiting the baseAction production.
	ExitBaseAction(c *BaseActionContext)

	// ExitAtomicAction is called when exiting the atomicAction production.
	ExitAtomicAction(c *AtomicActionContext)

	// ExitPassStringPrefix is called when exiting the passStringPrefix production.
	ExitPassStringPrefix(c *PassStringPrefixContext)

	// ExitPassCharPrefix is called when exiting the passCharPrefix production.
	ExitPassCharPrefix(c *PassCharPrefixContext)

	// ExitMayPassStringPrefix is called when exiting the mayPassStringPrefix production.
	ExitMayPassStringPrefix(c *MayPassStringPrefixContext)

	// ExitMayPassCharPrefix is called when exiting the mayPassCharPrefix production.
	ExitMayPassCharPrefix(c *MayPassCharPrefixContext)

	// ExitPassChars is called when exiting the passChars production.
	ExitPassChars(c *PassCharsContext)

	// ExitPassUntil is called when exiting the passUntil production.
	ExitPassUntil(c *PassUntilContext)

	// ExitMayPassUntil is called when exiting the mayPassUntil production.
	ExitMayPassUntil(c *MayPassUntilContext)

	// ExitTakeUntil is called when exiting the takeUntil production.
	ExitTakeUntil(c *TakeUntilContext)

	// ExitTakeUntilOrRest is called when exiting the takeUntilOrRest production.
	ExitTakeUntilOrRest(c *TakeUntilOrRestContext)

	// ExitTakeUntilRest is called when exiting the takeUntilRest production.
	ExitTakeUntilRest(c *TakeUntilRestContext)

	// ExitOptionalNamedArea is called when exiting the optionalNamedArea production.
	ExitOptionalNamedArea(c *OptionalNamedAreaContext)

	// ExitOptionalArea is called when exiting the optionalArea production.
	ExitOptionalArea(c *OptionalAreaContext)

	// ExitAtEnd is called when exiting the atEnd production.
	ExitAtEnd(c *AtEndContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitTargetLit is called when exiting the targetLit production.
	ExitTargetLit(c *TargetLitContext)

	// ExitBound is called when exiting the bound production.
	ExitBound(c *BoundContext)

	// ExitLimit is called when exiting the limit production.
	ExitLimit(c *LimitContext)

	// ExitExact is called when exiting the exact production.
	ExitExact(c *ExactContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)
}