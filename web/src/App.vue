<template>
  <a-config-provider :locale="locale">
    <section id="app">
      <router-view :key="$route.fullPath"/>
    </section>
  </a-config-provider>
</template>

<script setup lang="ts">
import locale from 'ant-design-vue/lib/locale-provider/zh_CN';
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
dayjs.locale('zh-cn');

</script>
