package evaluator

import "github.com/mpaliwoda/monkeylang/object"

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. want=1, got=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported. got %s", args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. want=1, got=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				if len(arg.Value) > 0 {
					return &object.String{Value: arg.Value[0:1]}
				}
				return NULL
			case *object.Array:
				if len(arg.Elements) > 0 {
					return arg.Elements[0]
				}
				return NULL
			default:
				return newError("argument to `first` not supported. got %s", args[0].Type())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. want=1, got=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				if len(arg.Value) > 0 {
					return &object.String{Value: arg.Value[len(arg.Value)-1 : len(arg.Value)]}
				}
				return NULL
			case *object.Array:
				if len(arg.Elements) > 0 {
					return arg.Elements[len(arg.Elements)-1]
				}
				return NULL
			default:
				return newError("argument to `last` not supported. got %s", args[0].Type())
			}
		},
	},
}
