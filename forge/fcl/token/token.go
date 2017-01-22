// Package token defines constants representing the lexical tokens for HCL
// (HashiCorp Configuration Language)
package token

import (
	"fmt"
	"strconv"
	"strings"

	hclstrconv "github.com/tyler-johnson/forge/forge/fcl/strconv"
)

// Token defines a single HCL token which can be obtained via the Scanner
type Token struct {
	Type Type
	Pos  Pos
	Text string
	JSON bool
}

// Type is the set of lexical tokens of the HCL (HashiCorp Configuration Language)
type Type int

// Special tokens
const (
	ILLEGAL Type = iota
	EOF
	COMMENT

	identifierBeg
	IDENT // literals
	literalBeg
	NULL    // null,nil
	NUMBER  // 123.45
	BOOL    // true,false
	STRING  // "abc"
	HEREDOC // <<FOO\nbar\nFOO
	literalEnd
	identifierEnd

	operatorBeg
	LBRACE // {
	LBRACK // [
	LPAREN // (
	RBRACE // }
	RBRACK // ]
	RPAREN // )

	PERIOD // .
	AT     // @
	BANG   // !
	DOLLAR // $

	operatorEnd
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:   "IDENT",
	NULL:    "NULL",
	NUMBER:  "NUMBER",
	BOOL:    "BOOL",
	STRING:  "STRING",
	HEREDOC: "HEREDOC",

	LBRACK: "LBRACK",
	LBRACE: "LBRACE",
	LPAREN: "LPAREN",
	RBRACK: "RBRACK",
	RBRACE: "RBRACE",
	RPAREN: "RPAREN",

	PERIOD: "PERIOD",
	AT:     "AT",
	BANG:   "BANG",
	DOLLAR: "DOLLAR",
}

// String returns the string corresponding to the token tok.
func (t Type) String() string {
	s := ""
	if 0 <= t && t < Type(len(tokens)) {
		s = tokens[t]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(t)) + ")"
	}
	return s
}

// IsIdentifier returns true for tokens corresponding to identifiers and basic
// type literals; it returns false otherwise.
func (t Type) IsIdentifier() bool { return identifierBeg < t && t < identifierEnd }

// IsLiteral returns true for tokens corresponding to basic type literals; it
// returns false otherwise.
func (t Type) IsLiteral() bool { return literalBeg < t && t < literalEnd }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
func (t Type) IsOperator() bool { return operatorBeg < t && t < operatorEnd }

// String returns the token's literal text. Note that this is only
// applicable for certain token types, such as token.IDENT,
// token.STRING, etc..
func (t Token) String() string {
	return fmt.Sprintf("%s %s %s", t.Pos.String(), t.Type.String(), t.Text)
}

// Value returns the properly typed value for this token. The type of
// the returned interface{} is guaranteed based on the Type field.
//
// This can only be called for literal types. If it is called for any other
// type, this will panic.
func (t Token) Value() interface{} {
	switch t.Type {
	case BOOL:
		if t.Text == "true" {
			return true
		} else if t.Text == "false" {
			return false
		}

		panic("unknown bool value: " + t.Text)
	case NUMBER:
		v, err := strconv.ParseFloat(t.Text, 64)
		if err != nil {
			panic(err)
		}

		return float64(v)
	case IDENT:
		return t.Text
	case HEREDOC:
		return unindentHeredoc(t.Text)
	case STRING:
		// Determine the Unquote method to use. If it came from JSON,
		// then we need to use the built-in unquote since we have to
		// escape interpolations there.
		f := hclstrconv.Unquote
		if t.JSON {
			f = strconv.Unquote
		}

		// This case occurs if json null is used
		if t.Text == "" {
			return ""
		}

		v, err := f(t.Text)
		if err != nil {
			panic(fmt.Sprintf("unquote %s err: %s", t.Text, err))
		}

		return v
	default:
		panic(fmt.Sprintf("unimplemented Value for type: %s", t.Type))
	}
}

// unindentHeredoc returns the string content of a HEREDOC if it is started with <<
// and the content of a HEREDOC with the hanging indent removed if it is started with
// a <<-, and the terminating line is at least as indented as the least indented line.
func unindentHeredoc(heredoc string) string {
	// We need to find the end of the marker
	idx := strings.IndexByte(heredoc, '\n')
	if idx == -1 {
		panic("heredoc doesn't contain newline")
	}

	unindent := heredoc[2] == '-'

	// We can optimize if the heredoc isn't marked for indentation
	if !unindent {
		return string(heredoc[idx+1 : len(heredoc)-idx+1])
	}

	// We need to unindent each line based on the indentation level of the marker
	lines := strings.Split(string(heredoc[idx+1:len(heredoc)-idx+2]), "\n")
	whitespacePrefix := lines[len(lines)-1]

	isIndented := true
	for _, v := range lines {
		if strings.HasPrefix(v, whitespacePrefix) {
			continue
		}

		isIndented = false
		break
	}

	// If all lines are not at least as indented as the terminating mark, return the
	// heredoc as is, but trim the leading space from the marker on the final line.
	if !isIndented {
		return strings.TrimRight(string(heredoc[idx+1:len(heredoc)-idx+1]), " \t")
	}

	unindentedLines := make([]string, len(lines))
	for k, v := range lines {
		if k == len(lines)-1 {
			unindentedLines[k] = ""
			break
		}

		unindentedLines[k] = strings.TrimPrefix(v, whitespacePrefix)
	}

	return strings.Join(unindentedLines, "\n")
}
