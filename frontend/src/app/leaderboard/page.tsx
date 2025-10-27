"use client";
import { useEffect, useState } from 'react';

interface Entry { name: string; score: number }

export default function LeaderboardPage() {
  const [entries, setEntries] = useState<Entry[]>([]);

  useEffect(() => {
    const raw = localStorage.getItem('orbchase_leaderboard');
    if (raw) setEntries(JSON.parse(raw));
    else setEntries([
      { name: 'Aditi', score: 21 },
      { name: 'Raghav', score: 18 },
      { name: 'Sam', score: 15 },
    ]);
  }, []);

  return (
    <main className="container py-10">
      <h1 className="text-4xl font-extrabold mb-6">Leaderboard</h1>
      <div className="card overflow-x-auto">
        <table className="w-full text-left">
          <thead className="text-slate">
            <tr>
              <th className="py-2 pr-4">#</th>
              <th className="py-2 pr-4">Name</th>
              <th className="py-2">Score</th>
            </tr>
          </thead>
          <tbody>
            {entries.map((e, i) => (
              <tr key={i} className="border-t border-white/10">
                <td className="py-2 pr-4">{i + 1}</td>
                <td className="py-2 pr-4">{e.name}</td>
                <td className="py-2">{e.score}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </main>
  );
}
