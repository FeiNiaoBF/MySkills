---
name: teach-go
description: >
  Teach Go (standard library) through the translate-webui project. Use when the user
  wants to learn or write Go, be taught Go, or work on translate-webui code —
  e.g. "教 Go", "教我写 Go", "write the Go code", "translate-webui", "Go 教学".
  Covers net/http, JSON, HTTP client, go:embed, error handling — one concept per
  lesson, minimal-first, each lesson ends with a working artifact.
---

# Teach Go — translate-webui

A lesson plan for teaching Go standard-library web development, using `translate-webui` as the vehicle project.

## Context (read first)

- Project: `~/Code/translate-webui` on pop (`ssh pop`), i.e. `/data/workspace/code/translate-webui`
- Read before teaching: `AGENTS.md` (hard constraints), `DESIGN.md` (architecture/API/prompt spec), `README.md` (run/deploy)
- Go 1.26.5 via mise: `mise exec go@1.26.5 -- go <cmd>`
- Inference backend already deployed: Ollama at `127.0.0.1:11434`, model `translategemma:4b`
- Service port: 8000. Historical files (`app.py`, `templates/`, `.venv/`) — do not touch, do not delete without consent.

## Hard constraints (from AGENTS.md)

- **Go standard library only** (`net/http`, `encoding/json`, `html/template`, `embed`). No third-party libs or frameworks.
- Frontend: plain HTML/CSS/JS, no CDN, no build step.
- Never modify Pop system config (Ollama systemd service, `OLLAMA_MODELS`, `/data` layout).
- Model dir `/data/tools/ollama/models` is read-only.

## Teaching protocol (inherits global AGENTS.md + APPEND_SYSTEM.md)

- Respond in clear English by default; use brief Chinese scaffold only when language blocks understanding.
- **从简到繁**: smallest working slice first; add a concept only when the learner meets its motivating pain. Never front-load theory.
- **One concept per lesson**; each lesson ends with a working artifact + exactly one exercise.
- **ZPD**: pick the next step the learner can complete with support but not yet alone. Fade support after repeated success.
- **Two dimensions** (programming difficulty / English difficulty) are never raised in the same step.
- **Retrieval practice**: ask the learner to recall/attempt before revealing the solution.
- **BLUF** output; code blocks always carry language tags; keep tables narrow (≤3 columns).

## Lesson plan (9 lessons, ~30–45 min each)

### L1 — Hello HTTP server
Concept: `net/http` — `http.HandleFunc`, `http.ListenAndServe`. Route `/` returns text.
Files: minimal `main.go`. Verify: `go run .` + `curl localhost:8000`.
Exercise: add `/healthz` returning 200.

### L2 — Method routing & handlers
Concept: Go 1.22 `ServeMux` patterns (`"GET /api/translate"`), handler signature, `ResponseWriter`/`Request`.
Verify: wrong method → 405.

### L3 — JSON in/out
Concept: `encoding/json` — decode request body, encode response; structs with json tags.
Exercise: echo endpoint returning `{text}`.

### L4 — Calling Ollama (the core)
Concept: `http.Client`; POST `/api/generate`; JSON payload; read/parse response; timeout ~300s; `context`.
Exercise: translate one hardcoded string via Ollama.

### L5 — Prompt construction & language map
Concept: `translate.go` — language map (`auto/zh/en/ja/ko/de/fr/es/ru`), prompt template (see DESIGN.md §4), `temperature 0.2`.
Verify with the 4 benchmark sentences: greeting / colloquial / zh→en / academic.

### L6 — go:embed static frontend
Concept: `embed.FS`; serve `web/index.html` at `/`; frontend fetches same-origin `POST /api/translate`.
Exercise: wire the existing UI design to the API.

### L7 — Error handling & status codes
Concept: 400 (empty input / invalid lang), 502 (Ollama down), 504 (timeout); consistent `{"error": ...}`.
Verify: all three branches via curl.

### L8 — Config via env vars
Concept: `PORT`, `OLLAMA_URL` from `os.Getenv` with sane defaults; keep `main.go` clean.

### L9 — Build & deploy
Concept: `go build -o translate-webui .`; swap the Flask systemd unit to the Go binary (unit already drafted in README); restart flow.
Verify: rerun the 4 benchmark sentences — results comparable to the Flask baseline.

## Verification checklist (every lesson)

- [ ] `go vet ./... && go build ./...` — zero errors
- [ ] lesson's curl smoke test passes
- [ ] one new concept explained + one exercise completed by the learner
- [ ] `DESIGN.md` / `README.md` updated if behavior changed

## Anti-patterns

- No front-loaded theory, no framework detours, no speculative features.
- Do not reveal full solutions before the learner attempts.
- Do not edit system files or the model dir.
- Do not delete historical Flask files without explicit consent.
