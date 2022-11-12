import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import PageLayout from "../components/PageLayout";
import HeaderNav from "../components/HeaderNav";

const Home: NextPage = ({business, entries}) => {
    const navigation = [
        { name: 'Daily Ledger', href: '/', current: true },
        { name: 'General Ledger', href: '/general-ledger', current: false },
    ]

  return (
      <PageLayout title={'ContaBlock - Daily Ledger'} navigation={navigation} business={business}>
          <div className="flex flex-col">
              <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                  <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
                      <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
                          <table className="min-w-full divide-y divide-gray-200">
                              <thead className="bg-gray-50">
                              <tr>
                                  <th
                                      scope="col"
                                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                                  >
                                      Name
                                  </th>
                                  <th
                                      scope="col"
                                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                                  >
                                      Title
                                  </th>
                                  <th
                                      scope="col"
                                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                                  >
                                      Status
                                  </th>
                                  <th
                                      scope="col"
                                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                                  >
                                      Role
                                  </th>
                                  <th scope="col" className="relative px-6 py-3">
                                      <span className="sr-only">Edit</span>
                                  </th>
                              </tr>
                              </thead>
                              <tbody className="bg-white divide-y divide-gray-200">
                              {entries.items.map((entry) => (
                                  <tr key={entry.email}>
                                      <td className="px-6 py-4 whitespace-nowrap">
                                          <div className="flex items-center">
                                              <div className="ml-4">
                                                  <div className="text-sm font-medium text-gray-900">{entry.accountId}</div>
                                              </div>
                                              <div className="ml-4">
                                                  <div className="text-sm font-medium text-gray-900">{entry.currency}</div>
                                                  <div className="text-sm text-gray-500">{entry.concept}</div>
                                              </div>
                                          </div>
                                      </td>
                                      <td className="px-6 py-4 whitespace-nowrap">
                                          <div className="text-sm text-gray-900">{entry.debitBalance}</div>
                                          <div className="text-sm text-gray-500">{entry.CreditBalance}</div>
                                      </td>
                                      <td className="px-6 py-4 whitespace-nowrap">
                      <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                        Active
                      </span>
                                      </td>
                                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{entry.state}</td>
                                      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                                          <a href="#" className="text-indigo-600 hover:text-indigo-900">
                                              Edit
                                          </a>
                                      </td>
                                  </tr>
                              ))}
                              </tbody>
                          </table>
                      </div>
                  </div>
              </div>
          </div>
      </PageLayout>
  )
}

export default Home

// static HTML page rendered at build time (SSG), most optimal technique for generating performant SEO metrics and Caching CDN
// TODO refactor to use both SSG and SSR
// export async function getStaticProps(){
//     const res = await fetch('http://localhost:4000/api/business/zara')
//     const business = await res.json()
//     return {
//         props: {
//             business
//         }
//     }
// }

// dynamically rendered page rendered at run time (SSR), frequently changing dynamic data, or data that needs to be private
export async function getServerSideProps(){
    const res1 = await fetch('http://localhost:4000/api/business/zara')
    const business = await res1.json()
    const res2 = await fetch('http://localhost:4000/api/entries')
    const entries = await res2.json()
    return {
        props: {
            business,
            entries
        }
    }
}