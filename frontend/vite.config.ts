import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import Unfonts from 'unplugin-fonts/vite'
import Icons from 'unplugin-icons/vite'
import UnoCSS from 'unocss/vite'
import IconsResolver from 'unplugin-icons/resolver'
import Components from 'unplugin-vue-components/vite'


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [
        IconsResolver({
          prefix: 'icon',
        }),
      ],
    }),
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
