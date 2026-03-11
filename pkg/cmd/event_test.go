// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/greenflash-public-api-cli/internal/mocktest"
)

func TestEventsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "events", "create",
			"--api-key", "string",
			"--event-type", "x",
			"--product-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--value", "value",
			"--conversation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--event-at", "'2019-12-27T18:11:19.117Z'",
			"--external-conversation-id", "externalConversationId",
			"--external-organization-id", "externalOrganizationId",
			"--external-user-id", "externalUserId",
			"--force-sample=true",
			"--influence", "positive",
			"--insert-id", "insertId",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--properties", "{foo: bar}",
			"--quality-impact-score", "-1",
			"--sample-rate", "0",
			"--user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--value-type", "currency",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"eventType: x\n" +
			"productId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"value: value\n" +
			"conversationId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"eventAt: '2019-12-27T18:11:19.117Z'\n" +
			"externalConversationId: externalConversationId\n" +
			"externalOrganizationId: externalOrganizationId\n" +
			"externalUserId: externalUserId\n" +
			"forceSample: true\n" +
			"influence: positive\n" +
			"insertId: insertId\n" +
			"organizationId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"properties:\n" +
			"  foo: bar\n" +
			"qualityImpactScore: -1\n" +
			"sampleRate: 0\n" +
			"userId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"valueType: currency\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "events", "create",
			"--api-key", "string",
		)
	})
}
