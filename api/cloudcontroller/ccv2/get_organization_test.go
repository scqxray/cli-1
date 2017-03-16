// generated by gen_get_by_guid.go

package ccv2_test

import (
	"net/http"

	. "code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Organization", func() {
	var client *Client

	BeforeEach(func() {
		client = NewTestClient()
	})

	Describe("GetOrganization", func() {
		Context("when the Organization exists", func() {
			BeforeEach(func() {
				response := `{
					"metadata": {
						"guid": "some-organization-guid"
					},
					"entity": {
						"name": "some-organization"
					}
				}`
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v2/organizations/some-organization-guid"),
						RespondWith(http.StatusOK, response, http.Header{"X-Cf-Warnings": {"warning-1, warning-2"}}),
					))
			})

			It("returns the organization and all warnings", func() {
				organization, warnings, err := client.GetOrganization("some-organization-guid")

				Expect(err).ToNot(HaveOccurred())
				Expect(organization).To(Equal(
					Organization{
						GUID: "some-organization-guid",
						Name: "some-organization",
					},
				))
				Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
			})
		})

		Context("when an error is encountered", func() {
			BeforeEach(func() {
				response := `{
					"code": 10001,
					"description": "Some Error",
					"error_code": "CF-SomeError"
				}`
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v2/organizations/some-organization-guid"),
						RespondWith(http.StatusTeapot, response, http.Header{"X-Cf-Warnings": {"warning-1, warning-2"}}),
					))
			})

			It("returns an error and all warnings", func() {
				_, warnings, err := client.GetOrganization("some-organization-guid")

				Expect(err).To(MatchError(UnexpectedResponseError{
					ResponseCode: http.StatusTeapot,
					CCErrorResponse: CCErrorResponse{
						Code:        10001,
						Description: "Some Error",
						ErrorCode:   "CF-SomeError",
					},
				}))
				Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
			})
		})
	})
})