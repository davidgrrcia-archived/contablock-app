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
import TreeSelect from 'primevue/treeselect';

const pinia = createPinia();
const app = createApp(App);

app.use(pinia);
app.use(router);
app.use(PrimeVue);

app.component('CalendarPrimeVue', Calendar);
app.component('TreeSelectPrimeVue', TreeSelect);

app.mount('#app');
