<script lang="ts" setup>
import { computed, onMounted, ref, watch } from 'vue'
import { DownloadFile, GetDcimFolder, PopulateImages } from '../wailsjs/go/main/App'
import PhotoThumbnail from './components/PhotoThumbnail.vue'
import Button from './components/Button.vue'
import { useDark, useSessionStorage, useToggle } from '@vueuse/core'
import { useLocalStorage } from '@vueuse/core'

const cachedSaveLocation = useLocalStorage('saveLocation', '')
const dcimFolder = useSessionStorage('dcimFolder', '')
const imageList = ref<string[]>([])
const ssid = ref<string>('')
const isDark = useDark()
const toggleDark = useToggle(isDark)
const downloading = ref(false)
const successCount = ref(0)
const failCount = ref(0)
const skippedCount = ref(0)

async function getImageList() {
  imageList.value = []
  const result = await PopulateImages()
  imageList.value = result
}

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

async function downloadImages() {
  imageList.value.forEach(async (filename) => {
    const { Done, Skipped } = await DownloadFile(filename, cachedSaveLocation.value)
    if (Done) {
      successCount.value += 1
    }
    else if (Skipped) {
      skippedCount.value += 1
    }
    else {
      failCount.value += 1
    }
  })

  downloading.value = false
}

onMounted(async () => {
  downloading.value = false
  await (window as any).window.runtime.EventsOn('ssid:update', (newSsid: string) => {
    ssid.value = newSsid
  })
})

watch(ssid, async (newSsid) => {
  if (newSsid.startsWith('E-M')) {
    await getImageList()
    dcimFolder.value = await GetDcimFolder()
  }
})

const downloadingText = computed(() => {
  return downloading.value ? 'Downloading..' : 'Download Images'
})

const pendingCount = computed(() => {
  return imageList.value.length - successCount.value - failCount.value
})

</script>

<template>
  <div class="bg-light dark:bg-dark text-dark dark:text-light h-screen overflow-auto">
    <div class="p-3 flex flex-col gap-2">
      <div class="flex flex-wrap gap-2">
        <Button @clicked="getImageList" :disabled="!dcimFolder">
          Refresh Thumbnails
        </Button>
        <Button @click="handleDirectorySelection">
          <span>Selected Folder: {{ cachedSaveLocation }}</span>
        </Button>
        <Button @click="downloadImages" :disabled="(!dcimFolder && imageList.length < 1) || downloading">
          <span>{{ downloadingText }}</span>
        </Button>
        <Button @clicked="toggleDark()">
          <icon-material-symbols-light:dark-mode-outline v-if="isDark" />
          <icon-material-symbols-light:light-mode-outline v-else />
        </Button>
        <Button @clicked="getImageList" :good="ssid.startsWith('uv')">
          <icon-material-symbols-light:wifi class="pr-2"/>
          <span>{{ ssid }}</span>
        </Button>
      </div>
      <Button v-if="failCount > 0 || successCount > 0">
        <icon-material-symbols-light:check-circle class="pr-2 text-green" v-if="failCount === 0 && successCount > 0" />
        <icon-material-symbols-light:warning class="pr-2 text-orange" v-if="failCount > 0 && successCount > 0" />
        <icon-material-symbols-light:error class="pr-2 text-red" v-if="failCount > 0 && successCount === 0" />
        <span>Pending: {{ pendingCount }}, Success: {{ successCount }}, Failed: {{ failCount }}, Skipped: {{ skippedCount }}</span>
      </Button>
      <div v-if="imageList.length > 0" class="flex-1 overflow-y-scroll p-4">
        <div class="flex flex-wrap gap-4">
          <PhotoThumbnail v-for="path, index in imageList" :key="index" :filename="path" />
        </div>
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
