<script lang="ts" setup>
import { computed, ref } from 'vue'
import { GetImageScreennail, PopulateImages } from '../wailsjs/go/main/App'

const imageList = ref()
const screennail = ref()
const orientation = ref()

async function getImageList() {
  const result = await PopulateImages()
  imageList.value = result
}

async function getScreennail(filename: string) {
  const response = await GetImageScreennail(filename)
  const base64String = response.Base64string
  orientation.value = response.Orientation
  screennail.value = `data:image/jpeg;base64,${base64String}`
}

const rotationClass = computed(() => {
  return orientation.value == "90" ? "rotate-270" : ""
})
</script>

<template>
  <div class="bg-neutral h-screen w-screen">
    <div class="p-3 flex flex-col gap-2">
      <button class="button mr-auto" @click="getImageList">Get Image List</button>
      <div class="flex flex-wrap gap-2">
        <button class="button h-8 w-30" v-if="imageList" v-for="path, index in imageList" :key="index" @click="getScreennail(path)">{{ path }}</button>
      </div>
      <img v-if="screennail" :src="screennail" alt="Screennail" class="w-9/10 object-cover" :class="rotationClass" />
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
