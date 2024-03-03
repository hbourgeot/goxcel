<script setup lang="ts">
import { ref, onBeforeMount, watch } from 'vue';
import InitialScreen from './components/InitialScreen.vue';
import CalendarScreen from './components/CalendarScreen.vue';
import { useLocalStorage } from './composables/localStorage';

const user = useLocalStorage('user', { name: "", mode: "light" });
const showInitialScreen = ref(user.name === '');

onBeforeMount(() => {
  if (user.mode === 'dark') {
    document.documentElement.classList.add('dark');
  } else {
    document.documentElement.classList.remove('dark');
  }
})
</script>

<template>
  <InitialScreen v-if="showInitialScreen" :show="showInitialScreen" />
  <CalendarScreen v-else />
</template>

