/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx,vue}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: 'Cantarell, sans-serif'
      },
      colors: {
        bg: '#f6f5f4'
      }
    },
  },
  plugins: [],
}
