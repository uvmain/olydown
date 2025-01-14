import { defineConfig } from 'unocss'

export default defineConfig({
  shortcuts: {
    button: 'border-1 border-neutral-700 rounded-sm border-solid px-4 py-2 outline-none dark:border-neutral-300 hover:bg-neutral-300 standard hover:text-dark hover:shadow-dark hover:shadow-md hover:shadow-op-20 hover:dark:bg-neutral-700 dark:hover:text-white dark:hover:shadow-white',
  },
  theme: {
    colors: {
      warning: {
        600: '#b86944'
      },
      okay: {
        600: '#2EA169'
      }
    },
  },
})
