import type { Config } from 'tailwindcss'

export default {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      colors: {
        midnight: '#0A192F',
        electric: '#8B5CF6',
        sky: '#38BDF8',
        soft: '#F8FAFC',
        slate: '#94A3B8',
        pink: '#EC4899',
        lime: '#84CC16',
      },
      boxShadow: {
        soft: '0 10px 30px -10px rgba(0,0,0,0.3)',
      },
      backgroundImage: {
        'gradient-accent': 'linear-gradient(90deg, #8B5CF6, #38BDF8)',
      },
    },
  },
  plugins: [],
} satisfies Config
