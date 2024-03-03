<script setup lang="ts">
import { ref, onBeforeMount } from 'vue';
import InitialScreen from './components/InitialScreen.vue';
import CalendarScreen from './components/CalendarScreen.vue';
import { useLocalStorage } from './composables/localStorage';
import Toaster from '@/components/ui/toast/Toaster.vue'

const user = useLocalStorage('user', { name: "", mode: "light" });
const initialScreen = ref({open:user.name === ''});

onBeforeMount(() => {
  if (user.mode === 'dark') {
    document.documentElement.classList.add('dark');
  } else {
    document.documentElement.classList.remove('dark');
  }
})

</script>

<template>
  <InitialScreen v-if="initialScreen.open" v-model:show="initialScreen.open"/>
  <CalendarScreen v-else />
  <Toaster />
</template>

