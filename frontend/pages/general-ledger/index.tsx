import Head from "next/head";
import PageLayout from "../../components/PageLayout";

export default function GeneralLedger({business}) {
    const navigation = [
        { name: 'Daily Ledger', href: '/', current: false },
        { name: 'General Ledger', href: '/general-ledger', current: true },
    ]

    return (
        <PageLayout title={'ContaBlock - General Ledger'} navigation={navigation} business={business}>
            <h1 className="text-6xl font-bold">
                General Ledger
            </h1>
            <p className="mt-3 text-2xl">
                This is the general ledger page.
            </p>
        </PageLayout>

    )
}

export async function getStaticProps(){
    const res = await fetch('http://localhost:4000/api/business/zara')
    const business = await res.json()
    return {
        props: {
            business
        }
    }
}