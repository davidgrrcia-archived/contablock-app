import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router';
import './assets/main.css';

import PrimeVue from 'primevue/config';
import 'primevue/resources/themes/lara-light-blue/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';
import Calendar from 'primevue/calendar';
import Dropdown from 'primevue/dropdown';
import AutoComplete from 'primevue/autocomplete';
import { useConfigStore } from './stores/config';

const pinia = createPinia();
const app = createApp(App);

app.use(pinia);
app.use(router);
app.use(PrimeVue);

//Activate pinia store before each route
router.beforeEach((to, from, next) => {
  const configStore = useConfigStore();

  if (typeof to.params.project === 'string') {
    configStore.activateCurrent(to.params.project, to.path);
  }

  next();
});

app.component('CalendarPrimeVue', Calendar);
app.component('DropdownPrimeVue', Dropdown);
app.component('AutoCompletePrimeVue', AutoComplete);

app.mount('#app');
