import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router';
import './assets/main.scss';

import PrimeVue from 'primevue/config';

import 'primevue/resources/themes/lara-light-blue/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';

import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import AutoComplete from 'primevue/autocomplete';
import Dropdown from 'primevue/dropdown';
import Divider from 'primevue/divider';
import Calendar from 'primevue/calendar';
import Textarea from 'primevue/textarea';
import InputNumber from 'primevue/inputnumber';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Message from 'primevue/message';

const pinia = createPinia();
const app = createApp(App);

app.use(router);
app.use(pinia);
app.use(PrimeVue);

app.component('DialogPrimeVue', Dialog);
app.component('InputTextPrimeVue', InputText);
app.component('ButtonPrimeVue', Button);
app.component('AutoCompletePrimeVue', AutoComplete);
app.component('DropdownPrimeVue', Dropdown);
app.component('DividerPrimeVue', Divider);
app.component('CalendarPrimeVue', Calendar);
app.component('TextareaPrimeVue', Textarea);
app.component('InputNumberPrimeVue', InputNumber);
app.component('DataTablePrimeVue', DataTable);
app.component('ColumnPrimeVue', Column);
app.component('MessagePrimeVue', Message);

app.mount('#app');
