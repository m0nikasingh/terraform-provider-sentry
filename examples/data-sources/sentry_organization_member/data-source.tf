# Retrieve a team
data "sentry_organization_member" "test" {
  organization = "my-organization"
  email = "test@example.com"
  role = "member"
  teams = ["my-team"]
}
