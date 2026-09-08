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

| Branch | Contract | Reality (2026-09-08) |
|---|---|---|
| `main` | Fast-forward-only mirror of `gastownhall/beads` `main`. Never a carry commit, never a PR branch. Refresh with `git push origin upstream/main:main` after fetching upstream; a personal fork's `main` is kept equal the same way. | `= upstream/main` (`2bb1e20de`, schema head 0067). |
| `carry/operational` | `upstream/main` plus the **marked documentation carries** in the ledger below — this file and `CONTRIBUTING.fork.md` — and nothing buildable. **Not a build source**: the fleet never builds `bd` from a branch that floats with upstream. Refreshed by rebasing the doc commit(s) onto the new `upstream/main` (force-with-lease on this branch only, never on `main`). | `= upstream/main + 1` (the commit that adds these two files). |
| `carry-v<schema>/<slug>[-<slug>…]` | The fleet build lineage: the upstream SHA that upstream `gastownhall/gascity` `main` pins in `go.mod` (principle 3, "never ahead of support") plus a short stack of cherry-picks — one branch per pick, one cumulative branch per window build. Built and tested at the pin before it is pushed; handed to the gascity pin owner; every pick has a bead and a drop condition. | Pin `bf97b73749ac` (2026-08-05, schema **v59**). LIVE on both Alex-town machines since the 2026-09-03 window: `carry-v59/be-qfm-be-4at-be-bs7` @ `3d0d62858` = pin + 5 (ledger below). Component branches: `carry-v59/be-qfm` (+3), `carry-v59/be-4at` (+1), `carry-v59/be-bs7` (+1), `carry-v59/be-qfm-be-4at` (+4). |
| `carry/operational-v1.1` | Cherub town's carry lineage on the same pin. Owned by Cherub's bd owner; Alex town never builds from, rebases or pushes to it. | @ `80f1605c3` = pin + 6; its ledger lives on Cherub's beads. |
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
row whose drop condition is met is removed with evidence — `git patch-id` or
`git range-diff` against upstream plus the behaviour check — never by
assumption.

### Code carries — on the pin (`carry-v59/*`, base `bf97b73749ac`)

| Commit | Bead | What / why | Upstream | Drop when |
|---|---|---|---|---|
| `448baad1f` perf(ready): recurse the descendant walk off indexed base tables | be-qfm | `bd ready --parent` walked descendants through a non-recursive CTE Dolt cannot index through (5.0 s vs 0.14 s on identical rows on the shared hub). | #6130 (approved 2026-09-03, unmerged) | The pin advances to a revision containing #6130's merge; prove by patch-id / range-diff, then drop. |
| `fba96444a` perf(ready): compute the --parent descendant set once per call | be-qfm | The same walk ran twice per call (issues leg and wisp leg). | #6131 (approved, unmerged) | As above, for #6131. |
| `4f62045df` test(embeddeddolt): gate the descendant-walk reference test on cgo | be-qfm | The reference test's file needed `//go:build cgo` for CI's boundary job. | Part of #6130 | With #6130. |
| `8f7471e01` fix(dolt): apply the pool read/write deadline knobs on every open path | be-4at | The configured pool deadline knobs were not applied on every open path, so a configured deadline did not reach some connections. Cherub carries the same pick on `carry/operational-v1.1`. | #6145 (Fixes #6144; approved, unmerged) | The pin advances past #6145's merge. |
| `3d0d62858` perf(recompute): full is_blocked repair runs unbatched semi-join statements (#5678) | be-bs7 | `bd recompute-blocked` could not finish inside its 5 m deadline on a hub-shaped store at the pin (19 m cold); upstream's #5678 makes it sub-second. Cherry-pick of a **merged** upstream commit (`60b15f2e6`, 2026-08-12). | #5678 (merged) | Any pin advance to a revision at or after `60b15f2e6` — expect patch-id equal; confirm and drop. |

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
- **Adding a code carry**: one cherry-pick per commit onto the pin on
  `carry-v<schema>/<slug>`, built and tested at the pin, then a cumulative
  window branch (`<live branch's slugs>-<new slug>`) — the live build's stack
  plus the new pick, never the new slug's branch alone. Add the ledger row
  here in the same change, with the bead and the drop condition, before the
  handover to the gascity pin owner.
- **Both towns read this ledger.** A row that is the other town's (their
  lineage, their bead) is theirs to change; a change to the rules in this
  section is agreed between the two bd owners on the bridge first.
