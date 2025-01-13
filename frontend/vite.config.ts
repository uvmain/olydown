import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import Unfonts from 'unplugin-fonts/vite'
import Icons from 'unplugin-icons/vite'
import UnoCSS from 'unocss/vite'


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Icons(),
    UnoCSS(),
    Unfonts({
      google: {
        families: [
          'Montserrat',
        ],
      },
    }),
  ]
})
