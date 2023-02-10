<script setup lang="ts">
import { useDailyLedgerStore } from '@/stores/dailyLedger';

const dailyLedgerStore = useDailyLedgerStore();
await dailyLedgerStore.loadDailyLedger();

const dailyLedger = dailyLedgerStore.state.storedDailyLedger;
</script>

<template>
  <div class="px-8">
    <div class="flex items-center">
      <div class="flex-auto">
        <h1 class="text-xl font-semibold text-gray-900">Libro Diario</h1>
        <p class="mt-2 text-sm text-gray-700">A list of all the users in your account including their name, title, email and role.</p>
      </div>
    </div>
    <div class="mt-8 flex flex-col">
      <div class="-my-2 -mx-8">
        <div class="inline-block min-w-full py-2 px-8 align-middle">
          <div class="shadow ring-1 ring-black ring-opacity-5">
            <table class="min-w-full">
              <thead class="bg-white">
                <tr>
                  <th scope="col" class="sticky top-14 z-10 rounded-tl-lg bg-gray-800 py-3.5 pl-4 pr-3 pl-6 text-left text-sm font-semibold text-white">Cuenta</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Tipo</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Divisa</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Fecha</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Trimestre</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Concepto</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Debe</th>
                  <th scope="col" class="sticky top-14 z-10 bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Haber</th>
                  <th scope="col" class="sticky top-14 z-10 rounded-tr-lg bg-gray-800 px-3 py-3.5 text-left text-sm font-semibold text-white">Estado</th>
                </tr>
              </thead>
              <tbody class="bg-white">
                <template v-for="(record, index) in dailyLedger" :key="index">
                  <tr class="sticky top-24 z-10 bg-gray-600">
                    <th colspan="8" scope="colgroup" class="px-4 px-6 text-left text-sm font-semibold text-white">Asiento {{ index }}</th>
                    <th>
                      <div class="flex space-x-4 py-2 px-2">
                        <div class="w-1/2">
                          <button type="button" class="btn bg-[var(--blue)]">Publicar</button>
                        </div>
                        <div class="w-1/2">
                          <button type="button" class="btn bg-[var(--red)]">Eliminar</button>
                        </div>
                      </div>
                    </th>
                  </tr>
                  <tr v-for="(entryExpanded, j) in dailyLedger[index].entriesExpanded" :key="j" :class="j % 2 === 0 ? undefined : 'bg-gray-50'" class="border-t border-gray-200">
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.account.displayName }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.account.type }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.entry.currency }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.entry.displayDate }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.entry.displayDate }}
                    </td>
                    <td class="px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.entry.concept }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.entry.debitAmount }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.entry.creditAmount }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-800">
                      {{ entryExpanded.account.type }}
                    </td>
                    <td></td>
                  </tr>
                </template>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
