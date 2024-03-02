<script setup lang="ts">
import { Input } from '@/components/ui/input';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue
} from '@/components/ui/select';
import { Button } from '@/components/ui/button';
import { ref, watch } from 'vue';
import { useLocalStorage } from '../composables/localStorage';

const props = defineProps<{
  show: boolean;
}>();

const show = ref(props.show);
const user = useLocalStorage('user', { name: "", mode: "light" });


</script>
<template>
  <div class="w-full p-4 flex justify-center items-center flex-col">
    <h1 class="text-2xl lg:text-4xl">¡Bienvenido!</h1>
    <p class="text-lg lg:text-2xl text-center">Para comenzar, por favor ingresa tu nombre de usuario:</p>
    <Input type="text" class="w-1/2 lg:w-1/4 p-2 my-4" v-model:model-value="user.name" />
    <p class="text-lg lg:text-2xl text-center">También puedes cambiar el modo de la aplicación:</p>
    <Select v-model:model-value="user.mode">
      <SelectTrigger class="max-w-40">
        <SelectValue placeholder="Seleccionar" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectItem v-model:model-value="user.mode" value="light">Claro</SelectItem>
          <SelectItem v-model:model-value="user.mode" value="dark">Oscuro</SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
    <Button variant="default" @click="show = false">Enviar</Button>
  </div>
</template>
<style></style>