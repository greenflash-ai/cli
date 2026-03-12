// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/greenflash-ai/cli/internal/mocktest"
)

func TestRatingsLog(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ratings", "log",
			"--api-key", "string",
			"--product-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--rating", "4",
			"--rating-max", "5",
			"--rating-min", "1",
			"--conversation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--external-conversation-id", "123e4567-e89b-12d3-a456-426614174000",
			"--external-message-id", "externalMessageId",
			"--feedback", "Helpful response!",
			"--message-id", "messageId",
			"--rated-at", "'2019-12-27'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"productId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"rating: 4\n" +
			"ratingMax: 5\n" +
			"ratingMin: 1\n" +
			"conversationId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"externalConversationId: 123e4567-e89b-12d3-a456-426614174000\n" +
			"externalMessageId: externalMessageId\n" +
			"feedback: Helpful response!\n" +
			"messageId: messageId\n" +
			"ratedAt: '2019-12-27'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ratings", "log",
			"--api-key", "string",
		)
	})
}
