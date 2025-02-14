package port

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccPortEntityUpdateProp(t *testing.T) {
	identifier := genID()
	var testAccActionConfigCreate = fmt.Sprintf(`
	resource "port-labs_blueprint" "tf_bp" {
		title = "TF Provider Test"
		icon = "Terraform"
		identifier = "%s"
		properties {
			identifier = "text"
			type = "string"
			title = "text"
		}
	}
	resource "port-labs_entity" "microservice" {
		title = "monolith"
		blueprint = port-labs_blueprint.tf_bp.id
		team = "Everyone"
		properties {
			name = "text"
			value = "hedwig"
		}
	}
`, identifier)
	var testAccActionConfigUpdate = fmt.Sprintf(`
	resource "port-labs_blueprint" "tf_bp" {
		title = "TF Provider Test"
		icon = "Terraform"
		identifier = "%s"
		properties {
			identifier = "text"
			type = "string"
			title = "text"
		}
	}
	resource "port-labs_entity" "microservice" {
		title = "monolith"
		team = "Everyone"
		blueprint = port-labs_blueprint.tf_bp.id
		properties {
			name = "text"
			value = "hedwig2"
		}
	}
`, identifier)
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"port-labs": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccActionConfigCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "title", "monolith"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "team", "Everyone"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "properties.0.name", "text"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "properties.0.value", "hedwig"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "blueprint", identifier),
				),
			},
			{
				Config: testAccActionConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "title", "monolith"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "team", "Everyone"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "properties.0.name", "text"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "properties.0.value", "hedwig2"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "blueprint", identifier),
				),
			},
		},
	})
}

func TestAccPortEntityUpdatePropWithTeamArray(t *testing.T) {
	identifier := genID()
	var testAccActionConfigCreate = fmt.Sprintf(`
	resource "port-labs_blueprint" "tf_bp" {
		title = "TF Provider Test"
		icon = "Terraform"
		identifier = "%s"
		properties {
			identifier = "text"
			type = "string"
			title = "text"
		}
	}
	resource "port-labs_entity" "microservice" {
		title = "monolith"
		blueprint = port-labs_blueprint.tf_bp.id
		teams = ["Everyone"]
		properties {
			name = "text"
			value = "hedwig"
		}
	}
`, identifier)
	var testAccActionConfigUpdate = fmt.Sprintf(`
	resource "port-labs_blueprint" "tf_bp" {
		title = "TF Provider Test"
		icon = "Terraform"
		identifier = "%s"
		properties {
			identifier = "text"
			type = "string"
			title = "text"
		}
	}
	resource "port-labs_entity" "microservice" {
		title = "monolith"
		teams = ["Everyone"]
		blueprint = port-labs_blueprint.tf_bp.id
		properties {
			name = "text"
			value = "hedwig2"
		}
	}
`, identifier)
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"port-labs": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccActionConfigCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "title", "monolith"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "teams.0", "Everyone"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "properties.0.name", "text"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "properties.0.value", "hedwig"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "blueprint", identifier),
				),
			},
			{
				Config: testAccActionConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "title", "monolith"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "teams.0", "Everyone"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "properties.0.name", "text"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "properties.0.value", "hedwig2"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "blueprint", identifier),
				),
			},
		},
	})
}

func TestAccPortEntity(t *testing.T) {
	identifier := genID()
	var testAccActionConfigCreate = fmt.Sprintf(`
	resource "port-labs_blueprint" "microservice" {
		title = "TF Provider Test BP0"
		icon = "Terraform"
		identifier = "%s"
		properties {
			identifier = "text"
			type = "string"
			title = "text"
		}
		properties {
			identifier = "bool"
			type = "boolean"
			title = "boolean"
		}
		properties {
			identifier = "num"
			type = "number"
			title = "number"
		}
		properties {
			identifier = "obj"
			type = "object"
			title = "object"
		}
		properties {
			identifier = "arr"
			type = "array"
			title = "array"
		}
	}
	resource "port-labs_entity" "microservice" {
		title = "monolith"
		blueprint = "${port-labs_blueprint.microservice.identifier}"
		teams = ["Everyone"]
		properties {
			name = "text"
			value = "hedwig"
		}
		properties {
			name = "bool"
			value = "true"
		}
		properties {
			name = "num"
			value = 123
		}
		properties {
			name = "arr"
			items = [1,2,3]
		}
		properties {
			name = "obj"
			value = jsonencode({"a":"b"})
		}
	}
`, identifier)
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"port-labs": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccActionConfigCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "teams.0", "Everyone"),
				),
			},
		},
	})
}

func TestAccPortEntityImport(t *testing.T) {
	blueprintIdentifier := genID()
	entityIdentifier := genID()
	var testAccActionConfigCreate = fmt.Sprintf(`
	resource "port-labs_blueprint" "microservice" {
		title = "TF Provider Test"
		icon = "Terraform"
		identifier = "%s"
		properties {
			identifier = "text"
			type = "string"
			title = "text"
		}
		properties {
			identifier = "bool"
			type = "boolean"
			title = "boolean"
		}
		properties {
			identifier = "num"
			type = "number"
			title = "number"
		}
		properties {
			identifier = "obj"
			type = "object"
			title = "object"
		}
		properties {
			identifier = "arr"
			type = "array"
			title = "array"
		}
	}
	resource "port-labs_entity" "microservice" {
		title = "monolith"
		blueprint = "${port-labs_blueprint.microservice.identifier}"
		teams = ["Everyone"]
		identifier = "%s"
		properties {
			name = "text"
			value = "hedwig"
		}
		properties {
			name = "bool"
			value = "true"
		}
		properties {
			name = "num"
			value = 123
		}
		properties {
			name = "arr"
			items = [1,2,3]
		}
		properties {
			name = "obj"
			value = jsonencode({"a":"b"})
		}
	}
`, blueprintIdentifier, entityIdentifier)
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"port-labs": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccActionConfigCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "title", "monolith"),
				),
			},
			{
				ResourceName:            "port-labs_entity.microservice",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           fmt.Sprintf("%s:%s", blueprintIdentifier, entityIdentifier),
				ImportStateVerifyIgnore: []string{"identifier"},
			},
		},
	})
}
func TestAccPortEntitiesRelation(t *testing.T) {
	identifier1 := genID()
	identifier2 := genID()
	var testAccActionConfigCreate = fmt.Sprintf(`
	resource "port-labs_blueprint" "microservice" {
		title = "TF Provider Test BP0"
		icon = "Terraform"
		identifier = "%s"
		properties {
			identifier = "text"
			type = "string"
			title = "text"
		}
		relations {
			identifier = "tf-relation"
			title = "Test Relation"
			target = port-labs_blueprint.microservice2.identifier
		}
	}
	resource "port-labs_blueprint" "microservice2" {
		title = "TF Provider Test BP0"
		icon = "Terraform"
		identifier = "%s"
		properties {
			identifier = "str"
			type = "string"
			title = "text"
		}
	}
	resource "port-labs_entity" "microservice" {
		title = "monolith"
		blueprint = port-labs_blueprint.microservice.id
		relations {
			name = "tf-relation"
			identifier = port-labs_entity.microservice2.id
		}
		properties {
			name = "text"
			value = "test-relation"
		}
	}
	resource "port-labs_entity" "microservice2" {
		title = "monolith2"
		blueprint = port-labs_blueprint.microservice2.id
		properties {
			name = "str"
			value = "test-relation"
		}
	}
`, identifier1, identifier2)
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"port-labs": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccActionConfigCreate,
			},
		},
	})
}

func TestAccPortEntitiesManyRelation(t *testing.T) {
	identifier1 := genID()
	identifier2 := genID()
	var testAccActionConfigCreate = fmt.Sprintf(`
	resource "port-labs_blueprint" "microservice" {
		title = "TF Provider Test BP0"
		icon = "Terraform"
		identifier = "%s"
		properties {
			identifier = "text"
			type = "string"
			title = "text"
		}
		relations {
			identifier = "tf-relation"
			title = "Test Relation"
			target = port-labs_blueprint.env.identifier
			many = true
		}
	}
	resource "port-labs_blueprint" "env" {
		title = "TF Provider Test BP0"
		icon = "Terraform"
		identifier = "%s"
		properties {
			identifier = "str"
			type = "string"
			title = "text"
		}
	}
	resource "port-labs_entity" "microservice" {
		title = "monolith"
		blueprint = port-labs_blueprint.microservice.id
		relations {
			name = "tf-relation"
			identifiers = ["production","staging"]
		}
		properties {
			name = "text"
			value = "test-relation"
		}
	}
	resource "port-labs_entity" "env" {
		title = "staging"
		identifier = "staging"
		blueprint = port-labs_blueprint.env.id
		properties {
			name = "str"
			value = "test-many-relation"
		}
	}

	resource "port-labs_entity" "env2" {
		title = "production"
		identifier = "production"
		blueprint = port-labs_blueprint.env.id
		properties {
			name = "str"
			value = "test-many-relation"
		}
	}
`, identifier1, identifier2)
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"port-labs": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccActionConfigCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "relations.0.identifiers.0", "production"),
					resource.TestCheckResourceAttr("port-labs_entity.microservice", "relations.0.identifiers.1", "staging"),
				),
			},
		},
	})
}
