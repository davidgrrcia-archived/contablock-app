<script setup lang="ts">
import { useEntriesStore } from '@/features/entries/application/entries-store';
import CreateEntryModal from '@/features/entries/adapters/views/CreateEntryModal.vue';
import ModifyEntryModal from '@/features/entries/adapters/views/ModifyEntryModal.vue';
import { isActionAllowed } from '@/features/utils/utils';
import { useConfigStore } from '@/stores/config';

const entriesStore = useEntriesStore();
const configStore = useConfigStore();

interface Props {
  title: string;
  description: string;
}

const cta = defineProps<{
  props: Props;
}>();
</script>

<template>
  <CreateEntryModal></CreateEntryModal>
  <ModifyEntryModal></ModifyEntryModal>
  <div class="m-8 flex items-center">
    <div class="flex-auto">
      <h1 class="text-3xl font-semibold text-gray-900">{{ cta.props.title }}</h1>
      <p class="mt-2 text-sm text-gray-700">{{ cta.props.description }}</p>
    </div>
    <div v-if="isActionAllowed(configStore.state.role, 'create')" class="w-1/2 text-right">
      <button type="button" @click="entriesStore.showCreationModal()" class="btn bg-blue-500 hover:bg-white hover:text-blue-500">Nuevo asiento</button>
    </div>
  </div>
</template>
