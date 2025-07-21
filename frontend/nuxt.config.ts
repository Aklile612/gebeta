import { fileURLToPath } from 'url'
import { dirname, resolve } from 'path'
import tsconfigPaths from 'vite-tsconfig-paths'

const __filename = fileURLToPath(import.meta.url)
const __dirname = dirname(__filename)

export default defineNuxtConfig({
  css: [resolve(__dirname, 'assets/css/main.css')],
  modules: [
    '@nuxt/image',
    '@nuxt/icon',
    '@nuxt/fonts',
    '@nuxtjs/tailwindcss', // ✅ Tailwind module handles everything
  ],
  vite: {
    plugins: [tsconfigPaths()],
  },
  compatibilityDate: '2025-07-21',
  devtools: { enabled: true },
})
