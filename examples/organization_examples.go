package examples

import (
	"fmt"
	"github.com/jonnonz/zed"
)

func organizationExamples() {

	client := zd.NewClient("zendesk_domain", "email/token", "token", nil)

	// get Organization
	existingOrg, _, err := client.Organizations.GetOrganizationById("org_id")

	if err != nil {
		fmt.Println("Oh no!")
	}

	// add tags
	existingOrg.Tags = []string{"test_tags"}

	// Change Name
	existingOrg.Name = "Oh hello!"

	// Update organization fields
	existingOrg.OrganizationFields = map[string]string{
		"key": "value",
	}

	// update the organization, handle any errors etc..
	updatedOrganization , _ := client.Organizations.UpdateOrganization(existingOrg)

	fmt.Println(updatedOrganization.UpdatedAt)

	// create a new org
	newOrg := zd.Organization{
		Name: "brand new org",
	}

	response, _ := client.Organizations.CreateOrganization(&newOrg)

	// print the id of new org.
	fmt.Println(response.ID)
}
