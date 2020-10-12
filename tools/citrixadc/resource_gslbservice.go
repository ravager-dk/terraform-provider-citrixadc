package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/gslb"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcGslbservice() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbserviceFunc,
		Read:          readGslbserviceFunc,
		Update:        updateGslbserviceFunc,
		Delete:        deleteGslbserviceFunc,
		Schema: map[string]*schema.Schema{
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cnameentry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookietimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hashid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healthmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxaaausers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxclient": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitornamesvc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"naptrdomainttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"naptrorder": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"naptrpreference": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"naptrreplacement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"naptrservices": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"publicip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publicport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitepersistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"siteprefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svrtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"viewip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"viewname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	var gslbserviceName string
	if v, ok := d.GetOk("servicename"); ok {
		gslbserviceName = v.(string)
	} else {
		gslbserviceName = resource.PrefixedUniqueId("tf-gslbservice-")
		d.Set("servicename", gslbserviceName)
	}
	gslbservice := gslb.Gslbservice{
		Appflowlog:       d.Get("appflowlog").(string),
		Cip:              d.Get("cip").(string),
		Cipheader:        d.Get("cipheader").(string),
		Clttimeout:       d.Get("clttimeout").(int),
		Cnameentry:       d.Get("cnameentry").(string),
		Comment:          d.Get("comment").(string),
		Cookietimeout:    d.Get("cookietimeout").(int),
		Downstateflush:   d.Get("downstateflush").(string),
		Hashid:           d.Get("hashid").(int),
		Healthmonitor:    d.Get("healthmonitor").(string),
		Ip:               d.Get("ip").(string),
		Ipaddress:        d.Get("ipaddress").(string),
		Maxaaausers:      d.Get("maxaaausers").(int),
		Maxbandwidth:     d.Get("maxbandwidth").(int),
		Maxclient:        d.Get("maxclient").(int),
		Monitornamesvc:   d.Get("monitornamesvc").(string),
		Monthreshold:     d.Get("monthreshold").(int),
		Naptrdomainttl:   d.Get("naptrdomainttl").(int),
		Naptrorder:       d.Get("naptrorder").(int),
		Naptrpreference:  d.Get("naptrpreference").(int),
		Naptrreplacement: d.Get("naptrreplacement").(string),
		Naptrservices:    d.Get("naptrservices").(string),
		Newname:          d.Get("newname").(string),
		Port:             d.Get("port").(int),
		Publicip:         d.Get("publicip").(string),
		Publicport:       d.Get("publicport").(int),
		Servername:       d.Get("servername").(string),
		Servicename:      d.Get("servicename").(string),
		Servicetype:      d.Get("servicetype").(string),
		Sitename:         d.Get("sitename").(string),
		Sitepersistence:  d.Get("sitepersistence").(string),
		Siteprefix:       d.Get("siteprefix").(string),
		State:            d.Get("state").(string),
		Svrtimeout:       d.Get("svrtimeout").(int),
		Viewip:           d.Get("viewip").(string),
		Viewname:         d.Get("viewname").(string),
		Weight:           d.Get("weight").(int),
	}

	_, err := client.AddResource(netscaler.Gslbservice.Type(), gslbserviceName, &gslbservice)
	if err != nil {
		return err
	}

	d.SetId(gslbserviceName)

	err = readGslbserviceFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbservice but we can't read it ?? %s", gslbserviceName)
		return nil
	}
	return nil
}

func readGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbserviceName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading gslbservice state %s", gslbserviceName)
	data, err := client.FindResource(netscaler.Gslbservice.Type(), gslbserviceName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservice state %s", gslbserviceName)
		d.SetId("")
		return nil
	}
	d.Set("servicename", data["servicename"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("cip", data["cip"])
	d.Set("cipheader", data["cipheader"])
	d.Set("clttimeout", data["clttimeout"])
	d.Set("cnameentry", data["cnameentry"])
	d.Set("comment", data["comment"])
	d.Set("cookietimeout", data["cookietimeout"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("hashid", data["hashid"])
	d.Set("healthmonitor", data["healthmonitor"])
	d.Set("ip", data["ip"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("maxaaausers", data["maxaaausers"])
	d.Set("maxbandwidth", data["maxbandwidth"])
	d.Set("maxclient", data["maxclient"])
	d.Set("monitornamesvc", data["monitornamesvc"])
	d.Set("monthreshold", data["monthreshold"])
	d.Set("naptrdomainttl", data["naptrdomainttl"])
	d.Set("naptrorder", data["naptrorder"])
	d.Set("naptrpreference", data["naptrpreference"])
	d.Set("naptrreplacement", data["naptrreplacement"])
	d.Set("naptrservices", data["naptrservices"])
	d.Set("newname", data["newname"])
	d.Set("port", data["port"])
	d.Set("publicip", data["publicip"])
	d.Set("publicport", data["publicport"])
	d.Set("servername", data["servername"])
	d.Set("servicename", data["servicename"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sitename", data["sitename"])
	d.Set("sitepersistence", data["sitepersistence"])
	d.Set("siteprefix", data["siteprefix"])
	d.Set("state", data["state"])
	d.Set("svrtimeout", data["svrtimeout"])
	d.Set("viewip", data["viewip"])
	d.Set("viewname", data["viewname"])
	d.Set("weight", data["weight"])

	return nil

}

func updateGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbserviceName := d.Get("servicename").(string)

	gslbservice := gslb.Gslbservice{
		Servicename: d.Get("servicename").(string),
	}
	hasChange := false
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowlog has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("cip") {
		log.Printf("[DEBUG]  citrixadc-provider: Cip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cip = d.Get("cip").(string)
		hasChange = true
	}
	if d.HasChange("cipheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Cipheader has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cipheader = d.Get("cipheader").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Clttimeout has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Clttimeout = d.Get("clttimeout").(int)
		hasChange = true
	}
	if d.HasChange("cnameentry") {
		log.Printf("[DEBUG]  citrixadc-provider: Cnameentry has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cnameentry = d.Get("cnameentry").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("cookietimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookietimeout has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cookietimeout = d.Get("cookietimeout").(int)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG]  citrixadc-provider: Downstateflush has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("hashid") {
		log.Printf("[DEBUG]  citrixadc-provider: Hashid has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Hashid = d.Get("hashid").(int)
		hasChange = true
	}
	if d.HasChange("healthmonitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Healthmonitor has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Healthmonitor = d.Get("healthmonitor").(string)
		hasChange = true
	}
	if d.HasChange("ip") {
		log.Printf("[DEBUG]  citrixadc-provider: Ip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Ip = d.Get("ip").(string)
		hasChange = true
	}
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipaddress has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Ipaddress = d.Get("ipaddress").(string)
		hasChange = true
	}
	if d.HasChange("maxaaausers") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxaaausers has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Maxaaausers = d.Get("maxaaausers").(int)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxbandwidth has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Maxbandwidth = d.Get("maxbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("maxclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxclient has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Maxclient = d.Get("maxclient").(int)
		hasChange = true
	}
	if d.HasChange("monitornamesvc") {
		log.Printf("[DEBUG]  citrixadc-provider: Monitornamesvc has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Monitornamesvc = d.Get("monitornamesvc").(string)
		hasChange = true
	}
	if d.HasChange("monthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Monthreshold has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Monthreshold = d.Get("monthreshold").(int)
		hasChange = true
	}
	if d.HasChange("naptrdomainttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Naptrdomainttl has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Naptrdomainttl = d.Get("naptrdomainttl").(int)
		hasChange = true
	}
	if d.HasChange("naptrorder") {
		log.Printf("[DEBUG]  citrixadc-provider: Naptrorder has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Naptrorder = d.Get("naptrorder").(int)
		hasChange = true
	}
	if d.HasChange("naptrpreference") {
		log.Printf("[DEBUG]  citrixadc-provider: Naptrpreference has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Naptrpreference = d.Get("naptrpreference").(int)
		hasChange = true
	}
	if d.HasChange("naptrreplacement") {
		log.Printf("[DEBUG]  citrixadc-provider: Naptrreplacement has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Naptrreplacement = d.Get("naptrreplacement").(string)
		hasChange = true
	}
	if d.HasChange("naptrservices") {
		log.Printf("[DEBUG]  citrixadc-provider: Naptrservices has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Naptrservices = d.Get("naptrservices").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("publicip") {
		log.Printf("[DEBUG]  citrixadc-provider: Publicip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Publicip = d.Get("publicip").(string)
		hasChange = true
	}
	if d.HasChange("publicport") {
		log.Printf("[DEBUG]  citrixadc-provider: Publicport has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Publicport = d.Get("publicport").(int)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  citrixadc-provider: Servername has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicename has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicetype has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sitename") {
		log.Printf("[DEBUG]  citrixadc-provider: Sitename has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Sitename = d.Get("sitename").(string)
		hasChange = true
	}
	if d.HasChange("sitepersistence") {
		log.Printf("[DEBUG]  citrixadc-provider: Sitepersistence has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Sitepersistence = d.Get("sitepersistence").(string)
		hasChange = true
	}
	if d.HasChange("siteprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Siteprefix has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Siteprefix = d.Get("siteprefix").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("svrtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Svrtimeout has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Svrtimeout = d.Get("svrtimeout").(int)
		hasChange = true
	}
	if d.HasChange("viewip") {
		log.Printf("[DEBUG]  citrixadc-provider: Viewip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Viewip = d.Get("viewip").(string)
		hasChange = true
	}
	if d.HasChange("viewname") {
		log.Printf("[DEBUG]  citrixadc-provider: Viewname has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Viewname = d.Get("viewname").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  citrixadc-provider: Weight has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Weight = d.Get("weight").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Gslbservice.Type(), gslbserviceName, &gslbservice)
		if err != nil {
			return fmt.Errorf("Error updating gslbservice %s", gslbserviceName)
		}
	}
	return readGslbserviceFunc(d, meta)
}

func deleteGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbserviceName := d.Id()
	err := client.DeleteResource(netscaler.Gslbservice.Type(), gslbserviceName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}