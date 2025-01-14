<script lang="ts" setup>
import { computed, onMounted, ref, watch } from 'vue'
import { GetDcimFolder, PopulateImages } from '../wailsjs/go/main/App'
import PhotoThumbnail from './components/PhotoThumbnail.vue'
import Button from './components/Button.vue'
import DirectorySelector from './components/DirectorySelector.vue'
import { useLocalStorage } from '@vueuse/core'

const dcimFolder = useLocalStorage('dcimFolder', '')
const imageList = ref<string[]>([])
const ssid = ref<string>('')

async function getImageList() {
  imageList.value = []
  const result = await PopulateImages()
  imageList.value = result
}

onMounted(async () => {
  await (window as any).window.runtime.EventsOn('ssid:update', (newSsid: string) => {
    ssid.value = newSsid
  })
})

const wifiClass = computed(() => {
  return ssid.value.startsWith('E-M') ? 'text-okay-600' : 'text-warning-600'
})

watch(ssid, async (newSsid) => {
  if (newSsid.startsWith('E-M')) {
    await getImageList()
    dcimFolder.value = await GetDcimFolder()
  }
})

</script>

<template>
  <div class="bg-neutral-300 h-screen w-screen">
    <div class="p-3 flex flex-col gap-2">
      <DirectorySelector />
      <div class="grid grid-cols-2 gap-2">
        <Button label="Refresh Thumbnails" @clicked="getImageList" />
        <div :class="wifiClass">
          <icon-material-symbols-light:wifi class="px-2"/>
          <span>{{ ssid }}</span>
        </div>
      </div>
      <div v-if="imageList" class="flex flex-wrap gap-x-2 gap-y-1 overflow-y-scroll">
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
