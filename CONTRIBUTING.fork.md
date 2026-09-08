> **FORK CARRY FILE — not upstream; lives on `carry/operational` of `quickserve-ai/beads`; dropped when upstream takes an equivalent contributor document or the two bd owners retire it (ledger: `CARRY.md`).**

# Contributing to gastownhall/beads from the quickserve-ai fork

This is the beads-repository half of the bd owner's rulebook: the mechanics
of getting a change into `gastownhall/beads` — their preflight, the rules that
decide whether a PR lands, the quality gates, commit and PR craft, issue
linking, how their merge regime actually behaves, what to do after opening,
and the craft anti-patterns. It is imported into the bd owner's
`CLAUDE.local.md` under qc-bridge `roles/bd-owner.md`, which keeps the
cross-town half — the four bars, the pin rule, the window and its handover,
the carry ledger's review, the shared-store rules and the change rule — and
records why these sections live here (§ *Placement — decision record*). Read
the bars first; this file says *how*, not *whether*.

**Owners:** the bd owners of both towns — `lyft` (Alex town) and `katya`
(Cherub town) — under the change rule stated once, in `roles/bd-owner.md`
§ *Change rule*; it is not restated here. **Adoption:** the role's — Alex
town's seat is bound; Cherub's seat is bound when its bd owner records
concurrence on the role's status line, and until then the owners line names
the seats, not two signatures. **Status:** v1.0 — 2026-09-08, moved verbatim from `roles/bd-owner.md`
v1.2, where these sections were drafted 2026-09-01 and measured against
`upstream/main` on the dates given. Re-measure before trusting a number older
than a month; the rules are what to keep. Upstream's own documents are cited
by path and line as of the measurement date; where a citation and the
document disagree, the document wins and the citation gets fixed here.

**Placement:** this file is a carry (`CARRY.md`, documentation ledger) because
the mirror's `main` is a fast-forward mirror of upstream and cannot hold a
fork-specific file. It leaves the day upstream takes an equivalent contributor
document, or when the two owners retire it; the ledger row in `CARRY.md` is
the record.

## Preflight — before you touch code (their protocol, non-negotiable)

```bash
scripts/pr-preflight.sh --search "<topic keywords>"        # blocks on any matching OPEN PR (scripts/pr-preflight.sh:162-180)
gh search prs    --repo gastownhall/beads "<phrases>" --state all
gh search issues --repo gastownhall/beads "<phrases>" --state all
git log upstream/main --oneline -S'<symbol>' -- <path>     # did upstream already touch this?
bd search "<keywords>"                                     # our own beads
```

Then read the doc that owns the area — `engdocs/PROJECT_CHARTER.md` before
any feature surface, `engdocs/TESTING.md` before tests, `engdocs/LINTING.md`
before touching lint, `engdocs/ICU-POLICY.md` before build tags,
`internal/storage/schema/migrations/README.md` + `scripts/check-migration-hygiene.sh`
before any migration, `engdocs/adr/0002-init-safety-invariants.md` before
`cmd/bd/init*.go`. (`docs/ARCHITECTURE.md`, `docs/LINTING.md`, `docs/ICU-POLICY.md`
no longer exist — architecture is `docs/architecture/index.md` + `engdocs/INTERNALS.md`.)

## Their rules that decide whether a PR lands (docs lane, cited)

- **One concern per PR; no riders, no "while I'm here"**
  (`CONTRIBUTING.md:116-119`; `PR_MAINTAINER_GUIDELINES.md:105`) — and **one
  layer per PR** (`CONTRIBUTING_PR_GUIDELINES.md:19-29`): a fix that needs the
  storage/issueops primitive *and* the `cmd/bd` wiring is two PRs in sequence,
  each citing the same issue (`Refs #N`, with `Fixes #N` on the one that
  closes it). The two rules agree; "one issue per PR" is about riders, not
  about splitting a fix across layers.
- **A few hundred lines is the cap** — "a 13K-line diff will not be reviewed"
  (`:23`); preflight warns at >30 files / >1000 additions.
- **Repro in stock beads; benchmark for perf claims** (`:11-15`).
- **CHANGELOG `[Unreleased]` entry for anything user-visible**, additive
  (Keep a Changelog) — the most frequent fix-round ask (8/14).
- **Enumerate the whole call-site surface**: direct vs proxied-server twins,
  embedded vs server paths, every flag path — the most frequent
  `REQUEST-CHANGES` driver (8/14).
- **Body must match the diff and name the riskiest thing it does**; not a
  draft, not WIP-titled when you want it merged (`PR_MAINTAINER_GUIDELINES.md:100-106`).
- **No `.beads/` data or generated junk in the diff** (`CONTRIBUTING.md:120-121`;
  CI "Check for .beads changes"; preflight blocks any `.beads/` path).
- **Comments go in the commit message, not the code** (`CONTRIBUTING_PR_GUIDELINES.md:33-35`).
- **ZFC**: no heuristics, keyword matching, scoring or thresholds in Go; the
  model decides (`CONTRIBUTING.md:128-130`). Metadata before schema; beads is
  not an orchestrator and not a storage engine (`engdocs/PROJECT_CHARTER.md:31-80`).
- **Storage boundary**: no `dolthub` imports outside `internal/storage/**`,
  no beads-side flocks/retry loops/engine introspection (`AGENTS.md:59-67`).
- **Migrations**: new files only, shipped ones frozen, no `UUID()/NOW()/RAND()`,
  ignored-plane twin for clone-local tables; migration/schema/sync paths need
  a filled template and a real human review (`PR_MAINTAINER_GUIDELINES.md:116`).
- **Lint is zero-new-issues, no tolerated baseline; moved code counts as new**
  (`engdocs/LINTING.md:16-50`). ICU never linked; `gms_pure_go` never removed
  (`engdocs/ICU-POLICY.md:5-9,147-151`).
- **Docs describe the pinned release** (`docs/cli-docs.pin`), so documenting an
  unreleased flag under `docs/` fails "Check doc flags freshness"; fork PRs
  must `git apply` the `cli-docs-freshness-patch` artifact themselves.
- **Visual**: never emoji icons; `○ ◐ ● ✓ ❄`; `P0`–`P4` colored, no glyph.
- **Never `bd edit`** (opens `$EDITOR`); `bd update <id> --body-file=-`.
- **Rebases are maintainer work** (`PR_MAINTAINER_GUIDELINES.md:138-144`):
  rebase on conflict, not on habit — every rebase risks turning a slice into a
  no-op against moved main.

## Quality gates — run them yourself; CI is the safety net

```bash
make build                                                  # never go build -o / go install (AGENT_INSTRUCTIONS.md:406-411)
./scripts/test.sh -run '^TestName$' ./internal/storage/...   # focused loop
./scripts/test.sh ./path/to/affected/...                    # affected packages
make test                                                   # ONE final baseline (engdocs/TESTING.md:18-28)
make fmt-check
BD_LINT_NEW_FROM_MERGE_BASE=origin/main make ci-pr-lint      # the exact PR lint lane
make ci-pr-policy                                           # policy wrapper incl. make api-check
make ci-pr-core                                             # -race -short, what "PR Core" runs
BEADS_TEST_ENV_RUN_DOLT=1 ./scripts/test.sh ./internal/storage/...   # Dolt-backed suites when storage changed
BEADS_BENCH_DOLT_PORT=<throwaway port> go test -tags gms_pure_go -bench=. ./internal/storage/dolt/...  # never the shared server
CGO_ENABLED=0 go vet -tags gms_pure_go ./<every package whose tests you touched>/   # the CI boundary job; a new file in a cgo-only test tier needs //go:build cgo (2026-09-01, #6130)
```

A red package at a pin-based carry head is classified against the LIVE build's
commit AND against stock `upstream/main` before the pick is blamed: at the
2026-09 pin, three `cmd/bd` backup tests are known-red (the viper-singleton
leak upstream fixed in #5194, merged after the pin) and the runner scrubs
neither `BEADS_BACKUP_ENABLED`/`BD_BACKUP_ENABLED` nor `BEADS_DIR` — run
gates with the town's ambient `BEADS_*`/`BD_*` variables unset, and document
the known-red on the handover rather than gating on it.

Gate every push on the suite's PARSED verdict (`tail -1` is `OK` / exit 0),
never on an eyeballed log; prove a new test both ways — red without the fix,
green with it. The same rule for every state claim: a "merged" report is read
from the PR object (`gh pr view --json state,mergeCommit`), never from the
merge command's exit text.

`make test-full-cgo` / `scripts/test-cgo.sh` are deprecated maintainer-only
ICU aliases — not a gate. Proportional validation: docs-only PRs run the docs
checks only. If a test fails on `upstream/main` too, it is not your PR: document
the diagnosis in the PR body. Never retry a deterministic failure into green;
retrigger a flake only with the flake issue cited.

## Commit and PR craft

**Commit subject** (measured: 87.5% of the last 200 upstream commits are
`type(scope): imperative`; squash-merge appends `(#NNNN)`; 39/200 carry a bead
id in the subject):

```
perf(ready): recurse the descendant walk off base tables (<your-rig>-<bead>)
```

The bead id goes **in parentheses at the end of the subject** — that is the
only place `bd orphans` parses (`cmd/bd/doctor/git.go:865-956`,
`AGENT_INSTRUCTIONS.md:94-99`); a body-only reference is invisible to it.
Body explains *why* (the diff shows what) and carries the measurement. Trailers:

```
Agent-Signature: claude-code-<model>-<reasoning> on behalf of <user>   (engdocs/AGENT_SIGNING.md:19-35)
Co-Authored-By: <the model that wrote it> <noreply@anthropic.com>
```

**PR body** — write to a file, `gh pr create --body-file`, run
`scripts/gh-body-lint <file>` first (rejects literal `\n` and `GH#123`; use
`#123`). Template is What / Why / Verification; reshape but keep the substance:

```markdown
## Summary
<symptom → root cause → fix, 1–3 bullets; name the riskiest thing the diff does>

## Why
<the principle / doc / prior PR this follows — cite paths; "Fixes #N">

## Verification
<commands + results; benchmark table for perf; which CI wrappers ran locally>

## Notes for reviewers
<what is deliberately not touched; follow-ups deferred; tradeoffs rejected and why>
```

Cite their own precedent (a prior PR that did the same thing). Tight,
technical, no salesmanship, no platitudes.

## Issue linking — best practice

1. **GitHub issue first.** There are no issue templates and labels are
   vestigial (`status/needs-repro|needs-info` have 0 uses; 81% of open issues
   older than a week have no reply; 52% of closures had zero comments), so
   the issue must carry its own evidence — the maintainers' literal ask, posted
   on ~60 issues in the 2026-06-16 repro sweep: *"the smallest command
   sequence or setup that triggers this, including the `bd version`, storage
   mode, OS, and the expected vs. actual output; a scratch-repo repro is
   ideal"*. What earned `has-repro` within days (#5427, #5712): repro on
   current `main` **with the commit hash**, root cause at `file:line` on the
   tip, measurements with ruled-out alternatives, a proposed fix shape. What
   gets pushed back: an old version ("Version first… main is ~1750 commits
   ahead", #5433), already-on-main (#4314, #4776), config-not-state (#4887),
   dependency direction (#723), scope (orchestration is Gas Town's, #850).
   Search duplicates before filing; reporters fix their own issue half the
   time and nothing discourages it. Features and design questions go as
   `RFC:` / `Proposal:` issues asking for direction.
2. **`Fixes #N` in the PR body's Why section** (the template's own suggestion);
   preflight warns when a closing reference is missing. `Refs #N` for related
   but not-closed issues. Never `GH#N` in anything posted to GitHub.
3. **Bead id at the end of the commit subject** `(<rig>-<id>)`; the bead carries
   the upstream issue and PR URLs in its notes, the merge SHA when it lands,
   and the drop-or-carry decision if a carry exists. Close the bead when the
   PR merges, not when it opens.
4. **Follow-up PRs name the parent in the title** — the measured house style is
   `... (#5820 follow-up)` — and link the parent in the body.
5. **Comments on upstream issues/PRs are signed** per `engdocs/AGENT_SIGNING.md:8-12`
   and say one thing: a finding, a measurement, a patch. No status chatter.

## The regime — how PRs actually get merged (measured 2026-09-01 over 200 merges, 151 open and 60 closed PRs; the fix-round observations below are dated 2026-09-03)

**Two pipelines.** Core PRs (julianknutsen, 106/200 merges) self-merge through
the `gc-review-queue` factory — not our pipeline. **External PRs are merged by
one operator, "bee" (`bee-ghosttrack`, also pushing as `steveyegge`; 39/40)**,
who posts an *executed* adversarial review ending in a verdict token —
`MERGE`, `MERGE on green CI`, `MERGE-AFTER-FIXES`, `REQUEST-CHANGES`. Reviews
reproduce the premise, re-run the body's verification commands and
mutation-test the tests (#5827, #5830, #5872, #5586). `maphew`, who merged our
#4117/#4119 in July, lost write access on Aug 7; standing approvals from him
no longer count.

**Cadence is batch sweeps.** 151/200 merges fell in 10 sessions; external
sweeps ran Aug 8/10/11, 19–20, 22, 27–28; 0/49 externals created Aug 24–31
had been touched by Sep 1 (#6031 is in that cohort). Once swept, externals
merge fast — **median 0.78 d, p75 3.4 d**, first response median 13 h. No
stale bot, no `status/*` label for externals, nothing to wait for or act on:
the job is to be mergeable when the sweep arrives and to answer the verdict
within hours. Expect a sweep, not a queue; never ping.

**What merged externals look like (n=40):** `fix` 26/40; **median 216 lines /
5 files, max 863**; 1–2 commits; 31/32 code PRs ship tests; 82% have
Summary/Why/Verification bodies; 65% disclose Claude co-authorship (never held
against them); 0 labels. Body hygiene does *not* discriminate — the open set
matches merged PRs on every template metric. What discriminates: **size**
(13/50 stale externals exceed 1,000 lines; 0/40 merged did), **missing a
sweep**, and **an unanswered verdict**.

**External PRs do not die by rejection** (closed-60: 0 maintainer rejections,
0 stale closes). They die by author withdrawal (10/14), duplication the
preflight would have caught (#5385, #5419, #5086), a fix already on `main`
(#5160), absorption into a maintainer PR with credit (#5124) — and mostly by
silent aging (50 open externals >30 d, 18 untouched for 30 d).

**What reviewers asked for before merging (14 fix rounds):**
1. **CHANGELOG `[Unreleased]` entry** for user-visible behavior — 8/14
   (#5549, #5586, #5799, #5802, #5813, #5822, #5830, #5859). Keep it additive;
   it is also the file that goes CONFLICTING.
2. **Missed sibling call sites / incomplete enumeration** — 8/14 (#5653
   `--watch` paths; #5826 "5 commands gated out of a ~120-site write
   surface"). Grep direct vs proxied-server twins, embedded vs server paths,
   every flag path, before pushing.
3. **Clear CONFLICTING by merging `main`** — 7/14; maintainers rescue
   CHANGELOG-only conflicts themselves, code conflicts wait on the author.
4. **Load-bearing tests** — no vacuous assertions (#5549, #5586, #5854).
5. Docs matching code (#5829), body matching diff (#5802), a triage of red CI
   in the thread (#5479, #5511), US spelling (`misspell`, #5854).
   Known-red signatures seen on our own PRs (2026-09-01): `Test (macos-latest)`
   is advisory and dies on its 10 m per-package timeout (`cmd/bd`,
   `embeddeddolt`, `uow`) with no failing assertion; `PR preflight process`
   jobs can come back `abandoned` (runner lost); `PR Core (wrapper timing)`
   (which IS in the gate) can hit Go's 10 m test timeout on a slow runner day
   after passing on the same code. None is the diff; say so in one signed
   comment with the run links and move on — no empty commits to retrigger, no
   pinging; maintainers rerun and fix-merge when the diff is clean.

**Fix rounds:** `CHANGES_REQUESTED` → author push median 7.1 h → merge 3.5 h
later; the one 221 h reply (#5859) turned a 3-day PR into 11 days.

**History worth keeping:** Alex town's #4117 and #4119 took 45 days each in
the pre-sweep regime and were praised for "the alternatives-considered
section, and the targeted tests"; a PR that touches `internal/storage/schema`
waits for a real human review after its sweep (schema PRs #5819/#5821 took
~3 days); an external PR without an issue reference misses a merged-external
norm. Measured on the five Alex-town PRs of 2026-09-03: one sweep read all five
in eleven minutes, every verdict reproduced the premise locally and asked for
one CHANGELOG sentence or one narrowed claim, and each fix push was approved
within the hour — the review is executed, not read, so the body's claims are
what get tested; answer the same day. An approved PR is not a merged one:
close the bead on the merge SHA, and follow-ups a reviewer lists under an
approval become their own `(#NNNN follow-up)` PRs, never pushes to the
approved branch.

Discord: `#beads` in Gas Town Hall (maintainers per the channel topic:
maphewyk, coffeegoddd, julianknutsen_gastown, steve). **No Discord pitching**
by default (operator ruling for gascity until after GC 1.5.0; applied here
unless told otherwise).

## After opening

- Verify the head repo: `gh pr view <N> --repo gastownhall/beads --json headRepositoryOwner --jq .headRepositoryOwner.login` → your personal fork's owner.
- Watch the whole matrix (incl. `PR Risk / CI Gate / Required`). **Triage your
  own red before the reviewer does**: one comment saying which lanes are yours
  (fix, push a follow-up commit — no force-push unless asked) and which are
  known-red (cite the flake issue). Fork PRs get the `docs-autofix` patch only
  as a comment recipe — apply it yourself.
- **Answer a verdict within hours**, every numbered ask in ONE push, replying
  in the reviewer's numbering with a per-ask status list; resolve threads
  after the fix lands. Take technical objections seriously, verify them, don't
  perform agreement, don't capitulate silently when you have evidence — default
  to deference.
- **Keep the branch mergeable by merging `upstream/main`, not rebasing**; if a
  force-push is unavoidable, say the tree is identical and show the
  `git range-diff`. Rebase on conflict, never on habit.
- Record the PR URL on the bead; hand the fleet's pin owner the cherry-pick SHA when the fix is
  needed before merge; close the bead on merge.

## Anti-patterns to refuse — the craft half

The cross-town anti-patterns (shared-store safety, the fleet build stamp,
deletion discipline, evidence discipline, the branch model) are in
`roles/bd-owner.md`. These are the beads-repository craft ones:

- "While I'm here" refactors, riders, cosmetic renames in a fix PR.
- `.beads/` data or regenerated artifacts in a diff; `GH#` in GitHub bodies.
- ICU flags, removing `gms_pure_go`, `go build -o bd`, `bd edit`, emoji icons.
- Skipping the preflight, `--no-verify`, retrying deterministic red into green.
