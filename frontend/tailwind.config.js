/** @type {import('tailwindcss').Config} */

const withMT = require('@material-tailwind/react/utils/withMT');
const colors = require('tailwindcss/colors');

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
        ...colors,
        'twitter-blue': '#1D9BF0',
        'twitter-black': '#000',
        'twitter-text': '#D9D9D9',
        'twitter-hover': '#161616',
        'twitter-grey': '#6E767D',
        'twitter-border-color': '#2F3336',
        'signin-background': '#242D34',
        'dark-hover': '#1B1B1B',
      },
    },
    screens: {
      // lg:w-10 means apply w-10 if screen is lg or smaller

      // desktop
      xl: { max: '1279.98px' },
      // => @media (max-width: 1279.98px) { ... }

      // laptop
      lg: { max: '1023.98px' },
      // => @media (max-width: 1023.98px) { ... }

      // tablet
      md: { max: '767.98px' },
      // => @media (max-width: 767.98px) { ... }

      // mobile
      sm: { max: '639.98px' },
      // => @media (max-width: 639.98px) { ... }
    },
  },
  plugins: [],
});
