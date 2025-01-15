<script setup lang="ts">
defineProps({
  disabled: { type: Boolean, required: false, default: false },
})

const emit = defineEmits(['clicked'])

</script>

<template>
  <div class="flex items-center">
    <button
      type="button"
      :disabled="disabled"
      :aria-disabled="disabled"
      class="inline-flex cursor-pointer items-center justify-center whitespace-nowrap rounded-md text-sm font-semibold leading-4 shadow-sm disabled:cursor-not-allowed disabled:opacity-40 outline-none active:ring-2 active:ring-offset-2 px-3 py-2 text-gray-700 dark:text-gray-200 bg-gray-100 dark:bg-gray-600 hover:bg-gray-200 dark:hover:bg-gray-500 active:ring-gray-400 dark:active:ring-gray-600 border-none ring-gray-200 dark:ring-gray-700 ring-1 hover:ring-offset-1"
      @click="emit('clicked')"
    >
      <slot />
    </button>
  </div>
</template>
