/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx,vue}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: 'Inter, sans-serif'
      },
      colors: {
        bg: '#f6f5f4'
      }
    },
  },
  plugins: [],
}
