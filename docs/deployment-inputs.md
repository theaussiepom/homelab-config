# BenBeast deployment inputs

Codex controller/tunnel source, their Compose project and the HA notification
package are deferred from this initial publication. References below describe
the retained deployment and do not imply those components are included here.

Maintained definitions live in `stacks/benbeast/<project>/compose.yaml`. The eight existing projects remain automation, codex, database, monitoring, music-assistant, network, scrypted and security. `stacks/benbeast/deployment-inputs.json` records the reviewed 35-service cohort and each project's required private input filenames.

Private inputs live outside Git at `/mnt/nas/docker/.homelab-private/benbeast/`: directory mode0700 and files mode0600, owned by ben. Existing project environment files are preserved separately; database backup literals use `database.db_backup.env`. Codex needs no project env file. Credentials, resolved Compose models, database data, service state and protected HA/editor configuration are not consolidation source outputs and must not be staged with this change.

On this NAS export, creation modes and umask alone do not enforce those permissions. Explicitly set and verify directory mode0700, then set and verify each newly opened private file mode0600 before writing any private bytes. Verify the final mode and content after writing.

## Operator workflow

Run from the BenBeast host as ben:

```sh
python3 /mnt/nas/docker/homelab/scripts/compose.py validate
python3 /mnt/nas/docker/homelab/scripts/compose.py validate database
python3 /mnt/nas/docker/homelab/scripts/compose.py command database phpmyadmin
```

`validate` parses the selected canonical definitions with their declared private inputs and prints only a sanitized result. `command` validates and prints a shell-quoted command for explicitly named services. The helper never executes that launch command or creates, stops, restarts, updates or removes containers. Review any printed command against the intended maintenance scope before independently authorizing and executing it; successful parsing is not a fresh image/runtime, backup, client-drain or application-health gate.

The helper fixes the local Docker socket, project identity, project directory and interpolation environment. It disables implicit `.env` discovery when a project needs no env file. Printed commands carry the same environment and local socket. Do not drop that prefix or add old override files to the maintained invocation. Private inputs must remain regular files with the documented permissions and without symbolic-link traversal.

Service names, not container names, are required. Selection must be nonempty and in the reviewed cohort. The printed command uses `--no-deps --no-build --pull never`; whole-project launches and arbitrary Compose command passthrough are unsupported. Source-only Scrypted `watchtower` remains absent and cannot be selected. Codex and codex-tunnel can be validated but the helper refuses their launch commands; their maintenance requires the separate host-controlled drain procedure. The manifest records an accepted cohort, not live permission to start a service that has since stopped.

## Preservation and recovery

This is source consolidation, not a data migration. Existing Docker project/service names, image identities, database SSD binds, NAS data, shared paths and protected HA/editor directory identities remain unchanged. Canonical resource owners keep their network/volume creation definitions; consumer references remain external where appropriate. Existing `pihole-ftl-db` is explicitly external by its confirmed live identity. Never delete or recreate a shared resource merely to make a source file appear uniform.

Existing containers retain historical Docker launch labels until an independently authorized recreation. Their old env/overlay files and protected consolidation snapshots remain in place. Those labels are a historical launch record; the maintained helper uses the new canonical definitions. Old inputs and legacy directories are not deletion candidates solely because the new model validates.

Canonical files and recovery metadata already consumed through single-file backup binds must retain their inodes, or have the affected backup source mount deliberately rebound. A normal replace-by-rename can leave a running backup reading old bytes. Verify both host content and container source views when updating those inputs.

Nonsecret database-recovery metadata is maintained under `services/database-backup/recovery/`, with existing container destination paths preserved. The reviewed recovery transition adds the canonical database project env, database backup service env and automation project env within the existing DB/HA configuration scope, for 25 source entries. Validate the actual manifest/container views and fresh native bundle before claiming that transition complete. Do not add unrelated security credentials to DB cloud bundles implicitly.

All old inputs and the targeted private NAS snapshot are retained. The new private directory stays on the same NAS storage class as the existing inputs. Neither that directory nor a second NAS copy proves independent backup; no general recurring/offsite configuration-backup coverage has been established. Do not replay historical cutover commands or restore stale pre-SSD database copies as an automatic rollback.

## Source review

Only explicit reviewed source paths belong in a later commit. The live checkout has unrelated edits and partially inaccessible protected paths; avoid repository-wide staging, reset, clean, branch changes or blanket ignore rules. Run the focused helper tests and privately compare complete resolved old/new service models before activation. Scan exactly the planned source outputs for known moved credential values on BenBeast. Keep private comparison models and scanner inputs on that host; sanitized counts and hashes are sufficient review evidence. A known-value scan is a bounded check, not a claim that every possible secret or historical credential has been discovered.
