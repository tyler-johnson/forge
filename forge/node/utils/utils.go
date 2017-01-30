package utils

import "github.com/tyler-johnson/forge/forge/fcl/ast"

func MatchKey(node ast.Node, key string, mods []string) bool {
	var keynode *ast.Identifier

	switch item := node.(type) {
	case *ast.Verb:
		keynode = item.Key
	}

	if keynode.Value != key {
		return false
	}

	if len(keynode.Modifiers) != len(mods) {
		return false
	}

	for _, mod := range mods {
		if !keynode.Modifier(mod) {
			return false
		}
	}

	return true
}

func ExtractName(verb *ast.Verb) (string, bool) {
	if !verb.Values.IsEmpty() {
		typedec, ok := verb.Values.Items[0].(*ast.MethodCall)
		if !ok || typedec.HasArguments() || typedec.Key.HasModifiers() {
			return "", false
		}

		return typedec.Key.Value, true
	}

	return "", false
}
