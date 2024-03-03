<template>
  <div>
    <DropdownMenu>
      <DropdownMenuTrigger v-slot="{ open }" asChild>
        <Button variant="outline" size="icon">
          <Sun class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
            v-if="user.mode == 'light'" />
          <Moon class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
            v-else />
          <span class="sr-only">Cambiar tema</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" v-show="open">
        <DropdownMenuItem @click="setTheme('light')">Light</DropdownMenuItem>
        <DropdownMenuItem @click="setTheme('dark')">Dark</DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { DropdownMenu, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuItem } from '@/components/ui/dropdown-menu';
import { Button } from '@/components/ui/button';
import { Moon, Sun } from 'lucide-vue-next';
import { useLocalStorage } from '@/composables/localStorage';

const user = useLocalStorage('user', { name: '', mode: 'light' });

const open = ref(false);
const setTheme = (theme: string) => {
  user.mode = theme;
};
</script>

<style>
/* Add your custom styles here */
</style>