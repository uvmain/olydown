<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useLocalStorage } from '@vueuse/core'

const props = defineProps({
  filename: { type: String, required: true },
})

const dcimFolder = useLocalStorage('dcimFolder', '')

const imageSource = computed(() => {
  return dcimFolder.value ? `http://192.168.0.10/get_thumbnail.cgi?DIR=/DCIM/${dcimFolder.value}/${props.filename}` : '/default-image.jpg'
})

</script>

<template>
  <div v-if="imageSource" class="relative">
    <img
      loading="lazy"
      :src="imageSource"
      :alt="filename"
      onerror="this.onerror=null;this.src='/default-image.jpg';"
      class="cursor-pointer border-2 border-white border-solid object-cover dark:border-neutral-500 size-40"
    />
    <p class="absolute left-2 bottom-0 text-white bg-neutral-600">{{ filename }}</p>
  </div>
</template>
