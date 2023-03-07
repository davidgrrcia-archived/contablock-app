<script setup lang="ts">
import { useConfigStore } from '@/stores/config';
import BusinessHeading from '@/features/business/adapters/views/BusinessHeading.vue';
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue';
import { UserCircleIcon } from '@heroicons/vue/24/outline';
import { getRoleColor } from '@/features/utils/utils';

const configStore = useConfigStore();
const themeColor = getRoleColor(configStore.state.role);
const businessName = configStore.state.name;

function logOut() {
  localStorage.setItem(configStore.state.project + 'Authentication', JSON.stringify(false));
  window.location.reload();
}
</script>

<template>
  <header class="fixed top-0 left-0 flex w-full items-center" :class="[themeColor]">
    <div class="flex w-1/2 items-center">
      <img alt="Vue logo" class="ml-5 object-cover" src="@/assets/logo.png" />
      <RouterLink class="text-lg" v-for="item in configStore.state.navigation" :key="item.name" :to="item.href" :class="['rounded-md p-2 font-normal text-white underline-offset-4 ' + (item.current ? 'underline' : 'hover:underline')]">{{ item.name }}</RouterLink>
    </div>
    <!-- Profile dropdown -->
    <Menu as="div" class="w-1/2 text-right">
      <div>
        <MenuButton class="btn mr-8 text-white hover:bg-gray-700">
          {{ businessName }}
          <UserCircleIcon class="ml-2 h-6 w-6" />
        </MenuButton>
      </div>
      <transition enter-active-class="transition ease-out duration-100" enter-from-class="transform opacity-0 scale-95" enter-to-class="transform opacity-100 scale-100" leave-active-class="transition ease-in duration-75" leave-from-class="transform opacity-100 scale-100" leave-to-class="transform opacity-0 scale-95">
        <MenuItems class="absolute right-8 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
          <MenuItem v-slot="{ active }">
            <a @click="logOut()" :class="[active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700']">Cerrar sesión</a>
          </MenuItem>
        </MenuItems>
      </transition>
    </Menu>
  </header>
  <suspense><BusinessHeading></BusinessHeading></suspense>
</template>
