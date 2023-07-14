/** @type {import('next').NextConfig} */
const nextConfig = {
  compiler: {
    relay: {
      src: './',
      artifactDirectory: './app/relay/__generated__',
      language: 'typescript',
    },
  },
};

module.exports = nextConfig;
