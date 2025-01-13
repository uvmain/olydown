<script lang="ts" setup>
import { ref } from 'vue'
import { GetImageScreennail, PopulateImages } from '../../wailsjs/go/main/App'

const data = ref()
const screennail = ref()

async function getImageList() {
  const result = await PopulateImages()
  data.value = result
}

async function getScreennail(filename: string) {
  const base64String = await GetImageScreennail(filename)
  screennail.value = `data:image/jpeg;base64,${base64String}`
}
</script>

<template>
  <main>
    <button @click="getImageList">Get Image List</button>
    <button v-if="data" v-for="path, index in data" :key="index" @click="getScreennail(path)">{{ path }}</button>
    <img v-if="screennail" :src="screennail" alt="Screennail" />
  </main>
</template>
