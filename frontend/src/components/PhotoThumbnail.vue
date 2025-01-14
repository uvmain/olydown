<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { GetImageThumbnail } from '../../wailsjs/go/main/App'

const props = defineProps({
  filename: { type: String, required: true },
})

const imgSource = ref('/default-image.jpg')
const orientation = ref("0")

async function getImageThumbnail(filename: string) {
  const response = await GetImageThumbnail(filename)
  const base64String = response.Base64string
  orientation.value = response.Orientation
  imgSource.value = `data:image/jpeg;base64,${base64String}`
}

const rotationClass = computed(() => {
  return orientation.value == "90" ? "rotate-270" : ""
})

onMounted(() => {
  getImageThumbnail(props.filename)
})
</script>

<template>
  <div v-if="imgSource" class="relative">
    <img
      :src="imgSource"
      :alt="filename"
      onerror="this.onerror=null;this.src='/default-image.jpg';"
      class="cursor-pointer border-2 border-white border-solid object-cover dark:border-neutral-500 size-40"
      :class="rotationClass"
    />
    <p class="absolute left-2 bottom-0 text-white bg-neutral-600">{{ filename }}</p>
  </div>
</template>
