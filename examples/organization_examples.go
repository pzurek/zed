package examples

import (
	"fmt"

	"github.com/pzurek/zed"
)

func organizationExamples() {

	client := zed.NewClient("zendesk_domain", "email/token", "token", nil)

	// get Organization
	existingOrg, _, err := client.Organizations.Get("org_id")

	if err != nil {
		fmt.Println("Oh no!")
	}

	// add tags
	existingOrg.Tags = []string{"test_tags"}

	// Change Name
	name := "Oh hello"
	existingOrg.Name = &name

	// Update organization fields
	existingOrg.OrganizationFields = map[string]string{
		"key": "value",
	}

	// update the organization, handle any errors etc..
	updatedOrganization, _ := client.Organizations.Update(existingOrg)

	fmt.Println(updatedOrganization.UpdatedAt)

	// create a new org
	orgName := "brand new org"
	newOrg := zed.Organization{
		Name: &orgName,
	}

	response, _ := client.Organizations.Create(&newOrg)

	// print the id of new org.
	fmt.Println(response.ID)
}
