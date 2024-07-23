/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    async redirects() {
      return [
        {
          source: '/error',
          destination: '/error',
          permanent: true,
        },
      ];
    },
};

export default nextConfig;

  