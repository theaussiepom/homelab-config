> Publication copy: deployment addresses and domains below are examples; restore private inputs before use.

# OpenBao retired from active deployment

On 17 October 2026 (Australia/Brisbane), Ben requested removal of unused OpenBao.
The service was initialized and sealed. Its container and active deployment entry
were removed. Ben additionally requested deletion of OpenBao storage.

The separate guarded storage cleanup covers:

- Raft data and logs: `/mnt/nas/docker/homelab-data/services/openbao/{data,logs}`.
- Both image-declared Docker volumes formerly mounted at `/openbao/file` and `/openbao/logs`.
- The OpenBao-only database archive identified by its accepted migration receipt.

Do not assume storage deletion completed from this note; consult the cleanup receipt.
The service-only result and original deployment definition, private resolved model,
exact volume names and runtime are under:
`/home/ben/.local/state/openbao-service-retirement/openbao-retire-20261017-84b3fb0b` (owner-only).
The non-database configuration at `services/openbao/openbao.hcl` and historical
acceptance/failure receipts remain as records, not a retained database backup.

The `secrets.site-2.example.invalid` and `secrets.site-1.example.invalid` OpenBao routes were provided
by container labels and are retired with the container. DNS itself was not changed.
Do not add the service back as routine stack maintenance. A future installation
requires a fresh deployment decision; no restore or unseal was verified here.
