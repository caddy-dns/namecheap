# namecheap module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with namecheap accounts. In order to use this module, you must build caddy with [xcaddy](https://github.com/caddyserver/xcaddy).

```
xcaddy build --with github.com/caddy-dns/namecheap
```

Visit [here](https://www.namecheap.com/support/api/intro/) to get started with the namecheap API. It is recommended to first use the sandbox environment to test your configuration.

## Caddy module name

```
dns.providers.namecheap
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "namecheap",
        "api_key": "{env.NAMECHEAP_API_KEY}",
        "user": "{env.NAMECHEAP_API_USER}",
        "api_endpoint": "https://api.namecheap.com/xml.response",
        "client_ip": "deduced-automatically-if-not-set"
      }
    }
  }
}
```

or with the Caddyfile:

```
tls {
    dns namecheap {
        api_key {env.NAMECHEAP_API_KEY}
        user {env.NAMECHEAP_API_USER}
        api_endpoint https://api.namecheap.com/xml.response
        client_ip <client_ip>
    }
}
```

You can replace `{env.NAMECHEAP_API_KEY}` with the actual auth token if you prefer to put it directly in your config instead of an environment variable.

A complete example

```
*.example.com {
    file_server
    tls {
        dns namecheap {
            api_key {env.NAMECHEAP_API_KEY}
            user namecheap_api_username
            api_endpoint https://api.namecheap.com/xml.response
            client_ip public_egress_ip
        }
    }
}
```
