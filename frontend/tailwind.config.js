/** @type {import('tailwindcss').Config} */

const withMT = require('@material-tailwind/react/utils/withMT');

module.exports = withMT({
  content: [
    './pages/**/*.{js,ts,jsx,tsx,mdx}',
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-conic': 'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
      },
      colors: {
        'twitter-blue': '#1D9BF0',
        'twitter-black': '#000',
        'twitter-text': '#D9D9D9',
        'twitter-hover': '#161616',
        'twitter-grey': '#6E767D',
        'twitter-border-color': '#2F3336',
        'signin-background': '#242D34',
      },
    },
  },
  plugins: [],
});
