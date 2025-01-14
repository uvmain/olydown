<script setup lang="ts">
import { useLocalStorage } from '@vueuse/core'
import { computed, ref } from 'vue'

defineProps({
  label: { type: String, required: true, default: '' },
})

const emit = defineEmits(['clicked'])

const cachedSaveLocation = useLocalStorage('saveLocation', '')

const buttonColour = computed(() => {
  return cachedSaveLocation.value ? 'bg-[#2EA169]' : 'bg-#b86944'
})
</script>

<template>
  <label>
    <div class="p-1rem text-center font-semibold rounded" :class="buttonColour" @click="emit('clicked')">
      <slot />
      <span>{{ label }}</span>
    </div>
  </label>
</template>
