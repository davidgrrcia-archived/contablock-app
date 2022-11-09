import Image from "next/image";

export default function Header({ business }){
    return (
        <header className="bg-white shadow">
            <div className="flex items-center mx-auto max-w-7xl py-6 px-8">
                <div className={"w-1/3 flex items-center"}>
                    <Image
                       src={business.logo}
                       alt=""
                       width={128}
                       height={64}>
                    </Image>
                    <div className={"px-5"}>
                        <h1 className="text-3xl font-bold tracking-tight text-gray-900 ">{business.socialDenomination}</h1>
                        <p>{business.registeredOffice}</p>
                    </div>
                </div>
                <div className={"w-1/3"}>
                    <p>CIF: {business.cif}</p>
                </div>
                <div className={"w-1/3"}>
                    <p>Ethereum Address: {business.ethAddress}</p>
                </div>
            </div>
        </header>
    )
}

