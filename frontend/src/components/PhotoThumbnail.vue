<script setup lang="ts">
import { computed } from 'vue'
import { useSessionStorage } from '@vueuse/core'

const props = defineProps({
  filename: { type: String, required: true },
})

const dcimFolder = useSessionStorage('dcimFolder', '')

const imageSource = computed(() => {
  return dcimFolder.value ? `http://192.168.0.10/get_thumbnail.cgi?DIR=/DCIM/${dcimFolder.value}/${props.filename}` : './assets/default-image.jpg'
  // thumbnail does not include exif data, so we get a resized 1024 image instead
  // return dcimFolder.value ? `http://192.168.0.10/get_resizeimg.cgi?DIR=/DCIM/${dcimFolder.value}/${props.filename}&size=1024` : './assets/default-image.jpg'
})

</script>

<template>
  <div v-if="imageSource" class="relative">
    <img
      loading="lazy"
      :src="imageSource"
      :alt="filename"
      onerror="this.onerror=null;this.src='./assets/default-image.jpg';"
      class="cursor-pointer border-2 border-white border-solid object-cover dark:border-neutral-500 size-40 orientation"
    />
    <p class="absolute left-2 bottom-0 text-white bg-neutral-600">{{ filename }}</p>
  </div>
</template>

<style scoped>
.orientation {
    image-orientation: from-image;
}
</style>
