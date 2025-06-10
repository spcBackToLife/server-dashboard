import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // string shorthand: /api -> http://localhost:8080/api
      '/api': {
        target: 'http://localhost:8080', // Your Go backend address
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '/api') // Ensure /api prefix is maintained if backend expects it
      }
    }
  }
})
