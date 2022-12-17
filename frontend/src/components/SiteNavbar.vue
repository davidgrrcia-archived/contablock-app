<template>
  <div class="flex bg-black">
    <div class="flex p-5">
      <img
        class="h-8 w-auto"
        src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
        alt="Your Company"
      />
      <div class="flex ml-6 space-x-4">
        <a
          v-for="item in navigation"
          :key="item.name"
          :href="item.href"
          :class="[
            item.current
              ? 'text-amber-400 underline decoration-2 underline-offset-8'
              : 'text-white hover:bg-amber-400 hover:text-black',
            'px-3 py-2 rounded-md text-sm font-medium',
          ]"
          :aria-current="item.current ? 'page' : undefined"
          >{{ item.name }}</a
        >
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const navigation = [
  { name: 'Dashboard', href: '#', current: true },
  { name: 'Team', href: '#', current: false },
  { name: 'Projects', href: '#', current: false },
  { name: 'Calendar', href: '#', current: false },
];
</script>
