<template>
  <div>
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="outline" size="icon">
          <Menu
            class="h-[1.2rem] w-[1.2rem] dark:text-white text-black" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" v-show="open">
        <DropdownMenuLabel>Menú</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuLabel>Temas</DropdownMenuLabel>
          <DropdownMenuItem @click="setTheme('light')">Claro</DropdownMenuItem>
          <DropdownMenuItem @click="setTheme('dark')">Oscuro</DropdownMenuItem>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuItem @click="reload">Salir</DropdownMenuItem>
      
      </DropdownMenuContent>
    </DropdownMenu>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { DropdownMenu, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuItem, DropdownMenuGroup, DropdownMenuLabel, DropdownMenuSeparator } from '@/components/ui/dropdown-menu';
import { Button } from '@/components/ui/button';
import { Menu } from 'lucide-vue-next';
import { useLocalStorage } from '@/composables/localStorage';

const user = useLocalStorage('user', { name: '', mode: 'light' });

const open = ref(false);
const setTheme = (theme: string) => {
  user.mode = theme;
};

const reload = () => {
  localStorage.clear();
  window.location.reload();
};
</script>

<style>
/* Add your custom styles here */
</style>