package belajargolangcontext

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	// Context Backeground
	background := context.Background()
	fmt.Println(background)

	// Context TODO
	todo := context.TODO()
	fmt.Println(todo)

}
