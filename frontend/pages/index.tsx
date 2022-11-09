import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import PageLayout from "../components/PageLayout";

const Home: NextPage = () => {
  return (
      <PageLayout title={'ContaBlock - Home'}>
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
