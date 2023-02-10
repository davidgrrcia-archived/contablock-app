<script setup lang="ts">
import {
  BriefcaseIcon,
  LinkIcon,
  MapPinIcon,
} from '@heroicons/vue/20/solid'
import {useBusinessStore} from "@/stores/business";

const businessState = useBusinessStore();
const state = await businessState.loadBusiness("zara")
console.log(state.business.logo)
</script>

<template>
  <div class="mt-18 mb-5 flex items-center justify-between">
    <div class="min-w-0 flex-1 bg-gray-100 pt-24 pb-10 text-center">
      <img class="inline-block h-20 w-40 rounded-lg mb-3" alt="" :src="state.business.logo"/>
      <h2 class="text-2xl font-bold leading-7 text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">
        {{ state.business.socialDenomination }}
      </h2>
      <div class="grid justify-items-center ">
        <div class="flex mt-0  space-x-6">
          <div class="mt-2 flex items-center text-sm text-gray-500">
            <BriefcaseIcon class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400" aria-hidden="true" />
            CIF: {{ state.business.cif }}
          </div>
          <div class="mt-2 flex items-center text-sm text-gray-500">
            <MapPinIcon class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400" aria-hidden="true" />
            Sede: {{ state.business.registeredOffice }}
          </div>
          <div class="mt-2 flex items-center text-sm text-gray-500">
            <LinkIcon class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400" aria-hidden="true" />
            Cuenta Ethereum: {{ state.business.ethAddress }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
