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

| Branch | Contract | Reality (2026-09-10) |
|---|---|---|
| `main` | Fast-forward-only mirror of `gastownhall/beads` `main`. Never a carry commit, never a PR branch. Refresh with `git push origin upstream/main:main` after fetching upstream; a personal fork's `main` is kept equal the same way. | `= upstream/main` (`a690b0a8c`, 2026-09-08, schema head 0067). |
| `carry/operational` | `upstream/main` plus the **marked documentation carries** in the ledger below — this file and `CONTRIBUTING.fork.md` — and nothing buildable. **Not a build source**: the fleet never builds `bd` from a branch that floats with upstream. Refreshed by rebasing the doc commit(s) onto the new `upstream/main` (force-with-lease on this branch only, never on `main`). | = `2bb1e20de` (the base is still the previous `upstream/main`, one behind the mirror; rebase at the next refresh) + the documentation commits in the ledger below (3 on 2026-09-10). Import checkouts follow a refresh by fetch + verified-clean `reset --keep`, never `pull --ff-only` (rules below). |
| `carry-v<schema>/<slug>[-<slug>…]` | The fleet build lineage: the upstream SHA that upstream `gastownhall/gascity` `main` pins in `go.mod` (principle 3, "never ahead of support") plus a short stack of cherry-picks — one branch per pick, one cumulative branch per window build. Built and tested at the pin before it is pushed; handed to the gascity pin owner; every pick has a bead and a drop condition. | **LIVE (2026-09-16): the `carry-v66` lineage.** Base `v1.3.0-rc.2` = `c185735c38e2` (schema **v66**), the revision upstream `gastownhall/gascity` `main` now requires in `go.mod` — so principle 3 moved the base, and this is a REBASE onto a new base, not a fast-forward (the v59 pin is not an ancestor of rc.2). Build source: `carry-v66/be-qfm-be-4at-be-xnr-be-vpc-ga-emfu6w-ga-inpgj6-be-c5h` @ `ec049b27c` = base + 11, released `v1.3.0-rc.2-fleet.20260915.3` (2026-09-16 04:30:41Z; sha256 darwin-arm64 `97e459c1870f18660fc36aeffde29784f7e1a3816b5529c5241720055cb5382c`, linux-amd64 `25c933cde4c8e47468490cf727b26f00dde8c33d46f776d22df775b1a0c160b6`) and pinned by gascity through a full module `replace`. Superseded tags on this lineage, all still published: `.0` @ `08e4650bd` (8 carries), `.1` @ `e640d174d` (+`ga-emfu6w`), `.2` @ `5febfb486` (+`ga-inpgj6`). **Two rows DROPPED at the rebase** — be-bs7 (#5678) and be-wgz (#5343), both ancestors of rc.2, proven by EMPTY CHERRY-PICK (not patch-id; see the drop-test rule below). be-vpc is picked from upstream's own `f7ecd0fc93`, not from our v59-rebased copy, because the original applies clean on rc.2 where ours conflicts. Cherub's `ga-fo8w65` and `ga-2xwhcz` are re-expressed off rc.2 by katya (`0b1b5f8ef`, `835feeccd`, unpushed at time of writing) and rejoin at the NEXT window — **do not mechanically pick the originals**: `ga-fo8w65`'s original is measured-WRONG on rc.2 (its `LongTimeoutDB` is byte-identical to rc.2's `openLongTimeoutConn`, and rc.2 wraps that in `withReadTxLongTimeout`, which also pins the store branch; planted-inconsistency measurement on an isolated server reads 1 pinned vs **0 unpinned WITH A NIL ERROR** — a confident all-clear from the check whose blindness opened the bead). **FROZEN as the v59 rollback point:** Pin `bf97b73749ac` (2026-08-05, schema **v59**); its last window source was `carry-v59/be-qfm-be-4at-be-bs7-be-xnr-ga-inpgj6-ga-fo8w65-ga-emfu6w-ga-2xwhcz-be-vpc` @ `52e292689` = pin + 13 — the pin + 12 joint head `bb6db8ceb` plus the incident-critical #6291 pick (be-vpc, 2026-09-15), pushed 2026-09-15 and tagged `v1.1.1-fleet.20260915` the same day. **Joint gate COMPLETE, both towns recorded.** Alex town: build green under the release workflow's own `CGO_ENABLED=1 -tags gms_pure_go`, `go vet` clean on the touched packages, 61 internal packages pass with 0 failures, the `Blocked|IsBlocked|Recompute` family green with 0 skips, and a cross-family adversarial read finding no carry-specific regression. Cherub town (katya, 2026-09-15, verdict **LINEAGE CLEAN, no `.1`**): the four `ga-*` rows unchanged and their ledger rows byte-identical to 2026-09-10, be-vpc's changed lines byte-identical to upstream #6291, issueops 109/109 and the embedded-Dolt bd suites 21/22 top-level with 0 skip at the tag. The single failure is pre-existing in Cherub's own `ga-inpgj6` row — it fails identically at the 2026-09-10 joint head and works on the `--force` path gc uses — filed Cherub-side, not a blocker for this tag. Cherub's durable record is `ga-lpaaf7`. Prior heads: `…-ga-2xwhcz` @ `bb6db8ceb` = pin + 12 (the joint head Cherub's 4 kept rows landed on, 2026-09-10); and the superseded window source `carry-v59/be-qfm-be-4at-be-bs7-be-xnr` @ `1370260e6` = pin + 8 (be-qfm ×3, be-4at ×2, be-bs7, be-wgz, be-xnr), tagged `v1.1.1-fleet.20260910`. Component branches: `carry-v59/be-qfm` (+3), `carry-v59/be-4at` (+2, refreshed 2026-09-10 to `9913fcb33`), `carry-v59/be-bs7` (+1), `carry-v59/be-xnr` (+1), `carry-v59/be-qfm-be-4at` (+4), `carry-v59/be-qfm-be-4at-be-bs7` (+5, the LIVE build on both Alex-town machines since the 2026-09-03 window), `carry-v59/be-vpc` (+1 **on `carry-v59/be-bs7`**, not on the bare pin — #6291 needs bd-t9ypt's union). |
| `carry/operational-v1.1` | Cherub town's carry lineage on the same pin. Owned by Cherub's bd owner; Alex town never builds from, rebases or pushes to it. | @ `011c3eed5` = pin + 7 as measured 2026-09-10 (was +6 on 2026-09-08); its ledger lives on Cherub's beads. Fold to the joint lineage decided 2026-09-10 (ga-lpaaf7): 4 rows cherry-picked into the joint window branch (ledger below), 3 dropped with evidence; freezes as a rollback pin once the fleet installs the joint tag. |
| `backup/*` | Pre-rewrite heads kept for rollback and forensics; never built, never deleted without both towns. | `backup/carry-operational-{local,origin}-pre-rebase-20260901` — the heads of the accumulated-carry model that the 2026-09-01 realignment ended. |
| `fix/*`, `feat/*`, `perf/*`, `test/*`, `docs/*` | Upstream contribution branches: cut from fetched `upstream/main`, pushed to a **personal** fork with *Allow edits by maintainers* on, PR against `gastownhall/beads` `main`. Never on this mirror. | One grandfathered exception: #6031's branch predates the rule and stays on the mirror until that PR merges or closes. |

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

### Code carries — on the pin (`carry-v59/*`, base `bf97b73749ac`)

| Commit | Bead | What / why | Upstream | Drop when |
|---|---|---|---|---|
| `448baad1f` perf(ready): recurse the descendant walk off indexed base tables | be-qfm | `bd ready --parent` walked descendants through a non-recursive CTE Dolt cannot index through (5.0 s vs 0.14 s on identical rows on the shared hub). | #6130 (approved 2026-09-03, unmerged) | The pin advances to a revision containing #6130's merge; prove by the empty-cherry-pick test above (`range-diff` only to read how the row drifted), then drop. |
| `fba96444a` perf(ready): compute the --parent descendant set once per call | be-qfm | The same walk ran twice per call (issues leg and wisp leg). | #6131 (approved, unmerged) | As above, for #6131. |
| `4f62045df` test(embeddeddolt): gate the descendant-walk reference test on cgo | be-qfm | The reference test's file needed `//go:build cgo` for CI's boundary job. | Part of #6130 | With #6130. |
| `8f7471e01` fix(dolt): apply the pool read/write deadline knobs on every open path | be-4at | The configured pool deadline knobs were not applied on every open path, so a configured deadline did not reach some connections. Cherub carries the same pick on `carry/operational-v1.1`. | #6145 (Fixes #6144; approved, unmerged) | The pin advances past #6145's merge. |
| `3d0d62858` perf(recompute): full is_blocked repair runs unbatched semi-join statements (#5678) | be-bs7 | `bd recompute-blocked` could not finish inside its 5 m deadline on a hub-shaped store at the pin (19 m cold); upstream's #5678 makes it sub-second. Cherry-pick of a **merged** upstream commit (`60b15f2e6`, 2026-08-12). | #5678 (merged) | Any pin advance to a revision at or after `60b15f2e6`. **Verified DROP against `v1.3.0-rc.2` 2026-09-15**: ancestor, and the cherry-pick comes out empty. Not patch-id equal — see the drop-test note above. |
| `327e0869c` fix(dolt): narrow the pool-deadline claim to DoltStore opens, disclose the import precedence | be-4at | #6145's head `a5f3fc491` — the config.yaml pool-deadline rung reaches every DoltStore open via a `GetStringFromDir` fallback (`poolTimeoutFromConfig`), the claim narrowed to DoltStore opens, the import precedence disclosed. | #6145 (approved 2026-09-03, unmerged; Cherub's #6475 covers the same hole at `applyResolvedConfig` only and becomes the max-conns rung after #6145 merges) | The pin advances past #6145's merge. |
| `05ba366c4` Fix: wisp delete never cascades to auxiliary tables (orphan row leak) (#5343) | be-wgz | wisp delete never cascaded to `wisp_labels` / `wisp_events` / `wisp_comments` / `wisp_child_counters` (the orphan rows hq purges weekly, be-ctw); gc's in-process Delete leaks the same rows (navani 2026-09-10). Cherry-pick of a **merged** upstream commit (`02c1d516e`, 2026-08-07). | #5343 (merged 2026-08-07, `02c1d516e`) | Any pin advance to a revision at or after `02c1d516e`. **Verified DROP against `v1.3.0-rc.2` 2026-09-15**: ancestor, and the cherry-pick comes out empty. Not patch-id equal — see the drop-test note above. |
| `1370260e6` ci(fleet): fleet-release workflow builds and publishes bd fleet units on v*-fleet.* tags | be-xnr | The fork's own bd fleet-release workflow, twin of quickserve-ai/gascity `.github/workflows/fleet-release.yml` at `dbfb52ec9` — builds and publishes `bd-<tag>-<platform>` on annotated `v*-fleet.*` tags whose message names the pin, stamping `main.Version=<tag without v>`, `main.Build=<pin7>+carry.<sha7>`. | none — a fork-only release mechanism (no config alternative for CI on a fork) | Kept at every refresh under Bar 3's own test — CI on a fork has no configuration alternative — reviewed like every row; rule wording proposed to Cherub's bd owner on the bridge 2026-09-10, concurrence pending; leaves when the fork ends or upstream ships a fleet-release mechanism both towns adopt. |
| `f49f2fdbd` feat(close): cascade-close molecule step-children when a molecule root closes | ga-inpgj6 | `bd close` cascades up (last step auto-closes the root) but had no downward inverse: a molecule/ephemeral root closed directly left its open parent-child step-children orphaned forever, accumulating as open-issue pollution on every store. Adds `cascadeCloseMoleculeSteps` (recursion-safe; plain epics untouched), embedded-dolt tests. bd-CLI surface only — gc's in-process close paths are unaffected. Cherub-town carry since 2026-06; kept at the 2026-09-10 fold with upstream verified absent (no cascade in `upstream/main` `cmd/bd/close.go`). | none yet — an upstream PR is the exit path (Cherub town to open) | Upstream ships the downward close cascade (our PR or equivalent) and the pin advances past it — the empty-cherry-pick test above, plus the behaviour check. |
| `a295277b9` fix(carry): doctor Blocked State check runs on a long-timeout handle | ga-fo8w65 | The blocked-consistency COUNT walks correlated EXISTS over every issue; against a remote hub it exceeds the pooled 10s read deadline and dies as 'invalid connection' — the one check that detects stale `is_blocked` rows was blind on exactly the shared multi-writer store where staleness is likeliest. Adds `DoltStore.LongTimeoutDB` (readTimeout=5m, one-shot) and routes the check through it. Upstream `main` (post-pin) now has `openLongTimeoutConn`/`withReadTxLongTimeout` machinery but still runs the doctor check on the pooled `UnderlyingDB()` — re-express this carry on that machinery at the next pin advance. | none yet | Upstream routes the blocked-consistency check through a long-timeout handle and the pin advances past it. |
| `e8ebc76ce` fix(carry): refuse jq/JS sentinel tokens (null, undefined, empty) as issue IDs | ga-emfu6w | `jq -r` prints the literal string `null` when a selector misses, and ID resolution substring-matches, so the fleet's most common shell idiom (`B=$(bd list --json … jq -r .id); bd update "$B"` piped) silently mutated an arbitrary unrelated issue whose hash contained the token. Refuses the bare sentinels before any lookup; prefixed or longer forms still resolve. Data-corruption-class guard, fleet-wide value. | none yet — upstream PR candidate (small, self-contained) | Upstream refuses bare sentinel tokens in ID resolution and the pin advances past it. |
| `bb6db8ceb` fix(carry): a client deadline kill must not blame the transport | ga-2xwhcz | The MySQL driver reports its own read-deadline kill as 'invalid connection' — the same string a dead server produces; during the 2026-08-31 stalls a slow server read as a broken transport, twice, at hours of cost. Adds the elapsed-vs-deadline discriminator (the deadline read back off the DSN, so it stays one value with the be-4at knobs) and rewrites such failures to name the real cause; retry classification deliberately unchanged. | upstream hole confirmed by #6483 (the deadline error never names the knob); upstream PR is gastownhall#6220 (lyft's enumeration; owns deadlines upstream per the 2026-09-10 #6475/#6145 split) | Upstream ships an equivalent deadline-kill discriminator and the pin advances past it. |
| `52e292689` perf(issueops): batched is_blocked mark/unmark decide membership through the batch-scoped should-be-blocked union (#6291) | be-vpc | The four batched `is_blocked` mark/unmark templates (issues/wisps × mark/unmark) spelled out five correlated EXISTS per outer row, re-executed per candidate row; `bd close` on a bead with 15-35 dependents took minutes on the westeros qcore store (61,267 beads), timing the fixer out at 120 s and stopping the review ladder for 31 minutes on 2026-09-14 (`recompute is_blocked (mark): context canceled`). Same statement as the westeros graph-apply faults. Cherry-pick of a **merged** upstream commit (`f7ecd0fc93`, 2026-09-12), in no release: `v1.3.0-rc.2` and `v1.1.1-fleet.20260910` both predate it. **Component branch bases on `carry-v59/be-bs7`, not the bare pin**: #6291 builds on `shouldBeBlockedIDsUnionSQL`, which bd-t9ypt (#5678, be-bs7) introduced and the pin lacks — on the bare pin the pick does not compile. One conflict on the pick, a comment block in `dolt/store.go` where upstream has a `withCircuitWrite` wrapper this base does not; resolved by keeping our function shape and taking upstream's corrected comment, which is true here because this base carries #5678. Resulting diffstat byte-identical to upstream's own #6291 diff. | #6291 (merged 2026-09-12, `f7ecd0fc93`; fixes #6288). The scope half of #5939 stays open upstream: `RecomputeIsBlockedInTxWithResult` still reruns the whole affected set to a fixpoint. | Any pin advance to a revision at or after `f7ecd0fc93` — confirm by empty cherry-pick, then drop. **Not in `v1.3.0-rc.2`**: on that base re-express by picking the UPSTREAM commit `f7ecd0fc93`, which applies clean there, never this fork copy, which conflicts on the `withCircuitWrite` comment block. |
| `ed6b91b8c` fix(schema): dolt_ignore the `__temp__` table-rebuild intermediates | be-c5h | The ignored-series table rebuilds stage every clone-local table through a `__temp__<table>` intermediate, but the intermediate names carried no `dolt_ignore` pattern, so against a `@@dolt_transaction_commit=1` server each `CREATE TABLE __temp__X` auto-commits a real tracked table at HEAD and the rename onto the ignored final name leaves an unstageable rename half, wedging every later open behind the dirty-table guard. **Window-safety row, not a nicety:** migration 0062 copies EVERY events row through `__temp__events_flip` and adds only `events` to `dolt_ignore`, so at `.2` that copy is born TRACKED — a migration killed inside 0062 leaves a full copy of the events table dirty, the retry is refused, and the remedy bd itself prints (`commit -Am`) would publish that copy into versioned history. On the hub 0062 is the single longest *schema* step in the pass (425,791 events rows, read 2026-09-16 17:00Z — **events, not beads**, is the cost axis; the 61,267-bead figure this row used to quote is retired as both stale and the wrong axis), so it is where an interruption inside the schema phase is most likely to land. Measured since: 0060-0066 including 0062's full copy is UNDER 20 s of a ~950 s pass, and the post-cursor rekey tail is ~98 % — so 0062 is the longest schema step but a small share of the whole window. | #6031 (approved 2026-09-03, unmerged; its branch is the grandfathered mirror exception above) | The pin advances to a revision containing #6031's merge — verify **BOTH** halves present in `internal/storage/schema/schema.go`: `__temp__%` in `doltIgnorePatterns`, **and** `--skip-empty` on the seed's `DOLT_COMMIT` in `commitSeededDoltIgnore`. Then drop. **This condition named only the first until 2026-09-17** (caught by syl reading #437): the pick carries two independent changes, and `--skip-empty` is the one that is load-bearing on the hub. Without it, on a `@@dolt_transaction_commit=1` server the seed's INSERTs are already dolt-committed at their own transaction boundaries, so the labeled commit dies with "nothing to commit" and takes the whole pass down — and the westeros hub IS such a server (`/data/dolt-hub/config.yaml` lines 18 and 22). A drop on the first half alone would silently reintroduce a pass-killing failure on exactly the store the window targets. |
| `8283d1c67` fix(init): the reinit preflight must not be the thing that migrates | ga-ylug59 | `bd init` preflight's countExistingIssues opened WRITABLE with a 5s deadline and its error ignored, so a slow server half-migrated a fresh hq and init then refused it (#5920) — the CI rest-smoke-2 flake. Two-part fix: read-only open + `GetStatisticsNoBlocked` (read-only alone would have been a silent-data-loss regression — a deterministic destroy-gate bypass on half-migrated DBs). Behavior-proven by woodhouse's isolated deadline sweep (5/5 at 100–1500ms, 2026-09-17) and settling runs of `bd init --reinit-local --server` on BOTH release contents — `.20260922.1` and the v1.3.0 joint head `3d86fb056` (exit 0 twice each, 32 tables asserted; cherub ga-ylug59). Stock beads defect, not carry-specific (woodhouse 2026-09-17). Ships in `v1.3.0-fleet.20260923.1`. | none yet — an upstream PR is the exit path (Cherub town to open) | Upstream ships a non-migrating read-only init preflight and the pin advances past it — empty-cherry-pick test plus the deadline-sweep behaviour check. |
| `522f564f4` + `c592205ed` fix(ga-knhu61): a gate created without await_type is invisible to every notifier | ga-knhu61 | `bd create -t gate` left await_type NULL, and every notifier filters on it — gates minted by anything but the explicit flag were invisible to the notify plane (9 null qc- gates measured). Defaults 'human' at the SHARED create seam, adds `bd update --await-type`, exposes await_type in list JSON (ga-49tby1 part 1). PR #3's merge to carry/operational violated that branch's doc-only contract and is REVERTED in this same change (lyft's catch, 2026-09-22) — the code carries live on the pinned-base lineage only: re-expressed as `fad1d3416` (+ this row's max-conns sibling `59bf6ca3f`) on the v1.3.0 joint head, WITH lyft's fixup `3d86fb056`: the review-round commit added AwaitType to IssuePatch but not nonCoordinationPatchSignals, so an await_type-only update still no-opped on the proxied route — the exact class the commit fixed; guard tests red at the superseded `.20260922.1`, green at the joint head (measured both towns). OPEN CONTRACT QUESTION at adoption (lyft 2026-09-22, cherub ga-knhu61): the shared-seam default also types gc's in-process formula wait gates 'human', flipping them from notifier-dropped to notifier-live — each town's gc owner answers (cherub: gc passes explicit await_type at its machinery create sites before installing the pairing) before the unit serves. | none yet — an upstream PR is the exit path (Cherub town to open) | Upstream ships a non-NULL await_type default (or mandatory await_type) at the gate create seam plus an update verb, and the pin advances past it — empty-cherry-pick test plus the notifier-visibility behaviour check. |
| `90465c803` review: carry max-conns through the same config.yaml fallback | ga-dwobeb | `dolt.max-conns` had the identical library-consumer hole the pool deadlines had: `config.GetString` reads a viper only cmd/bd populates, so gc (including the supervisor) silently ran the defaults. Carries max-conns through the same config.yaml fallback ladder; red-proved (reverting fails the subtest with MaxOpenConns 0, want 25). Re-expressed onto the rc.2 lineage as `3180d6674`, carried onto the v1.3.0 joint head as `59bf6ca3f` (ships in `v1.3.0-fleet.20260923.1`) — one conflict resolved by KEEPING rc.2's `applyPoolTimeouts`/`poolTimeoutFromConfig` shape (alex be-4at) and taking only the max-conns leg, so the pool-deadline shape survives byte-identical (recorded on be-4at, lyft 2026-09-22). | upstream PR gastownhall#6475 (the max-conns leg; its pool-deadline first commit is superseded by be-4at's shape) | Upstream ships the max-conns config fallback (#6475 or equivalent) and the pin advances past it — empty-cherry-pick test plus the MaxOpenConns red/green subtest. |

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
