package provider

import (
        "fmt"
        "github.com/antoineaugusti/updown"
        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func recipientsResource() *schema.Resource {
        return &schema.Resource{
                Description: "`updown_recipients` defines a recipients",

                Create: recipientsCreate,
                Read:   recipientsRead,
                Delete: recipientsDelete,
                Exists: recipientsExists,

                Importer: &schema.ResourceImporter{
                        State: schema.ImportStatePassthrough,
                },

                Schema: map[string]*schema.Schema{
                        "type": {
                                Type:        schema.TypeString,
                                Required:    true,
                                Description: "Type of recipient ('email', 'sms', 'webhook' or 'slack_compatible' only). The other integrations (slack, telegram, zapier, statuspage, etc.) require the web UI to setup.",
                        },
                        "value": {
                                Type:        schema.TypeString,
                                Required:    true,
                                Description: "The recipient value (email address, phone number or URL)",
                                ForceNew:    true,
                        },
                        "selected": {
                                Type:        schema.TypeBool,
                                Optional:    true,
                                Description: "Initial state for all checks: true = selected on all existing checks (default)",
                                ForceNew:    true,
                        },
                },
        }
}

func constructRecipientsPayload(d *schema.ResourceData) updown.Recipients {
        payload := updown.Recipients{}
        if v, ok := d.GetOk("type"); ok {
                payload.Type = v.(string)
        }

        if v, ok := d.GetOk("value"); ok {
                payload.Value = v.(string)
        }

        if v, ok := d.GetOk("selected"); ok {
                payload.Selected = v.(bool)
        }

        return payload
}

func recipientsCreate(d *schema.ResourceData, meta interface{}) error {
        client := meta.(*updown.Client)

        recipients, _, err := client.Recipients.Add(constructCheckPayload(d))
        if err != nil {
                return fmt.Errorf("creating recipients with the API")
        }

        d.SetId(recipients.Token)

        return recipientsRead(d, meta)
}

func recipientsRead(d *schema.ResourceData, meta interface{}) error {
        client := meta.(*updown.Client)
        recipients, _, err := client.Recipients.Get(d.Id())

        if err != nil {
                return fmt.Errorf("reading recipients from the API")
        }

        for k, v := range map[string]interface{}{
                "type":               recipients.Type,
                "value":              recipients.Value,
                "selected":           recipients.Selected,
        } {
                if err := d.Set(k, v); err != nil {
                        return err
                }
        }

        return nil
}

func recipientsUpdate(d *schema.ResourceData, meta interface{}) error {
        client := meta.(*updown.Client)

        _, _, err := client.Recipients.Update(d.Id(), constructRecipientsPayload(d))
        if err != nil {
                return fmt.Errorf("updating recipients with the API")
        }

        return nil
}

func recipientsDelete(d *schema.ResourceData, meta interface{}) error {
        client := meta.(*updown.Client)
        recipientsDeleted, _, err := client.Recipients.Remove(d.Id())

        if err != nil {
                return fmt.Errorf("removing recipients from the API")
        }

        if !recipientsDeleted {
                return fmt.Errorf("recipients couldn't be deleted")
        }

        return nil
}

func recipientsExists(d *schema.ResourceData, meta interface{}) (bool, error) {
        err := recipientsRead(d, meta)
        return err == nil, err
}
