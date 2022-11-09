import Header from "./Header";
import Image from "next/image";
import HeaderNav from "./HeaderNav";
import Head from "next/head";

export default function PageLayout({ children, title = 'ContaBlock', description = 'Accounting blockchain software', navigation }) {
  return (
    <>
      <Head>
        <title> {title} </title>
        <meta name="description" content={description} />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <HeaderNav navigation={navigation}/>
      <Header></Header>
      <main className="flex w-full flex-1 flex-col items-center justify-center px-20 text-center">
          <div className="flex min-h-screen flex-col items-center justify-center py-2">
              <div className="flex w-full flex-1 flex-col items-center justify-center px-20 text-center">
                  {children}
              </div>
          </div>
      </main>
      <footer className="flex h-24 w-full items-center justify-center border-t">
        <a
            className="flex items-center justify-center gap-2"
            href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
            target="_blank"
            rel="noopener noreferrer"
        >
            Powered by{' '}
            <Image src="/vercel.svg" alt="Vercel Logo" width={72} height={16} />
        </a>
      </footer>
    </>
  )
}