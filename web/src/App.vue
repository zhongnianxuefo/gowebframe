<template>
  <a-config-provider :locale="locale">
    <section id="app">
      <router-view />
    </section>
  </a-config-provider>
</template>


<script>
import locale from 'ant-design-vue/lib/locale-provider/zh_CN';
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
dayjs.locale('zh-cn');


export default {
  name: 'App',
  setup() {
    return { locale };
  },
};
</script>

