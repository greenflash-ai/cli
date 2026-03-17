// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/greenflash-ai/cli/internal/mocktest"
)

func TestOrganizationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "create",
			"--external-organization-id", "org-456",
			"--name", "Acme Corporation",
			"--properties", "{industry: bar, size: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"externalOrganizationId: org-456\n" +
			"name: Acme Corporation\n" +
			"properties:\n" +
			"  industry: bar\n" +
			"  size: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organizations", "create",
		)
	})
}

func TestOrganizationsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "update",
			"--organization-id", "organizationId",
			"--name", "Updated Organization Name",
			"--properties", "{industry: bar, size: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Updated Organization Name\n" +
			"properties:\n" +
			"  industry: bar\n" +
			"  size: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organizations", "update",
			"--organization-id", "organizationId",
		)
	})
}

func TestOrganizationsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "list",
			"--limit", "1",
			"--offset", "0",
			"--page", "1",
		)
	})
}

func TestOrganizationsGetOrganizationAnalytics(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "get-organization-analytics",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--mode", "simple",
			"--product-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
