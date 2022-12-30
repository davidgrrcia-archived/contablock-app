<script setup lang="ts">
import { useEntriesStore } from '@/stores/entries';

const entriesStore = useEntriesStore();
await entriesStore.loadEntries();

const groupBy = <T, K extends keyof any>(arr: T[], key: (i: T) => K) =>
  arr.reduce((groups, item) => {
    (groups[key(item)] ||= []).push(item);
    return groups;
  }, {} as Record<K, T[]>);

const results = groupBy(entriesStore.state.storedEntries, i => i.multipleEntryId);
</script>

<template>
  <table class="styled-table">
    <thead>
      <tr>
        <th>Cuenta</th>
        <th>Divisa</th>
        <th>Fecha</th>
        <th>Concepto</th>
        <th>Debe</th>
        <th>Haber</th>
      </tr>
    </thead>
    <tbody>
      <template v-for="(item, index) in results" :key="index">
        <tr class="divider">
          <td colspan="6">Asiento {{ index }}</td>
        </tr>
        <template v-for="(entry, j) in results[index]" :key="j">
          <tr>
            <td>{{ entry.accountId }}</td>
            <td>{{ entry.currency }}</td>
            <td>{{ entry.date }}</td>
            <td>{{ entry.concept }}</td>
            <td>{{ entry.debitAmount }}</td>
            <td>{{ entry.creditAmount }}</td>
          </tr>
        </template>
      </template>
    </tbody>
  </table>
</template>

<style scoped>
.styled-table {
  border-collapse: collapse;
  margin: 25px 0;
  font-size: 0.9em;
  font-family: sans-serif;
  width: 100%;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
  text-align: left;
}

.styled-table thead tr {
  background-color: var(--blue);
  color: #ffffff;

}

.styled-table th,
.styled-table td {
  padding: 12px 15px;
}

.styled-table tbody tr {
  border-bottom: 1px solid #dddddd;
}

.styled-table tbody tr:nth-of-type(even) {
  background-color: var(--gray);
}

.styled-table tbody tr:last-of-type {
  border-bottom: 2px solid var(--blue);
}

.styled-table tbody tr.divider {
  font-weight: bold;
  color: var(--blue);
}
</style>
