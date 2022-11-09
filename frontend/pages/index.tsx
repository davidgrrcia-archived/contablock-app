import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import PageLayout from "../components/PageLayout";
import HeaderNav from "../components/HeaderNav";

const Home: NextPage = ({business}) => {
    const navigation = [
        { name: 'Daily Ledger', href: '/', current: true },
        { name: 'General Ledger', href: '/general-ledger', current: false },
    ]

  return (
      <PageLayout title={'ContaBlock - Daily Ledger'} navigation={navigation} business={business}>
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

// static HTML page rendered at build time (SSG), most optimal technique for generating performant SEO metrics and Caching CDN
export async function getStaticProps(){
    const res = await fetch('http://localhost:4000/api/business/zara')
    const business = await res.json()
    return {
        props: {
            business
        }
    }
}

// dynamically rendered page rendered at run time (SSR), frequently changing dynamic data, or data that needs to be private
// export async function getServerSideProps(){
//     const res = await fetch('http://localhost:4000/api/business/zara')
//     const business = await res.json()
//     return {
//         props: {
//             business
//         }
//     }
// }