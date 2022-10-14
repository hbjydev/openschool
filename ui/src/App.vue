<script setup lang="ts">
import Header from './components/Header.vue';
import { RouterView } from 'vue-router';
</script>

<template>
  <Header />
  <div class="p-6">
    <RouterView />
  </div>
</template>
