# FastEnv
Zero-dependency Go CLI that eliminates Python venv creation latency using content-addressed caching and fast lockfile hashing.

## Current Status (v0.1)

- [x] **Content-Addressed Hashing:** Fast hashing of `requirements.txt` manifests via internal `hasher`.
- [x] **Centralized Cache Store:** Isolated cache management in `~/.fastenv/store/<hash>`.
- [x] **Lightweight Linker:** OS-native symlinking/junctions from project `.venv` directly to the store.
- [x] **Fallback Builder:** Automated environment creation on cache misses.

---

## v0.5 — State Synchronization & Scoping

The goal of v0.5 is to make `fastenv` context-aware, robust against unintended environment overwrites, and interactive.

### Core Objectives

#### 1. Interactive TUI Integration (`huh?`)
- [ ] **Cache Selection:** Interactive prompt to view and select existing cached builds on cache misses or collisions.
- [ ] **Confirmation Flow:** Prompt users before executing destructive operations (e.g., store pruning or link overwrites).
- [ ] **Status Dashboard:** Clean terminal UI overview showing the active `.venv` status, target manifest hash, and store location.

#### 2. Environment Reconciliation (`sync`)
- [ ] **Store Query Interface:** Implement `store.ReadStore()` to map existing cache hashes without tight coupling.
- [ ] **Automated Repair:** Detect broken or out-of-sync `.venv` pointers and instantly rewire them to matching cached builds.
- [ ] **Cache Pruning:** Add `fastenv sync --prune` to purge unlinked or stale environments.

#### 3. Project Scoping & Context (`context`)
- [ ] **Metadata Marker (`.fastenv`):** Store project-level scope rules, active manifest paths, and expected hash targets.
- [ ] **Directory Boundary Guard:** Validate working directory boundaries before executing link or build operations.
- [ ] **Isolated Context Hashing:** Combine project path metadata with manifest content hashes to preserve environment boundaries.

---

## Architecture Overview (v0.5)

