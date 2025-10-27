"use client";
import { useEffect, useMemo, useRef, useState } from 'react';
import { useGameStore } from '../store/gameStore';

export default function Target() {
  const addScore = useGameStore((s) => s.addScore);
  const running = useGameStore((s) => s.running);
  const [pos, setPos] = useState({ x: 50, y: 50 });
  const containerRef = useRef<HTMLDivElement>(null);

  const bounds = useMemo(() => ({ w: 800, h: 420 }), []);

  useEffect(() => {
    if (!running) return;
    const id = setInterval(() => {
      setPos({
        x: 10 + Math.random() * 80,
        y: 10 + Math.random() * 70,
      });
    }, 900);
    return () => clearInterval(id);
  }, [running]);

  return (
    <div ref={containerRef} className="relative w-full h-[420px] bg-white/5 border border-white/10 rounded-2xl overflow-hidden">
      <button
        className="absolute -translate-x-1/2 -translate-y-1/2 w-12 h-12 rounded-full bg-gradient-to-r from-electric to-sky shadow-soft hover:scale-105 active:scale-95 transition"
        style={{ left: `${pos.x}%`, top: `${pos.y}%` }}
        onClick={() => addScore(1)}
        aria-label="target"
      />
    </div>
  );
}
