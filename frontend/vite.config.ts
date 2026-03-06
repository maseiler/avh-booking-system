import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    host: true
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '~': path.resolve(__dirname, './node_modules/'),
      '@schemas': path.resolve(__dirname, '../schemas'),
    }
  },
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: ''
      }
    }
  }
})
