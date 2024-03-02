// Composable useCalendarData.ts
import { onBeforeMount, ref } from "vue";
import type { Month } from "@/types"; // Asumiendo que tus interfaces están aquí
import { getLocalStorage } from "./localStorage";

export function useCalendarData() {
  const user = getLocalStorage("user");  
  const data = ref<Month[]>([]);

  async function fetchData() {
    try {
      const response = await fetch(`http://localhost:8080/getGasIng/${user.name}`);
      const jsonData: Month[] = await response.json();
      data.value = jsonData;
    } catch (error) {
      console.error("Error fetching calendar data:", error);
      data.value = [];
    }
  }

  onBeforeMount(fetchData);

  return { data };
}
