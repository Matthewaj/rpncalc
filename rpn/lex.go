package rpn

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"strconv"
	"unicode"
)

//go:generate goyacc -l -o parser.go parser.y

const eof = 0
const errorCode = 256

func isParenthesis(token rune) bool {
	return token == ')' || token == '('
}

func Parse(input []byte) (float64, error) {
	l := newLex(input)
	_ = yyParse(l)
	return l.result, l.err
}

type lex struct {
	input  *bytes.Buffer
	result float64
	err    error
}

func newLex(input []byte) *lex {
	return &lex{
		input: bytes.NewBuffer(input),
	}
}

func (l *lex) Lex(lval *yySymType) int {
	token := l.nextToken()

	switch {
	case token == eof:
		return eof
	case unicode.IsLetter(token):
		identifier, err := l.nextIdentifier()
		if err != nil {
			l.err = err
			return errorCode
		}

		return identifier
	case unicode.IsDigit(token):
		val, err := l.nextDigit()
		if err != nil {
			l.err = err
			return errorCode
		}
		lval.val = val
		return DIGIT
	}
	return int(token)
}

func (l *lex) nextToken() rune {
	b, err := l.input.ReadByte()
	if err != nil {
		return 0
	}

	token := rune(b)
	for unicode.IsSpace(token) || isParenthesis(token) {
		token = l.nextToken()
	}
	return token
}

func (l *lex) nextDigit() (float64, error) {
	_ = l.input.UnreadByte()
	b, err := l.input.ReadByte()
	if err != nil {
		return 0, errors.Wrap(err, "no digit to read")
	}
	digit := string(rune(b))

	for {
		b, err = l.input.ReadByte()
		char := rune(b)
		if err == io.EOF {
			f, err := strconv.ParseFloat(digit, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid format for a digit: %s", digit)
			}
			return f, nil
		}

		if char != '.' && !unicode.IsDigit(char) {
			_ = l.input.UnreadByte()
			break
		}
		digit = digit + string(char)
	}

	f, err := strconv.ParseFloat(digit, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid format for a digit: %s", digit)
	}
	return f, nil
}

func (l *lex) nextIdentifier() (int, error) {
	_ = l.input.UnreadByte()
	identifier := l.nextIdent()

	switch identifier {
	case "sin":
		return sin, nil
	case "cos":
		return cos, nil
	case "tan":
		return tan, nil
	default:
		return 0, fmt.Errorf("undefined token: %s", identifier)
	}
}

func (l *lex) nextIdent() string {
	b, err := l.input.ReadByte()
	if err != nil {
		return ""
	}
	ident := string(rune(b))

	for {
		ch, err := l.input.ReadByte()
		if err == io.EOF {
			return ident
		}

		if !unicode.IsLetter(rune(ch)) && !unicode.IsNumber(rune(ch)) {
			_ = l.input.UnreadByte()
			break
		} else {
			ident = ident + string(rune(ch))
		}
	}
	return ident
}

// Error satisfies yyLexer.
func (l *lex) Error(s string) {
	if l.err != nil {
		l.err = errors.Wrap(l.err, s)
	} else {
		l.err = errors.New(s)
	}

}
