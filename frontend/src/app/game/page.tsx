"use client";
import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import Target from '@/components/Target';
import { useGameStore } from '@/store/gameStore';

export default function GamePage() {
  const router = useRouter();
  const { start, stop, tick, timeLeft, score, running } = useGameStore();

  useEffect(() => {
    start();
    return () => stop();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  useEffect(() => {
    if (!running) return;
    const id = setInterval(() => tick(), 1000);
    return () => clearInterval(id);
  }, [running, tick]);

  useEffect(() => {
    if (timeLeft === 0) {
      stop();
      router.push(`/game-over?score=${score}`);
    }
  }, [timeLeft, stop, router, score]);

  return (
    <main className="container py-10">
      <div className="flex items-center justify-between mb-4">
        <div className="card flex-1 mr-2">
          <div className="text-slate">Time Left</div>
          <div className="text-2xl font-extrabold">{timeLeft}s</div>
        </div>
        <div className="card flex-1 ml-2">
          <div className="text-slate">Score</div>
          <div className="text-2xl font-extrabold">{score}</div>
        </div>
      </div>
      <Target />
    </main>
  );
}
