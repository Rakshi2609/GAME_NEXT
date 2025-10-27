"use client";
import Link from 'next/link';
import { motion } from 'framer-motion';

export default function HomePage() {
  return (
    <main>
      <section className="relative overflow-hidden">
        <div className="absolute inset-0 pointer-events-none">
          <div className="h-[50vh] bg-gradient-to-r from-electric/20 to-sky/20 blur-3xl" />
        </div>
        <div className="container pt-28 pb-16 text-center">
          <motion.h1
            initial={{ opacity: 0, y: 12 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6 }}
            className="text-5xl/tight md:text-6xl/tight font-extrabold tracking-tight"
          >
            Orb Chase
          </motion.h1>
          <p className="mt-4 text-slate max-w-2xl mx-auto">
            A minimalist reflex game. Click the moving orb to score before the timer runs out.
          </p>
          <div className="mt-8 flex items-center justify-center gap-4">
            <Link href="/game" className="btn">Play Now</Link>
            <Link href="/leaderboard" className="btn bg-white/10 text-soft !bg-none border border-white/10 hover:bg-white/15">
              Leaderboard
            </Link>
          </div>
        </div>
      </section>
      <section className="container grid md:grid-cols-3 gap-4 pb-20">
        <div className="card"><h3 className="font-semibold">Responsive</h3><p>Looks great on all screens.</p></div>
        <div className="card"><h3 className="font-semibold">Smooth</h3><p>Subtle motion, zero clutter.</p></div>
        <div className="card"><h3 className="font-semibold">Clean Code</h3><p>Modular, typed, and readable.</p></div>
      </section>
    </main>
  );
}
