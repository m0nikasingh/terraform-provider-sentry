package sentry

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/jianyuan/go-sentry/v2/sentry"
)

func TestAccSentryOrganizationMemberDataSource_basic(t *testing.T) {
	var member sentry.OrganizationMember

	dn := "data.sentry_organization_member.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSentryOrganizationMemberDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSentryOrganizationMemberExists(dn, &member),
					resource.TestCheckResourceAttr(dn, "organization", testOrganization),
					resource.TestCheckResourceAttr(dn, "email", member.Email),
					resource.TestCheckResourceAttr(dn, "role", member.Role),
					resource.TestCheckResourceAttrSet(dn, "internal_id"),
					resource.TestCheckResourceAttrSet(dn, "pending"),
					resource.TestCheckResourceAttrSet(dn, "expired"),
				),
			},
		},
	})
}

var testAccSentryOrganizationMemberDataSourceConfig = fmt.Sprintf(`
data "sentry_organization_member" "test" {
	organization = "%[1]s"
	email = "%[2]s"
}
`, testOrganization, testMemberEmail)
