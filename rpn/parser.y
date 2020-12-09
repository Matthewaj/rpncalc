%{
package rpn

import (
	"math"
)

func setResult(l yyLexer, value float64) {
  l.(*lex).result = value
}
%}

%union{
  val    float64
  result float64
}

%type <val> expr

// same for terminals
%token <val> DIGIT
%token sin cos tan
%start main

%%

main:
	expr
	{
		setResult(yylex, $1)
	}

expr:
	DIGIT
|	'(' expr ')'
	{
		$$ = $2
	}
|	expr expr '+'
	{
		$$ = $1 + $2
	}
| 	expr expr '-'
	{
		$$ = $1 - $2
	}
| 	expr expr '*'
	{
		$$ = $1 * $2
	}
| 	expr expr '/'
	{
		$$ = $1 / $2
	}
| 	expr sin
	{
		$$ = math.Sin($1)
	}
| 	expr tan
	{
		$$ = math.Tan($1)
	}
| 	expr cos
	{
		$$ = math.Cos($1)
	}
|	expr expr
	{
		$$ = $1 * $2
	}
