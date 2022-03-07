package provider

import (
	"fmt"

	"github.com/antoineaugusti/updown"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func recipientResource() *schema.Resource {
	return &schema.Resource{
		Description: "`updown_recipient` defines a recipient",

		Create: recipientCreate,
		Read:   recipientRead,
		Delete: recipientDelete,
		Exists: recipientExists,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of recipient ('email', 'sms', 'webhook' or 'slack_compatible' only). The other integrations (slack, telegram, zapier, statuspage, etc.) require the web UI to setup.",
				ForceNew:    true,
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The recipient value (email address, phone number or URL)",
				ForceNew:    true,
			},
		},
	}
}

func constructRecipientPayload(d *schema.ResourceData) updown.RecipientItem {
	payload := updown.RecipientItem{}
	if v, ok := d.GetOk("type"); ok {
		payload.Type = v.(updown.RecipientType)
	}

	if v, ok := d.GetOk("value"); ok {
		payload.Value = v.(string)
	}

	return payload
}

func recipientCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)

	recipient, _, err := client.Recipient.Add(constructRecipientPayload(d))
	if err != nil {
		return fmt.Errorf("creating Recipient with the API")
	}

	d.SetId(recipient.ID)

	return recipientRead(d, meta)
}

func recipientRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)
	recipients, _, err := client.Recipient.List()

	if err != nil {
		return fmt.Errorf("reading recipients from the API")
	}

	for _, r := range recipients {
		if d.Id() == r.ID {
			for k, v := range map[string]interface{}{
				"type":  string(r.Type),
				"value": r.Name,
			} {
				if err := d.Set(k, v); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func recipientDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)
	RecipientDeleted, _, err := client.Recipient.Remove(d.Id())

	if err != nil {
		return fmt.Errorf("removing Recipient from the API")
	}

	if !RecipientDeleted {
		return fmt.Errorf("recipient couldn't be deleted")
	}

	return nil
}

func recipientExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	err := recipientRead(d, meta)
	return err == nil, err
}
