import Head from "next/head";
import PageLayout from "../../components/PageLayout";

export default function GeneralLedger() {
    return (
        <PageLayout title={'ContaBlock - General Ledger'}>
            <h1 className="text-6xl font-bold">
                General Ledger
            </h1>
            <p className="mt-3 text-2xl">
                This is the general ledger page.
            </p>
        </PageLayout>

    )
}