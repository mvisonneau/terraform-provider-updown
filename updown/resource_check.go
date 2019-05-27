package updown

import (
	"fmt"

	"github.com/antoineaugusti/updown"
	"github.com/hashicorp/terraform/helper/schema"
)

func checkResource() *schema.Resource {
	return &schema.Resource{
		Create: checkCreate,
		Read:   checkRead,
		Delete: checkDelete,
		Update: checkUpdate,
		Exists: checkExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL you want to monitor.",
			},
			"period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Interval in seconds (15, 30, 60, 120, 300, 600, 1800 or 3600).",
				Default:     60,
			},
			"apdex_t": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "APDEX threshold in seconds (0.125, 0.25, 0.5, 1.0 or 2.0).",
				Default:     0.5,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Is the check enabled (true or false).",
				Default:     true,
			},
			"published": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Shall the status page be public (true or false).",
				Default:     false,
			},
			"alias": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Human readable name.",
			},
			"string_match": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search for this string in the page.",
			},
			"mute_until": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mute notifications until given time, accepts a time, 'recovery' or 'forever'.",
			},
			"disabled_locations": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Disabled monitoring locations. It's a lsit of abbreviated location names.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"custom_headers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The HTTP headers you want in requests.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func contrusctPayload(d *schema.ResourceData) updown.CheckItem {
	payload := updown.CheckItem{}

	if v, ok := d.GetOk("url"); ok {
		payload.URL = v.(string)
	}

	if v, ok := d.GetOk("period"); ok {
		payload.Period = v.(int)
	}

	if v, ok := d.GetOk("apdex_t"); ok {
		payload.Apdex = v.(float64)
	}

	if v, ok := d.GetOk("enabled"); ok {
		payload.Enabled = v.(bool)
	}

	if v, ok := d.GetOk("published"); ok {
		payload.Published = v.(bool)
	}

	if v, ok := d.GetOk("alias"); ok {
		payload.Alias = v.(string)
	}

	if v, ok := d.GetOk("string_match"); ok {
		payload.StringMatch = v.(string)
	}

	if v, ok := d.GetOk("mute_until"); ok {
		payload.MuteUntil = v.(string)
	}

	if v, ok := d.GetOk("disabled_locations"); ok {
		interfaceSlice := v.(*schema.Set).List()
		var stringSlice []string
		for s := range interfaceSlice {
			stringSlice = append(stringSlice, interfaceSlice[s].(string))
		}
		payload.DisabledLocations = stringSlice
	}

	if m, ok := d.GetOk("custom_headers"); ok {
		payload.CustomHeaders = map[string]string{}
		for k, v := range m.(map[string]interface{}) {
			payload.CustomHeaders[k] = v.(string)
		}
	}

	return payload
}

func checkCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)

	check, _, err := client.Check.Add(contrusctPayload(d))
	if err != nil {
		return fmt.Errorf("Error creating check with the API.")
	}

	d.SetId(check.Token)

	return checkRead(d, meta)
}

func checkRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)
	check, _, err := client.Check.Get(d.Id())

	if err != nil {
		return fmt.Errorf("Error reading check from the API.")
	}

	d.Set("url", check.URL)
	d.Set("period", check.Period)
	d.Set("apdex_t", check.Apdex)
	d.Set("enabled", check.Enabled)
	d.Set("published", check.Published)
	d.Set("alias", check.Alias)
	d.Set("string_match", check.StringMatch)
	d.Set("mute_until", check.MuteUntil)
	d.Set("disabled_locations", check.DisabledLocations)
	d.Set("custom_headers", check.CustomHeaders)

	return nil
}

func checkUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)

	_, _, err := client.Check.Update(d.Id(), contrusctPayload(d))
	if err != nil {
		return fmt.Errorf("Error updating check with the API.")
	}

	return nil
}

func checkDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)
	checkDeleted, _, err := client.Check.Remove(d.Id())

	if err != nil {
		return fmt.Errorf("Error removing check from the API.")
	}

	if !checkDeleted {
		return fmt.Errorf("Check couldn't be deleted.")
	}

	return nil
}

func checkExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	err := checkRead(d, meta)
	return err == nil, err
}
