> Publication copy: deployment addresses and domains below are examples; restore private inputs before use.

# Traefik domain aliases

Deployment: 10 October 2026. This adds access through `site-1.example.invalid` while preserving `site-2.example.invalid`. It does not migrate application identities or stored canonical URLs.

## Domain and certificate configuration

- Keep `DOMAIN=site-2.example.invalid` as the original domain. `DOMAIN_ALIAS=site-1.example.invalid` supplies the second hostname in relevant Compose labels and dynamic templates.
- The alias variable is required in the private `automation`, `security`, `network`, `database`, `monitoring`, `music-assistant` and `scrypted` project environment files under `/mnt/nas/docker/.homelab-private/benbeast`. Traefik also receives it in its container environment for file-template rendering.
- The intended deployed scope is 29 service hostname aliases. Bitwarden is deliberately deferred: its router, CSP and application domain remain unchanged pending a separate decision.
- Host rules retain the original match and add an OR match for the alias. Existing path/header conditions, priorities, services and access-control middleware remain in place.
- The issued `site-1.example.invalid` certificate covers `*.site-1.example.invalid`. Existing original-domain certificates and their renewal configuration remain in place. DNS challenge credentials stay in private deployment inputs and must not be copied into repository examples.
- Split DNS and public DNS are separate controls. Preserve their intended internal/public exposure when adding further aliases; a certificate alone does not publish a service.
- Both internal Unbound servers resolve the alias zone to the existing Traefik address. Public Cloudflare exposure mirrors the original named records; only the missing `webdav` alias record needed adding. No public wildcard DNS record was added.

## Authentication and application URLs

- Thirteen active per-application Authentik forward-auth providers have equivalent alias providers and applications: `zigbee1`, `zigbee2`, `mqtt`, `code`, `maria`, `influx`, `proxy`, `scrypted`, `pihole1`, `pihole2`, `zwave`, `zigbeelens` and `threadlens`.
- Original providers, policy bindings, permission grants and outpost membership are preserved. Alias objects use fresh identities and equivalent authorization settings. Unused providers were not duplicated, and applications were not collapsed into one domain-wide authorization policy.
- Authentication still begins at canonical `https://sso.site-2.example.invalid`. Alias callbacks return to the requested application hostname. Authentik uses host-only session cookies rather than a cookie domain pinned to the original SSO host.
- Duplicate alias portal tiles are hidden through Authentik 2025.2's supported `blank://blank` launch URL. Original application tiles remain the normal portal entries.
- pgAdmin's strict OAuth callback allowlist includes both `https://pgadmin.site-2.example.invalid/oauth2/authorize` and `https://pgadmin.site-1.example.invalid/oauth2/authorize`; its original callback is retained.
- Zabbix and phpMyAdmin retain their existing canonical original-domain return behavior. Opening an alias can therefore lead back to the original hostname. This is not a migration of those applications' canonical URLs.

## File-provider reloads on NFS

Dynamic configuration is bound from `services/traefik/conf.d` on NFS into `/etc/traefik/conf.d`, with `providers.file.watch=true`. During this deployment, all 18 changed files were present and correct inside the container, but Traefik continued using its previous in-memory rules until a controlled restart of the existing container.

After future dynamic file changes, verify active `/api/rawdata` routing and middleware state; matching file hashes alone do not prove activation. If state remains stale, validate the complete candidate first, preserve the running container configuration, then perform a controlled restart and verify active state again. A container recreation or unrelated application restart is not required solely for this reload.

Traefik relies on filesystem notifications; fsnotify documents the lack of network notification support in NFS. See [Traefik file-provider limitations](https://doc.traefik.io/traefik/reference/install-configuration/providers/others/file/) and [fsnotify's NFS limitation](https://github.com/fsnotify/fsnotify#why-dont-notifications-work-with-nfs-smb-fuse-proc-or-sys).

## Code-server deployment provenance

Code-server's pre-existing live logging configuration was `json-file` with no logging options. The canonical Compose file contained size/file limits that were not active. Applying those limits was outside the domain-alias change.

The selected recreation therefore used the private override:

`/mnt/nas/docker/.homelab-private/benbeast/traefik-dual-domain-20261010/code-server-preserve-logging.yaml`

It uses `logging: !override` with `driver: json-file`, preserving the live behavior. The override is recorded in the container's Compose provenance. A later ordinary Compose deployment without this override would apply the canonical logging limits; review and reconcile that known drift deliberately before the next code-server deployment. The temporary preservation override is not a new repository logging policy.

## Verification boundary

The post-restart check verified all 20 file-provider routes against their explicit deployed fields and verified the three updated CSPs exactly. All 26 original/alias Authentik forward-auth endpoint checks returned the expected authentication redirect and matching callback host. Source hashes, original Authentik objects, alias authorization equivalence and outpost membership were checked separately.

pgAdmin's anonymous sign-in start produces its new strict callback, which Authentik accepts. Both SSO hosts issue host-only, Secure, HttpOnly session cookies. Origin TLS passed for all 60 tested names, including the deliberately unrouted Bitwarden alias. Public Cloudflare checks passed for Home Assistant and SSO on both domains.

These checks do not substitute for completing browser login and normal application workflows on the aliases. The configured Drive, Drive Server, Photos, WebDAV, Glances, NetAlertX and Thread upstream ports refuse connections; both domains return HTTP 502 for them. Their backend addresses were not changed in this deployment. Correcting those upstreams is separate work.
