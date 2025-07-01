import type { Metadata } from 'next';
import { Noto_Sans_JP as NotoSansJP } from 'next/font/google';
import './globals.css';

const notoSans = NotoSansJP({
  weight: ['400', '700'],
  subsets: ['latin', 'latin-ext'],
  display: 'swap',
});

export const metadata: Metadata = {
  title: 'Policy Evaluation System',
  description: '政策評価システムのホームページ',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ja" className={notoSans.className}>
      <body>{children}</body>
    </html>
  );
}