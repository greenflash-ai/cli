// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/greenflash-ai/greenflash-cli/internal/mocktest"
)

func TestInteractionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "interactions", "list",
			"--api-key", "string",
			"--limit", "1",
			"--offset", "0",
			"--page", "1",
			"--product-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestInteractionsGetInteractionAnalytics(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "interactions", "get-interaction-analytics",
			"--api-key", "string",
			"--interaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--mode", "simple",
		)
	})
}
