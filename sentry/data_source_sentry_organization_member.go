package sentry

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jianyuan/go-sentry/v2/sentry"
)

func dataSourceSentryOrganizationMember() *schema.Resource {
	return &schema.Resource{
		Description: "Sentry Organization member data source.",

		ReadContext: dataSourceSentryOrganizationMemberRead,

		Schema: map[string]*schema.Schema{
			"email": {
				Description: "The unique email address of the member.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"organization": {
				Description: "The slug of the organization the team should be created for.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Description: "The ID of this resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"internal_id": {
				Description: "The internal ID for this organization member.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"user": {
				Description: "The human readable user for this organization.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The human readable name for this organization.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"role": {
				Description: "The role of the organization member.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"teams": {
				Description: "The teams the organization member should be added to.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"expired": {
				Description: "The invite has expired.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"pending": {
				Description: "The invite is pending.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
}

func dataSourceSentryOrganizationMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*sentry.Client)

	org := d.Get("organization").(string)
	memberId := d.Get("id").(string)
	email := d.Get("email").(string)

	tflog.Debug(ctx, "Reading organization", map[string]interface{}{
		"org":   org,
		"id":    memberId,
		"email": email,
	})
	member, _, err := client.OrganizationMembers.Get(ctx, org, memberId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sentry.StringValue(&member.ID))
	retErr := multierror.Append(
		d.Set("organization", org),
		d.Set("name", member.Name),
		d.Set("internal_id", member.ID),
		d.Set("pending", member.Pending),
		d.Set("expired", member.Expired),
		d.Set("teams", member.Teams),
		d.Set("role", member.Role),
	)
	return diag.FromErr(retErr.ErrorOrNil())
}
