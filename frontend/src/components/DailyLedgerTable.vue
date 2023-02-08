<script setup lang="ts">
import { useDailyLedgerStore } from '@/stores/dailyLedger';

const dailyLedgerStore = useDailyLedgerStore();
await dailyLedgerStore.loadDailyLedger();

const dailyLedger = dailyLedgerStore.state.storedDailyLedger;
</script>

<template>
  <table class="styled-table">
    <thead>
      <tr>
        <th>Cuenta</th>
        <th>Divisa</th>
        <th>Fecha</th>
        <th>Trimestre</th>
        <th>Concepto</th>
        <th>Debe</th>
        <th>Haber</th>
        <th>Estado</th>
        <th>Acciones</th>
      </tr>
    </thead>
    <tbody>
      <template v-for="(item, index) in dailyLedger" :key="index">
        <tr class="divider">
          <td colspan="8">Asiento {{ index }}</td>
          <td>
            <button class="btn-publish" type="button">Publicar</button>
            <button class="btn-remove" type="button">Eliminar</button>
          </td>
        </tr>
        <template v-for="(entryExpanded, j) in dailyLedger[index].entriesExpanded" :key="j">
          <tr class="content">
            <td>{{ entryExpanded.account.displayName }}</td>
            <td>{{ entryExpanded.entry.currency }}</td>
            <td>{{ entryExpanded.entry.displayDate }}</td>
            <td>{{ entryExpanded.entry.displayDate }}</td>
            <td>{{ entryExpanded.entry.concept }}</td>
            <td>{{ entryExpanded.entry.debitAmount }}</td>
            <td>{{ entryExpanded.entry.creditAmount }}</td>
            <td>{{ entryExpanded.entry.displayDate }}</td>
            <td></td>
          </tr>
        </template>
        <tr>
          <td colspan="9" class="footer"></td>
        </tr>
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

.content:nth-of-type(even) {
  background-color: var(--gray);
}

.styled-table tbody tr:last-of-type {
  border-bottom: 2px solid var(--blue);
}

.styled-table tbody tr.divider {
  font-weight: bold;
  color: var(--blue);
}
.footer{
  background-color: var(--blue) !important;
  padding: 2px !important;
}
.btn-publish {
  border: none;
  padding: 5px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 15px;
  border-radius: 5px;
  background-color: var(--green);
  color: var(--white);
  margin-right: 10px;
  white-space: nowrap;
}
.btn-remove {
  border: none;
  padding: 5px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 15px;
  border-radius: 5px;
  background-color: var(--red);
  color: var(--white);
  margin-left: 10px;
  white-space: nowrap;
}
</style>
