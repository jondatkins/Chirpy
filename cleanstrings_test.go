package main

import "testing"

func TestDirtyString(t *testing.T) {
	dirtyString := "This is a kerfuffle opinion I need to share with the world"
	cleanString := "This is a **** opinion I need to share with the world"
	resultString, _ := stripDirtyWords(dirtyString)
	if cleanString != resultString {
		t.Errorf(`Expect resultString '%s' to equal cleanString '%s'`, resultString, cleanString)
	}
	// if !isDirty {
	// 	t.Errorf("String %s should be dirty", dirtyString)
	// }
}

func TestCleanString(t *testing.T) {
	cleanString := "This is a kerfuffle! opinion I need to share with the world"
	resultString, _ := stripDirtyWords(cleanString)

	if cleanString != resultString {
		t.Errorf(`Expect resultString '%s' to equal cleanString '%s'`, resultString, cleanString)
	}
	// if isDirty {
	// 	t.Errorf("String %s is not dirty", cleanString)
	// }
}
