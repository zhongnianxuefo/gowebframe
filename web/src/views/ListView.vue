<template>
  <main-layout>

    {{module}}
    <br/>
    List

    <br/>
    Content
    Content
    <br/>
    Content
    Content
    <br/>
    Content
    Content
    <br/>
    Content
    Content
    <br/>
    Content
    Content
    <br/>
    Content
    Content
    <br/>
    Content
    Content
    <br/>
    Content
    Content
    <br/>
    Content
    Content
    <br/>
    Content
    Content
    <br/>
    Content



  </main-layout>
</template>
<script setup>
import MainLayout from '@/components/MainLayout';
import {ref} from "vue";


import { useRoute } from 'vue-router'
const route = useRoute()



const module = ref( route.params["module"]);
// const selectedKey = 'list-'+route.params["module"];
// console.log(selectedKey,module)
</script>

