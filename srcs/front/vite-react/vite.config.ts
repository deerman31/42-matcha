import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0', // これが重要です
    port: 5173,
    //port: process.env.FRONT_PORT ? parseInt(process.env.FRONT_PORT) : 5173,
    watch: {
      usePolling: true
    }
  }
})
