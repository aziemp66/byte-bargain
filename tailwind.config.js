/** @type {import('tailwindcss').Config} */
const withMT = require("@material-tailwind/html/utils/withMT");

module.exports = withMT({
  content: ["./web/views/**/*.html"],
  theme: {
    extend: {
    },
  },
  plugins: [],
})

