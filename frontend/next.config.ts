import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  async rewrites() {
    return [
      {
        source: "/",
        destination: "/home",
      },
      {
        source: "/api/:path*",
        destination: "https://smash-voters.onrender.com/:path*",
      },
    ];
  },
};

export default nextConfig;
