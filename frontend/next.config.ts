import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  async redirects() {
    return [
      {
        source: '/',
        destination: '/login',
        permanent: false, // 一時的なリダイレクトの場合はfalse、恒久的なリダイレクトの場合はtrue
      },
    ];
  },
  /* config options here */
};

export default nextConfig;
