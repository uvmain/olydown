<script setup lang="ts">
import { useLocalStorage } from '@vueuse/core'
import { computed } from 'vue'

const cachedSaveLocation = useLocalStorage('saveLocation', '')

async function handleDirectorySelection() {
  await (window as any).window.go.main.App.OpenDirectoryDialog({
    defaultDirectory: "C:/",
    title: "Please select a folder to save Top and Bot CAD files..."
  }).then((path: string) => {
    cachedSaveLocation.value = path
  }).catch((err: any) => {
    console.log(err)
  });
}

const buttonColour = computed(() => {
  return cachedSaveLocation.value ? 'bg-[#2EA169]' : 'bg-#b86944'
})
</script>

<template>
  <label>
    <div class="p-1rem text-center font-semibold rounded" :class="buttonColour" @click="handleDirectorySelection">
      <span v-if="cachedSaveLocation">Selected Folder: {{ cachedSaveLocation }}</span>
      <span v-else>Select Folder</span>
    </div>
  </label>
</template>
