import "./assets/main.css";
import { createApp } from "vue";
import App from "./App.vue";

import { vuetify } from "./plugins/vuetify";
import { FontAwesomeIcon } from "./plugins/fontAwesome";
import router from "./router";
import { createPinia } from "pinia";
import HighChartsVue from 'highcharts-vue';
createApp(App)
  .component('font-awesome-icon', FontAwesomeIcon)
  .use(router)
  .use(vuetify)
  .use(HighChartsVue)
  .use(createPinia())
  .mount("#app");
