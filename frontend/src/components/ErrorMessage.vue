<script setup lang="ts">
import { useErrorMessagesStore } from '@/stores/errorMessages';

const errorMessagesStore = useErrorMessagesStore();
</script>

<template>
  <div>
    <transition-group class="message" name="p-message" tag="div">
      <MessagePrimeVue
        v-for="msg of errorMessagesStore.state.messages"
        :life="msg.life"
        :sticky="msg.sticky"
        :severity="msg.severity"
        :key="msg.id"
        >{{ msg.content }}</MessagePrimeVue
      >
    </transition-group>
  </div>
</template>

<style scoped>
.message {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 999999999;
}
</style>
