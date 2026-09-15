# Early Gate Pipeline

> Test document. This PR exists to exercise the early gate pipeline.

## What the early gate is

The *early gate* is the set of checks that run at **pull request time**, before a change is
allowed to merge. It is the shift-left counterpart to the
[post-codefreeze gatekeeper](.github/workflows/post-codefreeze-gatekeeper.yaml), which only
guards the release branches (`rhoai-3.5`, `rhoai-3.6-ea.1`) after code freeze.

The goal is the same in both cases — keep broken changes out of the release branches — but the
early gate catches them on the contributor's PR instead of at promotion time, when a fix is
much cheaper.

## What runs

| Check | Workflow | What it protects |
|-------|----------|------------------|
| Controller tests | `.github/workflows/controller-tests.yaml` | Reconciler behaviour (Ginkgo + envtest) |
| Module controller tests | `.github/workflows/trustyai-module-controller-tests.yaml` | Operator module reconcilers |
| Manifest policy | `.github/workflows/conftest.yaml` | OPA/Rego policies in `policy/` against every kustomize overlay |
| YAML lint | `.github/workflows/lint-yaml.yaml` | `config/**/*.yaml` formatting |
| Security scan | `.github/workflows/gosec.yaml`, `.github/workflows/security-scan.yaml` | Known insecure patterns, SARIF upload |
| Smoke | `.github/workflows/smoke.yaml` | Basic build/deploy sanity |
| Upgrade validation | `.github/workflows/operator-chaos.yml` | Breaking CRD schema and knowledge-model changes (`chaos/knowledge/trustyai.yaml`) |
| Disconnected readiness | `.github/workflows/disconnected-readiness.yaml` | Image references resolvable in disconnected installs |
| Konflux PR builds | `.tekton/*-pull-request.yaml` | Hermetic container builds for the operator and the LMES driver |

## Ordering

1. **Fast feedback** — lint, policy, and security scans finish first and fail loudly on
   formatting or RBAC violations.
2. **Correctness** — controller tests run against envtest with the CRDs from
   `config/crd/bases/`.
3. **Packaging** — Konflux pull-request pipelines build the images the same way the release
   pipeline will, so build breakage surfaces before merge.
4. **Upgrade safety** — `operator-chaos` diffs CRDs and the knowledge model against the base
   branch; a breaking schema change fails the gate without needing a cluster.

## When the gate is bypassed

- `.tekton/` files are synced from
  [`konflux-central`](https://github.com/red-hat-data-services/konflux-central) and must be
  changed there, not here — see [.tekton/README.md](.tekton/README.md).
- `.github/workflows/instant-merge.yaml` fast-paths automated sync PRs.
- The post-codefreeze gatekeeper can be re-triggered on demand with the `run-gatekeeper` label.
