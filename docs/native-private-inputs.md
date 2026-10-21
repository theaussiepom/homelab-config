# Native private deployment inputs

Production inventory keeps its stable host aliases and uses required Ansible
environment lookups for deployment addresses. Key names are in
`ansible/inventory/production/deployment.env.example`. Missing variables fail
instead of choosing fallback production values.

Traefik dynamic files use native Go templates. Their required environment keys
are passed through the network Compose service. Guards reject missing values;
formatted values are YAML-quoted. The static file uses the separate `.in`
template described in `rendered-private-config.md`.

Unbound keeps public DNS policy, including standard private-address ranges.
Site-local zones and records belong in the privately restored
`deployment.private.conf`. Set `HOMELAB_UNBOUND_LOCAL_RECORDS_FILE` to its host
path. The added read-only bind refuses to create a missing source directory.

The retired OpenBao deployment configuration is excluded from public source and
retained privately. This does not remove or stop any service.

These changes are prepared on the isolated branch. Private input files must be
installed and native application validation completed before deployment.

References:
- https://docs.ansible.com/ansible/latest/collections/ansible/builtin/env_lookup.html
- https://doc.traefik.io/traefik/reference/dynamic-configuration/file/
- https://www.nlnetlabs.nl/documentation/unbound/unbound.conf/
