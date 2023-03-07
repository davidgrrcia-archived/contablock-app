<script setup lang="ts">
import { useBalanceStore } from '@/features/balances/application/balances-store';
import { formatAmount } from '../../../utils/utils';

const balanceStore = useBalanceStore();
await balanceStore.loadBalance();
const balance = balanceStore.state.storedBalance;
</script>

<template>
  <div class="m-8 flex flex-col">
    <div class="-my-2 -mx-8">
      <div class="inline-block min-w-full py-2 px-8 align-middle">
        <div class="shadow ring-1 ring-black ring-opacity-5">
          <table class="min-w-full">
            <thead class="bg-white">
              <tr>
                <th scope="col" class="sticky top-14 z-10 rounded-tl-lg bg-gray-800 py-3.5 pr-3 pl-6 text-left text-sm font-semibold text-white">Naturaleza</th>
                <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">{{ balance.length != 0 ? balance[0].currentYear : '' }}</th>
                <th scope="col" class="sticky top-14 z-10 rounded-tr-lg bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">{{ balance.length != 0 ? balance[0].previousYear : '' }}</th>
              </tr>
            </thead>
            <tbody v-for="(record, i) in balance" :key="i" class="bg-white">
              <tr class="sticky top-24 z-10 bg-gray-700">
                <td class="whitespace-nowrap px-3 py-4 text-sm font-semibold text-white">
                  {{ record.name }}
                </td>
                <td class="whitespace-nowrap px-3 py-4 text-sm font-semibold text-white">
                  {{ formatAmount(record.currentAmount, 'EUR') }}
                </td>
                <td class="whitespace-nowrap px-3 py-4 text-sm font-semibold text-white">
                  {{ formatAmount(record.previousAmount, 'EUR') }}
                </td>
              </tr>
              <template v-for="(child, j) in record.child" :key="j">
                <tr :class="i % 2 === 0 ? undefined : 'bg-gray-50'" class="border-t border-gray-200">
                  <td class="whitespace-nowrap px-3 py-4 text-sm font-semibold text-gray-800">
                    {{ child.name }}
                  </td>
                  <td class="whitespace-nowrap px-3 py-4 text-sm font-semibold text-gray-800">
                    {{ formatAmount(child.currentAmount, 'EUR') }}
                  </td>
                  <td class="whitespace-nowrap px-3 py-4 text-sm font-semibold text-gray-800">
                    {{ formatAmount(child.previousAmount, 'EUR') }}
                  </td>
                </tr>
                <template v-for="(child2, k) in child.child" :key="k">
                  <tr :class="i % 2 === 0 ? undefined : 'bg-gray-50'" class="border-t border-gray-200">
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ child2.name }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ formatAmount(child2.currentAmount, 'EUR') }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ formatAmount(child2.previousAmount, 'EUR') }}
                    </td>
                  </tr>
                  <template v-for="(child3, l) in child2.child" :key="l">
                    <tr :class="i % 2 === 0 ? undefined : 'bg-gray-50'" class="border-t border-gray-200">
                      <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                        {{ child3.name }}
                      </td>
                      <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                        {{ formatAmount(child3.currentAmount, 'EUR') }}
                      </td>
                      <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                        {{ formatAmount(child3.previousAmount, 'EUR') }}
                      </td>
                    </tr>
                  </template>
                </template>
              </template>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
