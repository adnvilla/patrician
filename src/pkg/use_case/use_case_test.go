package use_case_test

import (
	"context"
	"errors"
	"testing"

	"github.com/adnvilla/patrician/src/pkg/use_case"
	"github.com/stretchr/testify/assert"
)

// Define a mock implementation of UseCase for testing
type MockUseCase struct {
	ShouldFail bool
}

// Input and output types for our mock
type MockInput struct {
	Value string
}

type MockOutput struct {
	Result string
}

// Implement the Handle method required by the UseCase interface
func (m *MockUseCase) Handle(ctx context.Context, in MockInput) (MockOutput, error) {
	if m.ShouldFail {
		return MockOutput{}, errors.New("mock error")
	}
	return MockOutput{Result: "processed: " + in.Value}, nil
}

func TestUseCaseInterface(t *testing.T) {
	// Create context
	ctx := context.Background()

	// Test successful case
	t.Run("SuccessfulUseCase", func(t *testing.T) {
		// Create our mock that implements the UseCase interface
		mockUseCase := &MockUseCase{ShouldFail: false}

		// Execute the use case
		input := MockInput{Value: "test input"}
		output, err := mockUseCase.Handle(ctx, input)

		// Assert success
		assert.NoError(t, err)
		assert.Equal(t, "processed: test input", output.Result)
	})

	// Test error case
	t.Run("FailingUseCase", func(t *testing.T) {
		// Create our mock that implements the UseCase interface
		mockUseCase := &MockUseCase{ShouldFail: true}

		// Execute the use case
		input := MockInput{Value: "test input"}
		output, err := mockUseCase.Handle(ctx, input)

		// Assert failure
		assert.Error(t, err)
		assert.Equal(t, "mock error", err.Error())
		assert.Equal(t, "", output.Result)
	})

	// Test type constraints
	t.Run("TypeConstraints", func(t *testing.T) {
		// This is a compile-time test
		// The following line declares a variable of the interface type
		// which our MockUseCase must satisfy
		var _ use_case.UseCase[MockInput, MockOutput] = &MockUseCase{}

		// If the above line compiles, it means our MockUseCase properly implements
		// the UseCase interface with the correct type parameters
		assert.True(t, true, "Type constraints are satisfied")
	})
}
