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
