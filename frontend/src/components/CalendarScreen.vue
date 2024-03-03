<script setup lang="ts">
import { Ref, ref, watch, watchEffect } from 'vue';
import { Calendar } from '@/components/ui/calendar';
import { useCalendarData } from '@/composables/calendarData';
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import ModeSwitch from './ModeSwitch.vue';
import { Month } from "@/types";
import { Button } from './ui/button';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger, DialogClose } from './ui/dialog';
import { Input } from './ui/input';
import { getLocalStorage } from '@/composables/localStorage';
import { useToast } from './ui/toast';

// Uso del composable para obtener los datos
const { data, fetchData } = useCalendarData();
const user = getLocalStorage('user');
const { toast } = useToast();


// ref constants
const dateSelected: Ref<Date> = ref(new Date());
const attributes: Ref<any[]> = ref([]);
const daySelected: Ref<number | null> = ref(null);
const monthSelected: Ref<number | null> = ref(null);
const form = ref({
  gastos: 0,
  ingresos: 0
});
const gastos = ref(0);
const ingresos = ref(0);

// normal constants
const monthDisplayNames = [
  'Enero', 'Febrero', 'Marzo', 'Abril', 'Mayo', 'Junio',
  'Julio', 'Agosto', 'Septiembre', 'Octubre', 'Noviembre', 'Diciembre'
];

const monthNames = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];

// Convertir los datos en atributos para el calendario
watchEffect(() => {
  attributes.value = data.value.flatMap(month =>
    month.days.map(day => {
      if (day.gastos > 0 || day.ingresos > 0) {
        return {
          key: `${month.month}-${day.day}`,
          dates: new Date(`${new Date().getFullYear()}-${month.month}-${day.day}`),
          highlight: true,
        }
      }

      return {
        key: `${month.month}-${day.day}`,
        dates: new Date(`${new Date().getFullYear()}-${month.month}-${day.day}`),
      }
    })
  );

  if (daySelected.value) {
    const month = data.value.find((month: Month) => month.month === monthNames[monthSelected.value ?? 0]);
    const day = month?.days.find((day: any) => day.day === dateSelected.value.getDate());

    gastos.value = day?.gastos ?? 0;
    ingresos.value = day?.ingresos ?? 0;
  }
});

watch(dateSelected, (newDate) => {
  if (newDate) {
    daySelected.value = newDate.getDate();
    monthSelected.value = newDate.getMonth();

    const month = data.value.find((month: Month) => month.month === monthNames[newDate.getMonth()]);
    const day = month?.days.find((day: any) => day.day === newDate.getDate());

    gastos.value = day?.gastos ?? 0;
    ingresos.value = day?.ingresos ?? 0;

  }
});

async function handleAdd() {
  const month = data.value.find((month: Month) => month.month === monthNames[monthSelected.value ?? 0]);
  const day = month?.days.find((day: any) => day.day === dateSelected.value.getDate());

  const response = await fetch(`/api/appendDay/${user.name}/${month?.month}-${day?.day}/-${form.value.gastos}-${form.value.ingresos}`, {
    method: 'POST'
  });

  if ([200, 201].includes(response.status)) {
    toast({
      title: '¡Listo!',
      description: "Se han agregado los gastos e ingresos correctamente",
    })
    fetchData();
  } else {
    toast({
      title: 'Error',
      description: "Hubo un error al agregar los gastos e ingresos",
      variant: 'destructive'
    })
  }

  form.value.gastos = 0;
  form.value.ingresos = 0;
  await fetchData();
}
</script>

<template>
  <div class="w-full p-4 flex justify-center gap-2 items-center flex-col">
    <Card class="bg-[#031030] p-1">
      <CardHeader>
        <CardTitle class="flex justify-between items-center">Seleccionar día
          <ModeSwitch />
        </CardTitle>
        <CardDescription>Selecciona un día para ingresar los gastos e ingresos</CardDescription>
      </CardHeader>
      <CardContent>
        <Calendar class="w-full" :attributes="attributes" :min-date="new Date(2024, 0, 1)"
          :max-date="new Date(2024, 11, 31)" locale="es_ES" v-model="dateSelected" />
      </CardContent>
    </Card>
    <Card class="bg-[#031030] p-1 mt-4 w-full" v-show="daySelected">
      <CardHeader>
        <CardTitle>{{ daySelected == 1 ? '1ero' : daySelected }} de {{ monthDisplayNames[monthSelected ?? 0] }}
        </CardTitle>
        <CardDescription>Resumen de los gastos e ingresos del mes</CardDescription>
      </CardHeader>
      <CardContent>
        <p class="text-white">Gastos: {{ gastos }}$</p>
        <p class="text-white">Ingresos: {{ ingresos }}$</p>
      </CardContent>
      <CardFooter>
        <Dialog>
          <DialogTrigger class="w-full">
            <Button variant="outline" class="p-2 rounded-md w-full hover:bg-white hover:text-black">Agregar
              Gastos/Ingresos</Button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Agregar</DialogTitle>
              <DialogDescription>
                Agrega nuevos gastos e ingresos para el día seleccionado, estos se sumarán a los ya existentes
              </DialogDescription>
            </DialogHeader>
            <div class="flex flex-col gap-2">
              <label for="gastos" class="text-white">Gastos</label>
              <Input type="number" v-model="form.gastos" />
            </div>
            <div class="flex flex-col gap-2">
              <label for="ingresos" class="text-white">Ingresos</label>
              <Input type="number" v-model="form.ingresos" />
            </div>

            <DialogFooter>
              <div class="flex gap-x-3 justify-between flex-nowrap">
                <DialogClose as-child class="w-1/2">
                  <Button variant="ghost"
                    class="p-2 rounded-md w-full">Cancelar</Button>
                </DialogClose>
                <DialogClose as-child class="w-1/2">
                  <Button @click="handleAdd" variant="secondary" class="p-2 rounded-md w-full">Agregar</Button>
                </DialogClose>
              </div>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </CardFooter>
    </Card>
  </div>
</template>