<template>
  <a-sub-menu :key="menuKey" v-if="hasSubmenu">
    <template #title>
      <span>
        <component :is="menuIcon" v-if="menuIcon"/>
        <span>{{ menuTitle }}</span>
      </span>
    </template>
    <slot></slot>
  </a-sub-menu>
  <a-menu-item :key="menuKey" v-else>
    <component :is="menuIcon" v-if="menuIcon"/>
    <span>{{ menuTitle }}</span>
  </a-menu-item>
</template>

<script setup lang="ts">

import { defineProps } from 'vue'

defineProps({
  menuTitle: {
    type: String,
    default: ""
  },
  menuKey: {
    type: String,
    default: ""
  },
  menuIcon: {
    type: String,
    default: ""
  },
  hasSubmenu:{
    type: Boolean,
    default: false
  }
})





</script>


