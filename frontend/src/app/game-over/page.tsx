"use client";
import Link from 'next/link';

export default function GameOver({ searchParams }: { searchParams: { score?: string } }) {
  const score = Number(searchParams?.score ?? 0);
  return (
    <main className="container py-20 text-center">
      <h1 className="text-5xl font-extrabold">Game Over</h1>
      <p className="mt-3 text-slate">Your score</p>
      <div className="mt-2 text-6xl font-extrabold bg-gradient-accent bg-clip-text text-transparent">
        {score}
      </div>
      <div className="mt-8 flex items-center justify-center gap-4">
        <Link href="/game" className="btn">Play Again</Link>
        <Link href="/" className="btn bg-white/10 text-soft !bg-none border border-white/10">Home</Link>
      </div>
    </main>
  );
}
