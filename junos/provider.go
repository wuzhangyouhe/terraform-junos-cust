package junos

import (
	"context"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	idSeparator    = "_-_"
	defaultWord    = "default"
	inetWord       = "inet"
	inet6Word      = "inet6"
	emptyWord      = "empty"
	matchWord      = "match"
	permitWord     = "permit"
	thenWord       = "then"
	prefixWord     = "prefix"
	actionNoneWord = "none"
	addWord        = "add"
	deleteWord     = "delete"
	setWord        = "set"
	setLineStart   = setWord + " "
	st0Word        = "st0"
	opsfV2         = "ospf"
	ospfV3         = "ospf3"
)

var (
	mutex = &sync.Mutex{}
)

// Provider junos for terraform.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_HOST", nil),
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_PORT", 830),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_USERNAME", "netconf"),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_PASSWORD", nil),
			},
			"sshkey_pem": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_KEYPEM", nil),
			},
			"sshkeyfile": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_KEYFILE", nil),
			},
			"keypass": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_KEYPASS", nil),
			},
			"group_interface_delete": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_GROUP_INTERFACE_DELETE", nil),
			},
			"cmd_sleep_short": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_SLEEP_SHORT", 100),
			},
			"cmd_sleep_lock": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_SLEEP_LOCK", 10),
			},
			"debug_netconf_log_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("JUNOS_LOG_PATH", ""),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"junos_interface": dataSourceInterface(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"junos_policyoptions_as_path_group":    resourcePolicyoptionsAsPathGroup(),
			"junos_policyoptions_as_path":          resourcePolicyoptionsAsPath(),
			"junos_policyoptions_community":        resourcePolicyoptionsCommunity(),
			"junos_policyoptions_policy_statement": resourcePolicyoptionsPolicyStatement(),
			"junos_policyoptions_prefix_list":      resourcePolicyoptionsPrefixList(),
			"junos_vlan":                           resourceVlan(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		junosIP:                  d.Get("ip").(string),
		junosPort:                d.Get("port").(int),
		junosUserName:            d.Get("username").(string),
		junosPassword:            d.Get("password").(string),
		junosSSHKeyPEM:           d.Get("sshkey_pem").(string),
		junosSSHKeyFile:          d.Get("sshkeyfile").(string),
		junosKeyPass:             d.Get("keypass").(string),
		junosGroupIntDel:         d.Get("group_interface_delete").(string),
		junosCmdSleepShort:       d.Get("cmd_sleep_short").(int),
		junosCmdSleepLock:        d.Get("cmd_sleep_lock").(int),
		junosDebugNetconfLogPath: d.Get("debug_netconf_log_path").(string),
	}

	return config.Session()
}
