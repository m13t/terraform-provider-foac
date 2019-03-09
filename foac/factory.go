package foac

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func generateDataSource(name string, opts []string) *schema.Resource {
	s := map[string]*schema.Schema{
		"message": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"subtitle": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	for _, opt := range opts {
		s[opt] = &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		}
	}

	return &schema.Resource{
		Schema: s,

		Read: func(d *schema.ResourceData, m interface{}) error {
			c := m.(*client)

			params := []string{name}

			for _, opt := range opts {
				params = append(params, d.Get(opt).(string))
			}

			res, err := c.makeRequest(params...)

			if err != nil {
				return err
			}

			d.SetId(time.Now().UTC().String())
			d.Set("message", res.Message)
			d.Set("subtitle", res.Subtitle)

			return nil
		},
	}
}
