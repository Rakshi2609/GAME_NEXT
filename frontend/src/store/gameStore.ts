import { create } from 'zustand'

interface GameState {
  score: number
  timeLeft: number
  running: boolean
  start: () => void
  stop: () => void
  addScore: (n?: number) => void
  tick: () => void
  reset: () => void
}

export const useGameStore = create<GameState>((set) => ({
  score: 0,
  timeLeft: 30,
  running: false,
  start: () => set({ running: true, score: 0, timeLeft: 30 }),
  stop: () => set({ running: false }),
  addScore: (n = 1) => set((s) => ({ score: s.score + n })),
  tick: () => set((s) => ({ timeLeft: Math.max(0, s.timeLeft - 1) })),
  reset: () => set({ score: 0, timeLeft: 30, running: false }),
}))
