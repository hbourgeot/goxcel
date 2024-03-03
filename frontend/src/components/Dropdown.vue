<template>
  <div>
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="outline" size="icon">
          <Menu class="h-[1.2rem] w-[1.2rem] dark:text-white text-black" />
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
        <DropdownMenuItem @click="downloadFile" class="flex gap-x-2 justify-between w-full flex-row-reverse">
          <DownloadCloud class="h-[1.2rem] w-[1.2rem] dark:text-white text-black" />
          Descargar Excel
        </DropdownMenuItem>
        <DropdownMenuItem @click="downloadTemplate" class="flex gap-x-2 justify-between w-full flex-row-reverse">
          <DownloadCloud class="h-[1.2rem] w-[1.2rem] dark:text-white text-black" />
          Descargar Plantilla
        </DropdownMenuItem>
        <DropdownMenuItem @click="reload">Salir</DropdownMenuItem>

      </DropdownMenuContent>
    </DropdownMenu>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { DropdownMenu, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuItem, DropdownMenuGroup, DropdownMenuLabel, DropdownMenuSeparator } from '@/components/ui/dropdown-menu';
import { Button } from '@/components/ui/button';
import { DownloadCloud, Menu } from 'lucide-vue-next';
import { useLocalStorage } from '@/composables/localStorage';
import { useToast } from './ui/toast';

const user = useLocalStorage('user', { name: '', mode: 'light' });
const { toast } = useToast();

const open = ref(false);
const setTheme = (theme: string) => {
  user.mode = theme;
};

const reload = () => {
  localStorage.clear();
  window.location.reload();
};

const downloadFile = async () => {
  const response = await fetch(`/api/downloadFile/${user.name}`, {
    method: 'GET',
  });

  if (response.ok) {
    const blob = await response.blob();
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `gastos_ingresos_${user.name}.xlsx`;
    a.click();
    window.URL.revokeObjectURL(url);
  } else {
    toast({
      title: 'Error al descargar el archivo',
      description: 'Ocurrió un error al descargar el archivo',
      variant: 'destructive'
    });
  }
};

const downloadTemplate = async () => {
  const response = await fetch(`/api/downloadTemplate/${user.name}`, {
    method: 'GET',
  });

  if (response.ok) {
    const blob = await response.blob();
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `plantilla_gastos_ingresos.xlsx`;
    a.click();
    window.URL.revokeObjectURL(url);
  } else {
    toast({
      title: 'Error al descargar el archivo',
      description: 'Ocurrió un error al descargar el archivo',
      variant: 'destructive'
    });
  }
};
</script>

<style>
/* Add your custom styles here */
</style>