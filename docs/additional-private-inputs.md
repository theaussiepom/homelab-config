# Additional private application inputs

Codex controller/tunnel source, their Compose project and the HA notification
package are deferred from this initial publication. References below describe
the retained deployment and do not imply those components are included here.

The existing Node-RED private-input loader also requires `HOMELAB_NR_UI_HOST`.
Its health check follows that exact configured address. No live bind changes
are made by this preparation.

phpMyAdmin receives its reviewed warning-host list as JSON through
`HOMELAB_PMA_SSL_WARNING_SAFE_HOSTS_JSON`. This preserves the existing list;
it does not enable additional TLS exceptions.

The Codex HA installer requires `HOMELAB_CODEX_HA_BASE_URL` before making
changes. The Zabbix reader requires `HOMELAB_ZABBIX_HOST` and
`HOMELAB_ZABBIX_BINDINGS_JSON`. Those trusted deployment inputs do not change
the fixed API path, allowed methods or metric validation. Missing inputs fail
closed. Public reader tests use synthetic host/item bindings.

The Ansible offline inventory checker pins the source-variable wiring per host.
The backup role separately requires `HOMELAB_ANSIBLE_EXPECTED_BACKUP_NFS_SOURCE`
and compares the configured target against it.

Original private values are retained with the preparation evidence on the host.
Install private inputs and complete native validation before deployment.
