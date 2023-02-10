import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router';
import './assets/main.scss';

import PrimeVue from 'primevue/config';

import 'primevue/resources/themes/lara-light-blue/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';

import Calendar from 'primevue/calendar';
import InputNumber from 'primevue/inputnumber';
import TreeSelect from 'primevue/treeselect';

const pinia = createPinia();
const app = createApp(App);

app.use(router);
app.use(pinia);
app.use(PrimeVue);

app.component('CalendarPrimeVue', Calendar);
app.component('InputNumberPrimeVue', InputNumber);
app.component('TreeSelectPrimeVue', TreeSelect);

app.mount('#app');
