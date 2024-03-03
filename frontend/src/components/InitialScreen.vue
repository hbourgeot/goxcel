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
import { ref } from 'vue';
import { useLocalStorage } from '../composables/localStorage';
import { useToast } from './ui/toast';

const {toast} = useToast();

const props = defineProps<{
  show: boolean;
}>();

const emit = defineEmits(['update:show']);

function close() {
  emit('update:show', false);
}

const show = ref(props.show);
const user = useLocalStorage('user', { name: "", mode: "light" });

const init = async () => {
  const response = await fetch(`http://localhost:8080/initGoxcel/${user.name}`, {
    method: 'POST'
  });

  if ([200, 201].includes(response.status)) {
    close();
    toast({
      title: '¡Bienvenido!',
      description: "Iniciaste sesión correctamente",
    })
  } else {
    toast({
      title: 'Error',
      description: "Hubo un error al iniciar sesión",
      variant: 'destructive'
    })
  }

  show.value = false;
};

</script>
<template>
  <div class="w-full h-screen p-4 flex justify-center items-center flex-col text-center gap-y-1">
    <h1 class="text-3xl">¡Bienvenido!</h1>
    <p class="text-sm mb-2">Esta es una aplicación para llevar el control de tus gastos e ingresos. Ten en cuenta que está pensada para dispositivos móviles, por ello puede que te encuentres con diseños no adaptables a pantallas de tablets o laptops.</p>
    <h2 class="text-lg text-center">Para comenzar, por favor ingresa tu nombre de usuario:</h2>
    <Input type="text" class="w-full lg:w-1/4 p-2 my-4 border-primary" v-model:model-value="user.name" />
    <h2 class="text-base text-center">También puedes cambiar aquí el modo de la aplicación:</h2>
    <Select v-model:model-value="user.mode">
      <SelectTrigger class="w-full">
        <SelectValue placeholder="Seleccionar" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectItem v-model:model-value="user.mode" value="light">Claro</SelectItem>
          <SelectItem v-model:model-value="user.mode" value="dark">Oscuro</SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
    <Button variant="default" @click="init" class="w-full my-2">Enviar</Button>
  </div>
</template>
<style></style>