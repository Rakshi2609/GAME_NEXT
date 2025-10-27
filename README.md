# CodeChef VIT — Web Development Department Onboarding Project

This repository contains a complete, production-ready scaffold for the three onboarding tracks:

- UI/UX: Static prototype of a Hackathon Event page following the shared Design System
- Frontend: Next.js + Tailwind CSS interactive game (Landing, Game, Game Over, Leaderboard)
- Backend: Go (Gin) + MongoDB API with JWT authentication, roles (Event Creators, Attendees), and event/user models

The goal is to give you a clean, modern, and consistent foundation with an Apple-like aesthetic and a minimal developer experience.

## Repository Structure

```
/ (root)
├─ uiux/                # Static, responsive UI prototype (HTML/CSS)
├─ frontend/            # Next.js + Tailwind app (TypeScript)
└─ backend/             # Go Gin REST API (JWT, MongoDB)
```

## Quick Start

- UI/UX: Open `uiux/index.html` in your browser.
- Frontend: Requires Node.js 18+ (Node 22 supported)
  - In `frontend/`: `npm install` then `npm run dev` (dev server at http://localhost:3000)
- Backend: Requires Go 1.21+ and a MongoDB instance
  - In `backend/`: Copy `.env.example` to `.env`, set `MONGO_URI` and `JWT_SECRET`, then `go run .`

If Go or MongoDB aren’t installed locally, you can still review the backend code and API structure.

## Design System (Shared)

- Colors: Midnight #0A192F, Electric #8B5CF6, Sky #38BDF8, Soft #F8FAFC, Slate #94A3B8, Pink #EC4899, Lime #84CC16
- Typography: Inter/SF system stacks; headings bold, generous spacing, subtle gradients

## Author

- Rakshith Ganjimut · Register Number: 24BRS1301

---

This scaffold is intentionally minimal, clean, and cohesive. Extend features as needed and keep code modular and documented.


# GAME_NEXT
