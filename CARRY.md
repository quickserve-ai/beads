> **FORK CARRY FILE — not upstream; lives on `carry/operational` of `quickserve-ai/beads`; dropped when the last carry leaves and the fork ends (ledger below).**

# CARRY.md — quickserve-ai/beads carry model

Fork-local operational layer on top of the qc-bridge doctrine: the joint
`shared/fork-upstream-operating-principles.v1.md` and its application to this
repository, `roles/bd-owner.md` (the four bars). The doctrine governs
judgment — classification before code, the pin rule, the carry ledger's
default drop and its cross-town review; this file records what is specific to
this fork: its branches, the carries that exist today and the condition under
which each one leaves. When they conflict, the doctrine wins on procedure;
Git wins on state. Live state below is dated — re-derive it from Git before
acting on it. The shared rules here are owned by the two towns' bd owners;
which seats each rule binds today is recorded on the bridge role's status line
(`roles/bd-owner.md`), not here.

## Branch state

| Branch | Contract | Reality (2026-10-05 unless the cell says otherwise) |
|---|---|---|
| `main` | Fast-forward-only mirror of `gastownhall/beads` `main`. Never a carry commit, never a PR branch. Refresh with `git push origin upstream/main:main` after fetching upstream; a personal fork's `main` is kept equal the same way. | `= upstream/main` (`e72cd8b31`, 2026-10-05, schema head 0069 / ignored 0027). Fast-forwarded 2026-10-05 23:43Z from `a690b0a8c` (2026-09-08, 369 behind); the Alex-town bd owner's personal fork `main` moved with it. |
| `carry/operational` | `upstream/main` plus the **marked documentation carries** in the ledger below — this file and `CONTRIBUTING.fork.md` — and nothing buildable. **Not a build source**: the fleet never builds `bd` from a branch that floats with upstream. Refreshed by rebasing the doc commit(s) onto the new `upstream/main` (force-with-lease on this branch only, never on `main`). | Refreshed 2026-10-06 ~17:00Z by Cherub's bd owner (katya): base `b1bee123b` (`upstream/main` of 2026-10-05) + the 17 documentation commits in the ledger below, cherry-picked in order; fork PR #3's two `ga-knhu61` code commits and their revert (fork PR #5) left the history at this refresh, as the previous cell said they would (that code lives only in the fleet lineage, `aa9acb5cf`, `735cc83d4`). Pre-refresh head `b3eace5ae` kept as `backup/carry-operational-pre-rebase-20261006`. Documentation bytes identical before and after (checked with `git diff` on both files). Import checkouts follow by fetch + verified-clean `reset --keep`, never `pull --ff-only` (rules below). |
| `carry-v<schema>/<slug>[-<slug>…]` | The fleet build lineage: the upstream SHA that upstream `gastownhall/gascity` `main` pins in `go.mod` (principle 3, "never ahead of support") plus a short stack of cherry-picks — one branch per pick, one cumulative branch per window build. Built and tested at the pin before it is pushed; handed to the gascity pin owner; every pick has a bead and a drop condition. | **LIVE lineage: `carry-v66-1.3.1/fleet-20261005` @ `3ecd84346` = `v1.3.1` (`c1c4b642ac`, schema **v66** / ignored 0026) + 18 cherry-picks** — the rows of the code ledger below, byte-identical to the 2026-09-23 unit's 18 (the range-diff differs only in `-x` trailer lines). Base: `v1.3.1` (the release-branch cut of 2026-09-30), the version upstream `gastownhall/gascity` `main` requires in `go.mod` (read 2026-10-05); `v1.3.0` is an ancestor and the schema did not move, so the swap is binary-only. Tagged `v1.3.1-fleet.20261005.1` (2026-10-05 08:57Z; release run 37286920196; sha256 darwin-arm64 `42fe7f8a1f1fa21720110ea8f4b6566185d84312ba9d81e36f758dd7daf4cd59`, linux-amd64 `08af5de6caaafa1d78a4ed1259430e4ad44b42542e3ed4db84d7171daea521f9`; `bd version` = `1.3.1-fleet.20261005.1 (c1c4b64+carry.3ecd843)`) and pinned by quickserve-ai/gascity re-sync #5 at `ff7c500fe` (2026-10-05 11:14Z, module `replace`). Installs are per machine and owned by each machine's installer: read `bd version` there; this ledger does not record them. Gates: Alex town's are in the tag message and on be-qfm (2026-10-05). **Cherub town gate (katya, 2026-10-06 01:1xZ): STATIC.** The 18 commits of `v1.3.1..3ecd84346` are byte-identical by `git patch-id --stable` to the 18 of `v1.3.0..3d86fb056` reviewed 2026-09-23 (18/18 identical; `git range-diff` shows all 18 as `!` only because the `-x` trailers changed). Base `v1.3.1` = `c1c4b642a` is upstream's release-prep merge #7048 on its release branch (41 commits not on `upstream/main`, 373 behind it). No build or test of this unit was run in Cherub town; the install (gc `1584a6116503` + this bd) is woodhouse's: deferred on the evening of 2026-10-05 (a supervisor restart would have re-armed a long sweep against the shared hub) and planned with a gc binary swap on the evening of 2026-10-06 Pacific, with rollback to the 09-23 twin by rename. Nothing read there says "do not install". *(Recorded on fork PR #6's review, 2026-10-06.)* **Rollback twin:** `v1.3.0-fleet.20260923.2` = `carry-v66-1.3.0/fleet-20260923` @ `3d86fb056` = `v1.3.0` (`f45b249ce`) + 18, release run 35859552460, sha256 darwin-arm64 `63ec9fccedea45ce6bdc067e4c23f33e5bc543f47643c5dc8266e38580afba3c`, linux-amd64 `fc80c7100245d7a4907114642f798976aa44e725ff7af35b3ba2fa57d406c071`; same schema, so the swap back is binary-only (record: be-oau). Its `.1` tag is dead: its pin line named `551e8b27c`, the v1.3.0 tag object — a fleet pin is always `<tag>^{commit}`. **Facts for the next window (read 2026-10-05, not a plan):** six of the rows' upstream merges — #6130 `176463709`, #6145 `916887c5a`, #6291 `f7ecd0fc93`, #6215 `ab88b2051`, #6993 `1f6f2c41f`, #6475 `ab258606f` — are on `upstream/main` and none is an ancestor of `v1.3.1`, so a pin past them drops up to nine rows by the empty-pick test; and a pin at today's `upstream/main` moves the schema 66 → **69** (ignored 26 → 27) — a migrating window, not the binary-only swap of the last two units. **History, none a build source:** `v1.3.0-rc.2-fleet.20260915.3` @ `ec049b27c` (rc.2 + 11, the first v66 unit; be-bs7 and be-wgz dropped at that rebase by the empty-pick test) and the v59 lineage, frozen at `v1.1.1-fleet.20260915` @ `52e292689` (pin `bf97b73749ac` + 13) — a v59 bd refuses a v66 store, so it rolls back no store in the fleet. Their gate records and the rc.2 re-expression lessons are on be-oau and Cherub's ga-lpaaf7; their branches stay on the mirror, every `carry-v59/*` branch included. |
| `carry/operational-v1.1` | Cherub town's carry lineage on the same pin. Owned by Cherub's bd owner; Alex town never builds from, rebases or pushes to it. | @ `011c3eed5` = pin + 7 as measured 2026-09-10 (was +6 on 2026-09-08); its ledger lives on Cherub's beads. Fold to the joint lineage decided 2026-09-10 (ga-lpaaf7): 4 rows cherry-picked into the joint window branch (ledger below), 3 dropped with evidence; freezes as a rollback pin once the fleet installs the joint tag. |
| `backup/*` | Pre-rewrite heads kept for rollback and forensics; never built, never deleted without both towns. | `backup/carry-operational-{local,origin}-pre-rebase-20260901` — the heads of the accumulated-carry model that the 2026-09-01 realignment ended. |
| `fix/*`, `feat/*`, `perf/*`, `test/*`, `docs/*` | Upstream contribution branches: cut from fetched `upstream/main`, pushed to a **personal** fork with *Allow edits by maintainers* on, PR against `gastownhall/beads` `main`. Never on this mirror. | No live PR branch on this mirror. The grandfathered exception ended when #6031 closed 2026-09-30 (replaced by #6993). Two stale contribution branches remain, both of closed PRs: `fix/dolt-ignore-temp-intermediates` @ `7bc34d519` (#6031) and `fix/graph-html-viewport-legend` @ `b6b728e4f` (#4117, merged 2026-07-07) — their author deletes them; the SHAs here are the way back. |

Remotes convention (verify with `git remote -v`, never assume): `upstream` =
`gastownhall/beads` (read-only, the PR target); the shared org mirror
`quickserve-ai/beads` (this repository: `main`, `carry/*`, `carry-v*/*`,
`backup/*`); and each contributor's personal fork for PR branches. Remote
names are local to a checkout; the bd owner's town header records them.

## Carry ledger

Authoritative enumeration of the build lineage: `git log --no-merges
<pin>..<live carry-v branch>`; of this branch: `git log --no-merges
upstream/main..carry/operational`. Every row is an intentional divergence with
a written drop-or-carry decision, reviewed at every refresh, default **drop**
(principle 4; Bar 3). A rebase that silently drops a row is a regression; a
row whose drop condition is met is removed with evidence, never by assumption.

**The drop test is an EMPTY CHERRY-PICK onto the new base, not patch-id equality.**
Measured 2026-09-15 while sizing the move to `v1.3.0-rc.2`: the two rows that
genuinely drop there are *not* patch-id equal to their upstream originals
(`3d0d62858` → `d6081ed24` vs `60b15f2e6` → `75983f29f`; `05ba366c4` →
`8512087829` vs `02c1d516e` → `1103a70bb`). That is context drift from the
original cherry-pick, not evidence against the drop, and a reviewer trusting
patch-id would have carried both rows forward for nothing. The authoritative
test is two steps: `git merge-base --is-ancestor <upstream sha> <new base>`,
then `git cherry-pick -n <upstream sha>` onto the new base leaving an EMPTY
working tree. `git range-diff` and `git patch-id` remain useful for *reading*
how a row drifted; they do not decide the drop. The behaviour check still
applies to every row. *(Rule corrected by the Alex-town bd owner 2026-09-15 on
measurement and sent to Cherub's bd owner the same hour; it narrows a test that
was producing false negatives rather than changing what the ledger requires.)*

**Audit a carry row against the FLEET HEAD, never against the base tag.** The
base tag is by definition the revision the carries are *not* in yet, so a
textual search there reports every live row as lost. Measured 2026-09-17 on
this lineage: the sentinel refusal `ga-emfu6w` adds reads **0 occurrences at
the base tag `c185735c3` and 1 at the fleet head `ec049b27c`**, and
`cmd/bd/close.go` is byte-identical to the tag while differing from the fleet
head by +58 lines (`5febfb486`, `ga-inpgj6`). Both rows are present; an audit
pointed at the tag says both are gone. This is not hypothetical — Cherub's bd
owner came within one filing of recording `ga-emfu6w` as a DROPPED KEEP on
exactly this reading, and **a false DROPPED KEEP is worse than a missed one**,
because it sends someone re-doing work that already exists and then landing it
twice. The cheap second signal, when a row is hard to grep for: **a row that
landed brought its test** — `id_parser_sentinel_test.go` is present at the
fleet head. *(Found by katya, Cherub town, 2026-09-17; verified on both rows
and written here by the Alex-town bd owner the same hour.)*

### Code carries — on the pin (`carry-v66-1.3.1/fleet-20261005`, base `v1.3.1` = `c1c4b642ac`)

Keyed by the live lineage's commit; `← <sha>` is the same row in the
2026-09-23 unit (its `-x` trailer), for reading older records. Rows marked †
are Cherub town's: updated here from their commit messages and fork PR #5's
rows, theirs to correct (see the last rule below).

| Commit | Bead | What / why | Upstream | Drop when |
|---|---|---|---|---|
| `27566d312` ← `c00a94649` perf(ready): recurse the descendant walk off indexed base tables | be-qfm | `bd ready --parent` walked descendants through a non-recursive CTE Dolt cannot index through (5.0 s vs 0.14 s on identical rows on the shared hub). | #6130 — **merged 2026-10-01, `176463709`**, after maintainer adoption: `706247f52` adds the Bazel `srcs` rows and a `depends_on_external` target arm this row lacks, so the merged form supersedes ours rather than equalling it. | A pin at or after `176463709` — empty-pick test above, then drop. Ours is not re-expressed; upstream's form replaces it. |
| `61e456cb9` ← `dc9c8e1bc` perf(ready): compute the --parent descendant set once per call | be-qfm | The same walk ran twice per call (issues leg and wisp leg). Narrowed upstream since: #6731 already walks once on the `bd ready --json` path; `GetReadyWorkInTx` still walks twice. | #6131 — open; rebased onto `main` 2026-10-01 (head `ba9f2728a2`, so not byte-equal to this row), re-approved, mergeable; awaiting a maintainer merge. | A pin past #6131's merge — empty-pick test, then drop. |
| `ac8fdee40` ← `5679250e4` test(embeddeddolt): gate the descendant-walk reference test on cgo | be-qfm | The reference test's file needed `//go:build cgo` for CI's boundary job. | Part of #6130 (merged `176463709`) | With `27566d312`. |
| `eb49ef907` ← `d49e4e2ef` fix(dolt): apply the pool read/write deadline knobs on every open path | be-4at | The configured pool deadline knobs were not applied on every open path, so a configured deadline did not reach some connections. | #6145 (Fixes #6144) — **merged 2026-10-01, `916887c5a`**, after maintainer adoption | A pin at or after `916887c5a` — empty-pick test, then drop. |
| `540b626c5` ← `a98fdeb68` fix(dolt): narrow the pool-deadline claim to DoltStore opens, disclose the import precedence | be-4at | #6145's earlier head `a5f3fc491` — the config.yaml pool-deadline rung reaches every DoltStore open via a `GetStringFromDir` fallback (`poolTimeoutFromConfig`), the claim narrowed to DoltStore opens, the import precedence disclosed. | #6145 — merged as `916887c5a`; the maintainer's adoption kept this rung. | With `eb49ef907`. |
| `1d7665b06` ← `3080ce722` ci(fleet): fleet-release workflow builds and publishes bd fleet units on v*-fleet.* tags | be-xnr | The fork's own bd fleet-release workflow, twin of quickserve-ai/gascity `.github/workflows/fleet-release.yml` — builds and publishes `bd-<tag>-<platform>` on annotated `v*-fleet.*` tags whose message names the pin (always `<tag>^{commit}`, never a tag object), stamping `main.Version=<tag without v>`, `main.Build=<pin7>+carry.<sha7>`. | none — a fork-only release mechanism (no config alternative for CI on a fork) | Kept at every refresh under Bar 3's own test — CI on a fork has no configuration alternative — reviewed like every row; leaves when the fork ends or upstream ships a fleet-release mechanism both towns adopt. |
| `09177bc02` ← `e1c2b8729` ci(fleet): the fleet-tag grammar accepts a prerelease base | be-xnr | The tag grammar accepted only a plain `v<X.Y.Z>` base, so the first v66 tag (`v1.3.0-rc.2-fleet.20260915`) was refused by its own release workflow (run 34986054096, "is not a fleet tag"). An optional prerelease segment, excluding `-`, keeps the `-fleet.` separator unambiguous. | none — fork-only, with `1d7665b06` | With `1d7665b06`. |
| `80e0f222c` ← `7126c2499` perf(issueops): batched is_blocked mark/unmark decide membership through the batch-scoped should-be-blocked union (#6291) | be-vpc | The four batched `is_blocked` mark/unmark templates spelled out five correlated EXISTS per outer row, re-executed per candidate row; `bd close` on a bead with 15-35 dependents took minutes on the westeros qcore store, timing the fixer out at 120 s and stopping the review ladder for 31 minutes on 2026-09-14. Cherry-pick of a **merged** upstream commit (`f7ecd0fc93`, 2026-09-12) that is in no release the fleet has pinned: `v1.3.0` and the `v1.3.1` release branch both lack it. The scope half of #5939 stays open upstream (`RecomputeIsBlockedInTxWithResult` reruns the whole affected set to a fixpoint). | #6291 (merged 2026-09-12, `f7ecd0fc93`; fixes #6288) | A pin at or after `f7ecd0fc93` — empty-pick test, then drop. |
| `e8e2f975e` ← `3a7d9158f` fix(carry): doctor Blocked State check runs pinned on a long-timeout handle † | ga-fo8w65 | The blocked-consistency COUNT walks correlated EXISTS over every issue; against a remote hub it exceeds the pooled 10 s read deadline and dies as 'invalid connection' — the one check that detects stale `is_blocked` rows was blind on exactly the shared multi-writer store where staleness is likeliest. Re-expressed by Cherub town on the v1.3.0 base's `withReadTxLongTimeout`, which also pins the store branch. **Never pick the v59-era original `a295277b9`**: on rc.2 and later it reads the unpinned branch and returns a confident all-clear (measured 2026-09-15, record be-oau). | #6216 (open) | Upstream routes the blocked-consistency check through a long-timeout handle and the pin advances past it. |
| `c8a421c6b` ← `ac2effb6d` fix(carry): a client deadline kill must not blame the transport † | ga-2xwhcz | The MySQL driver reports its own read-deadline kill as 'invalid connection' — the same string a dead server produces; during the 2026-08-31 stalls a slow server read as a broken transport, twice, at hours of cost. Adds the elapsed-vs-deadline discriminator (the deadline read back off the DSN, so it stays one value with the be-4at knobs) and rewrites such failures to name the real cause; retry classification deliberately unchanged. Re-expressed by Cherub town off rc.2. | upstream hole confirmed by #6483; upstream PR #6220 (open) | Upstream ships an equivalent deadline-kill discriminator and the pin advances past it. |
| `6c9ce71f5` ← `3d26163b9` fix(carry): refuse jq/JS sentinel tokens (null, undefined, empty) as issue IDs † | ga-emfu6w | `jq -r` prints the literal string `null` when a selector misses, and ID resolution substring-matches, so the fleet's most common shell idiom silently mutated an arbitrary unrelated issue whose hash contained the token. Refuses the bare sentinels before any lookup; prefixed or longer forms still resolve. | #6215 — **merged 2026-09-30, `ab88b2051`** | A pin at or after `ab88b2051` — empty-pick test, then drop. |
| `1c5d5c081` ← `0135f008a` feat(close): cascade-close molecule step-children when a molecule root closes † | ga-inpgj6 | `bd close` cascades up (last step auto-closes the root) but had no downward inverse: a molecule/ephemeral root closed directly left its open parent-child step-children orphaned forever. Adds `cascadeCloseMoleculeSteps` (recursion-safe; plain epics untouched). bd-CLI surface only — gc's in-process close paths are unaffected. | none yet — an upstream PR is the exit path (Cherub town to open) | Upstream ships the downward close cascade (our PR or equivalent) and the pin advances past it — the empty-pick test, plus the behaviour check. |
| `0e1e1aab4` ← `a1ef4c4c0` fix(schema): dolt_ignore the `__temp__` table-rebuild intermediates | be-c5h | The ignored-series table rebuilds stage every clone-local table through a `__temp__<table>` intermediate whose name carried no `dolt_ignore` pattern, so against a `@@dolt_transaction_commit=1` server each `CREATE TABLE __temp__X` auto-commits a real tracked table at HEAD and the rename onto the ignored final name leaves an unstageable rename half, wedging every later open behind the dirty-table guard. Two independent halves: `__temp__%` in `doltIgnorePatterns`, and `--skip-empty` on the seed's `DOLT_COMMIT` in `commitSeededDoltIgnore` — the second is the one load-bearing on the hub, where the seed's INSERTs are already committed at their own transaction boundaries and a labeled commit without it dies with "nothing to commit". | #6031 closed 2026-09-30 for the maintainer's replacement #6993, **merged 2026-10-01, `1f6f2c41f`**: both halves present there (`schema.go:458` and `:632`, read 2026-10-05), plus a forced stage of the #4356 untrack-scratch sweep. That sweep misses a scratch table present only at HEAD (be-zid; the fix is #7075, open). | A pin at or after `1f6f2c41f` with **both** halves verified at the pin — then drop. Until then, the next window replaces this row with a pick of upstream's `1f6f2c41f`, and of #7075's merged head once it lands (decided 2026-10-05 on be-qfm; not yet built or measured on a base). |
| `b3d3ad0f5` ← `b59bcd243` fix(init): the reinit preflight must not be the thing that migrates † | ga-ylug59 | `bd init`'s `countExistingIssues` preflight opened the store WRITABLE under a 5 s deadline and ignored the error, so against a slow server it half-migrated a fresh hq and then refused it as "already initialized" — the carry CI `rest-smoke-2` flake (#5920). The preflight now opens read-only and reports rather than swallows. | upstream advice on gastownhall/beads#5628 (comment 5712916752) recommends the same two-part shape; upstream PR to follow from Cherub town | Upstream ships a read-only, non-migrating init preflight (our PR or equivalent) and the pin advances past it — the empty-pick test. |
| `aa9acb5cf` ← `530d55356` fix(ga-knhu61): a gate created without await_type is invisible to every notifier — default 'human' at the create seam, add `bd update --await-type` † | ga-knhu61 | `bd create -t gate` left `await_type` NULL, and every notifier keys on it, so a machinery gate was invisible to all of them (ga-49tby1 part 1). Default `human` in the shared constructor, an `--await-type` verb on update, `await_type` in list JSON. Was merged to `carry/operational` as fork PR #3 on 2026-09-22 contrary to that branch's contract (doc-only); reverted there by fork PR #5 — the code lives ONLY in the fleet lineage. | none yet — upstream PR is the exit path (Cherub town to open) | Upstream defaults `await_type` at the gate create seam and exposes the update verb, and the pin advances past it. |
| `735cc83d4` ← `fad1d3416` fix(ga-knhu61): review round — the update verb was a silent no-op on the proxied route, and the seam was not single † | ga-knhu61 | `IssuePatch` gained `AwaitType` so `bd update --await-type` reaches the proxied route; the default moved to one seam. | with the row above | With `aa9acb5cf`. |
| `4620c6c19` ← `59bf6ca3f` review: carry max-conns through the same config.yaml fallback; fold tests into the ladder † | ga-knhu61 (review round; the knob itself is be-4at's ladder) | The `dolt.max-conns` rung reaches every DoltStore open through the same `poolTimeoutFromConfig` fallback the pool deadlines use, and its tests join the ladder test rather than standing alone. | #6145 family (see be-4at). Upstream #6475, **merged 2026-09-27, `ab258606f`**, is this row's upstream form — confirmed by Cherub town's bd owner 2026-10-06: `ab258606f`'s message carries this commit's subject and its `open.go` reads `poolCfg("dolt.max-conns")`. | A pin at or after both `916887c5a` (#6145) and `ab258606f` (#6475) — empty-pick test, plus the behaviour check on `dolt.max-conns`. |
| `3ecd84346` ← `3d86fb056` fix(ga-knhu61): classify an AwaitType patch as a non-coordination signal † | ga-knhu61 (Alex town's fixup on Cherub's row, reviewed by katya 2026-09-23) | `fad1d3416` added `AwaitType` to `IssuePatch` but not to `nonCoordinationPatchSignals`, so an update setting ONLY `await_type` still classified coordination-only and dropped — the guard tests (`TestNonCoordinationPatchSignalsCountMatchesStruct`, `TestHasNonCoordinationPatchClassifiesEveryScalarField`) were red at Cherub's fleet head `3180d6674`, so unit `.20260922.1` carries the defect; green at `3d86fb056`. | with `aa9acb5cf` | With `aa9acb5cf`. |

### Documentation carries — on `carry/operational` (base `upstream/main`)

| Files | Bead | What / why | Drop when |
|---|---|---|---|
| `CARRY.md` (this file), `CONTRIBUTING.fork.md` | be-vfh | The fork's branch contract and ledger, and the beads-repository contribution mechanics both towns' bd owners work to — placed in the repository they govern (qc-bridge `shared/repository-ownership.md`, placement test 1) by the Alex-town operator's decision of 2026-09-08 (option 2 of three; recorded on be-vfh and in qc-bridge `roles/bd-owner.md` § *Placement — decision record*). Banner-marked, minimal, documentation only. | `CONTRIBUTING.fork.md`: upstream takes an equivalent contributor document, or the two bd owners retire it under the change rule in `roles/bd-owner.md`. `CARRY.md`: the last carry leaves and the fork ends. Until then **exempt from the default drop**: rebased forward at every refresh (Bar 3's stated exemption for a minimal, clearly marked documentation carry). |

## Rules specific to this fork

- **A marked documentation carry is exempt from Bar 3's default drop.** It is
  rebased forward at every pin refresh and leaves only on its own file-specific
  condition in the documentation ledger above: the contributor file when
  upstream takes an equivalent document or the two owners retire it, this
  ledger when the last carry leaves and the fork ends. The exemption covers
  exactly the files in the documentation
  ledger above, each carrying the banner at its top; anything else on
  `carry/operational` is a mistake — move it to a `carry-v*` branch with a
  ledger row, or drop it.
- **Carry files live only here.** Never on a PR branch (cut from
  `upstream/main`), never on `main` (a fast-forward mirror), never on a
  `carry-v*` build branch (code picks only, so a window build's diff against
  the pin is the ledger's code rows and nothing else).
- **Refreshing `carry/operational`** (at a pin advance, or whenever the doc
  commit should sit on a newer upstream): fetch `upstream` and the mirror's
  `carry/operational` explicitly (`+refs/heads/carry/operational:refs/remotes/origin/carry/operational`),
  capture the expected OID, rebase the doc commit(s) onto `upstream/main`, then
  `git push origin HEAD:carry/operational --force-with-lease=carry/operational:<expected OID>`.
  Never bare `--force`; a rejected lease means unseen shared work — stop and
  reconcile. `main` is never force-pushed.
- **Consumers of this branch follow a refresh by reset, never by fast-forward.**
  Every nested import checkout of `carry/operational` (each bd owner's seat,
  both towns) is updated after a refresh by fetching the exact new OID and
  moving a verified-clean tree onto it:
  `git fetch origin '+refs/heads/carry/operational:refs/remotes/origin/carry/operational'`,
  then `test -z "$(git status --porcelain)"` — refuse a dirty or foreign-owned
  checkout rather than merge or discard anything — then
  `git reset --keep origin/carry/operational` (aborts by itself if a local
  change would be lost; the old head stays in the reflog). A rebase makes
  `git pull --ff-only` impossible here (measured by Cherub's bd owner
  2026-09-08 in an isolated repository: exit 128, "Not possible to
  fast-forward"); `--ff-only` fits `main`, the fast-forward mirror, and
  nothing else on this repository.
- **Adding a code carry**: one cherry-pick per commit onto the pin on
  `carry-v<schema>/<slug>`, built and tested at the pin, then a cumulative
  window branch (`<live branch's slugs>-<new slug>`) — the live build's stack
  plus the new pick, never the new slug's branch alone. Add the ledger row
  here in the same change, with the bead and the drop condition, before the
  handover to the gascity pin owner.
- **A fleet tag may be cut before both gates; a fleet tag may NOT be installed
  before both gates.** *(Ruled by Cherub's bd owner 2026-09-15 on the Alex-town
  bd owner's question, after the `v1.1.1-fleet.20260915` cut went ahead of
  Cherub's gate; binding on both towns, recorded here and on Cherub's
  `ga-lpaaf7`.)* The protection that matters is the INSTALL, not the tag: a tag
  is an artifact, an install is a fleet change. **Strict, always: no machine
  installs a fleet tag until both towns' gates are recorded on both ledgers.**
  Cutting ahead of the other town's gate is legitimate when three things hold
  together — a live incident the tag fixes, a same-day `.n` offer standing, and
  the tag installed nowhere until both gates land. The 2026-09-15 cut met all
  three. Absent an incident, the gates come first.

- **Both towns read this ledger.** A row that is the other town's (their
  lineage, their bead) is theirs to change; a change to the rules in this
  section is agreed between the two bd owners on the bridge first.
