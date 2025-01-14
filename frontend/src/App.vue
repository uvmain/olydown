<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { GetSSID, PopulateImages } from '../wailsjs/go/main/App'
import PhotoThumbnail from './components/PhotoThumbnail.vue'
import Button from './components/Button.vue'
import DirectorySelector from './components/DirectorySelector.vue'

const imageList = ref()
const ssid = ref()

async function getImageList() {
  const result = await PopulateImages()
  imageList.value = result
}

async function getSSID() {
  const result = await GetSSID()
  ssid.value = result
}

onMounted(async () => {
  await (window as any).window.runtime.EventsOn('ssid:update', (newSsid: string) => {
    ssid.value = newSsid
  })
})


</script>

<template>
  <div class="bg-neutral-300 h-screen w-screen">
    <div class="p-3 flex flex-col gap-2">
      <DirectorySelector />
      <div class="grid grid-cols-2 gap-2">
        <Button label="Get Image List" @clicked="getImageList" />
        <Button label="GetSSID" @clicked="getSSID"/>
      </div>
      {{ ssid }}
      <div v-if="imageList" class="flex flex-wrap gap-x-2 gap-y-1">
        <PhotoThumbnail v-for="path, index in imageList" :key="index" :filename="path" />
      </div>
    </div>
  </div>
</template>

<style>
html, body, #app {
  margin: 0;
  padding: 0;
  border: 0;
  font-family: 'Montserrat', sans-serif;
  min-height: 100%;
}
</style>
