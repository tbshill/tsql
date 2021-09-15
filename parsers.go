package tsql

import . "github.com/tbshill/goparsec"

const (
	LineCommentLiteral = "--"
	BlockCommentStart  = "/*"
	BlockCommentEnd    = "*/"
)

var (
	ExpectComma                 = ExpectRune(',')
	ExpectLineComment           = ExpectString(LineCommentLiteral)
	ExpectMultiLineCommentStart = ExpectString(BlockCommentStart)
	ExpectMultiLineCommentEnd   = ExpectString(BlockCommentEnd)
	ExpectStar                  = ExpectRune('*')
	ExpectPlus                  = ExpectRune('+')
	ExpectMinus                 = ExpectRune('-')
	ExpectDiv                   = ExpectRune('/')
	ExpectSemicolon             = ExpectRune(';')
	ExpectSqLeft                = ExpectRune('[')
	ExpectSqRight               = ExpectRune(']')
	ExpectParenLeft             = ExpectRune('(')
	ExpectParenRight            = ExpectRune(')')
	ExpectCurlyLeft             = ExpectRune('{')
	ExpectCurlyRight            = ExpectRune('}')
	ExpectPeriod                = ExpectRune('.')
	ExpectUnderscore            = ExpectRune('_')
	ExpectEquals                = ExpectRune('=')
	ExpectValidIdentifierRunes  = Or(ExpectDigit, ExpectLetter, ExpectUnderscore)
	ExpectIdentifier            = And(ExpectLetter, Repeat(ExpectValidIdentifierRunes))
	ExpectOptionalWhitespace    = Optional(Repeat(ExpectWhiteSpace))
	ExpectInteger               = Repeat(ExpectDigit)
)
