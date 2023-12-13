package evaluator

import (
	"Ahmadi/object"
	"fmt"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{
					Value: int64(len(arg.Value)),
				}

			case *object.Array:
				return &object.Integer{
					Value: int64(len(arg.Elements)),
				}

			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},

	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},

	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},

	"pop_front": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments in 'pop_front' function. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `pop_front` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{
					Elements: newElements,
				}
			}

			return NULL
		},
	},

	"pop_back": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments to 'pop_back' function. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `pop_back` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[:length-1])
				return &object.Array{
					Elements: newElements,
				}
			}

			return NULL
		},
	},

	"push_back": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments to 'push_back' function. got=%d, want=2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'push_back' must be ARRAY. got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			newElements := make([]object.Object, len(arr.Elements))
			copy(newElements, arr.Elements)
			newElements = append(newElements, args[1])
			return &object.Array{
				Elements: newElements,
			}
		},
	},

	"push_front": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments to 'push_front' function. got=%d, want=2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'push_front' must be ARRAY. got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			newElements := make([]object.Object, 0)
			newElements = append(newElements, args[1])
			newElements = append(newElements, arr.Elements...)
			return &object.Array{
				Elements: newElements,
			}
		},
	},

	"merge": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments to 'merge' function. it should be at least %d. got=%d", 2, len(args))
			}

			for _, arg := range args {
				if arg.Type() != object.ARRAY_OBJ {
					return newError("arguments to 'merge' must be ARRAY. got %s", arg.Type())
				}
			}

			newElements := make([]object.Object, 0)

			for _, arg := range args {
				newElements = append(newElements, arg.(*object.Array).Elements...)
			}

			return &object.Array{
				Elements: newElements,
			}
		},
	},

	"echo": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return &object.Integer{
				Value: int64(len(args)),
			}
		},
	},
}
