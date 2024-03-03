<script setup lang="ts">
import { Ref, ref, watch, watchEffect } from 'vue';
import { Calendar } from '@/components/ui/calendar';
import { useCalendarData } from '@/composables/calendarData';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import ModeSwitch from './ModeSwitch.vue';
import { Month } from "@/types";

// Uso del composable para obtener los datos
const { data } = useCalendarData();


// ref constants
const dateSelected: Ref<Date> = ref(new Date());
const attributes: Ref<any[]> = ref([]);
const daySelected: Ref<number | null> = ref(null);
const monthSelected: Ref<number | null> = ref(null);
const form = ref({
  gastos: 0,
  ingresos: 0
});

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
          customData: {
            gastos: day.gastos,
            ingresos: day.ingresos
          },
          popover: {
            label: `Gastos: ${day.gastos}, Ingresos: ${day.ingresos}`,
          },
        }
      }

      return {
        key: `${month.month}-${day.day}`,
        dates: new Date(`${new Date().getFullYear()}-${month.month}-${day.day}`),
      }
    })
  );
});

watch(dateSelected, (newDate) => {
  console.log(newDate);
  if (newDate) {
    daySelected.value = newDate.getDate();
    monthSelected.value = newDate.getMonth();

    const month = data.value.find((month: Month) => month.month === monthNames[newDate.getMonth()]);
    const day = month.days.find((day: any) => day.day === newDate.getDate());

    form.value.gastos = day.gastos;
    form.value.ingresos = day.ingresos;

  }
});
</script>

<template>
  <div class="w-full p-4 flex justify-center gap-2 items-center flex-col">
    <Card class="bg-[#031030] p-1">
      <CardHeader>
        <CardTitle class="flex justify-between items-center">Seleccionar día <ModeSwitch/> </CardTitle>
        <CardDescription>Selecciona un día para ingresar los gastos e ingresos</CardDescription>
      </CardHeader>
      <CardContent>
        <Calendar class="w-full" :attributes="attributes" :min-date="new Date(2024, 0, 1)"
          :max-date="new Date(2024, 11, 31)" locale="es_ES" v-model="dateSelected" />
      </CardContent>
    </Card>
    <Card class="bg-[#031030] p-1 mt-4 w-full" v-show="dateSelected">
      <CardHeader>
        <CardTitle>{{ daySelected == 1 ? '1ero' : daySelected }} de {{ monthDisplayNames[monthSelected ?? 0] }}</CardTitle>
        <CardDescription>Resumen de los gastos e ingresos del mes</CardDescription>
      </CardHeader>
      <CardContent>
        <p class="text-white">Gastos: {{ form.gastos }}$</p>
        <p class="text-white">Ingresos: {{ form.ingresos }}$</p>
      </CardContent>
    </Card>
  </div>
</template>