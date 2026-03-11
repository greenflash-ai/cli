// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/greenflash-public-api-cli/internal/mocktest"
)

func TestUsersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users", "create",
			"--api-key", "string",
			"--external-user-id", "user-123",
			"--anonymized=false",
			"--email", "alice@example.com",
			"--external-organization-id", "org-456",
			"--name", "Alice Example",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--phone", "+15551234567",
			"--properties", "{plan: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"externalUserId: user-123\n" +
			"anonymized: false\n" +
			"email: alice@example.com\n" +
			"externalOrganizationId: org-456\n" +
			"name: Alice Example\n" +
			"organizationId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"phone: '+15551234567'\n" +
			"properties:\n" +
			"  plan: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "users", "create",
			"--api-key", "string",
		)
	})
}

func TestUsersUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users", "update",
			"--api-key", "string",
			"--user-id", "userId",
			"--anonymized=true",
			"--email", "alice.updated@example.com",
			"--external-organization-id", "externalOrganizationId",
			"--name", "Alice Updated",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--phone", "phone",
			"--properties", "{plan: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"anonymized: true\n" +
			"email: alice.updated@example.com\n" +
			"externalOrganizationId: externalOrganizationId\n" +
			"name: Alice Updated\n" +
			"organizationId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"phone: phone\n" +
			"properties:\n" +
			"  plan: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "users", "update",
			"--api-key", "string",
			"--user-id", "userId",
		)
	})
}

func TestUsersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users", "list",
			"--api-key", "string",
			"--limit", "1",
			"--offset", "0",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page", "1",
		)
	})
}

func TestUsersGetUserAnalytics(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users", "get-user-analytics",
			"--api-key", "string",
			"--user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--mode", "simple",
			"--product-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
