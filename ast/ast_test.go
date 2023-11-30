package ast

import (
	"Ahmadi/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&DefStatement{
				Token: token.Token{
					Type:    token.DEF,
					Literal: "def",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.ID,
						Literal: "x",
					},
					Value: "x",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.ID,
						Literal: "y",
					},
					Value: "y",
				},
			},
		},
	}

	if program.String() != "def x = y;" {
		t.Errorf("program.String() wrong. got='%q'", program.String())
	}
}

func TestString2(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&DefStatement{
				Token: token.Token{
					Type:    token.DEF,
					Literal: "def",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.ID,
						Literal: "variable",
					},
					Value: "variable",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.INT,
						Literal: "t",
					},
					Value: "12",
				},
			},
			&ReturnStatement{
				Token: token.Token{
					Type:    token.RETURN,
					Literal: "return",
				},
				ReturnValue: &Identifier{
					Token: token.Token{
						Type:    token.ID,
						Literal: "variable",
					},
					Value: "variable",
				},
			},
		},
	}

	if program.String() != "def variable = 12;return variable;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
