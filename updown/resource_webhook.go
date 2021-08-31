package updown

import (
	"fmt"

	"github.com/antoineaugusti/updown"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func webhookResource() *schema.Resource {
	return &schema.Resource{
		Create: webhookCreate,
		Read:   webhookRead,
		Delete: webhookDelete,
		Exists: webhookExists,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL of the webhook you want to trigger on updown events.",
				ForceNew:    true,
			},
		},
	}
}

func constructWebhookPayload(d *schema.ResourceData) (webhook updown.Webhook) {
	if v, ok := d.GetOk("url"); ok {
		webhook.URL = v.(string)
	}
	return
}

func webhookCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)

	webhook, _, err := client.Webhook.Add(constructWebhookPayload(d))
	if err != nil {
		return fmt.Errorf("creating webhook with the API: %s", err.Error())
	}

	d.SetId(webhook.ID)

	return webhookRead(d, meta)
}

func webhookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)
	webhooks, _, err := client.Webhook.List()

	if err != nil {
		return fmt.Errorf("reading webhooks from the API")
	}

	for _, w := range webhooks {
		if d.Id() == w.ID {
			for k, v := range map[string]interface{}{
				"url": w.URL,
			} {
				if err := d.Set(k, v); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func webhookDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)
	webhookDeleted, _, err := client.Webhook.Remove(d.Id())

	if err != nil {
		return fmt.Errorf("removing webhook from the API")
	}

	if !webhookDeleted {
		return fmt.Errorf("webhook couldn't be deleted")
	}

	return nil
}

func webhookExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	err := webhookRead(d, meta)
	return err == nil, err
}
