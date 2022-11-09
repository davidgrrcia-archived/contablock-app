import { Disclosure } from '@headlessui/react'
import {useRouter} from "next/router";

const navigation = [
    { name: 'Daily Ledger', href: '/', current: false },
    { name: 'General Ledger', href: '/general-ledger', current: false },
]

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

export default function HeaderNav({navigation}) {
    return (
        <Disclosure as="nav" className="bg-gray-100">
            {({ open }) => (
                <>
                    <div className="mx-auto max-w-7xl px-2">
                        <div className="relative flex h-16 items-center justify-between">
                            <div className="flex flex-1 items-stretch justify-start">
                                <div className="flex flex-shrink-0 items-center">
                                    <img
                                        className="h-8 w-auto block"
                                        src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
                                        alt="Your Company"
                                    />
                                </div>
                                <div className="ml-6 block">
                                    <div className="flex space-x-4">
                                        {navigation.map((item) => (
                                            <a
                                                key={item.name}
                                                href={item.href}
                                                className={classNames(
                                                    item.current ? 'bg-gray-200 text-gray-700' : 'text-gray-500 hover:bg-gray-200 hover:text-black',
                                                    'px-3 py-2 rounded-md text-sm font-medium'
                                                )}
                                                aria-current={item.current ? 'page' : undefined}
                                            >
                                                {item.name}
                                            </a>
                                        ))}
                                    </div>
                                </div>
                            </div>

                        </div>
                    </div>
                </>
            )}
        </Disclosure>
    )
}
