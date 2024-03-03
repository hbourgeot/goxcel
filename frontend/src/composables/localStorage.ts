import { watch, reactive } from "vue";

export function useLocalStorage(key: string, defaultValue: any) {
  const storedValue = localStorage.getItem(key);
  let data = reactive(storedValue ? JSON.parse(storedValue) : defaultValue);

  watch(
    () => data,
    (newValue) => {
      localStorage.setItem(
        key,
        JSON.stringify(
          newValue instanceof Map ? Object.fromEntries(newValue) : newValue
        )
      );

      if (newValue.mode === "dark") {
        document.documentElement.classList.add("dark");
      } else {
        document.documentElement.classList.remove("dark");
      }
    },
    { deep: true }
  );

  return data;
}

// get from local storage
export function getLocalStorage(key: string) {
  const storedValue = localStorage.getItem(key);
  return storedValue ? JSON.parse(storedValue) : null;
}