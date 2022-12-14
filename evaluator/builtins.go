package evaluator

import (
	"strings"

	"github.com/mpaliwoda/monkeylang/object"
)

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
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. want=1, got=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				if len(arg.Value) > 1 {
					return &object.String{Value: arg.Value[1 : len(arg.Value)]}
				}
				return &object.String{Value: ""}
			case *object.Array:
				if len(arg.Elements) > 0 {
					newElements := make([]object.Object, len(arg.Elements)-1, len(arg.Elements)-1)
					copy(newElements, arg.Elements[1:len(arg.Elements)])
					return &object.Array{Elements: newElements}
				}
				return &object.Array{Elements: []object.Object{}}
			default:
				return newError("argument to `last` not supported. got %s", args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. want>=2, got=%d", len(args))
			}

			switch dst:=args[0].(type) {
			case *object.String:
				elements := []string{dst.Value}
				for ix, arg := range args[1:] {
					switch arg := arg.(type) {
					case *object.String:
						elements = append(elements, arg.Value)
					default:
						return newError("argument %d not supported for dst type String. got %s",ix, arg.Type())
					}
				}
				return &object.String{Value: strings.Join(elements, "")}
			case *object.Array:
				newElements := make([]object.Object, len(dst.Elements))
				copy(newElements, dst.Elements)
				for _, arg := range args[1:] {
					newElements = append(newElements, arg)	
				}

				return &object.Array{Elements: newElements}
			default:
				return newError("first argument to `push` not supported. got %s", args[0].Type())
			}
		},
	},
}
