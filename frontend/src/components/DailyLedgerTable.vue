<script setup lang="ts">
import { useDailyLedgerStore } from '@/stores/dailyLedger';

const dailyLedgerStore = useDailyLedgerStore();
await dailyLedgerStore.loadDailyLedger();

const dailyLedger = dailyLedgerStore.state.storedDailyLedger;
</script>

<template>
  <div class="px-4 sm:px-6 lg:px-8">
    <div class="sm:flex sm:items-center">
      <div class="sm:flex-auto">
        <h1 class="text-xl font-semibold text-gray-900">Libro Diario</h1>
        <p class="mt-2 text-sm text-gray-700">A list of all the users in your account including their name, title, email and role.</p>
      </div>
    </div>
    <div class="mt-8 flex flex-col">
      <div class="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
          <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
            <table class="min-w-full">
              <thead class="bg-white">
                <tr>
                  <th scope="col" class="bg-[var(--blue)] py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-white sm:pl-6">Cuenta</th>
                  <th scope="col" class="bg-[var(--blue)] px-3 py-3.5 text-left text-sm font-semibold text-white">Tipo</th>
                  <th scope="col" class="bg-[var(--blue)] px-3 py-3.5 text-left text-sm font-semibold text-white">Divisa</th>
                  <th scope="col" class="bg-[var(--blue)] px-3 py-3.5 text-left text-sm font-semibold text-white">Fecha</th>
                  <th scope="col" class="bg-[var(--blue)] px-3 py-3.5 text-left text-sm font-semibold text-white">Trimestre</th>
                  <th scope="col" class="bg-[var(--blue)] px-3 py-3.5 text-left text-sm font-semibold text-white">Concepto</th>
                  <th scope="col" class="bg-[var(--blue)] px-3 py-3.5 text-left text-sm font-semibold text-white">Debe</th>
                  <th scope="col" class="bg-[var(--blue)] px-3 py-3.5 text-left text-sm font-semibold text-white">Haber</th>
                  <th scope="col" class="bg-[var(--blue)] px-3 py-3.5 text-left text-sm font-semibold text-white">Estado</th>
                  <th scope="col" class="relative bg-[var(--blue)] py-3.5 pl-3 pr-4 sm:pr-6">
                    <span class="sr-only">Acciones</span>
                  </th>
                </tr>
              </thead>
              <tbody class="bg-white">
                <template v-for="(record, index) in dailyLedger" :key="index">
                  <tr class="border-t border-gray-200">
                    <th colspan="9" scope="colgroup" class="bg-[var(--white)] px-4 py-2 text-left text-sm font-semibold text-[var(--blue)] sm:px-6">Asiento {{ index }}</th>
                    <th>
                      <button type="button" class="btn bg-[var(--green)] hover:bg-indigo-700 focus:ring-indigo-500">Publicar</button>
                      <button type="button" class="btn bg-[var(--red)] hover:bg-indigo-700 focus:ring-indigo-500">Eliminar</button>
                    </th>
                  </tr>
                  <tr v-for="(entryExpanded, j) in dailyLedger[index].entriesExpanded" :key="j" :class="[j === 0 ? 'border-gray-300' : 'border-gray-200', 'border-t']">
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                      {{ entryExpanded.account.displayName }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                      {{ entryExpanded.account.type }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                      {{ entryExpanded.entry.currency }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                      {{ entryExpanded.entry.displayDate }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                      {{ entryExpanded.entry.displayDate }}
                    </td>
                    <td class="px-3 py-4 text-sm text-gray-500">
                      {{ entryExpanded.entry.concept }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                      {{ entryExpanded.entry.debitAmount }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                      {{ entryExpanded.entry.creditAmount }}
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
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
