package namecheap

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"

	"github.com/libdns/namecheap"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *namecheap.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.namecheap",
		New: func() caddy.Module { return &Provider{new(namecheap.Provider)} },
	}
}

// Before using the provider config, resolve placeholders in the API token.
// Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIKey = caddy.NewReplacer().ReplaceAll(p.Provider.APIKey, "")
	p.Provider.User = caddy.NewReplacer().ReplaceAll(p.Provider.User, "")
	p.Provider.APIEndpoint = caddy.NewReplacer().ReplaceAll(p.Provider.APIEndpoint, "")
	p.Provider.ClientIP = caddy.NewReplacer().ReplaceAll(p.Provider.ClientIP, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	namecheap {
//	    api_key <api_key>
//	    user <user>
//	    api_endpoint https://api.namecheap.com/xml.response
//	    client_ip <client_ip>
//	}
//
// Expansion of placeholders is left to the JSON config caddy.Provisioner (above).
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_key":
				if p.Provider.APIKey != "" {
					return d.Err("API key already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.APIKey = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			case "user":
				if p.Provider.User != "" {
					return d.Err("user already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.User = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			case "api_endpoint":
				if p.Provider.APIEndpoint != "" {
					return d.Err("API endpoint already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.APIEndpoint = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			case "client_ip":
				if p.Provider.ClientIP != "" {
					return d.Err("client IP already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.ClientIP = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIKey == "" {
		return d.Err("missing API key")
	}
	if p.Provider.User == "" {
		return d.Err("missing user")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
