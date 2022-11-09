import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import PageLayout from "../components/PageLayout";
import HeaderNav from "../components/HeaderNav";

const Home: NextPage = () => {
    const navigation = [
        { name: 'Daily Ledger', href: '/', current: true },
        { name: 'General Ledger', href: '/general-ledger', current: false },
    ]

  return (
      <PageLayout title={'ContaBlock - Daily Ledger'} navigation={navigation}>
            <h1 className="text-6xl font-bold">
              Welcome to{' '}
              <a className="text-blue-600" href="https://nextjs.org">
                Next.js!
              </a>
            </h1>
      </PageLayout>
  )
}

export default Home
