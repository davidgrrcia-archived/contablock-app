export default function Header(){
    return (
        <header className="bg-white shadow">
            <div className="flex items-center mx-auto max-w-7xl py-6 px-8">
                <div className={"w-1/3 flex items-center"}>
                    <img className="h-16 w-32"
                         src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/fd/Zara_Logo.svg/1024px-Zara_Logo.svg.png"
                         alt="">
                    </img>
                    <div className={"px-5"}>
                        <h1 className="text-3xl font-bold tracking-tight text-gray-900 ">Zara</h1>
                        <p>Technical advisor</p>
                    </div>
                </div>
                <div className={"w-1/3"}>
                    <p>CIF: A209304302</p>
                </div>
                <div className={"w-1/3"}>
                    <p>Ethereum Address: 0x0000000000000000000000000000000000000000</p>
                </div>
            </div>
        </header>
    )
}