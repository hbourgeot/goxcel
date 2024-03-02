<script setup lang="ts">
import { Calendar } from '@/components/ui/calendar';
import { useCalendarData } from '@/composables/calendarData';
import { Day, Month } from '@/types';
import { Ref, ref, watchEffect } from 'vue';
const {data} = useCalendarData();
const calendarEvents: Ref<{ date: Date, title: string, details: string }[]> = ref([]);

// Observa cambios en `data` y recalcula los eventos del calendario
watchEffect(() => {
  calendarEvents.value = [];

  data.value.forEach((month: Month) => {
    month.days.forEach((day: Day) => {
      if (day.gastos > 0 || day.ingresos > 0) {
        const date = new Date(`${month.month}-${day.day}`); // Asegúrate de que el formato de fecha sea correcto
        calendarEvents.value.push({
          date: date as Date,
          title: 'Datos disponibles',
          details: `Gastos: ${day.gastos}, Ingresos: ${day.ingresos}`,
        } as { date: Date, title: string, details: string });
      }
    });
  });
});
</script>
<template>
  <div class="w-full p-2 flex justify-center items-center flex-col">
    <Calendar class="w-screen" :events="calendarEvents"/>
  </div>
</template>