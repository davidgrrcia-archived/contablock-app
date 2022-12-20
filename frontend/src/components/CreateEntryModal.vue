<script setup lang="ts">
import { useCreateEntryModalStore } from '@/stores/createEntryModal';

const modalState = useCreateEntryModalStore();

function hideCreateEntryModal() {
  modalState.hideModal();
}
</script>

<template>
  <div v-if="modalState.isOpen" class="modal" id="modal">
    <div class="modal-content">
      <div class="modal-header">
        <h2>Crear Nuevo Asiento</h2>
        <button @click="hideCreateEntryModal()">&times;</button>
      </div>
      <div class="modal-body">
        <form>
          <label for="title">Titulo:</label><br />
          <input type="text" id="title" name="title" /><br />
          <label for="description">Descripcion:</label><br />
          <textarea id="description" name="description"></textarea><br />
          <button class="btn-submit" type="submit">Crear</button>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
.modal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 1;
}

.modal-content {
  background-color: white;
  border: 1px solid #888;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  border-radius: 5px;
  overflow: hidden;
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px;
}

.modal-header button {
  color: red;
}

.modal-body {
  padding: 20px;
}

.close-button {
  appearance: none;
  background-color: transparent;
}
</style>
