<script setup lang="ts">
import { BriefcaseIcon, LinkIcon, MapPinIcon } from '@heroicons/vue/20/solid';
import { useBusinessStore } from '@/features/business/application/business-store.js';
import { getRoleColor } from '@/features/utils/utils';
import { useConfigStore } from '@/stores/config';

const businessState = useBusinessStore();
const configStore = useConfigStore();
const themeColor = getRoleColor(configStore.state.role);

const state = await businessState.loadBusiness();
</script>

<template>
  <div class="-mt-8 mb-5 flex items-center justify-between">
    <div class="min-w-0 flex-1 pt-24 pb-10 text-center" :class="[themeColor]">
      <img class="mb-3 inline-block h-auto w-32 rounded-lg bg-gray-100 p-3" alt="" :src="state.business.logo" />
      <h2 class="text-2xl font-bold leading-7 text-gray-100 sm:truncate sm:text-3xl sm:tracking-tight">
        {{ state.business.socialDenomination }}
      </h2>
      <div class="grid justify-items-center">
        <div class="mt-0 flex space-x-6">
          <div class="mt-2 flex items-center text-sm text-gray-100">
            <BriefcaseIcon class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-100" aria-hidden="true" />
            <p class="font-semibold">CIF:</p>
            <p class="pl-1">{{ state.business.cif }}</p>
          </div>
          <div class="mt-2 flex items-center text-sm text-gray-100">
            <MapPinIcon class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-100" aria-hidden="true" />
            <p class="font-semibold">Sede:</p>
            <p class="pl-1">{{ state.business.registeredOffice }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
