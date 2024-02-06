/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pkg/template/**/*.{html,js,templ,go}"
  ],
  theme: {
    extend: {},
  },
  //plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};